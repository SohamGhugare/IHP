package contract

import (
	"log"
	"math/big"
	"math/rand"
	"os"

	ihpApi "github.com/SohamGhugare/IHP/blockchain/ihpApi"

	"github.com/SohamGhugare/IHP/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ihpDeployedAddr string
var ihpClient *ethclient.Client
var ihpConn *ihpApi.Api

var ihpInt *big.Int

func ConnectIHPContract() {
	// contractAddressHex := os.Getenv("CONTRACT_ADDRESS")
	// contractABIStr := `[{"constant":false,"inputs":[{"name":"uhpId","type":"uint256"},{"name":"uri","type":"string"}],"name":"storeProfile","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"uhpId","type":"uint256"}],"name":"getProfile","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`

	// ihpClient, err := ethihpClient.Dial("https://eth-sepolia.g.alchemy.com/v2/f1thP4VKMvnfX9LaGREjrMogP4Vnplo0")
	var err error
	ihpClient, err = ethclient.Dial("http://127.0.0.1:7545")

	if err != nil {
		log.Fatal("error creating ihpClient:", err)
	}

	pvtKey := os.Getenv("PVT_KEY")
	auth := utility.GetAccountAuth(ihpClient, pvtKey)
	deployedContractAddr, tx, instance, err := ihpApi.DeployApi(auth, ihpClient)
	if err != nil {
		log.Fatal("error deploying:", err)
	}
	_, _ = tx, instance
	ihpDeployedAddr = deployedContractAddr.Hex()

}

func GetConnection() {
	var err error
	ihpConn, err = ihpApi.NewApi(common.HexToAddress(ihpDeployedAddr), ihpClient)
	if err != nil {
		log.Fatal("error while creating api: ", err.Error())
	}
	ihpID := rand.Intn(99999999999) + 10000000000
	ihpInt = big.NewInt(int64(ihpID))
}

func CreateProfile() {

	tx, err := ihpConn.StoreProfile(&bind.TransactOpts{}, ihpInt, "test123")
	if err != nil {
		log.Fatal("error storing: ", err)
	}
	log.Println("successfully created profile:", tx)
}

func GetProfile() {
	tx, err := ihpConn.GetProfile(&bind.CallOpts{}, ihpInt)
	if err != nil {
		log.Fatal("error fetching profile: ", err)
	}
	log.Println("fetched string:", tx)
}
