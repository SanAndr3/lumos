package keeper_test

import (
	"github.com/stretchr/testify/suite"

	"github.com/lumoschain/lumos/v20/testutil/integration/lumos/factory"
	"github.com/lumoschain/lumos/v20/testutil/integration/lumos/grpc"
	"github.com/lumoschain/lumos/v20/testutil/integration/lumos/keyring"
	"github.com/lumoschain/lumos/v20/testutil/integration/lumos/network"
	"github.com/lumoschain/lumos/v20/x/inflation/v1/types"
)

var denomMint = types.DefaultInflationDenom

type KeeperTestSuite struct {
	suite.Suite

	network *network.UnitTestNetwork
	handler grpc.Handler
	keyring keyring.Keyring
	factory factory.TxFactory
}

func (suite *KeeperTestSuite) SetupTest() {
	keys := keyring.New(2)
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keys.GetAllAccAddrs()...),
	)
	gh := grpc.NewIntegrationHandler(nw)
	tf := factory.New(nw, gh)

	suite.network = nw
	suite.factory = tf
	suite.handler = gh
	suite.keyring = keys
}
