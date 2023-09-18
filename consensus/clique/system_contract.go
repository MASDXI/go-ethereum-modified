package clique

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type systemContract struct {
	address    common.Address
	deployedBytecode    []byte
	packFun func() ([]byte, error)
}

func getSystemContracts(abi map[string]abi.ABI, initCommittees []common.Address, admin common.Address, voteDelay *big.Int, votePeriod *big.Int) []*systemContract {

	return []*systemContract{
		{
			// Committee Contract
			address: committeeAddress,
			deployedBytecode: committeeCode,
			packFun: func() ([]byte, error) {
				return abi[committeeContractName].Pack("initialize",
					voteDelay,// voteDelay
					votePeriod,// votePeriod
					initCommittees,// committee []
					admin, // admin
				)
			},
		},
		{
			address: supplyControlAddress,
			deployedBytecode: supplyControlCode,
			packFun: func() ([]byte, error) {
				return abi[supplyControlContractName].Pack("initialize",
					voteDelay,// voteDelay
					votePeriod,// votePeriod
					committeeAddress,
				)
			},
		},
	}
}

var committeeCode = common.FromHex("")
var supplyControlCode = common.FromHex("")