package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	// Example: ensure required variables are set
	// requiredVars := []string{"DB_HOST", "DB_USER", "DB_PASS"}
	requiredVars := []string{"DATABASE_URL", "JWT_SECRET"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Printf("Warning: environment variable %s is not set\n", v)
		}
	}
}
