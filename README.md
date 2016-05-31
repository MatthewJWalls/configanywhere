
# ConfigAnywhere

ConfigAnywhere is a convenience library that wraps up multiple ways of getting configuration from various
sources into a single package with an easy to use API.

#### Configuration Formats Supported

* Json
* key=value pairs
* XML
* Yaml
    
#### Configuration Providers Supported

* ZooKeeper
* Etcd
* Files
* Strings
* Environment Variables

## Example Usage

#### Simple API usage
    
```go
    import (
        "github.com/MatthewJWalls/configanywhere"
    )
    
    type JsonAppConfig struct {
        Name string `json:"name"`
    }
        
    func jsonExamples() {
    
        // Example of loading JSON configuration from various sources.
    
        appconfig := JsonAppConfig{}
    
        // configuration from string
    
        configanywhere.Json(&appconfig).FromString(`{"name":"testing"}`)
        
        // configuration from file
    
        configanywhere.Json(&appconfig).FromFile("test.txt")
    
        // configuration from Zookeeper
    
        configanywhere.Json(&appconfig).FromZookeeper(
            []string{"127.0.0.1"},
            "/testing",
        )
    
    }

```

#### Using Choose() to try many providers

```go
    import (
        "github.com/MatthewJWalls/configanywhere"
        "github.com/MatthewJWalls/configanywhere/providers"    
    )
    
    type JsonAppConfig struct {
        Name string `json:"name"`
    }
        
    func ChooseFromManyExample() {
    
        // You can specify several providers, and Choose will
        // try each one in turn until it finds one that works.
        // err will be returned only if *all* of the providers failed.
    
        appconfig := JsonAppConfig{}    
    
        err := configanywhere.Json(&appconfig).Choose(
            providers.NewFileProvider("doesnotexist.txt"),
            providers.NewZookeeperProvider([]string{"127.0.0.1"}, "/doesnotexist"),
            providers.NewStringProvider(`{"name":"testing"}`),
        )
    
        if err != nil {
            t.Errorf("Error returned from Choose method")
        }
    
    }

```
