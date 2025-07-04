// Copyright LCT . (LumosChainTeam)

package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/lumoschain/lumos/v20/x/evm/types"
)

var (
	//go:embed compiled_contracts/WLUMOS.json
	WLUMOSJSON []byte

	// WLUMOSContract is the compiled contract of WLUMOS
	WLUMOSContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WLUMOSJSON, &WLUMOSContract)
	if err != nil {
		panic(err)
	}

	if len(WLUMOSContract.Bin) == 0 {
		panic("failed to load WLUMOS smart contract")
	}
}
