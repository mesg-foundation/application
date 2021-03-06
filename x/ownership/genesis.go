package ownership

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mesg-foundation/engine/x/ownership/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters and the keeper's address to pubkey map.
func InitGenesis(ctx sdk.Context, k Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	if err := k.Import(ctx, data.Ownerships); err != nil {
		panic(err)
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis writes the current store values // to a genesis file,
// which can be imported again with InitGenesis.
func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	ownerships, err := k.List(ctx)
	if err != nil {
		panic(err)
	}
	return types.NewGenesisState(ownerships)
}
