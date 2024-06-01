package getEnv

import (
	"log"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	// load .env file
	if err := godotenv.Load(filepath.Join(cwd, "./.env")); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
