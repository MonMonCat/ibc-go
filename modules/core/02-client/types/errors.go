package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IBC client sentinel errors
var (
	ErrClientExists                           = sdkerrors.Register(SubModuleName, 217, "light client already exists")
	ErrInvalidClient                          = sdkerrors.Register(SubModuleName, 317, "light client is invalid")
	ErrClientNotFound                         = sdkerrors.Register(SubModuleName, 417, "light client not found")
	ErrClientFrozen                           = sdkerrors.Register(SubModuleName, 517, "light client is frozen due to misbehaviour")
	ErrInvalidClientMetadata                  = sdkerrors.Register(SubModuleName, 617, "invalid client metadata")
	ErrConsensusStateNotFound                 = sdkerrors.Register(SubModuleName, 717, "consensus state not found")
	ErrInvalidConsensus                       = sdkerrors.Register(SubModuleName, 817, "invalid consensus state")
	ErrClientTypeNotFound                     = sdkerrors.Register(SubModuleName, 917, "client type not found")
	ErrInvalidClientType                      = sdkerrors.Register(SubModuleName, 1017, "invalid client type")
	ErrRootNotFound                           = sdkerrors.Register(SubModuleName, 1117, "commitment root not found")
	ErrInvalidHeader                          = sdkerrors.Register(SubModuleName, 1217, "invalid client header")
	ErrInvalidMisbehaviour                    = sdkerrors.Register(SubModuleName, 1317, "invalid light client misbehaviour")
	ErrFailedClientStateVerification          = sdkerrors.Register(SubModuleName, 1417, "client state verification failed")
	ErrFailedClientConsensusStateVerification = sdkerrors.Register(SubModuleName, 1517, "client consensus state verification failed")
	ErrFailedConnectionStateVerification      = sdkerrors.Register(SubModuleName, 1617, "connection state verification failed")
	ErrFailedChannelStateVerification         = sdkerrors.Register(SubModuleName, 1717, "channel state verification failed")
	ErrFailedPacketCommitmentVerification     = sdkerrors.Register(SubModuleName, 1817, "packet commitment verification failed")
	ErrFailedPacketAckVerification            = sdkerrors.Register(SubModuleName, 1917, "packet acknowledgement verification failed")
	ErrFailedPacketReceiptVerification        = sdkerrors.Register(SubModuleName, 2017, "packet receipt verification failed")
	ErrFailedNextSeqRecvVerification          = sdkerrors.Register(SubModuleName, 2117, "next sequence receive verification failed")
	ErrSelfConsensusStateNotFound             = sdkerrors.Register(SubModuleName, 2217, "self consensus state not found")
	ErrUpdateClientFailed                     = sdkerrors.Register(SubModuleName, 2317, "unable to update light client")
	ErrInvalidUpdateClientProposal            = sdkerrors.Register(SubModuleName, 2417, "invalid update client proposal")
	ErrInvalidUpgradeClient                   = sdkerrors.Register(SubModuleName, 2517, "invalid client upgrade")
	ErrInvalidHeight                          = sdkerrors.Register(SubModuleName, 2617, "invalid height")
	ErrInvalidSubstitute                      = sdkerrors.Register(SubModuleName, 2717, "invalid client state substitute")
	ErrInvalidUpgradeProposal                 = sdkerrors.Register(SubModuleName, 2817, "invalid upgrade proposal")
	ErrClientNotActive                        = sdkerrors.Register(SubModuleName, 2917, "client is not active")
)
