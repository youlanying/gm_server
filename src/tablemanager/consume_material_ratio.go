package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Consume_material_ratio 原表为主人公—英雄角色.xlsx 的子表 英雄升级材料表
type Consume_material_ratio struct {
	Id         int32       `json:"id"`         //英雄等级
	Profession []int32     `json:"profession"` //职业
	Item_id    [][]int32   `json:"item_id"`    //升级材料
	Item_radio [][]float32 `json:"item_radio"` //升级系数
	Cost_gold  [][]int32   `json:"cost_gold"`  //升级金币
}

type Consume_material_ratioMgr struct {
}

var (
	Consume_material_ratio_Model Consume_material_ratioMgr
	consume_material_ratioDic    map[int32]*Consume_material_ratio
	// Consume_material_ratio_All 主人公—英雄角色.xlsx (英雄升级材料表)
	Consume_material_ratio_All []*Consume_material_ratio
)

// Consume_material_ratio_Get 主人公—英雄角色.xlsx (英雄升级材料表)
func Consume_material_ratio_Get(Id int32) (*Consume_material_ratio, bool) {
	data, ok := consume_material_ratioDic[Id]
	if !ok {
		PROTO_ERROR_ID = "主人公—英雄角色.xlsx\nconsume_material_ratio not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Consume_material_ratioMgr) PrintArr() {
	vLen := len(Consume_material_ratio_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Consume_material_ratioMgr) PrintArrOne(index int) {
	logger.Logf("==Consume_material_ratio==:%+v", Consume_material_ratio_All[index])
}

func (*Consume_material_ratioMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := consume_material_ratioDic[int32Key]
		if !ok {
			logger.LogWarn("Consume_material_ratio PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Consume_material_ratio==:%+v", vData)
	}
}

func (*Consume_material_ratioMgr) Load(buffer []byte) bool {
	Consume_material_ratio_All = make([]*Consume_material_ratio, 0)
	err := json.Unmarshal(buffer, &Consume_material_ratio_All)
	if err != nil {
		logger.LogErr("Consume_material_ratio JsonFailed:", err)
		return false
	}
	vLen := len(Consume_material_ratio_All)
	consume_material_ratioDic = make(map[int32]*Consume_material_ratio, vLen)
	for _, mem := range Consume_material_ratio_All {
		consume_material_ratioDic[mem.Id] = mem
	}
	return true
}
