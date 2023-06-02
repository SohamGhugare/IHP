package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	log.SetPrefix("[IHP] ")
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load env file...")
	}
}
