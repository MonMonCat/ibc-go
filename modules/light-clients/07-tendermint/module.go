package tendermint

import (
	"github.com/MonMonCat/ibc-go/modules/light-clients/07-tendermint/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
