package clique

import (
	"math"
	"math/big"
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
	// Contract ABI map
	abiMap := make(map[string]abi.ABI, 0)

	abiMap[committeeContractName], _ = abi.JSON(strings.NewReader(committeesInteractiveABI))
	abiMap[supplyControlContractName], _ = abi.JSON(strings.NewReader(supplyControlInteractiveABI))

	// Contract Addresses map
	addrs := make(map[string]common.Address, 0)

	// v1 addresses
	addrs[committeeContractName] = committeeAddress
	addrs[supplyControlContractName] = supplyControlAddress

	return abiMap, addrs
}

// executeMsg executes transaction sent to system contracts.
func executeMsg(msg *core.Message, state *state.StateDB, header *types.Header, chainContext core.ChainContext, chainConfig *params.ChainConfig) (ret []byte, err error) {
	// Set gas price to zero
	context := core.NewEVMBlockContext(header, chainContext, nil)
	txContext := core.NewEVMTxContext(msg)
	vmenv := vm.NewEVM(context, txContext, state, chainConfig, vm.Config{})
	ret, _, err = vmenv.Call(vm.AccountRef(msg.From), *msg.To, msg.Data, msg.GasLimit, msg.Value)
	// msg.GasPrice
	state.Finalise(true)
	return ret, err
}

func MessageType(caller common.Address, to *common.Address, data []byte) (*core.Message) {
	msg := &core.Message{
		From:              caller,
		To:                to,
		Value:             big.NewInt(0),
		GasLimit:          math.MaxUint64,
		GasPrice:          big.NewInt(0),
		GasFeeCap:         big.NewInt(0),
		GasTipCap:         big.NewInt(0),
		Data:              data,
		AccessList:        nil,
		SkipAccountChecks: true,
	}
	return msg
}