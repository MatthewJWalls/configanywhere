package providers

import (
	"time"
	"errors"
	
	"github.com/samuel/go-zookeeper/zk"
)

type ZookeeperProvider struct {
	servers []string
	nodePath string
}

func (this ZookeeperProvider) GetBytes() ([]byte, error) {
	
	conn, _, err := zk.Connect(this.servers, time.Second)

	if err != nil {
		return nil, err
	}
	
	defer conn.Close()
	
	if exists, _, _ := conn.Exists(this.nodePath); ! exists {
		return nil, errors.New("ZooKeeper node does not exist")
	}

	bytes, _, err := conn.Get(this.nodePath)

	if err != nil {
		return nil, err
	}

	return bytes, nil
	
}

func NewZookeeperProvider(servers []string, nodePath string) ZookeeperProvider {
	return ZookeeperProvider { servers, nodePath  }
}
