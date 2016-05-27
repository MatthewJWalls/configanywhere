package configanywhere

import (
	"testing"
)

type Example struct {
	Name string `json:"name"`
}

func TestFiles(t *testing.T) {

	thing := Example{}

	Json(&thing).FromFile("test.txt")

	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected testing")
	}
	
}

func TestZookeeper(t *testing.T) {

	thing := Example{}

	Json(&thing).FromZookeeper([]string{"127.0.0.1"}, "/testing")

	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected testing")
	}
	
}

func TestString(t *testing.T) {

	thing := Example{}
	
	Json(&thing).FromString(`{"name":"testing"}`)
	
	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected testing")
	}

}
