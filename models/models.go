package models

import (
	"log"
	"we/device_admin/models/mock"
	"we/device_admin/models/usage"
	"we/device_admin/models/zk"
)

var clusterUsage *usage.ClusterUsage

func GetClusterUsage() *usage.ClusterUsage {
	return clusterUsage
}

type ModelConfig struct {
	// mock
	Mock          int    `json:"mock"`
	MockFilePath  string `json:"mock_file_path"`
	UsePlayground int    `json:"use_playground"`

	ZK zk.ZKConfig `json:"zookeeper"`

	RootName        string `json:"root_name"`
	MaxHistoryUnits int    `json:"max_history_units"`
	UpdateInterval  int    `json:"update_interval"`
}

func Initialize(cfg ModelConfig) {
	if cfg.Mock != 0 {
		var mocker *usage.UsageZKMocker
		var err error
		if len(cfg.MockFilePath) > 0 {
			mocker, err = usage.UsageZKMockerFromJsonFile(cfg.MockFilePath)
		} else if cfg.UsePlayground > 0 {
			mocker = usage.NewUsageMocker()
		} else {
			mocker = usage.DefaultUsageZKMocker()
		}

		if err != nil {
			log.Fatalf("Error in mocking ZK, %v", err)
		}

		zk.MockZookeeper(mocker)

		if cfg.UsePlayground > 0 {
			mock.InitializeDefault()
		}
	} else {
		zk.InitZookeeper(cfg.ZK)
	}

	clusterUsage = usage.NewClusterUsage(cfg.RootName, cfg.MaxHistoryUnits, cfg.UpdateInterval)
}
