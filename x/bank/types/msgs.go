package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = MsgMint{}
	_ sdk.Msg = MsgBurn{}
	_ sdk.Msg = MsgUpdateMinter{}
	_ sdk.Msg = MsgRemoveMinter{}
)

type MsgMint struct {
	Minter sdk.AccAddress
	To     sdk.AccAddress
	Amount sdk.Coin
}

// NewMsgMint creates a new MsgMint instance
func NewMsgMint(minter, to sdk.AccAddress, amount sdk.Coin) MsgMint {
	return MsgMint{
		Minter: minter,
		To:     to,
		Amount: amount,
	}
}

// Route implements Msg
func (msg MsgMint) Route() string { return ModuleName }

// Type implements Msg
func (msg MsgMint) Type() string { return ModuleName }

// ValidateBasic implements Msg
func (msg MsgMint) ValidateBasic() error {
	if msg.Minter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "minter address is empty")
	}
	if msg.To.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "to address is empty")
	}
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%s is an invalid amount", msg.Amount.String())
	}
	return nil
}

// GetSignBytes implements Msg
func (msg MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Minter}
}

type MsgBurn struct {
	Burner sdk.AccAddress
	Amount sdk.Coin
}

// NewMsgBurn creates a new MsgMint instance
func NewMsgBurn(burner sdk.AccAddress, amount sdk.Coin) MsgBurn {
	return MsgBurn{
		Burner: burner,
		Amount: amount,
	}
}

// Route implements Msg
func (msg MsgBurn) Route() string { return ModuleName }

// Type implements Msg
func (msg MsgBurn) Type() string { return ModuleName }

// ValidateBasic implements Msg
func (msg MsgBurn) ValidateBasic() error {
	if msg.Burner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "burner address is empty")
	}
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%s is an invalid amount", msg.Amount.String())
	}
	return nil
}

// GetSignBytes implements Msg
func (msg MsgBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Burner}
}

// MsgUpdateMinter is used to add a minter or update their allowance
type MsgUpdateMinter struct {
	Authority sdk.AccAddress
	Minter    sdk.AccAddress
	Allowance sdk.Coin
}

// NewMsgUpdateMinter creates a new MsgMsgUpdateMinter instance
func NewMsgUpdateMinter(authority, minter sdk.AccAddress, allowance sdk.Coin) MsgUpdateMinter {
	return MsgUpdateMinter{
		Authority: authority,
		Minter:    minter,
		Allowance: allowance,
	}
}

// Route implements Msg
func (msg MsgUpdateMinter) Route() string { return ModuleName }

// Type implements Msg
func (msg MsgUpdateMinter) Type() string { return ModuleName }

// ValidateBasic implements Msg
func (msg MsgUpdateMinter) ValidateBasic() error {
	if msg.Authority.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "authority address is empty")
	}
	if msg.Minter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "minter address is empty")
	}
	if !msg.Allowance.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%s is an invalid allowance", msg.Allowance.String())
	}
	return nil
}

// GetSignBytes implements Msg
func (msg MsgUpdateMinter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgUpdateMinter) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Authority}
}

type MsgRemoveMinter struct {
	Authority sdk.AccAddress
	Minter    sdk.AccAddress
}

// NewMsgRemoveMinter creates a new MsgRemoveMinter instance
func NewMsgRemoveMinter(authority, minter sdk.AccAddress) MsgRemoveMinter {
	return MsgRemoveMinter{
		Authority: authority,
		Minter:    minter,
	}
}

// Route implements Msg
func (msg MsgRemoveMinter) Route() string { return ModuleName }

// Type implements Msg
func (msg MsgRemoveMinter) Type() string { return ModuleName }

// ValidateBasic implements Msg
func (msg MsgRemoveMinter) ValidateBasic() error {
	if msg.Authority.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "authority address is empty")
	}
	if msg.Minter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "minter address is empty")
	}
	return nil
}

// GetSignBytes implements Msg
func (msg MsgRemoveMinter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgRemoveMinter) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Authority}
}
