package cosmos

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/node"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	tenderminttypes "github.com/tendermint/tendermint/types"
)

// Client is a tendermint client with helper functions.
type Client struct {
	rpcclient.Client
	cdc     *codec.Codec
	kb      keys.Keybase
	chainID string
}

// NewClient returns a rpc tendermint client.
func NewClient(node *node.Node, cdc *codec.Codec, kb keys.Keybase, chainID string) *Client {
	return &Client{
		Client:  rpcclient.NewLocal(node),
		cdc:     cdc,
		kb:      kb,
		chainID: chainID,
	}
}

// Query is abci.query wrapper with errors check and decode data.
func (c *Client) Query(path string, data cmn.HexBytes, ptr interface{}) error {
	result, err := c.ABCIQuery(path, data)
	if err != nil {
		return err
	}
	if !result.Response.IsOK() {
		return errors.New(result.Response.Log)
	}
	return c.cdc.UnmarshalJSON(result.Response.Value, ptr)
}

// BuildAndBroadcastMsg builds and signs message and broadcast it to node.
func (c *Client) BuildAndBroadcastMsg(msg sdktypes.Msg, accName, accPassword string) (*abci.ResponseDeliverTx, error) {
	info, err := c.kb.Get(accName)
	if err != nil {
		return nil, err
	}

	acc, err := auth.NewAccountRetriever(c).GetAccount(info.GetAddress())
	if err != nil {
		return nil, err
	}

	txBuilder := NewTxBuilder(c.cdc, acc.GetAccountNumber(), acc.GetSequence(), c.kb, c.chainID)

	// TODO: cannot sign 2 tx at the same time. Maybe keybase cannot be access at the same time. Add a lock?
	signedTx, err := txBuilder.BuildAndSignStdTx(msg, accName, accPassword)
	if err != nil {
		return nil, err
	}

	encodedTx, err := txBuilder.Encode(signedTx)
	if err != nil {
		return nil, err
	}

	txres, err := c.BroadcastTxSync(encodedTx)
	if err != nil {
		return nil, err
	}

	if txres.Code != abci.CodeTypeOK {
		return nil, fmt.Errorf("transaction returned with invalid code %d", txres.Code)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	subscriber := "engine"
	query := tenderminttypes.EventQueryTxFor(encodedTx).String()
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
		return nil, errors.New("i/o timeout")
	}
}

// QueryWithData performs a query to a Tendermint node with the provided path
// and a data payload. It returns the result and height of the query upon success
// or an error if the query fails.
func (c *Client) QueryWithData(path string, data []byte) ([]byte, int64, error) {
	return c.query(path, data)
}

// query performs a query to a Tendermint node with the provided store name
// and path. It returns the result and height of the query upon success
// or an error if the query fails. In addition, it will verify the returned
// proof if TrustNode is disabled. If proof verification fails or the query
// height is invalid, an error will be returned.
func (c *Client) query(path string, key cmn.HexBytes) (res []byte, height int64, err error) {
	result, err := c.ABCIQuery(path, key)
	if err != nil {
		return res, height, err
	}

	resp := result.Response
	if !resp.IsOK() {
		return res, resp.Height, errors.New(resp.Log)
	}

	return resp.Value, resp.Height, nil
}
