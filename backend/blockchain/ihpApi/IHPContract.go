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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uhpId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"record_cid\",\"type\":\"string\"}],\"name\":\"addRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uhpId\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uhpId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRecord\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uhpId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"storeProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061072b806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806333406502146100515780634131d0e31461007a5780637a38ac351461008f578063f08f4f64146100a2575b600080fd5b61006461005f366004610376565b6100c3565b60405161007191906103de565b60405180910390f35b61008d61008836600461049b565b6101d8565b005b61008d61009d3660046104e2565b61020a565b6100b56100b036600461054f565b610239565b604051610071929190610568565b600082815260208190526040902060020154606090821061011a5760405162461bcd60e51b815260206004820152600d60248201526c092dcecc2d8d2c84092dcc8caf609b1b604482015260640160405180910390fd5b600083815260208190526040902060020180548390811061013d5761013d610596565b906000526020600020018054610152906105ac565b80601f016020809104026020016040519081016040528092919081815260200182805461017e906105ac565b80156101cb5780601f106101a0576101008083540402835291602001916101cb565b820191906000526020600020905b8154815290600101906020018083116101ae57829003601f168201915b5050505050905092915050565b600082815260208181526040822060028101805460018101825590845291909220016102048382610635565b50505050565b6000838152602081905260409020806102238482610635565b50600181016102328382610635565b5050505050565b600081815260208190526040902080546060918291819060018201908290610260906105ac565b80601f016020809104026020016040519081016040528092919081815260200182805461028c906105ac565b80156102d95780601f106102ae576101008083540402835291602001916102d9565b820191906000526020600020905b8154815290600101906020018083116102bc57829003601f168201915b505050505091508080546102ec906105ac565b80601f0160208091040260200160405190810160405280929190818152602001828054610318906105ac565b80156103655780601f1061033a57610100808354040283529160200191610365565b820191906000526020600020905b81548152906001019060200180831161034857829003601f168201915b505050505090509250925050915091565b6000806040838503121561038957600080fd5b50508035926020909101359150565b6000815180845260005b818110156103be576020818501810151868301820152016103a2565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006103f16020830184610398565b9392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261041f57600080fd5b813567ffffffffffffffff8082111561043a5761043a6103f8565b604051601f8301601f19908116603f01168101908282118183101715610462576104626103f8565b8160405283815286602085880101111561047b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156104ae57600080fd5b82359150602083013567ffffffffffffffff8111156104cc57600080fd5b6104d88582860161040e565b9150509250929050565b6000806000606084860312156104f757600080fd5b83359250602084013567ffffffffffffffff8082111561051657600080fd5b6105228783880161040e565b9350604086013591508082111561053857600080fd5b506105458682870161040e565b9150509250925092565b60006020828403121561056157600080fd5b5035919050565b60408152600061057b6040830185610398565b828103602084015261058d8185610398565b95945050505050565b634e487b7160e01b600052603260045260246000fd5b600181811c908216806105c057607f821691505b6020821081036105e057634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561063057600081815260208120601f850160051c8101602086101561060d5750805b601f850160051c820191505b8181101561062c57828155600101610619565b5050505b505050565b815167ffffffffffffffff81111561064f5761064f6103f8565b6106638161065d84546105ac565b846105e6565b602080601f83116001811461069857600084156106805750858301515b600019600386901b1c1916600185901b17855561062c565b600085815260208120601f198616915b828110156106c7578886015182559484019460019091019084016106a8565b50858210156106e55787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220897f08c359591fcfa0e7752c8cb38bb72819ab93f7284c1fc3bc0dc6d13503ad64736f6c63430008130033",
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

// GetRecord is a free data retrieval call binding the contract method 0x33406502.
//
// Solidity: function getRecord(uint256 uhpId, uint256 index) view returns(string)
func (_Api *ApiCaller) GetRecord(opts *bind.CallOpts, uhpId *big.Int, index *big.Int) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRecord", uhpId, index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetRecord is a free data retrieval call binding the contract method 0x33406502.
//
// Solidity: function getRecord(uint256 uhpId, uint256 index) view returns(string)
func (_Api *ApiSession) GetRecord(uhpId *big.Int, index *big.Int) (string, error) {
	return _Api.Contract.GetRecord(&_Api.CallOpts, uhpId, index)
}

// GetRecord is a free data retrieval call binding the contract method 0x33406502.
//
// Solidity: function getRecord(uint256 uhpId, uint256 index) view returns(string)
func (_Api *ApiCallerSession) GetRecord(uhpId *big.Int, index *big.Int) (string, error) {
	return _Api.Contract.GetRecord(&_Api.CallOpts, uhpId, index)
}

// AddRecord is a paid mutator transaction binding the contract method 0x4131d0e3.
//
// Solidity: function addRecord(uint256 uhpId, string record_cid) returns()
func (_Api *ApiTransactor) AddRecord(opts *bind.TransactOpts, uhpId *big.Int, record_cid string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addRecord", uhpId, record_cid)
}

// AddRecord is a paid mutator transaction binding the contract method 0x4131d0e3.
//
// Solidity: function addRecord(uint256 uhpId, string record_cid) returns()
func (_Api *ApiSession) AddRecord(uhpId *big.Int, record_cid string) (*types.Transaction, error) {
	return _Api.Contract.AddRecord(&_Api.TransactOpts, uhpId, record_cid)
}

// AddRecord is a paid mutator transaction binding the contract method 0x4131d0e3.
//
// Solidity: function addRecord(uint256 uhpId, string record_cid) returns()
func (_Api *ApiTransactorSession) AddRecord(uhpId *big.Int, record_cid string) (*types.Transaction, error) {
	return _Api.Contract.AddRecord(&_Api.TransactOpts, uhpId, record_cid)
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
