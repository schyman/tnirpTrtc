package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	configuration := readConfig()
	startDatabase(configuration.Database)
	startRestServer(configuration.Rest)
}

func readConfig() Configuration {
	input, readError := ioutil.ReadFile("config.json")
	if readError != nil {
		log.Fatal(readError)
	}
	var data Configuration
	err := json.Unmarshal(input, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
