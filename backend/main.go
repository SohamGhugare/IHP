package main

import (
	"github.com/SohamGhugare/IHP/contract"
	"github.com/SohamGhugare/IHP/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	contract.ConnectContract()
	initializers.LoadEnvVars()
}

func main() {

	r := gin.Default()

	r.Run()
}
