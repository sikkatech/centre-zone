package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the contract module
	ModuleName = "bank"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// TStoreKey is the string transient store representation
	TStoreKey = "transient_" + ModuleName

	// QuerierRoute is the querier route for the staking module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the staking module
	RouterKey = ModuleName
)

func GetRolesKey(denom string, role AuthorityRole) []byte {
	return []byte(fmt.Sprintf("roles/%s/%s", denom, role))
}

func GetAllowancesKey(denom string, addr sdk.AccAddress) []byte {
	return append([]byte(fmt.Sprintf("allowances/%s/%s", denom)), addr.Bytes()...)
}
