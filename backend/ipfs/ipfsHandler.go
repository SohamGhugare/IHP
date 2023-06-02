package ipfs

import (
	"fmt"
	"log"
	"os"

	"github.com/SohamGhugare/IHP/initializers"
)

type User struct {
	IHP  int    `json:"ihp"`
	Name string `json:"name"`
}

func StoreFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error while storing file: %v", err.Error())
	}
	defer file.Close()

	cid, err := initializers.Sh.Add(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	return cid
}
