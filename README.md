
# ConfigAnywhere

ConfigAnywhere is a convenience library that wraps up multiple ways of getting configuration from various
sources into a single package with an easy to use API.

#### Configuration Formats Supported

* Json
* key=value pairs

#### Configuration Providers Supported

* ZooKeeper
* Files
* Strings
* Environment Variables

## Example Usage

```go
    import (
        ca "github.com/MatthewJWalls/configanywhere"
    )
    
    type JsonAppConfig struct {
        Name string `json:"name"`
    }
        
    func jsonExamples() {
    
        // Example of loading JSON configuration from various sources.
    
        appconfig := JsonAppConfig{}
    
        // configuration from string
    
        ca.Json(&appconfig).FromString(`{"name":"testing"}`)
        
        // configuration from file
    
        ca.Json(&appconfig).FromFile("test.txt")
    
        // configuration from Zookeeper
    
        ca.Json(&appconfig).FromZookeeper(
            []string{"127.0.0.1"},
            "/testing",
        )
    
    }

    type KeyValueAppConfig struct {
	    Name string `key:"name"`
    }

    func keyValueExamples() {

        // Example of loading key=value pairs from the environment

        appconfig := KeyValueAppConfig{}

        // configuration from environment variables

        ca.KeyValue(&appconfig).FromEnvironment()

    }
```
