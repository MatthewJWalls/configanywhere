package providers

import (
	"io/ioutil"
)

type FileProvider struct {
	filename string
}

func (this FileProvider) GetBytes() []byte {
	
	bytes, err := ioutil.ReadFile(this.filename)
	
	if err != nil {
		panic(err)
	}

	return bytes
	
}

func NewFileProvider(filename string) FileProvider {
	return FileProvider{ filename }
}
