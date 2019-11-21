package usage

import (
	"math"
	"sync"
	"testing"
	"time"
	"we/device_admin/models/zk"
)

func roundEql(a, b float32) bool {
	d := float64(a - b)
	return math.Abs(d) < 0.0001
}

func initMock() {
	zk.MockZookeeper(NewUsageZKMocker())
}

var initOnce sync.Once

func TestClusterUsageAll(t *testing.T) {
	initOnce.Do(initMock)
	uu := NewClusterUsage("root", 100, 1000)

	time.Sleep(time.Second * 2)

	mp := uu.GetDeiveStatusMap()
	if len(mp) != 2 {
		t.Fatalf("incorrect map size")
	}

	if stat, suc := mp["abcdefg"]; !suc {
		t.Fatalf("node is not in")
	} else {
		if stat.DeviceTypeID != RaspberryPI_Zero {
			t.Errorf("incorrect device type")
		}

		if stat.GPU > 0 {
			t.Errorf("pi does not have a cuda GPU")
		}

		if !roundEql(stat.CPU, 0.5) {
			t.Errorf("incorrect cpu for first device")
		}
	}

	hist := uu.GetStatusHistory()
	if len(hist) == 0 {
		t.Fatalf("empty history")
	}

	snap := hist[0]
	if !roundEql(snap.CPU, 0.7) {
		t.Errorf("invalid cpu snap: %f", snap.CPU)
	}

	if !roundEql(snap.GPU, 0.9) {
		t.Errorf("invalid gpu snap, %f", snap.GPU)
	}

	if snap.GPUDeviceNum != 1 {
		t.Errorf("invalid device num")
	}
}

func TestClusterUsageHistory(t *testing.T) {
	initOnce.Do(initMock)

	uu := &ClusterUsage{
		RootName:        "root",
		MaxHistoryUnits: 100,
		UpdateInterval:  100,
	}

	for i := 0; i < HistoryCapacityThresh+1234; i++ {
		uu.update()
	}

	if len(uu.history) != 100 {
		t.Errorf("invalid len history")
	}

	if cap(uu.history) > HistoryCapacityThresh {
		t.Errorf("invalid cap history")
	}
}
