package simulation

import (
	"math/rand"

	"github.com/MonMonCat/ibc-go/modules/core/04-channel/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// GenChannelGenesis returns the default channel genesis state.
func GenChannelGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
