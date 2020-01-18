package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/wasm module errors
var (
	// ErrAuthorityNotFound is for when an authority could not be found
	ErrAuthorityInvalid       = sdkerrors.Register(ModuleName, 1, "authority invalid")
	ErrAuthorityAlreadyExists = sdkerrors.Register(ModuleName, 2, "authority already exists")
)
