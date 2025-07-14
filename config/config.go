package config

import (
	"log"

	"github.com/joho/godotenv"
)

// load environmental variables from .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}
}
