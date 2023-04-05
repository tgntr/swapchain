package interchaintest

import (
	"encoding/json"
	"fmt"

	"github.com/icza/dyno"
	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
)

var (
	numVals      = 1
	numFullNodes = 0
)

var swapchainSpec = &interchaintest.ChainSpec{
	ChainName:     "swapchain",
	Version:       "latest",
	NumValidators: &numVals,
	NumFullNodes:  &numFullNodes,
	ChainConfig: ibc.ChainConfig{
		Type:           "cosmos",
		ChainID:        "swapchain-1",
		Bin:            "swapchaind",
		Denom:          "atom",
		Bech32Prefix:   "cosmos",
		CoinType:       "118",
		GasPrices:      "0.0atom",
		GasAdjustment:  1.1,
		TrustingPeriod: "500h",
		Images: []ibc.DockerImage{{
			Repository: "swapchaind",
			Version:    "latest",
		}},
	},
}

var osmosisSpec = &interchaintest.ChainSpec{
	Name:          "osmosis",
	Version:       "v15.0.0",
	NumValidators: &numVals,
	NumFullNodes:  &numFullNodes,
	ChainConfig: ibc.ChainConfig{
		ChainID:       "osmosis-1",
		GasPrices:     "0.0025uosmo",
		ModifyGenesis: modifyGenesisIcqAndIcaParams,
	},
}

// adds "/osmosis.gamm.v2.Query/SpotPrice" to ICQ allowed queries
// adds "/osmosis.gamm.v1beta1.MsgSwapExactAmountIn" to ICA allowed messages
func modifyGenesisIcqAndIcaParams(chainConfig ibc.ChainConfig, genbz []byte) ([]byte, error) {
	g := make(map[string]interface{})
	if err := json.Unmarshal(genbz, &g); err != nil {
		return nil, fmt.Errorf("failed to unmarshal genesis file: %w", err)
	}

	if err := dyno.Set(g, []string{"/osmosis.gamm.v2.Query/SpotPrice"}, "app_state", "interchainquery", "params", "allow_queries"); err != nil {
		return nil, fmt.Errorf("failed to set allowed interchain queries in genesis json: %w", err)
	}

	if err := dyno.Set(g, []string{"/osmosis.gamm.v1beta1.MsgSwapExactAmountIn"}, "app_state", "interchainaccounts", "host_genesis_state", "params", "allow_messages"); err != nil {
		return nil, fmt.Errorf("failed to set allowed interchain messages in genesis json: %w", err)
	}

	out, err := json.Marshal(g)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal genesis bytes to json: %w", err)
	}

	return out, nil
}
