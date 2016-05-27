package main

import (
	ca "github.com/MatthewJWalls/configanywhere"
)

type AppConfig struct {
	Name string `json:"name"`
}

func main() {

	// Example of loading JSON configuration from various sources.

	appconfig := AppConfig{}
	
	// configuration from file

	ca.Json(&appconfig).FromFile("test.txt")

	// configuration from Zookeeper

	ca.Json(&appconfig).FromZookeeper(
		[]string{"127.0.0.1"},
		"/testing",
	)

}
