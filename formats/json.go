package formats

import (
	"encoding/json"
)

type JsonFormat struct {
	target interface{}
}

func (this JsonFormat) Using(bytes []byte) {
	
	err := json.Unmarshal(bytes, this.target)
	
	if err != nil {
		panic(err)
	}
	
}

func NewJsonFormat(target interface{}) JsonFormat {
	return JsonFormat{ target }
}
