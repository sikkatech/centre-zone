package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkbank "github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/sikkatech/centre-zone/x/bank/types"
)

// Keeper defines the governance module Keeper
type Keeper struct {
	// The reference to the Paramstore to get and set gov specific params
	paramSpace types.ParamSubspace

	// The SupplyKeeper to reduce the supply of the network
	supplyKeeper types.SupplyKeeper

	// The baseBankKeeper that actually handles balances
	sendKeeper sdkbank.SendKeeper

	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey

	// The codec codec for binary encoding/decoding.
	cdc *codec.Codec
}

// NewKeeper returns a new centre-bank keeper.
// CONTRACT: the parameter Subspace must have the param key table already initialized
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, paramSpace types.ParamSubspace,
	supplyKeeper types.SupplyKeeper, sendKeeper sdkbank.SendKeeper, rtr types.Router,
) Keeper {
	return Keeper{
		storeKey:     key,
		paramSpace:   paramSpace,
		supplyKeeper: supplyKeeper,
		sendKeeper:   sendKeeper,
		cdc:          cdc,
	}
}

// Logger returns a module-specific logger.
func (keeper Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// MINTERS

func (keeper Keeper) GetMinters(ctx sdk.Context, denom string) (minters []types.Minter) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.GetMintersKey(denom))
	if bz == nil {
		return minters
	}
	keeper.cdc.MustUnmarshalBinaryBare(bz, &minters)
	return minters
}

// CONTRACT:  minters allowance denoms must be equal to denom paramter
func (keeper Keeper) SetMinters(ctx sdk.Context, denom string, minters []types.Minter) {
	store := ctx.KVStore(keeper.storeKey)
	if len(minters) == 0 {
		store.Delete(types.GetMintersKey(denom))
		return
	}
	bz := keeper.cdc.MustMarshalBinaryBare(minters)
	store.Set(types.GetMintersKey(denom), bz)
}

func (keeper Keeper) UpdateMinter(ctx sdk.Context, minter sdk.AccAddress, allowance sdk.Coin) {
	minters := keeper.GetMinters(ctx, allowance.Denom)

	for i, existing := range minters {
		if existing.Address.Equals(minter) {
			minters[i].Allowance = allowance
			keeper.SetMinters(ctx, allowance.Denom, minters)
			return
		}
	}

	minters = append(minters, types.Minter{
		Address:   minter,
		Allowance: allowance,
	})
	keeper.SetMinters(ctx, allowance.Denom, minters)
}

func (keeper Keeper) RemoveMinter(ctx sdk.Context, denom string, minter sdk.AccAddress) error {
	minters := keeper.GetMinters(ctx, denom)
	index := -1
	for i, existing := range minters {
		if existing.Address.Equals(minter) {
			index = i
			break
		}
	}
	if index == -1 {
		return sdkerrors.Wrap(types.ErrInvalidMinter, minter.String())
	}

	minters[index] = minters[len(minters)-1]
	minters = minters[:len(minters)-1]

	keeper.SetMinters(ctx, denom, minters)
}
