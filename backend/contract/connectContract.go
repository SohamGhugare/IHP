package contract

import (
	"log"
	"math/big"
	"math/rand"
	"os"

	"github.com/SohamGhugare/IHP/blockchain/api"
	"github.com/SohamGhugare/IHP/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var deployedAddr string
var client *ethclient.Client
var conn *api.Api

var ihpInt *big.Int

func ConnectContract() {
	// contractAddressHex := os.Getenv("CONTRACT_ADDRESS")
	// contractABIStr := `[{"constant":false,"inputs":[{"name":"uhpId","type":"uint256"},{"name":"uri","type":"string"}],"name":"storeProfile","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"uhpId","type":"uint256"}],"name":"getProfile","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`

	// client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/f1thP4VKMvnfX9LaGREjrMogP4Vnplo0")
	var err error
	client, err = ethclient.Dial("http://127.0.0.1:7545")

	if err != nil {
		log.Fatal("error creating client:", err)
	}

	pvtKey := os.Getenv("PVT_KEY")
	auth := utility.GetAccountAuth(client, pvtKey)
	deployedContractAddr, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		log.Fatal("error deploying:", err)
	}
	_, _ = tx, instance
	deployedAddr = deployedContractAddr.Hex()

}

func GetConnection() {
	var err error
	conn, err = api.NewApi(common.HexToAddress(deployedAddr), client)
	if err != nil {
		log.Fatal("error while creating api: ", err.Error())
	}
	ihpID := rand.Intn(99999999999) + 10000000000
	ihpInt = big.NewInt(int64(ihpID))
}

func CreateProfile() {

	tx, err := conn.StoreProfile(&bind.TransactOpts{}, ihpInt, "test123")
	if err != nil {
		log.Fatal("error storing: ", err)
	}
	log.Println("successfully created profile:", tx)
}

func GetProfile() {
	tx, err := conn.GetProfile(&bind.CallOpts{}, ihpInt)
	if err != nil {
		log.Fatal("error fetching profile: ", err)
	}
	log.Println("fetched string:", tx)
}
