package types

import (
	clienttypes "github.com/MonMonCat/ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/MonMonCat/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/MonMonCat/ibc-go/modules/core/04-channel/types"
	commitmenttypes "github.com/MonMonCat/ibc-go/modules/core/23-commitment/types"
	solomachinetypes "github.com/MonMonCat/ibc-go/modules/light-clients/06-solomachine/types"
	ibctmtypes "github.com/MonMonCat/ibc-go/modules/light-clients/07-tendermint/types"
	localhosttypes "github.com/MonMonCat/ibc-go/modules/light-clients/09-localhost/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterInterfaces registers x/ibc interfaces into protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	clienttypes.RegisterInterfaces(registry)
	connectiontypes.RegisterInterfaces(registry)
	channeltypes.RegisterInterfaces(registry)
	solomachinetypes.RegisterInterfaces(registry)
	ibctmtypes.RegisterInterfaces(registry)
	localhosttypes.RegisterInterfaces(registry)
	commitmenttypes.RegisterInterfaces(registry)
}
