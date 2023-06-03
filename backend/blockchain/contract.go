package blockchain

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ProfileContract represents the contract instance
type ProfileContract struct {
	contract *bind.BoundContract // Generic contract wrapper for the low-level calls
	abi      abi.ABI             // Parsed ABI definition of the contract
	address  common.Address      // Ethereum address of the contract
	auth     *bind.TransactOpts  // Transact options for contract calls
	callOpts *bind.CallOpts
}

// NewProfileContract creates a new instance of the ProfileContract
func NewProfileContract(address common.Address, client bind.ContractBackend, privateKey *ecdsa.PrivateKey, chainID *big.Int) (*ProfileContract, error) {
	contractABI, err := abi.JSON(strings.NewReader(profileContractABI))
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	contract := &ProfileContract{
		contract: bind.NewBoundContract(address, contractABI, client, client, nil),
		abi:      contractABI,
		address:  address,
		auth:     auth,
		callOpts: &bind.CallOpts{},
	}

	return contract, nil
}

// StoreProfile stores the profile URI for a given UHP ID
func (c *ProfileContract) StoreProfile(uhpId *big.Int, uri string) (*types.Transaction, error) {
	tx, err := c.contract.Transact(c.auth, "storeProfile", uhpId, uri)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// ProfileContractABI is the ABI specification for the Profile contract
const profileContractABI = `[{"constant":false,"inputs":[{"name":"uhpId","type":"uint256"},{"name":"uri","type":"string"}],"name":"storeProfile","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"uhpId","type":"uint256"}],"name":"getProfile","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`
