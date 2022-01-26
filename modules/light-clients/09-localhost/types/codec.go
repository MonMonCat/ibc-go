package types

import (
	"github.com/MonMonCat/ibc-go/modules/core/exported"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}
