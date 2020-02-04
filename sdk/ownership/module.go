package ownershipsdk

import (
	"fmt"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/mesg-foundation/engine/cosmos"
	"github.com/mesg-foundation/engine/hash"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ModuleName is the name of this module.
const ModuleName = "ownership"

// NewModule returns the module of this sdk.
func NewModule(k *Keeper) module.AppModule {
	return cosmos.NewAppModule(cosmos.NewAppModuleBasic(ModuleName), handler(k), querier(k))
}

func handler(k *Keeper) cosmos.Handler {
	return func(request cosmostypes.Request, msg cosmostypes.Msg) (hash.Hash, error) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Unrecognized ownership Msg type: %v", msg.Type())
	}
}

func querier(k *Keeper) cosmos.Querier {
	return func(request cosmostypes.Request, path []string, req abci.RequestQuery) (res interface{}, err error) {
		switch path[0] {
		case "list":
			return k.List(request)
		default:
			return nil, fmt.Errorf("unknown ownership query endpoint %s", path[0])
		}
	}
}
