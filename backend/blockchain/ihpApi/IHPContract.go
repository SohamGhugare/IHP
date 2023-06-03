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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uhpId\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uhpId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"storeProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506103f4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063b2bb8f841461003b578063f08f4f6414610050575b600080fd5b61004e610049366004610153565b610079565b005b61006361005e36600461020e565b610098565b6040516100709190610227565b60405180910390f35b60008281526020819052604090208061009283826102fe565b50505050565b600081815260208190526040902080546060919081906100b790610275565b80601f01602080910402602001604051908101604052809291908181526020018280546100e390610275565b80156101305780601f1061010557610100808354040283529160200191610130565b820191906000526020600020905b81548152906001019060200180831161011357829003601f168201915b5050505050915050919050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561016657600080fd5b82359150602083013567ffffffffffffffff8082111561018557600080fd5b818501915085601f83011261019957600080fd5b8135818111156101ab576101ab61013d565b604051601f8201601f19908116603f011681019083821181831017156101d3576101d361013d565b816040528281528860208487010111156101ec57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60006020828403121561022057600080fd5b5035919050565b600060208083528351808285015260005b8181101561025457858101830151858201604001528201610238565b506000604082860101526040601f19601f8301168501019250505092915050565b600181811c9082168061028957607f821691505b6020821081036102a957634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156102f957600081815260208120601f850160051c810160208610156102d65750805b601f850160051c820191505b818110156102f5578281556001016102e2565b5050505b505050565b815167ffffffffffffffff8111156103185761031861013d565b61032c816103268454610275565b846102af565b602080601f83116001811461036157600084156103495750858301515b600019600386901b1c1916600185901b1785556102f5565b600085815260208120601f198616915b8281101561039057888601518255948401946001909101908401610371565b50858210156103ae5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220c65c98e8505a236bf794147632f7ac1866a6b064f8567cbea15fdefc8ffa257f64736f6c63430008130033",
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
// Solidity: function getProfile(uint256 uhpId) view returns(string)
func (_Api *ApiCaller) GetProfile(opts *bind.CallOpts, uhpId *big.Int) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getProfile", uhpId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 uhpId) view returns(string)
func (_Api *ApiSession) GetProfile(uhpId *big.Int) (string, error) {
	return _Api.Contract.GetProfile(&_Api.CallOpts, uhpId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 uhpId) view returns(string)
func (_Api *ApiCallerSession) GetProfile(uhpId *big.Int) (string, error) {
	return _Api.Contract.GetProfile(&_Api.CallOpts, uhpId)
}

// StoreProfile is a paid mutator transaction binding the contract method 0xb2bb8f84.
//
// Solidity: function storeProfile(uint256 uhpId, string uri) returns()
func (_Api *ApiTransactor) StoreProfile(opts *bind.TransactOpts, uhpId *big.Int, uri string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "storeProfile", uhpId, uri)
}

// StoreProfile is a paid mutator transaction binding the contract method 0xb2bb8f84.
//
// Solidity: function storeProfile(uint256 uhpId, string uri) returns()
func (_Api *ApiSession) StoreProfile(uhpId *big.Int, uri string) (*types.Transaction, error) {
	return _Api.Contract.StoreProfile(&_Api.TransactOpts, uhpId, uri)
}

// StoreProfile is a paid mutator transaction binding the contract method 0xb2bb8f84.
//
// Solidity: function storeProfile(uint256 uhpId, string uri) returns()
func (_Api *ApiTransactorSession) StoreProfile(uhpId *big.Int, uri string) (*types.Transaction, error) {
	return _Api.Contract.StoreProfile(&_Api.TransactOpts, uhpId, uri)
}
