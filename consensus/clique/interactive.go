package clique

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

type chainContext struct {
	chainReader consensus.ChainHeaderReader
	engine      consensus.Engine
}

func newChainContext(chainReader consensus.ChainHeaderReader, engine consensus.Engine) *chainContext {
	return &chainContext{
		chainReader: chainReader,
		engine:      engine,
	}
}

// Engine retrieves the chain's consensus engine.
func (cc *chainContext) Engine() consensus.Engine {
	return cc.engine
}

// GetHeader returns the hash corresponding to their hash.
func (cc *chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return cc.chainReader.GetHeader(hash, number)
}

func getInteractiveABIAndAddrs() (map[string]abi.ABI, map[string]common.Address) {
	abiMap := make(map[string]abi.ABI, 0)
	tmpABI, _ := abi.JSON(strings.NewReader(nativeMintInteractiveABI))
	abiMap[nativeMintContractName] = tmpABI

	// Contract Addresses
	addrs := make(map[string]common.Address, 0)

	// v1 addresses
	addrs[nativeMintContractName] = nativeMintAddress
	return abiMap, addrs // TODO
}

// executeMsg executes transaction sent to system contracts.
func executeMsg(msg *core.Message, state *state.StateDB, header *types.Header, chainContext core.ChainContext, chainConfig *params.ChainConfig) (ret []byte, err error) {
	// Set gas price to zero
	context := core.NewEVMBlockContext(header, chainContext, nil)
	txContext := core.NewEVMTxContext(msg)
	vmenv := vm.NewEVM(context, txContext, state, chainConfig, vm.Config{})

	// msg.GasPrice

	ret, _, err = vmenv.Call(vm.AccountRef(msg.From), *msg.To, msg.Data, msg.GasLimit, msg.Value)

	if err != nil {
		return ret, err
	}

	return ret, nil
}