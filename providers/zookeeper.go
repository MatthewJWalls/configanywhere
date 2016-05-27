package providers

import (
	"time"
	
	"github.com/samuel/go-zookeeper/zk"
)

type ZookeeperProvider struct {
	servers []string
	nodePath string
}

func (this ZookeeperProvider) GetBytes() []byte {
	
	conn, _, err := zk.Connect(this.servers, time.Second)

	if err != nil {
		panic(err)
	}
	
	defer conn.Close()
	
	if exists, _, _ := conn.Exists(this.nodePath); ! exists {
		panic("Zookeeper node does not exist")
	}

	bytes, _, err := conn.Get(this.nodePath)

	if err != nil {
		panic(err)
	}

	return bytes
	
}

func NewZookeeperProvider(servers []string, nodePath string) ZookeeperProvider {
	return ZookeeperProvider { servers, nodePath  }
}
