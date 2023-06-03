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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ihpDeployedAddr string
var ihpClient *ethclient.Client
var ihpConn *ihpApi.Api

func ConnectIHPContract() {
	// contractAddressHex := os.Getenv("CONTRACT_ADDRESS")
	// contractABIStr := `[{"constant":false,"inputs":[{"name":"uhpId","type":"uint256"},{"name":"uri","type":"string"}],"name":"storeProfile","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"uhpId","type":"uint256"}],"name":"getProfile","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`

	// ihpClient, err := ethihpClient.Dial("https://eth-sepolia.g.alchemy.com/v2/f1thP4VKMvnfX9LaGREjrMogP4Vnplo0")
	var err error
	ihpClient, err = ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/f1thP4VKMvnfX9LaGREjrMogP4Vnplo0")

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

func GetIHPConnection() {
	var err error
	ihpConn, err = ihpApi.NewApi(common.HexToAddress(os.Getenv("IHP_CONTRACT")), ihpClient)
	if err != nil {
		log.Fatal("error while creating api: ", err.Error())
	}

}

func CreateIHPProfile(cid string, name string) (int, common.Hash) {
	ihpID := rand.Intn(99999999999) + 10000000000
	ihpInt := big.NewInt(int64(ihpID))
	addr := os.Getenv("DOCTOR_ADDRESS")
	pvt := os.Getenv("PVT_KEY")
	log.Println("attempting to store using", addr)

	pvtKey, err := crypto.HexToECDSA(pvt)
	if err != nil {
		log.Fatal("failed to convert pvt key: ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pvtKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal("failed to create transactor: ", err)
	}

	tx, err := ihpConn.StoreProfile(auth, ihpInt, cid, name)
	if err != nil {
		log.Fatal("error storing: ", err)
	}
	return ihpID, tx.Hash()
}

func GetIHPProfile(ihp int) (string, string) {

	pvt := os.Getenv("PVT_KEY")

	pvtKey, err := crypto.HexToECDSA(pvt)
	if err != nil {
		log.Fatal("failed to convert pvt key: ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pvtKey, big.NewInt(11155111))

	if err != nil {
		log.Fatal("failed to create transactor: ", err)
	}

	callOpts := &bind.CallOpts{
		From:    auth.From,
		Context: auth.Context,
	}
	uri, name, err := ihpConn.GetProfile(callOpts, big.NewInt(int64(ihp)))
	if err != nil {

		log.Fatal("error fetching profile: ", err)
	}
	return uri, name
}

func AddRecord(ihp int, cid string) common.Hash {
	pvt := os.Getenv("PVT_KEY")

	pvtKey, err := crypto.HexToECDSA(pvt)
	if err != nil {
		log.Fatal("failed to convert pvt key: ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pvtKey, big.NewInt(11155111))

	if err != nil {
		log.Fatal("failed to create transactor: ", err)
	}

	tx, err := ihpConn.AddRecord(auth, big.NewInt(int64(ihp)), cid)
	if err != nil {
		log.Fatal("error adding record: ", err)
	}
	return tx.Hash()
}
