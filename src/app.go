package main

import (
	"fmt"
)

func main() {
	configuration := readConfig()
	fmt.Println(configuration.Rest.Port)
	startRestServer(configuration.Rest)
}
