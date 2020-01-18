package types

// AuthorityRole defines an role for different authority types
type AuthorityRole string

const (
	Minter       = AuthorityRole("minter")
	MasterMinter = AuthorityRole("masterminter")
	Pauser       = AuthorityRole("pauser")
	Blacklister  = AuthorityRole("blacklister")
	Admin        = AuthorityRole("admin")
)
