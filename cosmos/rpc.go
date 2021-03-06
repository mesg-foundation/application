package cosmos

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authutils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authExported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/mesg-foundation/engine/ext/xreflect"
	abci "github.com/tendermint/tendermint/abci/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	tenderminttypes "github.com/tendermint/tendermint/types"
)

// RPC is a tendermint rpc client with helper functions.
type RPC struct {
	rpcclient.Client
	cdc         *codec.Codec
	kb          keys.Keybase
	chainID     string
	accName     string
	accPassword string
	gasPrices   sdktypes.DecCoins

	// Local state
	acc            authExported.Account
	accountMutex   sync.Mutex
	broadcastMutex sync.Mutex
}

// NewRPC returns a rpc tendermint client.
func NewRPC(client rpcclient.Client, cdc *codec.Codec, kb keys.Keybase, chainID, accName, accPassword, gasPrices string) (*RPC, error) {
	gasPricesDecoded, err := sdktypes.ParseDecCoins(gasPrices)
	if err != nil {
		return nil, err
	}
	return &RPC{
		Client:      client,
		cdc:         cdc,
		kb:          kb,
		chainID:     chainID,
		accName:     accName,
		accPassword: accPassword,
		gasPrices:   gasPricesDecoded,
	}, nil
}

// Codec returns the codec used by RPC.
func (c *RPC) Codec() *codec.Codec {
	return c.cdc
}

// QueryJSON is abci.query wrapper with errors check and decode data.
func (c *RPC) QueryJSON(path string, qdata, ptr interface{}) error {
	var data []byte
	if !xreflect.IsNil(qdata) {
		b, err := c.cdc.MarshalJSON(qdata)
		if err != nil {
			return err
		}
		data = b
	}

	result, _, err := c.QueryWithData(path, data)
	if err != nil {
		return err
	}
	return c.cdc.UnmarshalJSON(result, ptr)
}

// QueryWithData performs a query to a Tendermint node with the provided path
// and a data payload. It returns the result and height of the query upon success
// or an error if the query fails.
func (c *RPC) QueryWithData(path string, data []byte) ([]byte, int64, error) {
	result, err := c.ABCIQuery(path, data)
	if err != nil {
		return nil, 0, err
	}
	resp := result.Response
	if !resp.IsOK() {
		return nil, resp.Height, errors.New(resp.Log)
	}
	return resp.Value, resp.Height, nil
}

// BuildAndBroadcastMsg builds and signs message and broadcast it to node.
func (c *RPC) BuildAndBroadcastMsg(msg sdktypes.Msg) (*abci.ResponseDeliverTx, error) {
	signedTx, err := c.buildAndBroadcastMsgNoResult(msg)
	if err != nil {
		return nil, err
	}
	return c.waitForTxResult(signedTx)
}

func (c *RPC) buildAndBroadcastMsgNoResult(msg sdktypes.Msg) (tenderminttypes.Tx, error) {
	// Lock the getAccount + create and sign tx + broadcast
	c.broadcastMutex.Lock()
	defer c.broadcastMutex.Unlock()

	acc, err := c.GetAccount()
	if err != nil {
		return nil, err
	}

	// create and sign the tx
	signedTx, err := c.createAndSignTx([]sdktypes.Msg{msg}, acc)
	if err != nil {
		return nil, err
	}
	txres, err := c.BroadcastTxSync(signedTx)
	if err != nil {
		return nil, err
	}
	if txres.Code != abci.CodeTypeOK {
		return nil, fmt.Errorf("transaction returned with invalid code %d: %s", txres.Code, txres.Log)
	}

	// only increase sequence if no error during broadcast of tx
	if err := c.setAccountSequence(acc.GetSequence() + 1); err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (c *RPC) waitForTxResult(signedTx tenderminttypes.Tx) (*abci.ResponseDeliverTx, error) {
	// TODO: 20*time.Second should not be hardcoded here
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	subscriber := "engine"
	query := tenderminttypes.EventQueryTxFor(signedTx).String()
	out, err := c.Subscribe(ctx, subscriber, query)
	if err != nil {
		return nil, err
	}
	defer c.Unsubscribe(ctx, subscriber, query)

	select {
	case result := <-out:
		data, ok := result.Data.(tenderminttypes.EventDataTx)
		if !ok {
			return nil, errors.New("result data is not the right type")
		}
		if data.TxResult.Result.IsErr() {
			return nil, fmt.Errorf("an error occurred in transaction: %s", data.TxResult.Result.Log)
		}
		return &data.TxResult.Result, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("reach timeout for listening for transaction result: %w", ctx.Err())
	}
}

// GetAccount returns the local account.
func (c *RPC) GetAccount() (authExported.Account, error) {
	c.accountMutex.Lock()
	defer c.accountMutex.Unlock()
	if c.acc == nil {
		accKb, err := c.kb.Get(c.accName)
		if err != nil {
			return nil, err
		}
		c.acc = auth.NewBaseAccount(
			accKb.GetAddress(),
			nil,
			accKb.GetPubKey(),
			0,
			0,
		)
	}
	accR, err := auth.NewAccountRetriever(c).GetAccount(c.acc.GetAddress())
	if err != nil {
		return nil, err
	}
	// replace seq if sup
	if c.acc.GetSequence() > accR.GetSequence() {
		accR.SetSequence(c.acc.GetSequence())
	}
	c.acc = accR
	return c.acc, nil
}

// setAccountSequence sets the sequence on the local account.
func (c *RPC) setAccountSequence(seq uint64) error {
	c.accountMutex.Lock()
	defer c.accountMutex.Unlock()
	if c.acc == nil {
		return fmt.Errorf("c.acc should not be nil. use GetAccount first")
	}
	return c.acc.SetSequence(seq)
}

// createAndSignTx build and sign a msg with client account.
func (c *RPC) createAndSignTx(msgs []sdktypes.Msg, acc authExported.Account) (tenderminttypes.Tx, error) {
	// Create TxBuilder
	txBuilder := authtypes.NewTxBuilder(
		authutils.GetTxEncoder(c.cdc),
		acc.GetAccountNumber(),
		acc.GetSequence(),
		flags.DefaultGasLimit,
		GasAdjustment,
		true,
		c.chainID,
		"",
		nil,
		c.gasPrices,
	).WithKeybase(c.kb)

	// calculate gas
	if txBuilder.SimulateAndExecute() {
		txBytes, err := txBuilder.BuildTxForSim(msgs)
		if err != nil {
			return nil, err
		}
		_, adjusted, err := authutils.CalculateGas(c.QueryWithData, c.cdc, txBytes, txBuilder.GasAdjustment())
		if err != nil {
			return nil, err
		}
		txBuilder = txBuilder.WithGas(adjusted)
	}

	// create StdSignMsg
	stdSignMsg, err := txBuilder.BuildSignMsg(msgs)
	if err != nil {
		return nil, err
	}

	// create StdTx
	stdTx := authtypes.NewStdTx(stdSignMsg.Msgs, stdSignMsg.Fee, nil, stdSignMsg.Memo)

	// sign StdTx
	signedTx, err := txBuilder.SignStdTx(c.accName, c.accPassword, stdTx, false)
	if err != nil {
		return nil, err
	}

	return txBuilder.TxEncoder()(signedTx)
}
