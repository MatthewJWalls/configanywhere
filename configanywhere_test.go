package configanywhere

import (
	"testing"
)

type Example struct {
	Name string `json:"name"`
}

// Test using the dispatch API

func TestAPI(t *testing.T) {

	thing := Example{}

	Json(&thing).FromFile("test.txt")

	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected %s", )
	}
	
}

