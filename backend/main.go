package main

import (
	"github.com/SohamGhugare/IHP/contract"
	"github.com/SohamGhugare/IHP/controllers"
	"github.com/SohamGhugare/IHP/initializers"
	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVars()

	// IHP Contract Code
	contract.ConnectIHPContract()
	contract.GetIHPConnection()

	// Doctor Contract Code
	contract.ConnectDoctorContract()
	contract.GetDoctorConnection()
	// contract.CreateDoctorProfile(123456, "Doctor123", "doctor@test.com")
	// contract.GetDoctorProfile(123456)

	initializers.ConnectIPFS()
}

func main() {

	r := gin.Default()

	r.POST("/api/v1/doctors/signup", controllers.CreateDoctor)
	r.POST("/api/v1/doctors/login", controllers.LoginDoctor)

	r.POST("/api/v1/users/signup", controllers.CreateUser)
	r.POST("/api/v1/users/login", controllers.GetUser)
	r.POST("/api/v1/users/add-report", controllers.AddRecord)

	r.POST("/api/v1/reports/upload", controllers.StoreFile)
	r.POST("/api/v1/reports/download", controllers.GetFile)

	r.Run()
}
