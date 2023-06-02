package main

import (
	"github.com/SohamGhugare/IHP/contract"
	"github.com/SohamGhugare/IHP/initializers"
	"github.com/SohamGhugare/IHP/ipfs"
	"github.com/gin-gonic/gin"
)

func init() {
	contract.ConnectContract()
	initializers.LoadEnvVars()
	initializers.ConnectIPFS()
}

func main() {

	ipfs.StoreFile()

	r := gin.Default()

	r.Run()
}
