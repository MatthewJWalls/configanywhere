package configanywhere

import (
	"github.com/MatthewJWalls/configanywhere/formats"
	"github.com/MatthewJWalls/configanywhere/providers"	
)

type Format interface {
	Using([]byte)
}

type Dispatch struct {
	F Format
}

func (this Dispatch) FromFile(filename string) {
	this.F.Using(providers.NewFileProvider(filename).GetBytes())
}

func (this Dispatch) FromZookeeper(servers []string, nodePath string) {
	this.F.Using(providers.NewZookeeperProvider(servers, nodePath).GetBytes())
}

func Json(target interface{}) Dispatch {
	return Dispatch{ formats.NewJsonFormat(target) }
}

