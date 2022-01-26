package simulation

import (
	"math/rand"

	"github.com/MonMonCat/ibc-go/modules/core/03-connection/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// GenConnectionGenesis returns the default connection genesis state.
func GenConnectionGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
