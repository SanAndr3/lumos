// Copyright LCT . (LumosChainTeam)

package testdata

import (
	contractutils "github.com/lumoschain/lumos/v20/contracts/utils"
	evmtypes "github.com/lumoschain/lumos/v20/x/evm/types"
)

func LoadERC20MinterV5Contract() (evmtypes.CompiledContract, error) {
	return contractutils.LegacyLoadContractFromJSONFile("ERC20Minter_OpenZeppelinV5.json")
}
