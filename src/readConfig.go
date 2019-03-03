package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
