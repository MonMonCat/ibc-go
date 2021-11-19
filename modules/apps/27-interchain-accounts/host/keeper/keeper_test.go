package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/ibc-go/v2/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v2/testing"
)

var (
	// TestAccAddress defines a resuable bech32 address for testing purposes
	// TODO: update crypto.AddressHash() when sdk uses address.Module()
	TestAccAddress = types.GenerateAddress(sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName))), TestPortID)
	// TestOwnerAddress defines a reusable bech32 address for testing purposes
	TestOwnerAddress = "cosmos17dtl0mjt3t77kpuhg2edqzjpszulwhgzuj9ljs"
	// TestPortID defines a resuable port identifier for testing purposes
	TestPortID, _ = types.GeneratePortID(TestOwnerAddress, ibctesting.FirstConnectionID, ibctesting.FirstConnectionID)
	// TestVersion defines a resuable interchainaccounts version string for testing purposes
	TestVersion = types.NewAppVersion(types.VersionPrefix, TestAccAddress.String())
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
	chainC *ibctesting.TestChain
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainC = suite.coordinator.GetChain(ibctesting.GetChainID(2))
}

func NewICAPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = types.PortID
	path.EndpointB.ChannelConfig.PortID = types.PortID
	path.EndpointA.ChannelConfig.Order = channeltypes.ORDERED
	path.EndpointB.ChannelConfig.Order = channeltypes.ORDERED
	path.EndpointA.ChannelConfig.Version = types.VersionPrefix
	path.EndpointB.ChannelConfig.Version = TestVersion

	return path
}

// SetupICAPath invokes the InterchainAccounts entrypoint and subsequent channel handshake handlers
func SetupICAPath(path *ibctesting.Path, owner string) error {
	if err := InitInterchainAccount(path.EndpointA, owner); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenTry(); err != nil {
		return err
	}

	if err := path.EndpointA.ChanOpenAck(); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenConfirm(); err != nil {
		return err
	}

	return nil
}

// InitInterchainAccount is a helper function for starting the channel handshake
func InitInterchainAccount(endpoint *ibctesting.Endpoint, owner string) error {
	portID, err := types.GeneratePortID(owner, endpoint.ConnectionID, endpoint.Counterparty.ConnectionID)
	if err != nil {
		return err
	}

	channelSequence := endpoint.Chain.App.GetIBCKeeper().ChannelKeeper.GetNextChannelSequence(endpoint.Chain.GetContext())

	if err := endpoint.Chain.GetSimApp().ICAControllerKeeper.InitInterchainAccount(endpoint.Chain.GetContext(), endpoint.ConnectionID, endpoint.Counterparty.ConnectionID, owner); err != nil {
		return err
	}

	// commit state changes for proof verification
	endpoint.Chain.App.Commit()
	endpoint.Chain.NextBlock()

	// update port/channel ids
	endpoint.ChannelID = channeltypes.FormatChannelIdentifier(channelSequence)
	endpoint.ChannelConfig.PortID = portID

	return nil
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestIsBound() {
	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	isBound := suite.chainB.GetSimApp().ICAHostKeeper.IsBound(suite.chainB.GetContext(), path.EndpointB.ChannelConfig.PortID)
	suite.Require().True(isBound)
}

func (suite *KeeperTestSuite) TestGetAllPorts() {
	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	expectedPorts := []string{types.PortID}

	ports := suite.chainB.GetSimApp().ICAHostKeeper.GetAllPorts(suite.chainB.GetContext())
	suite.Require().Len(ports, len(expectedPorts))
	suite.Require().Equal(expectedPorts, ports)
}

func (suite *KeeperTestSuite) TestGetInterchainAccountAddress() {
	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	counterpartyPortID := path.EndpointA.ChannelConfig.PortID
	expectedAddr := authtypes.NewBaseAccountWithAddress(types.GenerateAddress(suite.chainA.GetSimApp().AccountKeeper.GetModuleAddress(types.ModuleName), counterpartyPortID)).GetAddress()

	retrievedAddr, found := suite.chainB.GetSimApp().ICAHostKeeper.GetInterchainAccountAddress(suite.chainB.GetContext(), counterpartyPortID)
	suite.Require().True(found)
	suite.Require().Equal(expectedAddr.String(), retrievedAddr)

	retrievedAddr, found = suite.chainB.GetSimApp().ICAHostKeeper.GetInterchainAccountAddress(suite.chainB.GetContext(), "invalid port")
	suite.Require().False(found)
	suite.Require().Empty(retrievedAddr)
}

func (suite *KeeperTestSuite) TestGetAllActiveChannels() {
	var (
		expectedChannelID string = "test-channel"
		expectedPortID    string = "test-port"
	)

	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	suite.chainB.GetSimApp().ICAHostKeeper.SetActiveChannelID(suite.chainB.GetContext(), expectedPortID, expectedChannelID)

	expectedChannels := []*types.ActiveChannel{
		{
			PortId:    path.EndpointB.ChannelConfig.PortID,
			ChannelId: path.EndpointB.ChannelID,
		},
		{
			PortId:    expectedPortID,
			ChannelId: expectedChannelID,
		},
	}

	activeChannels := suite.chainB.GetSimApp().ICAHostKeeper.GetAllActiveChannels(suite.chainB.GetContext())
	suite.Require().Len(activeChannels, len(expectedChannels))
	suite.Require().Equal(expectedChannels, activeChannels)
}

func (suite *KeeperTestSuite) TestGetAllInterchainAccounts() {
	var (
		expectedAccAddr string = "test-acc-addr"
		expectedPortID  string = "test-port"
	)

	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	suite.chainB.GetSimApp().ICAHostKeeper.SetInterchainAccountAddress(suite.chainB.GetContext(), expectedPortID, expectedAccAddr)

	expectedAccounts := []*types.RegisteredInterchainAccount{
		{
			PortId:         TestPortID,
			AccountAddress: TestAccAddress.String(),
		},
		{
			PortId:         expectedPortID,
			AccountAddress: expectedAccAddr,
		},
	}

	interchainAccounts := suite.chainB.GetSimApp().ICAHostKeeper.GetAllInterchainAccounts(suite.chainB.GetContext())
	suite.Require().Len(interchainAccounts, len(expectedAccounts))
	suite.Require().Equal(expectedAccounts, interchainAccounts)
}

func (suite *KeeperTestSuite) TestIsActiveChannel() {
	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	isActive := suite.chainB.GetSimApp().ICAHostKeeper.IsActiveChannel(suite.chainB.GetContext(), path.EndpointB.ChannelConfig.PortID)
	suite.Require().True(isActive)
}

func (suite *KeeperTestSuite) TestSetInterchainAccountAddress() {
	var (
		expectedAccAddr string = "test-acc-addr"
		expectedPortID  string = "test-port"
	)

	suite.chainB.GetSimApp().ICAHostKeeper.SetInterchainAccountAddress(suite.chainB.GetContext(), expectedPortID, expectedAccAddr)

	retrievedAddr, found := suite.chainB.GetSimApp().ICAHostKeeper.GetInterchainAccountAddress(suite.chainB.GetContext(), expectedPortID)
	suite.Require().True(found)
	suite.Require().Equal(expectedAccAddr, retrievedAddr)
}