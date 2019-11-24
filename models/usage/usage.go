package usage

import (
	"log"
	"sync"
	"time"
	"we/device_admin/models/zk"
)

const HistoryCapacityThresh = 200000

type SnapShot struct {
	Loc time.Time

	DeviceNum    int
	GPUDeviceNum int
	SensorNum    int
	EdgeNum      int

	CPU float32
	GPU float32
	Mem float32

	//device setting
	CPUCores  int
	GPUCores  int
	MemoryCap int
}

func SnapshotDiff(a, b SnapShot) (ret SnapShot) {
	ret.DeviceNum = b.DeviceNum - a.DeviceNum
	ret.GPUDeviceNum = b.GPUDeviceNum - a.GPUDeviceNum
	ret.SensorNum = b.SensorNum - a.SensorNum
	ret.EdgeNum = b.EdgeNum - a.EdgeNum
	ret.CPU = b.CPU - a.CPU
	ret.GPU = b.GPU - b.GPU
	ret.Mem = a.Mem - b.Mem
	ret.CPUCores = a.CPUCores - b.CPUCores
	ret.GPUCores = a.GPUCores - b.GPUCores
	ret.MemoryCap = a.MemoryCap - b.MemoryCap

	return
}

type DeviceStatus struct {
	DeviceTypeStr string `zk:"device_type"`
	DeviceTypeID  int

	DeviceRoleStr string `zk:"device_role"`
	DeviceRole    int

	CPU float32 `zk:"cpu"`
	GPU float32 `zk:"gpu"`
	Mem float32 `zk:"mem"`

	//device setting
	CPUCores  int `zk:"cpu_cores"`
	GPUCores  int `zk:"gpu_cores"`
	MemoryCap int `zk:"mem_cap"`

	// heart beat
	LastHeartBeat time.Time `zk:"heartbeat"`
}

func (d *DeviceStatus) HealthyScore() (ret int) {
	if time.Now().Sub(d.LastHeartBeat) > time.Second*20 {
		return -1
	}

	if d.GPU > 0.95 {
		return 1
	} else if d.GPU > 0.85 {
		return 2
	}

	if d.CPU > 0.95 {
		return 1
	} else if d.CPU > 0.85 {
		return 2
	}

	if d.Mem > 0.9 {
		return 1
	} else if d.Mem > 0.8 {
		return 2
	}

	return 3
}

type ClusterUsage struct {
	sync.RWMutex

	// status
	history []SnapShot
	devices map[string]DeviceStatus

	//setting
	RootName        string
	MaxHistoryUnits int
	UpdateInterval  int
}

func (c *ClusterUsage) GetDeiveStatusMap() map[string]DeviceStatus {
	c.RLock()
	defer c.RUnlock()

	return c.devices
}

func (c *ClusterUsage) GetStatusHistory() []SnapShot {
	c.RLock()
	defer c.RUnlock()

	return c.history
}

func (c *ClusterUsage) getDeviceList() (ret []string, reterr error) {
	ret, _, reterr = zk.ZKAgent.Children(c.RootName)
	return
}

func (c *ClusterUsage) loadDeviceStatusMap() (ret map[string]DeviceStatus, reterr error) {
	deviceNames, reterr := c.getDeviceList()
	if reterr != nil {
		return
	}

	ret = make(map[string]DeviceStatus)
	for _, name := range deviceNames {
		path := c.RootName + "/" + name
		var status DeviceStatus
		if err := zk.LoadAtributes(path, &status); err != nil {
			log.Printf("Error in accquiring status of Node %s, %v, node ignored", path, err)
			continue
		}

		status.DeviceTypeID = StrToDevType(status.DeviceTypeStr)
		status.DeviceRole = StrToRole(status.DeviceRoleStr)
		ret[name] = status
	}

	return
}

func (c *ClusterUsage) summarize(s map[string]DeviceStatus) (ret SnapShot) {
	ret.DeviceNum = len(s)
	ret.Loc = time.Now()

	for dname, ss := range s {
		log.Printf("%s | %v", dname, ss)
		if ss.DeviceTypeID == JetsonNano {
			ret.GPUDeviceNum++
		}

		if ss.DeviceRole == EdgeRole {
			ret.EdgeNum++
		} else if ss.DeviceRole == SensorRole {
			ret.SensorNum++
		}

		ret.CPU += ss.CPU
		ret.GPU += ss.GPU
		ret.Mem += ss.Mem

		ret.GPUCores += ss.GPUCores
		ret.CPUCores += ss.CPUCores
		ret.MemoryCap += ss.MemoryCap
	}

	ret.CPU /= float32(ret.DeviceNum)
	ret.GPU /= float32(ret.GPUDeviceNum)
	ret.Mem /= float32(ret.DeviceNum)

	return
}

func (c *ClusterUsage) update() (reterr error) {
	devices, reterr := c.loadDeviceStatusMap()
	if reterr != nil {
		return
	}

	snp := c.summarize(devices)
	nst := c.history
	if cap(nst) > HistoryCapacityThresh {
		log.Printf("Cleaning history slice...")

		nst = make([]SnapShot, len(nst))
		copy(nst, c.history)
	}

	nst = append(nst, snp)
	if len(nst) > c.MaxHistoryUnits {
		nst = nst[len(nst)-c.MaxHistoryUnits:]
	}

	c.Lock()

	c.history = nst
	c.devices = devices

	c.Unlock()

	return
}

func (c *ClusterUsage) start() {
	timer := time.NewTicker(time.Millisecond * time.Duration(c.UpdateInterval))

	go func() {
		for range timer.C {
			log.Printf("Start Updating....")
			if err := c.update(); err != nil {
				log.Printf("Error in updating cluter usage, %v", err)
			}
		}
	}()
}

func NewClusterUsage(rootName string, maxHistoryUnits, updateInterval int) (ret *ClusterUsage) {
	ret = &ClusterUsage{
		RootName:        rootName,
		MaxHistoryUnits: maxHistoryUnits,
		UpdateInterval:  updateInterval,
	}

	ret.start()
	return
}
