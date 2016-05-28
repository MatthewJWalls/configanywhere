package providers

import (
	"io/ioutil"
)

type FileProvider struct {
	filename string
}

func (this FileProvider) GetBytes() ([]byte, error) {
	
	bytes, err := ioutil.ReadFile(this.filename)
	return bytes, err

}

func NewFileProvider(filename string) FileProvider {
	return FileProvider{ filename }
}
