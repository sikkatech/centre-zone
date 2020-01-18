package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all staking state that must be provided at genesis
type GenesisState struct {
	AssetAuthorities []AssetAuthoritySet `json:"asset_authorities" yaml:"asset_authorities"`
}

type AssetAuthoritySet struct {
	Denom         string
	Admins        []sdk.AccAddress
	Pausers       []sdk.AccAddress
	BlackListers  []sdk.AccAddress
	MasterMinters []sdk.AccAddress
	Minters       []Minter
}

// NewGenesisState creates a new genesis state for the governance module
func NewGenesisState(assetAuthorities []AssetAuthoritySet) GenesisState {
	return GenesisState{
		AssetAuthorities: assetAuthorities,
	}
}

// DefaultGenesisState defines the default governance genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState([]AssetAuthoritySet{})
}

// Equal checks whether two gov GenesisState structs are equivalent
func (data GenesisState) Equal(data2 GenesisState) bool {
	b1 := ModuleCdc.MustMarshalBinaryBare(data)
	b2 := ModuleCdc.MustMarshalBinaryBare(data2)
	return bytes.Equal(b1, b2)
}

// IsEmpty returns true if a GenesisState is empty
func (data GenesisState) IsEmpty() bool {
	return data.Equal(GenesisState{})
}

// ValidateGenesis checks if parameters are within valid ranges
func ValidateGenesis(data GenesisState) error {
	return nil
}
