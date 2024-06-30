package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	env := os.Getenv("GO_ENV")
	if env == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			log.Println(".env file loaded successfully")
		}
	}
	log.Printf("Environment: %s", env)
	// log.Printf("Sample Env Var: %s", os.Getenv("DB_URL"))
}
