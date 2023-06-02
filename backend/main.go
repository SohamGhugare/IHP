package main

import (
	"log"

	"github.com/SohamGhugare/IHP/initializers"
	"github.com/SohamGhugare/IHP/ipfs"
	"github.com/gin-gonic/gin"
)

func init() {
	// contract.ConnectContract()
	initializers.LoadEnvVars()
	initializers.ConnectIPFS()
}

func main() {

	cid := ipfs.StoreFile("dummy/in/test.png")
	file, err := ipfs.GetStoredFile(cid)
	if err != nil {
		log.Fatalf("error while fetching file: %v", err.Error())
	}
	log.Printf("Fetched file: %v", file.Name())

	r := gin.Default()

	r.Run()
}
