package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mesg-foundation/engine/x/credit/internal/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the credit store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	ak         types.AccountKeeper
	paramstore types.ParamSubspace
}

// NewKeeper creates a credit keeper
func NewKeeper(cdc *codec.Codec, ak types.AccountKeeper, key sdk.StoreKey, paramstore types.ParamSubspace) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		ak:         ak,
		paramstore: paramstore.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Add a number of credits to an address
func (k Keeper) Add(ctx sdk.Context, address sdk.AccAddress, amount sdk.Int) (sdk.Int, error) {
	value, err := k.Get(ctx, address)
	if err != nil {
		return sdk.Int{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	res := value.Add(amount)
	// emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventType,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeActionAdded),
			sdk.NewAttribute(types.AttributeKeyAddress, address.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	)
	return k.set(ctx, address, res)
}

// Sub a number of credits to an address
func (k Keeper) Sub(ctx sdk.Context, address sdk.AccAddress, amount sdk.Int) (sdk.Int, error) {
	value, err := k.Get(ctx, address)
	if err != nil {
		return sdk.Int{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	res := value.Sub(amount)
	// emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventType,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeActionSubtracted),
			sdk.NewAttribute(types.AttributeKeyAddress, address.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	)
	return k.set(ctx, address, res)
}

// Set a number of credit to a specific address
func (k Keeper) set(ctx sdk.Context, address sdk.AccAddress, balance sdk.Int) (sdk.Int, error) {
	if k.ak.GetAccount(ctx, address) == nil {
		k.ak.SetAccount(ctx, k.ak.NewAccountWithAddress(ctx, address))
	}
	store := ctx.KVStore(k.storeKey)
	encoded, err := k.cdc.MarshalBinaryLengthPrefixed(balance)
	if err != nil {
		return sdk.Int{}, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, err.Error())
	}
	store.Set(address.Bytes(), encoded)
	return balance, nil
}

// Get the balance of a specific address
func (k Keeper) Get(ctx sdk.Context, address sdk.AccAddress) (sdk.Int, error) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(address.Bytes()) {
		return sdk.ZeroInt(), nil
	}
	value := store.Get(address.Bytes())
	var balance sdk.Int
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &balance); err != nil {
		return sdk.Int{}, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return balance, nil
}

// Export all credits from the keeper.
func (k *Keeper) Export(ctx sdk.Context) (map[string]sdk.Int, error) {
	var (
		credits = make(map[string]sdk.Int)
		iter    = ctx.KVStore(k.storeKey).Iterator(nil, nil)
	)
	for iter.Valid() {
		var cred sdk.Int
		value := iter.Value()
		if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &cred); err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, err.Error())
		}
		address := sdk.AccAddress(iter.Key())
		credits[address.String()] = cred
		iter.Next()
	}
	iter.Close()
	return credits, nil
}

// Import all credits into the keeper.
func (k *Keeper) Import(ctx sdk.Context, credits map[string]sdk.Int) error {
	for address, balance := range credits {
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return err
		}
		k.set(ctx, addr, balance)
	}
	return nil
}
