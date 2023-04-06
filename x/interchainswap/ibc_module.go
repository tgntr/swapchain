package interchainswap

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"

	"github.com/tgntr/swapchain/x/interchainswap/keeper"
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

var _ porttypes.IBCModule = IBCModule{}

type IBCModule struct {
	k keeper.Keeper
}

func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{k: k}
}

func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	isICA := strings.HasPrefix(portID, icatypes.PortPrefix)
	isValidICQ := portID == im.k.GetPort(ctx) && version == types.Version
	if !isICA && !isValidICQ {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "port %s, version %s", portID, version)
	}

	return version, im.k.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID))
}

func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	boundPort := im.k.GetPort(ctx)
	if portID != boundPort {
		return "", sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if counterpartyVersion != types.Version {
		return "", sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: got: %s, expected %s", counterpartyVersion, types.Version)
	}

	return types.Version, im.k.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID))
}

func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	if counterpartyVersion != types.Version && !strings.HasPrefix(portID, icatypes.PortPrefix) {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: %s, expected %s", counterpartyVersion, types.Version)
	}

	return nil
}

func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "users cannot close channel")
}

func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	return channeltypes.NewErrorAcknowledgement(sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "RecvPacket not supported"))
}

func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	if strings.HasPrefix(modulePacket.SourcePort, icatypes.PortPrefix) {
		return nil
	}

	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet acknowledgement: %v", err)
	}

	return im.k.OnICQAcknowledgementPacket(ctx, modulePacket, ack)
}

func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	im.k.Logger(ctx).Error(types.ErrInvalidPacketTimeout.Error(), types.AttributeKeySequence, modulePacket.Sequence)

	return nil
}
