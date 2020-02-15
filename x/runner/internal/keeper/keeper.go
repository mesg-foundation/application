package keeper

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/runner"
	"github.com/mesg-foundation/engine/x/runner/internal/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the runner store
type Keeper struct {
	storeKey       sdk.StoreKey
	cdc            *codec.Codec
	instanceKeeper types.InstanceKeeper
}

// NewKeeper creates a runner keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, instanceKeeper types.InstanceKeeper) Keeper {
	keeper := Keeper{
		storeKey:       key,
		cdc:            cdc,
		instanceKeeper: instanceKeeper,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Create creates a new runner.
func (k Keeper) Create(ctx sdk.Context, msg *types.MsgCreateRunner) (*runner.Runner, error) {
	store := ctx.KVStore(k.storeKey)
	inst, err := k.instanceKeeper.FetchOrCreate(ctx, msg.ServiceHash, msg.EnvHash)
	if err != nil {
		return nil, err
	}

	r := &runner.Runner{
		Address:      msg.Address.String(),
		InstanceHash: inst.Hash,
	}
	r.Hash = hash.Dump(r)
	if store.Has(r.Hash) {
		return nil, fmt.Errorf("runner %q already exists", r.Hash)
	}

	value, err := k.cdc.MarshalBinaryLengthPrefixed(r)
	if err != nil {
		return nil, err
	}
	store.Set(r.Hash, value)
	return r, nil
}

// Delete deletes a runner.
func (k Keeper) Delete(ctx sdk.Context, msg *types.MsgDeleteRunner) error {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(msg.RunnerHash) {
		return fmt.Errorf("runner %q not found", msg.RunnerHash)
	}

	value := store.Get(msg.RunnerHash)
	var r *runner.Runner
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &r); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}
	if r.Address != msg.Address.String() {
		return errors.New("only the runner owner can remove itself")
	}
	store.Delete(msg.RunnerHash)
	return nil
}

// Get returns the runner that matches given hash.
func (k Keeper) Get(ctx sdk.Context, hash hash.Hash) (*runner.Runner, error) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(hash) {
		return nil, fmt.Errorf("runner %q not found", hash)
	}
	value := store.Get(hash)
	var r *runner.Runner
	return r, k.cdc.UnmarshalBinaryLengthPrefixed(value, &r)
}

// List returns all runners.
func (k Keeper) List(ctx sdk.Context) ([]*runner.Runner, error) {
	var (
		runners []*runner.Runner
		iter    = ctx.KVStore(k.storeKey).Iterator(nil, nil)
	)
	for iter.Valid() {
		var r *runner.Runner
		if err := k.cdc.UnmarshalBinaryLengthPrefixed(iter.Value(), &r); err != nil {
			return nil, err
		}
		runners = append(runners, r)
		iter.Next()
	}
	iter.Close()
	return runners, nil
}