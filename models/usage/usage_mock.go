package usage

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/samuel/go-zookeeper/zk"
)

type UsageZKMocker struct {
	sync.RWMutex
	children map[string][]string
	content  map[string]string
}

func (u *UsageZKMocker) Set(path, val string) {
	u.Lock()
	defer u.Unlock()

	u.content[path] = val
}

func (u *UsageZKMocker) AddChild(path, name string) {
	u.Lock()
	defer u.Unlock()

	u.children[path] = append(u.children[path], name)
}

func (u *UsageZKMocker) RemoveChild(path, name string) {
	u.Lock()
	defer u.Unlock()

	var id int
	l := u.children[path]
	for id = 0; id < len(l); id++ {
		if l[id] == name {
			break
		}
	}

	nl := make([]string, len(l)-1)
	if id > 0 {
		copy(nl[0:id], l[0:id])
	}

	copy(nl[id:len(nl)], l[id+1:len(l)])
}

func (u *UsageZKMocker) Get(path string) (ret []byte, stat *zk.Stat, reterr error) {
	u.RLock()
	defer u.RUnlock()

	return []byte(u.content[path]), nil, nil
}

func (u *UsageZKMocker) Children(path string) (ret []string, stat *zk.Stat, reterr error) {
	u.RLock()
	defer u.RUnlock()

	return u.children[path], nil, nil
}

func DefaultUsageZKMocker() (ret *UsageZKMocker) {
	return &UsageZKMocker{
		children: map[string][]string{
			"root": []string{"abcdefg", "hijklmn"},
		},
		content: map[string]string{
			"root/abcdefg/cpu":         "0.5",
			"root/abcdefg/gpu":         "0",
			"root/abcdefg/mem":         "0.1",
			"root/abcdefg/device_type": "pi0",
			"root/abcdefg/device_role": "sensor",
			"root/abcdefg/cpu_cores":   "1",
			"root/abcdefg/gpu_cores":   "0",
			"root/abcdefg/mem_cap":     "512",
			"root/hijklmn/cpu":         "0.9",
			"root/hijklmn/gpu":         "0.9",
			"root/hijklmn/mem":         "0.34",
			"root/hijklmn/device_type": "jetson_nano",
			"root/hijklmn/device_role": "edge",
			"root/hijklmn/cpu_cores":   "4",
			"root/hijklmn/gpu_cores":   "128",
			"root/hijklmn/mem_cap":     "4096",
		},
	}
}

func NewUsageMocker() *UsageZKMocker {
	return &UsageZKMocker{
		children: make(map[string][]string),
		content:  make(map[string]string),
	}
}

func UsageZKMockerFromJsonFile(s string) (ret *UsageZKMocker, reterr error) {
	tmp := struct {
		Children map[string][]string `json:"children"`
		Content  map[string]string   `json:"content"`
	}{}

	data, reterr := ioutil.ReadFile(s)
	if reterr != nil {
		return
	}

	if reterr = json.Unmarshal(data, &tmp); reterr != nil {
		return
	}

	ret = &UsageZKMocker{
		children: tmp.Children,
		content:  tmp.Content,
	}

	return
}
