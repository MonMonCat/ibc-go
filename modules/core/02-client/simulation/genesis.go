package simulation

import (
	"math/rand"

	"github.com/MonMonCat/ibc-go/modules/core/02-client/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
