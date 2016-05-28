package formats

import (
    "gopkg.in/yaml.v2"	
)

type YamlFormat struct {
	target interface{}
}

func (this YamlFormat) Using(bytes []byte) {

	err := yaml.Unmarshal(bytes, this.target)

	if err != nil {
		panic(err)
	}
	
}

func NewYamlFormat(target interface{}) YamlFormat {
	return YamlFormat{ target }
}
