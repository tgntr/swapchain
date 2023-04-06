package interchainswap

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
)

var _ porttypes.Middleware = IBCMiddleware{}

type IBCMiddleware struct {
	app porttypes.IBCModule
	im  IBCModule
}

// middleware for icacontrollerstack that routes out ICQ messages
// app is icacontroller, ibcModule is interchainswap
func NewIBCMiddleware(app porttypes.IBCModule, ibcModule IBCModule) IBCMiddleware {
	return IBCMiddleware{
		app: app,
		im:  ibcModule,
	}
}

func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	if strings.HasPrefix(portID, icatypes.PortPrefix) {
		return im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
	}

	return im.im.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
}

func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	if strings.HasPrefix(portID, icatypes.PortPrefix) {
		return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, counterpartyVersion)
	}

	return im.im.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, counterpartyVersion)
}

func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	if strings.HasPrefix(portID, icatypes.PortPrefix) {
		return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
	}

	return im.im.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

func (im IBCMiddleware) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	if strings.HasPrefix(portID, icatypes.PortPrefix) {
		return im.app.OnChanOpenConfirm(ctx, portID, channelID)
	}

	return im.im.OnChanOpenConfirm(ctx, portID, channelID)
}

func (im IBCMiddleware) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.im.OnChanCloseInit(ctx, portID, channelID)
}

func (im IBCMiddleware) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	if strings.HasPrefix(portID, icatypes.PortPrefix) {
		return im.app.OnChanCloseConfirm(ctx, portID, channelID)
	}

	return im.im.OnChanCloseConfirm(ctx, portID, channelID)
}

func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	return im.im.OnRecvPacket(ctx, modulePacket, relayer)
}

func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	if strings.HasPrefix(modulePacket.SourcePort, icatypes.PortPrefix) {
		return im.app.OnAcknowledgementPacket(ctx, modulePacket, acknowledgement, relayer)
	}

	return im.im.OnAcknowledgementPacket(ctx, modulePacket, acknowledgement, relayer)
}

func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	if strings.HasPrefix(modulePacket.SourcePort, icatypes.PortPrefix) {
		return im.app.OnTimeoutPacket(ctx, modulePacket, relayer)
	}

	return im.im.OnTimeoutPacket(ctx, modulePacket, relayer)
}

func (im IBCMiddleware) GetAppVersion(
	ctx sdk.Context,
	portID string,
	channelID string,
) (string, bool) {
	return im.im.k.GetAppVersion(ctx, portID, channelID)
}

func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
) error {
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "SendPacket not supported")
}

func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "WriteAcknowledgement not supported")
}
