package tablemanager

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ModelInfo struct {
	Name string  `json:"Name"`
	Size float64 `json:"Size"`
}

type ModelInfoMgr struct {
	Dic map[string]*ModelInfo
	Arr []*ModelInfo
}

func (this *ModelInfoMgr) Load(vbuffer *[]byte) {
	buffer := *vbuffer
	err := json.Unmarshal(buffer, &this.Arr)
	if err != nil {
		fmt.Println("JsonFailed", err)
		return
	}
	vLen := len(this.Arr)
	this.Dic = make(map[string]*ModelInfo, vLen)
	for _, mem := range this.Arr {
		this.Dic[strings.ToLower(mem.Name)] = mem
	}
}

var ModelInfoModel ModelInfoMgr
