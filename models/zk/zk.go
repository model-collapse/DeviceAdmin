package zk

import (
	"log"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type ZookeeperAgent interface {
	Get(path string) (ret []byte, stat *zk.Stat, reterr error)
	Children(path string) (ret []string, stat *zk.Stat, reterr error)
}

var ZKAgent ZookeeperAgent

type ZKConfig struct {
	Servers          []string `json:"servers"`
	Timeout          time.Duration
	TimeoutMilliSecs int `json:"timeout"`
}

func InitZookeeper(cfg ZKConfig) (reterr error) {
	cfg.Timeout = time.Duration(cfg.TimeoutMilliSecs) * time.Millisecond

	var evt <-chan zk.Event
	ZKAgent, evt, reterr = zk.Connect(cfg.Servers, cfg.Timeout)
	go func() {
		for e := range evt {
			log.Printf("[ERR ZK] %v", e.Err)
		}
	}()

	return
}

func MockZookeeper(agent ZookeeperAgent) {
	ZKAgent = agent
}
