package configanywhere

import (
	"github.com/MatthewJWalls/configanywhere/formats"
	"github.com/MatthewJWalls/configanywhere/providers"	
)

// Dispatch and the convenience methods in this file
// allow us to support an API in the short format:
//
//     configanywhere.Json(xs).FromFile(ys)
//
// which due to Go's lack of support for Traits means
// we have to do this somewhat clumsily as below.

type Format interface {
	Using([]byte)
}

type Dispatch struct {
	F Format
}

// Convenience methods for providers

func (this Dispatch) FromFile(filename string) {
	this.F.Using(providers.NewFileProvider(filename).GetBytes())
}

func (this Dispatch) FromZookeeper(servers []string, nodePath string) {
	this.F.Using(providers.NewZookeeperProvider(servers, nodePath).GetBytes())
}

func (this Dispatch) FromString(text string) {
	this.F.Using(providers.NewStringProvider(text).GetBytes())
}

func (this Dispatch) FromEnvironment() {
	this.F.Using(providers.NewEnvironmentProvider().GetBytes())
}

// convenience methods for formats

func Json(target interface{}) Dispatch {
	return Dispatch{ formats.NewJsonFormat(target) }
}

func KeyValue(target interface{}) Dispatch {
	return Dispatch{ formats.NewKeyValueFormat(target) }
}
