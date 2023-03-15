package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

// SetQueryRequest saves the query request
func (k Keeper) SetQueryRequest(ctx sdk.Context, packetSequence uint64, req types.QueryOsmosisSpotPriceRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.QueryRequestStoreKey(packetSequence), k.cdc.MustMarshal(&req))
}

// GetQueryRequest returns the query request by packet sequence
func (k Keeper) GetQueryRequest(ctx sdk.Context, packetSequence uint64) (types.QueryOsmosisSpotPriceRequest, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.QueryRequestStoreKey(packetSequence))
	if bz == nil {
		return types.QueryOsmosisSpotPriceRequest{}, sdkerrors.Wrapf(types.ErrSample,
			"GetQueryRequest: Result for packet sequence %d is not available.", packetSequence,
		)
	}
	var req types.QueryOsmosisSpotPriceRequest
	k.cdc.MustUnmarshal(bz, &req)
	return req, nil
}

// SetQueryResponse saves the query response
func (k Keeper) SetQueryResponse(ctx sdk.Context, packetSequence uint64, resp types.QueryOsmosisSpotPriceResponse) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.QueryResponseStoreKey(packetSequence), k.cdc.MustMarshal(&resp))
}

// GetQueryResponse returns the query response by packet sequence
func (k Keeper) GetQueryResponse(ctx sdk.Context, packetSequence uint64) (types.QueryOsmosisSpotPriceResponse, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.QueryResponseStoreKey(packetSequence))
	if bz == nil {
		return types.QueryOsmosisSpotPriceResponse{}, sdkerrors.Wrapf(types.ErrSample,
			"GetQueryResponse: Result for packet sequence %d is not available.", packetSequence,
		)
	}
	var resp types.QueryOsmosisSpotPriceResponse
	k.cdc.MustUnmarshal(bz, &resp)
	return resp, nil
}

// GetLastQueryPacketSeq return the id from the last query request
func (k Keeper) GetLastQueryPacketSeq(ctx sdk.Context) uint64 {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastQueryPacketSeqKey))
	uintV := gogotypes.UInt64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &uintV)
	return uintV.GetValue()
}

// SetLastQueryPacketSeq saves the id from the last query request
func (k Keeper) SetLastQueryPacketSeq(ctx sdk.Context, packetSequence uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.LastQueryPacketSeqKey),
		k.cdc.MustMarshalLengthPrefixed(&gogotypes.UInt64Value{Value: packetSequence}))
}
