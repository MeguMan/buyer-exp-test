package main

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/apiserver"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	databaseURL, exists := os.LookupEnv("DATABASE_URL")

	if exists {
		if err := apiserver.Start(databaseURL); err != nil {
			log.Fatal(err)
		}
	}
}
