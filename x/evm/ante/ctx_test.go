// Copyright LCT . (LumosChainTeam)

package ante_test

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/lumoschain/lumos/v20/testutil/integration/lumos/network"
	evmante "github.com/lumoschain/lumos/v20/x/evm/ante"
)

func (suite *EvmAnteTestSuite) TestBuildEvmExecutionCtx() {
	network := network.New()

	ctx := evmante.BuildEvmExecutionCtx(network.GetContext())

	suite.Equal(storetypes.GasConfig{}, ctx.KVGasConfig())
	suite.Equal(storetypes.GasConfig{}, ctx.TransientKVGasConfig())
}
