package configanywhere

import (
	"os"
	"testing"
)

type JsonConf struct {
	Name string `json:"name"`
}

type KeyValueConf struct {
	Name string `key:"name"`
	Age int `key:"age"`
	Alive bool `key:"alive"`
}

func TestFiles(t *testing.T) {

	thing := JsonConf{}

	Json(&thing).FromFile("test.txt")

	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected testing")
	}
	
}

func TestZookeeper(t *testing.T) {

	thing := JsonConf{}

	Json(&thing).FromZookeeper([]string{"127.0.0.1"}, "/testing")

	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected testing")
	}
	
}

func TestString(t *testing.T) {

	thing := JsonConf{}
	
	Json(&thing).FromString(`{"name":"testing"}`)
	
	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected testing")
	}

}

func TestEnvironment(t *testing.T) {

	os.Setenv("name", "testing")
	os.Setenv("age", "10")
	os.Setenv("alive", "true")
	
	thing := KeyValueConf{}
	
	KeyValue(&thing).FromEnvironment()
	
	if thing.Name != "testing" {
		t.Errorf("Wrong name, expected testing")
	}

	if thing.Age != 10 {
		t.Errorf("Wrong age, expected 10")
	}

	if ! thing.Alive {
		t.Errorf("Wrong Alive, expected true")
	}	

}

