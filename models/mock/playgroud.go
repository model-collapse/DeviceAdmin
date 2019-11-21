package mock

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"we/device_admin/models/usage"
	"we/device_admin/models/zk"
)

const MockRoot = "root"
const Interval = 3

type Device struct {
	usage.DeviceStatus
	Name      string
	ActiveCnt int
}

func (d *Device) updateStateRandom() {
	d.LastHeartBeat = time.Now()
	d.CPU = rand.Float32()
	d.Mem = rand.Float32()
	if d.GPUCores > 0 {
		d.GPU = rand.Float32()
	}

	if rand.Float32() < 0.01 {
		d.ActiveCnt = rand.Int() % 30
	}
}

func (d *Device) writeZKMock(writeBasic bool) {
	mocker, suc := zk.ZKAgent.(*usage.UsageZKMocker)
	if !suc {
		log.Fatalf("Fail to initialize playground because the ZK agent is not for mocking")
		return
	}

	log.Printf("Writing stat of device %s", d.Name)
	prefix := fmt.Sprintf("%s/%s", MockRoot, d.Name)
	mocker.Set(prefix+"/cpu", fmt.Sprintf("%f", d.CPU))
	mocker.Set(prefix+"/gpu", fmt.Sprintf("%f", d.GPU))
	mocker.Set(prefix+"/mem", fmt.Sprintf("%f", d.Mem))
	mocker.Set(prefix+"/heatbeat", fmt.Sprintf("%d", d.LastHeartBeat.Unix()))

	if writeBasic {
		mocker.Set(prefix+"/device_type", d.DeviceTypeStr)
		mocker.Set(prefix+"/device_role", d.DeviceRoleStr)
		mocker.Set(prefix+"/cpu_cores", fmt.Sprintf("%d", d.CPUCores))
		mocker.Set(prefix+"/gpu_cores", fmt.Sprintf("%d", d.GPUCores))
		mocker.Set(prefix+"/mem_cap", fmt.Sprintf("%d", d.MemoryCap))
	}
}

func (d *Device) start() {
	ticker := time.NewTicker(time.Second * Interval)
	d.updateStateRandom()
	d.writeZKMock(true)
	go func() {
		for range ticker.C {
			if d.ActiveCnt <= 0 {
				d.updateStateRandom()
				d.writeZKMock(false)
			} else {
				d.ActiveCnt--
			}
		}
	}()
}

func randomName() string {
	return fmt.Sprintf("%x", rand.Int31())
}

var deviceTempl = map[string]usage.DeviceStatus{
	"jetson_nano": usage.DeviceStatus{
		DeviceTypeStr: "jetson_nano",
		DeviceRoleStr: "edge",
		CPUCores:      4,
		GPUCores:      128,
		MemoryCap:     4096,
	},
	"raspberry_pi3": usage.DeviceStatus{
		DeviceTypeStr: "pi3",
		DeviceRoleStr: "edge",
		CPUCores:      4,
		GPUCores:      0,
		MemoryCap:     1024,
	},
	"raspberry_pi3_sensor": usage.DeviceStatus{
		DeviceTypeStr: "pi3",
		DeviceRoleStr: "sensor",
		CPUCores:      4,
		GPUCores:      0,
		MemoryCap:     4096,
	},
	"raspberry_pi4": usage.DeviceStatus{
		DeviceTypeStr: "pi4",
		DeviceRoleStr: "edge",
		CPUCores:      4,
		GPUCores:      0,
		MemoryCap:     2048,
	},
	"raspberry_zero": usage.DeviceStatus{
		DeviceTypeStr: "pi0",
		DeviceRoleStr: "sensor",
		CPUCores:      1,
		GPUCores:      0,
		MemoryCap:     512,
	},
	"esp32cam": usage.DeviceStatus{
		DeviceTypeStr: "esp32cam",
		DeviceRoleStr: "sensor",
		CPUCores:      2,
		GPUCores:      0,
		MemoryCap:     16,
	},
	"sparkfun": usage.DeviceStatus{
		DeviceTypeStr: "sparkfun_ble",
		DeviceRoleStr: "sensor",
		CPUCores:      1,
		GPUCores:      0,
		MemoryCap:     2,
	},
}

func Initialize(devCnt map[string]int) {
	mocker, suc := zk.ZKAgent.(*usage.UsageZKMocker)
	if !suc {
		log.Fatalf("Fail to initialize playground because the ZK agent is not for mocking")
	}

	for name, cnt := range devCnt {
		for i := 0; i < cnt; i++ {
			dev := Device{
				DeviceStatus: deviceTempl[name],
				Name:         randomName(),
				ActiveCnt:    0,
			}

			dev.start()
			mocker.AddChild(MockRoot, dev.Name)
		}
	}
}

func InitializeDefault() {

	Initialize(map[string]int{
		"jetson_nano":          2,
		"raspberry_pi3":        1,
		"raspberry_pi3_sensor": 1,
		"raspberry_pi4":        1,
		"raspberry_zero":       1,
		"esp32cam":             1,
		"sparkfun":             1,
	})
}
