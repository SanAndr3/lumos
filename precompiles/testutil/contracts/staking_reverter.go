// Copyright LCT . (LumosChainTeam)

package contracts

import (
	contractutils "github.com/lumoschain/lumos/v20/contracts/utils"
	evmtypes "github.com/lumoschain/lumos/v20/x/evm/types"
)

func LoadStakingReverterContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("StakingReverter.json")
}
