package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Minter struct {
	Address   sdk.AccAddress
	Allowance sdk.Coin
}
