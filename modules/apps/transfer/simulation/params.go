package simulation

import (
	"fmt"
	"math/rand"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/MonMonCat/ibc-go/modules/apps/transfer/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeySendEnabled),
			func(r *rand.Rand) string {
				sendEnabled := RadomEnabled(r)
				return fmt.Sprintf("%s", types.ModuleCdc.MustMarshalJSON(&gogotypes.BoolValue{Value: sendEnabled}))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyReceiveEnabled),
			func(r *rand.Rand) string {
				receiveEnabled := RadomEnabled(r)
				return fmt.Sprintf("%s", types.ModuleCdc.MustMarshalJSON(&gogotypes.BoolValue{Value: receiveEnabled}))
			},
		),
	}
}
