// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uhpId\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uhpId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"storeProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610502806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637a38ac351461003b578063f08f4f6414610050575b600080fd5b61004e610049366004610289565b61007a565b005b61006361005e3660046102f6565b6100a9565b604051610071929190610355565b60405180910390f35b600083815260208190526040902080610093848261040c565b50600181016100a2838261040c565b5050505050565b6000818152602081905260409020805460609182918190600182019082906100d090610383565b80601f01602080910402602001604051908101604052809291908181526020018280546100fc90610383565b80156101495780601f1061011e57610100808354040283529160200191610149565b820191906000526020600020905b81548152906001019060200180831161012c57829003601f168201915b5050505050915080805461015c90610383565b80601f016020809104026020016040519081016040528092919081815260200182805461018890610383565b80156101d55780601f106101aa576101008083540402835291602001916101d5565b820191906000526020600020905b8154815290600101906020018083116101b857829003601f168201915b505050505090509250925050915091565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261020d57600080fd5b813567ffffffffffffffff80821115610228576102286101e6565b604051601f8301601f19908116603f01168101908282118183101715610250576102506101e6565b8160405283815286602085880101111561026957600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561029e57600080fd5b83359250602084013567ffffffffffffffff808211156102bd57600080fd5b6102c9878388016101fc565b935060408601359150808211156102df57600080fd5b506102ec868287016101fc565b9150509250925092565b60006020828403121561030857600080fd5b5035919050565b6000815180845260005b8181101561033557602081850181015186830182015201610319565b506000602082860101526020601f19601f83011685010191505092915050565b604081526000610368604083018561030f565b828103602084015261037a818561030f565b95945050505050565b600181811c9082168061039757607f821691505b6020821081036103b757634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561040757600081815260208120601f850160051c810160208610156103e45750805b601f850160051c820191505b81811015610403578281556001016103f0565b5050505b505050565b815167ffffffffffffffff811115610426576104266101e6565b61043a816104348454610383565b846103bd565b602080601f83116001811461046f57600084156104575750858301515b600019600386901b1c1916600185901b178555610403565b600085815260208120601f198616915b8281101561049e5788860151825594840194600190910190840161047f565b50858210156104bc5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220e24e861058d259b9555da94abd5501c74267ee24f66e3bec0f0cbe10499f500264736f6c63430008130033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 uhpId) view returns(string, string)
func (_Api *ApiCaller) GetProfile(opts *bind.CallOpts, uhpId *big.Int) (string, string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getProfile", uhpId)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 uhpId) view returns(string, string)
func (_Api *ApiSession) GetProfile(uhpId *big.Int) (string, string, error) {
	return _Api.Contract.GetProfile(&_Api.CallOpts, uhpId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 uhpId) view returns(string, string)
func (_Api *ApiCallerSession) GetProfile(uhpId *big.Int) (string, string, error) {
	return _Api.Contract.GetProfile(&_Api.CallOpts, uhpId)
}

// StoreProfile is a paid mutator transaction binding the contract method 0x7a38ac35.
//
// Solidity: function storeProfile(uint256 uhpId, string uri, string name) returns()
func (_Api *ApiTransactor) StoreProfile(opts *bind.TransactOpts, uhpId *big.Int, uri string, name string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "storeProfile", uhpId, uri, name)
}

// StoreProfile is a paid mutator transaction binding the contract method 0x7a38ac35.
//
// Solidity: function storeProfile(uint256 uhpId, string uri, string name) returns()
func (_Api *ApiSession) StoreProfile(uhpId *big.Int, uri string, name string) (*types.Transaction, error) {
	return _Api.Contract.StoreProfile(&_Api.TransactOpts, uhpId, uri, name)
}

// StoreProfile is a paid mutator transaction binding the contract method 0x7a38ac35.
//
// Solidity: function storeProfile(uint256 uhpId, string uri, string name) returns()
func (_Api *ApiTransactorSession) StoreProfile(uhpId *big.Int, uri string, name string) (*types.Transaction, error) {
	return _Api.Contract.StoreProfile(&_Api.TransactOpts, uhpId, uri, name)
}
