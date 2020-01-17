package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Bech32MainPrefix is the primary bech32 prefix that distinguishes usdc chain addresses
const Bech32MainPrefix = "usdc"

// SetBech32AddressPrefixes sets the bech32 prefixes for accounts, validators, and consensus addresses
// Uses default secondary prefixes from Cosmos SDK
func SetBech32AddressPrefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+sdk.PrefixPublic)
	config.SetBech32PrefixForValidator(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic)
}
