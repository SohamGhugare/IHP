package contract

import (
	"log"
	"math/big"
	"math/rand"
	"os"

	doctorApi "github.com/SohamGhugare/IHP/blockchain/doctorApi"
	"github.com/SohamGhugare/IHP/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var doctorDeployedAddr string
var doctorClient *ethclient.Client
var doctorConn *doctorApi.Api

func ConnectDoctorContract() {
	var err error
	// doctorClient, err = ethclient.Dial("http://127.0.0.1:7545")
	doctorClient, err = ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/f1thP4VKMvnfX9LaGREjrMogP4Vnplo0")

	if err != nil {
		log.Fatal("error creating doctorClient:", err)
	}

	pvtKey := os.Getenv("PVT_KEY")
	auth := utility.GetAccountAuth(doctorClient, pvtKey)
	deployedContractAddr, tx, instance, err := doctorApi.DeployApi(auth, doctorClient)
	if err != nil {
		log.Fatal("error deploying:", err)
	}
	_, _ = tx, instance
	doctorDeployedAddr = deployedContractAddr.Hex()

}

func GetDoctorConnection() {
	var err error
	doctorConn, err = doctorApi.NewApi(common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), doctorClient)
	if err != nil {
		log.Fatal("error while creating api: ", err.Error())
	}

}

func CreateDoctorProfile(license int, name string, email string) (int, common.Hash) {

	pin := rand.Intn(99999) + 10000
	pinInt := big.NewInt(int64(pin))
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

	tx, err := doctorConn.StoreProfile(auth, big.NewInt(int64(license)), name, email, pinInt)
	if err != nil {
		log.Fatal("error storing: ", err)
	}
	return int(pinInt.Int64()), tx.Hash()
}

func GetDoctorProfile(license int) (string, string, int) {
	callOpts := &bind.CallOpts{
		From: common.HexToAddress(os.Getenv("DOCTOR_ADDRESS")),
	}
	name, email, pin, err := doctorConn.GetProfile(callOpts, big.NewInt(int64(license)))
	if err != nil {
		log.Fatal("error fetching profile: ", err)
	}
	// log.Println("details:", name, email, pin)
	return name, email, int(pin.Int64())
}
