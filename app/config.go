// Copyright LCT . (LumosChainTeam)

package app

import (
	"github.com/lumoschain/lumos/v20/app/eips"
	evmconfig "github.com/lumoschain/lumos/v20/x/evm/config"
	"github.com/lumoschain/lumos/v20/x/evm/core/vm"
)

// The init function of the config file allows to setup the global
// configuration for the EVM, modifying the custom ones defined in evmOS.
func init() {
	err := evmconfig.NewEVMConfigurator().
		WithExtendedEips(lumosActivators).
		Configure()
	if err != nil {
		panic(err)
	}
}

// LumosActivators defines a map of opcode modifiers associated
// with a key defining the corresponding EIP.
var lumosActivators = map[string]func(*vm.JumpTable){
	"lumos_0": eips.Enable0000,
	"lumos_1": eips.Enable0001,
	"lumos_2": eips.Enable0002,
}
