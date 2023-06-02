package main

import (
	"log"
	"os"

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
	outputPath := "dummy/out/test.png"
	err := initializers.Sh.Get(cid, outputPath)
	if err != nil {
		log.Fatalf("error fetching file from ipfs daemon: %v", err.Error())
	}
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		log.Println("File not found!")
	} else {
		log.Println("File fetched and saved successfully.")
	}

	r := gin.Default()

	r.Run()
}
