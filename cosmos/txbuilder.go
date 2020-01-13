package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authutils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/mesg-foundation/engine/codec"
	"github.com/tendermint/tendermint/types"
)

// TxBuilder implements a transaction context created in SDK modules.
type TxBuilder struct {
	authtypes.TxBuilder
}

// NewTxBuilder returns a new initialized TxBuilder.
func NewTxBuilder(accSeq uint64, kb keys.Keybase, chainID string, minGasPrices sdktypes.DecCoins) TxBuilder {
	return TxBuilder{
		authtypes.NewTxBuilder(
			authutils.GetTxEncoder(codec.Codec),
			AccNumber,
			accSeq,
			flags.DefaultGasLimit,
			flags.DefaultGasAdjustment,
			true,
			chainID,
			"",
			sdktypes.NewCoins(),
			minGasPrices,
		).WithKeybase(kb),
	}
}

// BuildAndSignStdTx a signed transaction from a message.
func (b TxBuilder) BuildAndSignStdTx(msg sdktypes.Msg, accountName, accountPassword string) (authtypes.StdTx, error) {
	signedMsg, err := b.BuildSignMsg([]sdktypes.Msg{msg})
	if err != nil {
		return authtypes.StdTx{}, err
	}
	stdTx := authtypes.NewStdTx(signedMsg.Msgs, signedMsg.Fee, []authtypes.StdSignature{}, signedMsg.Memo)
	return b.SignStdTx(accountName, accountPassword, stdTx, false)
}

// Encode a transaction.
func (b TxBuilder) Encode(tx sdktypes.Tx) (types.Tx, error) {
	return b.TxEncoder()(tx)
}
