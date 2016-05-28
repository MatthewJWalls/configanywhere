package examples

import (
	"log"
	
	"github.com/MatthewJWalls/configanywhere"
	"github.com/MatthewJWalls/configanywhere/providers"	
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

	configanywhere.Json(&appconfig).FromString(`{"name":"testing"}`)
	
	// configuration from file

	configanywhere.Json(&appconfig).FromFile("test.txt")

	// configuration from Zookeeper

	configanywhere.Json(&appconfig).FromZookeeper(
		[]string{"127.0.0.1"},
		"/testing",
	)

}

func keyValueExamples() {

	// Example of loading key=value pairs from the environment

	appconfig := KeyValueAppConfig{}

	// configuration from environment variables

	configanywhere.KeyValue(&appconfig).FromEnvironment()

}

func ChooseFromManyExample() {

	// You configanywheren specify several providers, and configanywhere will
	// try each one in turn until it finds one that works.
	// err will be returned only if *all* of the providers failed.

	appconfig := JsonAppConfig{}	

	err := configanywhere.Json(&appconfig).Choose(
		providers.NewFileProvider("doesnotexist.txt"),
		providers.NewZookeeperProvider([]string{"127.0.0.1"}, "/doesnotexist"),
		providers.NewStringProvider(`{"name":"testing"}`),
	)

	if err != nil {
		log.Println("Error returned from Choose method")
	}

}
