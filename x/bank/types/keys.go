package types

import (
	"fmt"
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

func GetMintersKey(denom string) []byte {
	return []byte(fmt.Sprintf("%s/minters", denom))
}

func GetMasterMintersKey(denom string) []byte {
	return []byte(fmt.Sprintf("%s/masterminters", denom))
}

func GetBlacklistersKey(denom string) []byte {
	return []byte(fmt.Sprintf("%s/blacklisters", denom))
}

func GetPausersKey(denom string) []byte {
	return []byte(fmt.Sprintf("%s/pausers", denom))
}

func GetAdminsKey(denom string) []byte {
	return []byte(fmt.Sprintf("%s/admins", denom))
}
