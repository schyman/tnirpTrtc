package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	Database DatabaseConfig `json:"database"`
	Rest     RestConfig     `json:"rest"`
}

type DatabaseConfig struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type RestConfig struct {
	Port int `json:"port"`
}

func main() {
	configuration := readConfig()
	fmt.Println(configuration.Rest.Port)
	startRestServer()
}

func readConfig() Configuration {
	input, _ := ioutil.ReadFile("config.json")
	var data Configuration
	json.Unmarshal(input, &data)
	return data
}

func startRestServer() {
	fmt.Println("Started Rest server")
}
