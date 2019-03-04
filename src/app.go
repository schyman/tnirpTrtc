package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	configuration := readConfig()

	// Connect to Postgres database
	startDatabase(configuration.Database)

	// Start the rest server
	startRestServer(configuration.Rest)
}

// Read 'config.json' file for database connection details and rest server details
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
