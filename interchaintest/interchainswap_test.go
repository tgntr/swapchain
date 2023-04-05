package interchaintest

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"testing"

	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	"github.com/strangelove-ventures/interchaintest/v4/testreporter"
	"github.com/strangelove-ventures/interchaintest/v4/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestInterchainSwap(t *testing.T) {
	t.Parallel()

	// setup docker client
	client, network := interchaintest.DockerSetup(t)

	// build relayer
	r := interchaintest.
		NewBuiltinRelayerFactory(ibc.CosmosRly, zaptest.NewLogger(t)).
		Build(t, client, network)

	// setup chains
	chainSpecs := []*interchaintest.ChainSpec{swapchainSpec, osmosisSpec}
	chains, err := interchaintest.
		NewBuiltinChainFactory(zaptest.NewLogger(t), chainSpecs).
		Chains(t.Name())
	require.NoError(t, err)

	swapchain, osmosis := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	const transferPath, icqPath = "transfer-path", "icq-path"
	ic := interchaintest.NewInterchain().
		AddChain(swapchain).
		AddChain(osmosis).
		AddRelayer(r, "relayer").
		AddLink(interchaintest.InterchainLink{
			Chain1:  swapchain,
			Chain2:  osmosis,
			Relayer: r,
			Path:    transferPath,
		}).
		AddLink(interchaintest.InterchainLink{
			Chain1:  swapchain,
			Chain2:  osmosis,
			Relayer: r,
			Path:    icqPath,
			CreateChannelOpts: ibc.CreateChannelOptions{
				SourcePortName: "interchainswap",
				DestPortName:   "icqhost",
				Version:        "icq-1",
				Order:          ibc.Unordered,
			},
		})

	// create log reporter
	eRep := testreporter.NewNopReporter().RelayerExecReporter(t)

	ctx := context.Background()

	// build chains
	err = ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,
	})
	require.NoError(t, err)

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// swapchain common cli flags
	swapchainFlags := []string{
		"--node", swapchain.GetRPCAddress(),
		"--home", swapchain.HomeDir(),
		"--chain-id", swapchain.Config().ChainID,
		"--output", "json",
	}

	// start relayer
	err = r.StartRelayer(ctx, eRep, transferPath, icqPath)
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := r.StopRelayer(ctx, eRep); err != nil {
			t.Logf("an error occured while stopping the relayer: %s", err)
		}
	})

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// get connections/channels
	channels, err := r.GetChannels(ctx, eRep, swapchain.Config().ChainID)
	require.NoError(t, err)
	require.Len(t, channels, 2)

	var transferCh, icqCh, icaConn string
	for _, ch := range channels {
		switch ch.PortID {
		case transfertypes.PortID:
			transferCh = ch.ChannelID
			icaConn = ch.ConnectionHops[0]
		case "interchainswap":
			icqCh = ch.ChannelID
		default:
			t.FailNow()
		}
	}
	require.NotEmpty(t, transferCh)
	require.NotEmpty(t, icqCh)
	require.NotEmpty(t, icaConn)

	// create and fund user wallets
	users := interchaintest.GetAndFundTestUsers(t, ctx, "default", int64(200_000_000_000), swapchain, osmosis)
	swapchainUser, osmosisUser := users[0], users[1]

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// IBC transfer swapchain tokens to osmosisUser
	osmosisUserAddr := osmosisUser.Bech32Address(osmosis.Config().Bech32Prefix)
	require.NoError(t, err)

	transferDetails := ibc.WalletAmount{
		Address: string(osmosisUserAddr),
		Denom:   swapchain.Config().Denom,
		Amount:  int64(1_000_000),
	}
	ibcTransferResult, err := swapchain.SendIBCTransfer(ctx, transferCh, swapchainUser.KeyName, transferDetails, ibc.TransferOptions{})
	require.NoError(t, err)
	require.NoError(t, ibcTransferResult.Validate())

	// wait for transfer to be relayed
	swapchainHeight, err := swapchain.Height(ctx)
	require.NoError(t, err)

	_, err = testutil.PollForAck(ctx, swapchain, swapchainHeight, swapchainHeight+30, ibcTransferResult.Packet)
	require.NoError(t, err)

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// get IBC denom for swapchain token
	prefixedDenom := transfertypes.GetPrefixedDenom(transfertypes.PortID, transferCh, swapchain.Config().Denom)
	swapchainTokenIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()

	// verify osmosisUser balance
	osmosisUserBalance, err := osmosis.GetBalance(ctx, transferDetails.Address, swapchainTokenIbcDenom)
	require.NoError(t, err)
	require.Equal(t, transferDetails.Amount, osmosisUserBalance)

	// create osmosis liquidity pool for swapchain token - uosmo with 1:1 ratio
	poolID, err := cosmos.OsmosisCreatePool(osmosis, ctx, osmosisUser.KeyName, cosmos.OsmosisPoolParams{
		Weights:        fmt.Sprintf("1%s,1%s", swapchainTokenIbcDenom, osmosis.Config().Denom),
		InitialDeposit: fmt.Sprintf("1000%s,1000%s", swapchainTokenIbcDenom, osmosis.Config().Denom),
		SwapFee:        "0.01",
		ExitFee:        "0.01",
	})
	require.NoError(t, err)
	require.Equal(t, "1", poolID)

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// register ICA account
	cmdRegisterICA := append([]string{swapchain.Config().Bin, "tx",
		"interchainswap", "register-interchain-account", icaConn,
		"--from", swapchainUser.KeyName,
		"--keyring-backend", "test",
		"-y",
	}, swapchainFlags...)
	registerICAStdout, _, err := swapchain.Exec(ctx, cmdRegisterICA, nil)
	require.NoError(t, err)

	registerICAResponse := cosmos.CosmosTx{}
	err = json.Unmarshal([]byte(registerICAStdout), &registerICAResponse)
	require.NoError(t, err)
	require.Zero(t, registerICAResponse.Code)

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// query ICA account address
	creatorAddr := swapchainUser.Bech32Address(swapchain.Config().Bech32Prefix)

	cmdQueryICAAddr := append([]string{swapchain.Config().Bin, "q",
		"interchain-accounts", "controller", "interchain-account", creatorAddr, icaConn,
	}, swapchainFlags...)
	queryICAAddrStdout, _, err := swapchain.Exec(ctx, cmdQueryICAAddr, nil)
	require.NoError(t, err)

	var queryICAAddrResponse struct{ Address string }
	err = json.Unmarshal([]byte(queryICAAddrStdout), &queryICAAddrResponse)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(queryICAAddrResponse.Address, osmosis.Config().Bech32Prefix))

	// fund ICA address with uosmo
	fundDetails := ibc.WalletAmount{
		Address: queryICAAddrResponse.Address,
		Denom:   osmosis.Config().Denom,
		Amount:  int64(1_000_000),
	}
	err = osmosis.SendFunds(ctx, osmosisUser.KeyName, fundDetails)
	require.NoError(t, err)

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// verify ica address balance
	icaAddrBalance, err := osmosis.GetBalance(ctx, queryICAAddrResponse.Address, osmosis.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, fundDetails.Amount, icaAddrBalance)

	// ICA swap uosmo for swapchain token
	const swapAmount, tokenOutMinAmount = int64(10), int64(1)
	msgOsmosisSwapJson := fmt.Sprintf(`{
		"@type": "/osmosis.gamm.v1beta1.MsgSwapExactAmountIn",
		"sender": "%s",
		"routes": [{"poolId": "%s", "tokenOutDenom": "%s"}],
		"tokenIn": {"denom": "%s","amount": "%d"},
		"tokenOutMinAmount": "%d"
	}`, queryICAAddrResponse.Address, poolID, swapchainTokenIbcDenom, osmosis.Config().Denom, swapAmount, tokenOutMinAmount)

	cmdICAOsmosisSwap := append([]string{swapchain.Config().Bin, "tx",
		"interchainswap", "send-msg-osmosis-swap", msgOsmosisSwapJson, icaConn,
		"--from", swapchainUser.KeyName,
		"--keyring-backend", "test",
		"-y",
	}, swapchainFlags...)
	icaOsmosisSwapStdout, _, err := swapchain.Exec(ctx, cmdICAOsmosisSwap, nil)
	require.NoError(t, err)

	icaSwapOsmosisResponse := cosmos.CosmosTx{}
	err = json.Unmarshal([]byte(icaOsmosisSwapStdout), &icaSwapOsmosisResponse)
	require.NoError(t, err)
	require.Zero(t, icaSwapOsmosisResponse.Code)

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// verify ICA address balances after swap
	icaAddrUosmoBalance, err := osmosis.GetBalance(ctx, queryICAAddrResponse.Address, osmosis.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, fundDetails.Amount-swapAmount, icaAddrUosmoBalance)

	icaAddrSwapchainTokenBalance, err := osmosis.GetBalance(ctx, queryICAAddrResponse.Address, swapchainTokenIbcDenom)
	require.NoError(t, err)
	require.GreaterOrEqual(t, icaAddrSwapchainTokenBalance, tokenOutMinAmount)

	// ICQ osmosis spot price for swapchain token
	cmdICQOsmosisSpotPrice := append([]string{swapchain.Config().Bin, "tx",
		"interchainswap", "send-query-osmosis-spot-price", icqCh, poolID, osmosis.Config().Denom, swapchainTokenIbcDenom,
		"--from", swapchainUser.KeyName,
		"--keyring-backend", "test",
		"-y",
	}, swapchainFlags...)
	icqOsmosisSpotPriceStdout, _, err := swapchain.Exec(ctx, cmdICQOsmosisSpotPrice, nil)
	require.NoError(t, err)

	icqOsmosisSpotPriceResponse := cosmos.CosmosTx{}
	err = json.Unmarshal([]byte(icqOsmosisSpotPriceStdout), &icqOsmosisSpotPriceResponse)
	require.NoError(t, err)
	require.Zero(t, icqOsmosisSpotPriceResponse.Code)

	err = testutil.WaitForBlocks(ctx, 5, swapchain, osmosis)
	require.NoError(t, err)

	// query ICQ state
	cmdQueryICQState := append([]string{swapchain.Config().Bin, "q",
		"interchainswap", "query-state", "1",
	}, swapchainFlags...)
	queryICQStateStdout, _, err := swapchain.Exec(ctx, cmdQueryICQState, nil)
	require.NoError(t, err)

	var queryICQStateResponse struct {
		Request struct {
			PoolID          string `json:"pool_id"`
			BaseAssetDenom  string `json:"base_asset_denom"`
			QuoteAssetDenom string `json:"quote_asset_denom"`
		}
		Response struct {
			SpotPrice float64 `json:"spot_price,string"`
		}
	}
	err = json.Unmarshal([]byte(queryICQStateStdout), &queryICQStateResponse)
	require.NoError(t, err)
	require.Equal(t, poolID, queryICQStateResponse.Request.PoolID)
	require.Equal(t, osmosis.Config().Denom, queryICQStateResponse.Request.BaseAssetDenom)
	require.Equal(t, swapchainTokenIbcDenom, queryICQStateResponse.Request.QuoteAssetDenom)

	// pool no longer has 1:1 ratio, therefore spot price for swapchain token should be less than 1uosmo
	require.Less(t, queryICQStateResponse.Response.SpotPrice, float64(1))
	require.Greater(t, queryICQStateResponse.Response.SpotPrice, float64(0))
}
