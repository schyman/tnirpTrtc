package main

import (
	"fmt"
)

func main() {
	configuration := readConfig()
	fmt.Println(configuration.Rest.Port)
	startDatabase(configuration.Database)
	startRestServer(configuration.Rest)
}
