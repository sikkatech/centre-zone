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
	// // The reference to the Paramstore to get and set gov specific params
	// paramSpace types.ParamSubspace

	// // The SupplyKeeper to reduce the supply of the network
	// supplyKeeper types.SupplyKeeper

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
	cdc *codec.Codec, key sdk.StoreKey, sendKeeper sdkbank.SendKeeper,
) Keeper {
	return Keeper{
		storeKey: key,
		// paramSpace:   paramSpace,
		// supplyKeeper: supplyKeeper,
		sendKeeper: sendKeeper,
		cdc:        cdc,
	}
}

// Logger returns a module-specific logger.
func (keeper Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (keeper Keeper) GetAuthorities(ctx sdk.Context, denom string, role types.AuthorityRole) (authorities []sdk.AccAddress) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.GetRolesKey(denom, role))
	if bz == nil {
		return authorities
	}
	keeper.cdc.MustUnmarshalBinaryBare(bz, &authorities)
	return authorities
}

func (keeper Keeper) SetAuthorities(ctx sdk.Context, denom string, role types.AuthorityRole, authorities []sdk.AccAddress) {
	store := ctx.KVStore(keeper.storeKey)
	if len(authorities) == 0 {
		store.Delete(types.GetRolesKey(denom, role))
		return
	}
	bz := keeper.cdc.MustMarshalBinaryBare(authorities)
	store.Set(types.GetRolesKey(denom, role), bz)
}

func (keeper Keeper) IsAuthority(ctx sdk.Context, denom string, role types.AuthorityRole, addr sdk.AccAddress) bool {
	authorities := keeper.GetAuthorities(ctx, denom, role)
	for _, existing := range authorities {
		if existing.Equals(addr) {
			return true
		}
	}
	return false
}

func (keeper Keeper) AddAuthority(ctx sdk.Context, denom string, role types.AuthorityRole, addr sdk.AccAddress) error {
	if keeper.IsAuthority(ctx, denom, role, addr) {
		return sdkerrors.Wrapf(types.ErrAuthorityAlreadyExists, "%s - %s", role, addr)
	}
	authorities := keeper.GetAuthorities(ctx, denom, role)
	authorities = append(authorities, addr)
	keeper.SetAuthorities(ctx, denom, role, authorities)
	return nil
}

func (keeper Keeper) RemoveAuthority(ctx sdk.Context, denom string, role types.AuthorityRole, addr sdk.AccAddress) error {
	authorities := keeper.GetAuthorities(ctx, denom, role)
	index := -1
	for i, existing := range authorities {
		if existing.Equals(addr) {
			index = i
			break
		}
	}
	if index == -1 {
		return sdkerrors.Wrapf(types.ErrAuthorityInvalid, "%s - %s", role, addr)
	}

	authorities[index] = authorities[len(authorities)-1]
	authorities = authorities[:len(authorities)-1]
	keeper.SetAuthorities(ctx, denom, role, authorities)

	if role == types.Minter {
		keeper.UpdateMinterAllowance(ctx, addr, sdk.NewInt64Coin(denom, 0))
	}

	return nil
}

func (keeper Keeper) GetMinterAllowance(ctx sdk.Context, denom string, addr sdk.AccAddress) (allowance sdk.Coin) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.GetAllowancesKey(denom, addr))
	if bz == nil {
		return sdk.NewInt64Coin(denom, 0)
	}
	keeper.cdc.MustUnmarshalBinaryBare(bz, &allowance)
	return allowance
}

func (keeper Keeper) UpdateMinterAllowance(ctx sdk.Context, addr sdk.AccAddress, allowance sdk.Coin) {
	store := ctx.KVStore(keeper.storeKey)
	if !allowance.IsValid() || allowance.IsZero() {
		store.Delete(types.GetAllowancesKey(allowance.Denom, addr))
		return
	}
	bz := keeper.cdc.MustMarshalBinaryBare(allowance)
	store.Set(types.GetAllowancesKey(allowance.Denom, addr), bz)
}
