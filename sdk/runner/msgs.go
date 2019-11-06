package runnersdk

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/mesg-foundation/engine/hash"
)

// msgCreateRunner defines a state transition to create a runner.
type msgCreateRunner struct {
	Address     cosmostypes.AccAddress `json:"address"`
	ServiceHash hash.Hash              `json:"serviceHash"`
	EnvHash     hash.Hash              `json:"envHash"`
	cdc         *codec.Codec
}

// newMsgCreateRunner is a constructor function for msgCreateRunner.
func newMsgCreateRunner(cdc *codec.Codec, address cosmostypes.AccAddress, serviceHash hash.Hash, envHash hash.Hash) *msgCreateRunner {
	return &msgCreateRunner{
		Address:     address,
		ServiceHash: serviceHash,
		EnvHash:     envHash,
		cdc:         cdc,
	}
}

// Route should return the name of the module.
func (msg msgCreateRunner) Route() string {
	return backendName
}

// Type returns the action.
func (msg msgCreateRunner) Type() string {
	return "create_runner"
}

// ValidateBasic runs stateless checks on the message.
func (msg msgCreateRunner) ValidateBasic() cosmostypes.Error {
	if msg.ServiceHash.IsZero() {
		return cosmostypes.ErrInternal("serviceHash is missing")
	}
	if msg.EnvHash.IsZero() {
		return cosmostypes.ErrInternal("envHash is missing")
	}
	if msg.Address.Empty() {
		return cosmostypes.ErrInvalidAddress("address is missing")
	}
	return nil
}

// GetSignBytes encodes the message for signing.
func (msg msgCreateRunner) GetSignBytes() []byte {
	return cosmostypes.MustSortJSON(msg.cdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg msgCreateRunner) GetSigners() []cosmostypes.AccAddress {
	return []cosmostypes.AccAddress{msg.Address}
}

// msgDeleteRunner defines a state transition to delete a runner.
type msgDeleteRunner struct {
	Address    cosmostypes.AccAddress `json:"address"`
	RunnerHash hash.Hash              `json:"runnerHash"`
	cdc        *codec.Codec
}

// newMsgDeleteRunner is a constructor function for msgDeleteRunner.
func newMsgDeleteRunner(cdc *codec.Codec, address cosmostypes.AccAddress, runnerHash hash.Hash) *msgDeleteRunner {
	return &msgDeleteRunner{
		Address:    address,
		RunnerHash: runnerHash,
		cdc:        cdc,
	}
}

// Route should return the name of the module.
func (msg msgDeleteRunner) Route() string {
	return backendName
}

// Type returns the action.
func (msg msgDeleteRunner) Type() string {
	return "delete_runner"
}

// ValidateBasic runs stateless checks on the message.
func (msg msgDeleteRunner) ValidateBasic() cosmostypes.Error {
	if msg.RunnerHash.IsZero() {
		return cosmostypes.ErrInternal("runnerHash is missing")
	}
	if msg.Address.Empty() {
		return cosmostypes.ErrInvalidAddress("address is missing")
	}
	return nil
}

// GetSignBytes encodes the message for signing.
func (msg msgDeleteRunner) GetSignBytes() []byte {
	return cosmostypes.MustSortJSON(msg.cdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg msgDeleteRunner) GetSigners() []cosmostypes.AccAddress {
	return []cosmostypes.AccAddress{msg.Address}
}