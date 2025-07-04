// Copyright LCT . (LumosChainTeam)

package testdata

import (
	contractutils "github.com/lumoschain/lumos/v20/contracts/utils"
	evmtypes "github.com/lumoschain/lumos/v20/x/evm/types"
)

func LoadStakingCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("StakingCaller.json")
}
