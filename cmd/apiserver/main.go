package main

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/apiserver"
	"github.com/MeguMan/buyer-exp-test/internal/app/emailsender"
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
	sender := emailsender.Sender{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		TLSPort:  os.Getenv("TLSPORT"),
	}

	if exists {
		if err := apiserver.Start(databaseURL, &sender); err != nil {
			log.Fatal(err)
		}
	}
}
