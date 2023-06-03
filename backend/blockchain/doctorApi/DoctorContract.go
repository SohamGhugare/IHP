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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"license\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"license\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pin\",\"type\":\"uint256\"}],\"name\":\"storeProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610526806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063079a227c1461003b578063f08f4f6414610050575b600080fd5b61004e61004936600461029d565b61007b565b005b61006361005e366004610312565b6100ad565b60405161007293929190610371565b60405180910390f35b6000848152602081905260409020806100948582610430565b50600281016100a38482610430565b5060030155505050565b6060806000806000808681526020019081526020016000209050806000018160020182600301548280546100e0906103a7565b80601f016020809104026020016040519081016040528092919081815260200182805461010c906103a7565b80156101595780601f1061012e57610100808354040283529160200191610159565b820191906000526020600020905b81548152906001019060200180831161013c57829003601f168201915b5050505050925081805461016c906103a7565b80601f0160208091040260200160405190810160405280929190818152602001828054610198906103a7565b80156101e55780601f106101ba576101008083540402835291602001916101e5565b820191906000526020600020905b8154815290600101906020018083116101c857829003601f168201915b50505050509150935093509350509193909250565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261022157600080fd5b813567ffffffffffffffff8082111561023c5761023c6101fa565b604051601f8301601f19908116603f01168101908282118183101715610264576102646101fa565b8160405283815286602085880101111561027d57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080608085870312156102b357600080fd5b84359350602085013567ffffffffffffffff808211156102d257600080fd5b6102de88838901610210565b945060408701359150808211156102f457600080fd5b5061030187828801610210565b949793965093946060013593505050565b60006020828403121561032457600080fd5b5035919050565b6000815180845260005b8181101561035157602081850181015186830182015201610335565b506000602082860101526020601f19601f83011685010191505092915050565b606081526000610384606083018661032b565b8281036020840152610396818661032b565b915050826040830152949350505050565b600181811c908216806103bb57607f821691505b6020821081036103db57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561042b57600081815260208120601f850160051c810160208610156104085750805b601f850160051c820191505b8181101561042757828155600101610414565b5050505b505050565b815167ffffffffffffffff81111561044a5761044a6101fa565b61045e8161045884546103a7565b846103e1565b602080601f831160018114610493576000841561047b5750858301515b600019600386901b1c1916600185901b178555610427565b600085815260208120601f198616915b828110156104c2578886015182559484019460019091019084016104a3565b50858210156104e05787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea264697066735822122079ddf76185f1f98bf124794cb420057fae3f5378b1334279b6004960c279b33064736f6c63430008130033",
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
// Solidity: function getProfile(uint256 license) view returns(string, string, uint256)
func (_Api *ApiCaller) GetProfile(opts *bind.CallOpts, license *big.Int) (string, string, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getProfile", license)

	if err != nil {
		return *new(string), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 license) view returns(string, string, uint256)
func (_Api *ApiSession) GetProfile(license *big.Int) (string, string, *big.Int, error) {
	return _Api.Contract.GetProfile(&_Api.CallOpts, license)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 license) view returns(string, string, uint256)
func (_Api *ApiCallerSession) GetProfile(license *big.Int) (string, string, *big.Int, error) {
	return _Api.Contract.GetProfile(&_Api.CallOpts, license)
}

// StoreProfile is a paid mutator transaction binding the contract method 0x079a227c.
//
// Solidity: function storeProfile(uint256 license, string name, string email, uint256 pin) returns()
func (_Api *ApiTransactor) StoreProfile(opts *bind.TransactOpts, license *big.Int, name string, email string, pin *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "storeProfile", license, name, email, pin)
}

// StoreProfile is a paid mutator transaction binding the contract method 0x079a227c.
//
// Solidity: function storeProfile(uint256 license, string name, string email, uint256 pin) returns()
func (_Api *ApiSession) StoreProfile(license *big.Int, name string, email string, pin *big.Int) (*types.Transaction, error) {
	return _Api.Contract.StoreProfile(&_Api.TransactOpts, license, name, email, pin)
}

// StoreProfile is a paid mutator transaction binding the contract method 0x079a227c.
//
// Solidity: function storeProfile(uint256 license, string name, string email, uint256 pin) returns()
func (_Api *ApiTransactorSession) StoreProfile(license *big.Int, name string, email string, pin *big.Int) (*types.Transaction, error) {
	return _Api.Contract.StoreProfile(&_Api.TransactOpts, license, name, email, pin)
}
