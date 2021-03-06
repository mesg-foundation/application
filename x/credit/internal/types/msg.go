package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mesg-foundation/engine/ext/xvalidator"
)

// Route should return the name of the module.
func (msg MsgAdd) Route() string {
	return ModuleName
}

// Type returns the action.
func (msg MsgAdd) Type() string {
	return "add"
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgAdd) ValidateBasic() error {
	if err := xvalidator.Struct(msg); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return nil
}

// GetSignBytes encodes the message for signing.
func (msg MsgAdd) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgAdd) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
