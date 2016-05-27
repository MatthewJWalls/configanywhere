package configanywhere

import (
	"testing"
)

type Example struct {
	Name string `json:"name"`
}

// Test using the dispatch API

func TestFiles(t *testing.T) {

	thing := Example{}

	Json(&thing).FromFile("test.txt")

	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected %s", )
	}
	
}

func TestZookeeper(t *testing.T) {

	thing := Example{}

	Json(&thing).FromZookeeper([]string{"127.0.0.1"}, "/testing")

	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected %s", )
	}
	
}

