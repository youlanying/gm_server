package tablemanager

import (
	"encoding/json"
	"fmt"
)

type Collider struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Radius float32 `json:"radius"`
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Z      float32 `json:"z"`
}

type ColliderInfoMgr struct {
	Dic map[string]*Collider
	Arr []*Collider
}

func (this *ColliderInfoMgr) PrintArr()                     {}
func (this *ColliderInfoMgr) PrintArrOne(index int)         {}
func (this *ColliderInfoMgr) PrintMapByKey(key interface{}) {}

func (this *ColliderInfoMgr) Load(buffer []byte) bool {
	err := json.Unmarshal(buffer, &this.Arr)
	if err != nil {
		fmt.Println("JsonFailed", err)
		return false
	}
	vLen := len(this.Arr)
	this.Dic = make(map[string]*Collider, vLen)
	for _, mem := range this.Arr {
		temp := &Collider{}
		temp.Name = mem.Name
		this.Dic[mem.Name] = mem
	}
	return true
}

var ColliderInfoModel ColliderInfoMgr
