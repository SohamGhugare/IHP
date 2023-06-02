package initializers

import (
	"log"

	shell "github.com/ipfs/go-ipfs-api"
)

var Sh *shell.Shell

func ConnectIPFS() {
	log.Println("Connecting to IPFS Daemon...")
	Sh = shell.NewShell("localhost:5001")
	log.Println("Successfully established connection to IPFS Daemon...")
}
