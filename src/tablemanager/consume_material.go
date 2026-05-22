package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Consume_material 原表为主人公—英雄角色.xlsx 的子表 英雄升级材料基础表
type Consume_material struct {
	Id       int32     `json:"id"`       //物品ID
	Rare_num [][]int32 `json:"rare_num"` //品质:0，金1，银2，铜
}

type Consume_materialMgr struct {
}

var (
	Consume_material_Model Consume_materialMgr
	consume_materialDic    map[int32]*Consume_material
	// Consume_material_All 主人公—英雄角色.xlsx (英雄升级材料基础表)
	Consume_material_All []*Consume_material
)

// Consume_material_Get 主人公—英雄角色.xlsx (英雄升级材料基础表)
func Consume_material_Get(Id int32) (*Consume_material, bool) {
	data, ok := consume_materialDic[Id]
	if !ok {
		PROTO_ERROR_ID = "主人公—英雄角色.xlsx\nconsume_material not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Consume_materialMgr) PrintArr() {
	vLen := len(Consume_material_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Consume_materialMgr) PrintArrOne(index int) {
	logger.Logf("==Consume_material==:%+v", Consume_material_All[index])
}

func (*Consume_materialMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := consume_materialDic[int32Key]
		if !ok {
			logger.LogWarn("Consume_material PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Consume_material==:%+v", vData)
	}
}

func (*Consume_materialMgr) Load(buffer []byte) bool {
	Consume_material_All = make([]*Consume_material, 0)
	err := json.Unmarshal(buffer, &Consume_material_All)
	if err != nil {
		logger.LogErr("Consume_material JsonFailed:", err)
		return false
	}
	vLen := len(Consume_material_All)
	consume_materialDic = make(map[int32]*Consume_material, vLen)
	for _, mem := range Consume_material_All {
		consume_materialDic[mem.Id] = mem
	}
	return true
}
