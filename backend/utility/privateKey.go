package utility

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetPrivateKey(client *ethclient.Client, addr string) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(addr)
	if err != nil {
		log.Fatal("failed to convert to ecdsa: ", err)
	}
	return privateKey
}

func GetAccountAuth(client *ethclient.Client, addr string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(addr)
	if err != nil {
		log.Fatal("failed to convert to ecdsa: ", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000000)

	return auth
}
