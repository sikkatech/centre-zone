package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/wasm module errors
var (
	// ErrCreateFailed error for wasm code that has already been uploaded or failed
	ErrInvalidMinter = sdkerrors.Register(ModuleName, 1, "minter not found")
)
