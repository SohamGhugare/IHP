package contract

import (
	"log"
	"os"

	"github.com/SohamGhugare/IHP/blockchain/api"
	"github.com/SohamGhugare/IHP/utility"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectContract() {
	// contractAddressHex := os.Getenv("CONTRACT_ADDRESS")
	// contractABIStr := `[{"constant":false,"inputs":[{"name":"uhpId","type":"uint256"},{"name":"uri","type":"string"}],"name":"storeProfile","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"uhpId","type":"uint256"}],"name":"getProfile","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`

	// client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/f1thP4VKMvnfX9LaGREjrMogP4Vnplo0")
	client, err := ethclient.Dial("http://127.0.0.1:7545")

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
	log.Println("hash: ", deployedContractAddr.Hex())

}
