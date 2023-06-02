package main

import (
	"log"

	"github.com/SohamGhugare/IHP/initializers"
	"github.com/SohamGhugare/IHP/ipfs"
	"github.com/SohamGhugare/IHP/utility"
)

func init() {
	// contract.ConnectContract()
	initializers.LoadEnvVars()
	initializers.ConnectIPFS()
}

func main() {

	cid := ipfs.StoreFile("dummy/in/test.png")
	log.Println("CID:", cid)
	// file, err := ipfs.GetStoredFile(cid)
	// if err != nil {
	// 	log.Fatalf("error while fetching file: %v", err.Error())
	// }
	// log.Printf("Fetched file: %v", file.Name())

	key := "0123456789ABCDEF" // 16-byte key for AES-128

	// Encrypt the plaintext
	encryptedText, err := utility.EncryptString(key, cid)
	if err != nil {
		log.Fatalln("Encryption error:", err)
		return
	}

	log.Println("Encrypted:", encryptedText)

	// Decrypt the encrypted text
	decryptedText, err := utility.DecryptString(key, encryptedText)
	if err != nil {
		log.Fatalln("Decryption error:", err)
		return
	}

	log.Println("Decrypted:", decryptedText)

	// r := gin.Default()

	// r.Run()
}
