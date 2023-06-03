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
	// contract.ConnectIHPContract()
	// contract.GetConnection()
	// contract.CreateProfile()
	// contract.GetProfile()

	// Doctor Contract Code
	contract.ConnectDoctorContract()
	contract.GetDoctorConnection()
	// contract.CreateDoctorProfile(123456, "Doctor123", "doctor@test.com")
	// contract.GetDoctorProfile(123456)

	// initializers.ConnectIPFS()
}

func main() {

	// cid := ipfs.StoreFile("dummy/in/test.png")
	// log.Println("CID:", cid)

	// Driver codef for fetching files
	// file, err := ipfs.GetStoredFile(cid)
	// if err != nil {
	// 	log.Fatalf("error while fetching file: %v", err.Error())
	// }
	// log.Printf("Fetched file: %v", file.Name())

	// Driver code for encryption
	// key := "0123456789ABCDEF"

	// //  Encrypt the CID
	// encryptedText, err := utility.EncryptString(key, cid)
	// if err != nil {
	// 	log.Fatalln("Encryption error:", err)
	// 	return
	// }

	// log.Println("Encrypted:", encryptedText)

	// // Decrypt the encrypted text
	// decryptedText, err := utility.DecryptString(key, encryptedText)
	// if err != nil {
	// 	log.Fatalln("Decryption error:", err)
	// 	return
	// }

	// log.Println("Decrypted:", decryptedText)

	r := gin.Default()

	r.POST("/api/v1/doctors/signup", controllers.CreateDoctor)

	r.Run()
}
