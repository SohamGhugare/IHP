package contract

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type DoctorWallet struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

func CreateDoctorWallet(privateKey *ecdsa.PrivateKey) *DoctorWallet {
	address := common.BytesToAddress(crypto.FromECDSA(privateKey))
	return &DoctorWallet{
		Address:    address,
		PrivateKey: privateKey,
	}
}
