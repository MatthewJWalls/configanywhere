package providers

import (
	"time"

    "golang.org/x/net/context"
	"github.com/coreos/etcd/client"	
)

type EtcdProvider struct {
	servers []string
	nodePath string
}

func (this EtcdProvider) GetBytes() ([]byte, error) {

	// create a config
	
    cfg := client.Config{
		Endpoints:               this.servers,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	// start the etcd client & get the keys API
	
    c, err := client.New(cfg)
	
	if err != nil {
		return nil, err
	}

	keys := client.NewKeysAPI(c)

	// get the value from the path

	resp, err := keys.Get(context.Background(), this.nodePath, nil)
	
	if err != nil {
		return nil, err
	}

	return []byte(resp.Node.Value), nil
	
}

func NewEtcdProvider(servers []string, nodePath string) EtcdProvider {
	return EtcdProvider { servers, nodePath  }
}
