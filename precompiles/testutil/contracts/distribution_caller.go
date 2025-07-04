// Copyright LCT . (LumosChainTeam)

package contracts

import (
	contractutils "github.com/lumoschain/lumos/v20/contracts/utils"
	evmtypes "github.com/lumoschain/lumos/v20/x/evm/types"
)

func LoadDistributionCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("DistributionCaller.json")
}
