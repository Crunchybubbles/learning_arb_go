// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UniswapFlashQuery

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UniswapFlashQueryABI is the input ABI used to generate the binding from.
const UniswapFlashQueryABI = "[{\"inputs\":[{\"internalType\":\"contractUniswapV2Factory\",\"name\":\"_uniswapFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stop\",\"type\":\"uint256\"}],\"name\":\"getPairsByIndexRange\",\"outputs\":[{\"internalType\":\"address[3][]\",\"name\":\"\",\"type\":\"address[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Pair[]\",\"name\":\"_pairs\",\"type\":\"address[]\"}],\"name\":\"getReservesByPairs\",\"outputs\":[{\"internalType\":\"uint256[3][]\",\"name\":\"\",\"type\":\"uint256[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UniswapFlashQueryBin is the compiled bytecode used for deploying new contracts.
var UniswapFlashQueryBin = "0x608060405234801561001057600080fd5b50610fd2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634dbf0f391461003b578063ab2217e41461006b575b600080fd5b61005560048036038101906100509190610823565b61009b565b6040516100629190610bac565b60405180910390f35b6100856004803603810190610080919061089d565b61029d565b6040516100929190610b8a565b60405180910390f35b606060008383905067ffffffffffffffff8111156100bc576100bb610ea1565b5b6040519080825280602002602001820160405280156100f557816020015b6100e26106c9565b8152602001906001900390816100da5790505b50905060005b848490508110156102925784848281811061011957610118610e72565b5b905060200201602081019061012e9190610870565b73ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561017357600080fd5b505afa158015610187573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101ab91906108f0565b826dffffffffffffffffffffffffffff169250816dffffffffffffffffffffffffffff1691508063ffffffff1690508484815181106101ed576101ec610e72565b5b602002602001015160006003811061020857610207610e72565b5b6020020185858151811061021f5761021e610e72565b5b602002602001015160016003811061023a57610239610e72565b5b6020020186868151811061025157610250610e72565b5b602002602001015160026003811061026c5761026b610e72565b5b60200201838152508381525083815250505050808061028a90610dfa565b9150506100fb565b508091505092915050565b606060008473ffffffffffffffffffffffffffffffffffffffff1663574f2ba36040518163ffffffff1660e01b815260040160206040518083038186803b1580156102e757600080fd5b505afa1580156102fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031f9190610943565b90508083111561032d578092505b83831015610370576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036790610bce565b60405180910390fd5b6000848461037e9190610d3c565b905060008167ffffffffffffffff81111561039c5761039b610ea1565b5b6040519080825280602002602001820160405280156103d557816020015b6103c26106eb565b8152602001906001900390816103ba5790505b50905060005b828110156106bb5760008873ffffffffffffffffffffffffffffffffffffffff16631e3dd18b838a61040d9190610ce6565b6040518263ffffffff1660e01b81526004016104299190610bee565b60206040518083038186803b15801561044157600080fd5b505afa158015610455573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047991906107f6565b90508073ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b1580156104c157600080fd5b505afa1580156104d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f991906107f6565b83838151811061050c5761050b610e72565b5b602002602001015160006003811061052757610526610e72565b5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b1580156105a457600080fd5b505afa1580156105b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105dc91906107f6565b8383815181106105ef576105ee610e72565b5b602002602001015160016003811061060a57610609610e72565b5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508083838151811061065557610654610e72565b5b60200260200101516002600381106106705761066f610e72565b5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505080806106b390610dfa565b9150506103db565b508093505050509392505050565b6040518060600160405280600390602082028036833780820191505090505090565b6040518060600160405280600390602082028036833780820191505090505090565b60008151905061071c81610f12565b92915050565b60008083601f84011261073857610737610ed5565b5b8235905067ffffffffffffffff81111561075557610754610ed0565b5b60208301915083602082028301111561077157610770610eda565b5b9250929050565b60008135905061078781610f29565b92915050565b60008135905061079c81610f40565b92915050565b6000815190506107b181610f57565b92915050565b6000813590506107c681610f6e565b92915050565b6000815190506107db81610f6e565b92915050565b6000815190506107f081610f85565b92915050565b60006020828403121561080c5761080b610ee4565b5b600061081a8482850161070d565b91505092915050565b6000806020838503121561083a57610839610ee4565b5b600083013567ffffffffffffffff81111561085857610857610edf565b5b61086485828601610722565b92509250509250929050565b60006020828403121561088657610885610ee4565b5b600061089484828501610778565b91505092915050565b6000806000606084860312156108b6576108b5610ee4565b5b60006108c48682870161078d565b93505060206108d5868287016107b7565b92505060406108e6868287016107b7565b9150509250925092565b60008060006060848603121561090957610908610ee4565b5b6000610917868287016107a2565b9350506020610928868287016107a2565b9250506040610939868287016107e1565b9150509250925092565b60006020828403121561095957610958610ee4565b5b6000610967848285016107cc565b91505092915050565b600061097c83836109d0565b60208301905092915050565b600061099483836109df565b60608301905092915050565b60006109ac8383610af2565b60608301905092915050565b60006109c48383610b6c565b60208301905092915050565b6109d981610d70565b82525050565b6109e881610c3d565b6109f28184610c9d565b92506109fd82610c09565b8060005b83811015610a2e578151610a158782610970565b9650610a2083610c69565b925050600181019050610a01565b505050505050565b6000610a4182610c48565b610a4b8185610ca8565b9350610a5683610c13565b8060005b83811015610a87578151610a6e8882610988565b9750610a7983610c76565b925050600181019050610a5a565b5085935050505092915050565b6000610a9f82610c53565b610aa98185610cb9565b9350610ab483610c23565b8060005b83811015610ae5578151610acc88826109a0565b9750610ad783610c83565b925050600181019050610ab8565b5085935050505092915050565b610afb81610c5e565b610b058184610cca565b9250610b1082610c33565b8060005b83811015610b41578151610b2887826109b8565b9650610b3383610c90565b925050600181019050610b14565b505050505050565b6000610b56602083610cd5565b9150610b6182610ee9565b602082019050919050565b610b7581610de0565b82525050565b610b8481610de0565b82525050565b60006020820190508181036000830152610ba48184610a36565b905092915050565b60006020820190508181036000830152610bc68184610a94565b905092915050565b60006020820190508181036000830152610be781610b49565b9050919050565b6000602082019050610c036000830184610b7b565b92915050565b6000819050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050919050565b600060039050919050565b600081519050919050565b600081519050919050565b600060039050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610cf182610de0565b9150610cfc83610de0565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610d3157610d30610e43565b5b828201905092915050565b6000610d4782610de0565b9150610d5283610de0565b925082821015610d6557610d64610e43565b5b828203905092915050565b6000610d7b82610dc0565b9050919050565b6000610d8d82610d70565b9050919050565b6000610d9f82610d70565b9050919050565b60006dffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b6000610e0582610de0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610e3857610e37610e43565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b7f73746172742063616e6e6f7420626520686967686572207468616e2073746f70600082015250565b610f1b81610d70565b8114610f2657600080fd5b50565b610f3281610d82565b8114610f3d57600080fd5b50565b610f4981610d94565b8114610f5457600080fd5b50565b610f6081610da6565b8114610f6b57600080fd5b50565b610f7781610de0565b8114610f8257600080fd5b50565b610f8e81610dea565b8114610f9957600080fd5b5056fea264697066735822122029471bb47bbab65098cea05c20b3492ec6bf1400c2ed23f5071bc42f2280245064736f6c63430008060033"

// DeployUniswapFlashQuery deploys a new Ethereum contract, binding an instance of UniswapFlashQuery to it.
func DeployUniswapFlashQuery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniswapFlashQuery, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapFlashQueryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UniswapFlashQueryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapFlashQuery{UniswapFlashQueryCaller: UniswapFlashQueryCaller{contract: contract}, UniswapFlashQueryTransactor: UniswapFlashQueryTransactor{contract: contract}, UniswapFlashQueryFilterer: UniswapFlashQueryFilterer{contract: contract}}, nil
}

// UniswapFlashQuery is an auto generated Go binding around an Ethereum contract.
type UniswapFlashQuery struct {
	UniswapFlashQueryCaller     // Read-only binding to the contract
	UniswapFlashQueryTransactor // Write-only binding to the contract
	UniswapFlashQueryFilterer   // Log filterer for contract events
}

// UniswapFlashQueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapFlashQueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapFlashQueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapFlashQueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapFlashQueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapFlashQueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapFlashQuerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapFlashQuerySession struct {
	Contract     *UniswapFlashQuery // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniswapFlashQueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapFlashQueryCallerSession struct {
	Contract *UniswapFlashQueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniswapFlashQueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapFlashQueryTransactorSession struct {
	Contract     *UniswapFlashQueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniswapFlashQueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapFlashQueryRaw struct {
	Contract *UniswapFlashQuery // Generic contract binding to access the raw methods on
}

// UniswapFlashQueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapFlashQueryCallerRaw struct {
	Contract *UniswapFlashQueryCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapFlashQueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapFlashQueryTransactorRaw struct {
	Contract *UniswapFlashQueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapFlashQuery creates a new instance of UniswapFlashQuery, bound to a specific deployed contract.
func NewUniswapFlashQuery(address common.Address, backend bind.ContractBackend) (*UniswapFlashQuery, error) {
	contract, err := bindUniswapFlashQuery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapFlashQuery{UniswapFlashQueryCaller: UniswapFlashQueryCaller{contract: contract}, UniswapFlashQueryTransactor: UniswapFlashQueryTransactor{contract: contract}, UniswapFlashQueryFilterer: UniswapFlashQueryFilterer{contract: contract}}, nil
}

// NewUniswapFlashQueryCaller creates a new read-only instance of UniswapFlashQuery, bound to a specific deployed contract.
func NewUniswapFlashQueryCaller(address common.Address, caller bind.ContractCaller) (*UniswapFlashQueryCaller, error) {
	contract, err := bindUniswapFlashQuery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapFlashQueryCaller{contract: contract}, nil
}

// NewUniswapFlashQueryTransactor creates a new write-only instance of UniswapFlashQuery, bound to a specific deployed contract.
func NewUniswapFlashQueryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapFlashQueryTransactor, error) {
	contract, err := bindUniswapFlashQuery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapFlashQueryTransactor{contract: contract}, nil
}

// NewUniswapFlashQueryFilterer creates a new log filterer instance of UniswapFlashQuery, bound to a specific deployed contract.
func NewUniswapFlashQueryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapFlashQueryFilterer, error) {
	contract, err := bindUniswapFlashQuery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapFlashQueryFilterer{contract: contract}, nil
}

// bindUniswapFlashQuery binds a generic wrapper to an already deployed contract.
func bindUniswapFlashQuery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapFlashQueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapFlashQuery *UniswapFlashQueryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapFlashQuery.Contract.UniswapFlashQueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapFlashQuery *UniswapFlashQueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapFlashQuery.Contract.UniswapFlashQueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapFlashQuery *UniswapFlashQueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapFlashQuery.Contract.UniswapFlashQueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapFlashQuery *UniswapFlashQueryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapFlashQuery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapFlashQuery *UniswapFlashQueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapFlashQuery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapFlashQuery *UniswapFlashQueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapFlashQuery.Contract.contract.Transact(opts, method, params...)
}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_UniswapFlashQuery *UniswapFlashQueryCaller) GetPairsByIndexRange(opts *bind.CallOpts, _uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	var out []interface{}
	err := _UniswapFlashQuery.contract.Call(opts, &out, "getPairsByIndexRange", _uniswapFactory, _start, _stop)

	if err != nil {
		return *new([][3]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]common.Address)).(*[][3]common.Address)

	return out0, err

}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_UniswapFlashQuery *UniswapFlashQuerySession) GetPairsByIndexRange(_uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	return _UniswapFlashQuery.Contract.GetPairsByIndexRange(&_UniswapFlashQuery.CallOpts, _uniswapFactory, _start, _stop)
}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_UniswapFlashQuery *UniswapFlashQueryCallerSession) GetPairsByIndexRange(_uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	return _UniswapFlashQuery.Contract.GetPairsByIndexRange(&_UniswapFlashQuery.CallOpts, _uniswapFactory, _start, _stop)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapFlashQuery *UniswapFlashQueryCaller) GetReservesByPairs(opts *bind.CallOpts, _pairs []common.Address) ([][3]*big.Int, error) {
	var out []interface{}
	err := _UniswapFlashQuery.contract.Call(opts, &out, "getReservesByPairs", _pairs)

	if err != nil {
		return *new([][3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]*big.Int)).(*[][3]*big.Int)

	return out0, err

}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapFlashQuery *UniswapFlashQuerySession) GetReservesByPairs(_pairs []common.Address) ([][3]*big.Int, error) {
	return _UniswapFlashQuery.Contract.GetReservesByPairs(&_UniswapFlashQuery.CallOpts, _pairs)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapFlashQuery *UniswapFlashQueryCallerSession) GetReservesByPairs(_pairs []common.Address) ([][3]*big.Int, error) {
	return _UniswapFlashQuery.Contract.GetReservesByPairs(&_UniswapFlashQuery.CallOpts, _pairs)
}
