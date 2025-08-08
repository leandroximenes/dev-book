package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	StringConnection string = ""
	Port             int    = 0
)

func Load() {
	var err error
	appEnv := os.Getenv("APP_ENV")

	// Only load .env file if we are in "dev" or "test"
	if appEnv != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Could not load .env file")
		} else {
			log.Println(".env file loaded")
		}
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 3000
	}

	StringConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
