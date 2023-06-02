package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/SohamGhugare/IHP/initializers"
)

type User struct {
	IHP  int    `json:"ihp"`
	Name string `json:"name"`
}

func StoreFile() {
	user := &User{
		IHP:  12345,
		Name: "Soham Ghugare",
	}

	userBin, _ := json.Marshal(user)
	reader := bytes.NewReader(userBin)

	cid, err := initializers.Sh.Add(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)
}
