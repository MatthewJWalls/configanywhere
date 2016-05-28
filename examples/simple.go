package main

import (
	ca "github.com/MatthewJWalls/configanywhere"
)

type JsonAppConfig struct {
	Name string `json:"name"`
}

type KeyValueAppConfig struct {
	Name string `key:"name"`
}

func jsonExamples() {

	// Example of loading JSON configuration from various sources.

	appconfig := JsonAppConfig{}

	// configuration from string

	ca.Json(&appconfig).FromString(`{"name":"testing"}`)
	
	// configuration from file

	ca.Json(&appconfig).FromFile("test.txt")

	// configuration from Zookeeper

	ca.Json(&appconfig).FromZookeeper(
		[]string{"127.0.0.1"},
		"/testing",
	)

}

func keyValueExamples() {

	// Example of loading key=value pairs from the environment

	appconfig := KeyValueAppConfig{}

	// configuration from environment variables

	ca.KeyValue(&appconfig).FromEnvironment()


}
