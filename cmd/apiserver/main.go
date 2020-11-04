package main

import (
	"encoding/json"
	"github.com/MeguMan/buyer-exp-test/internal/app/apiserver"
	"github.com/MeguMan/buyer-exp-test/internal/app/emailsender"
	"log"
	"os"
)

func main() {
	dbConfig := apiserver.NewConfig()
	emailConfig := emailsender.NewConfig()
	ParseJsonToConfig(dbConfig, "db.json")
	ParseJsonToConfig(emailConfig, "email.json")

	if err := apiserver.Start(dbConfig, emailConfig); err != nil {
		log.Fatal(err)
	}
}

func ParseJsonToConfig(i interface{}, configName string) interface{} {
	configFile, err := os.Open("configs/" + configName)
	if err != nil {
		log.Print(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&i); err != nil {
		log.Print(err.Error())
	}

	return i
}
