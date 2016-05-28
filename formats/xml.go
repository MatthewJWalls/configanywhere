package formats

import (
	"encoding/xml"
)

type XMLFormat struct {
	target interface{}
}

func (this XMLFormat) Using(bytes []byte) {
	
	err := xml.Unmarshal(bytes, this.target)
	
	if err != nil {
		panic(err)
	}
	
}

func NewXMLFormat(target interface{}) XMLFormat {
	return XMLFormat{ target }
}
