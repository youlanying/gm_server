package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// EquipmentCompound 原表为合成表.xlsx 的子表 Sheet1
type EquipmentCompound struct {
	Id                 int32     `json:"id"`                 //可合成装备ID
	Craftedcost        [][]int32 `json:"craftedcost"`        //消耗代币
	Conditionequipment [][]int32 `json:"conditionequipment"` //消耗材料
}

type EquipmentCompoundMgr struct {
}

var (
	EquipmentCompound_Model EquipmentCompoundMgr
	EquipmentCompoundDic    map[int32]*EquipmentCompound
	// EquipmentCompound_All 合成表.xlsx (Sheet1)
	EquipmentCompound_All []*EquipmentCompound
)

// EquipmentCompound_Get 合成表.xlsx (Sheet1)
func EquipmentCompound_Get(Id int32) (*EquipmentCompound, bool) {
	data, ok := EquipmentCompoundDic[Id]
	if !ok {
		PROTO_ERROR_ID = "合成表.xlsx\nEquipmentCompound not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *EquipmentCompoundMgr) PrintArr() {
	vLen := len(EquipmentCompound_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*EquipmentCompoundMgr) PrintArrOne(index int) {
	logger.Logf("==EquipmentCompound==:%+v", EquipmentCompound_All[index])
}

func (*EquipmentCompoundMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := EquipmentCompoundDic[int32Key]
		if !ok {
			logger.LogWarn("EquipmentCompound PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==EquipmentCompound==:%+v", vData)
	}
}

func (*EquipmentCompoundMgr) Load(buffer []byte) bool {
	EquipmentCompound_All = make([]*EquipmentCompound, 0)
	err := json.Unmarshal(buffer, &EquipmentCompound_All)
	if err != nil {
		logger.LogErr("EquipmentCompound JsonFailed:", err)
		return false
	}
	vLen := len(EquipmentCompound_All)
	EquipmentCompoundDic = make(map[int32]*EquipmentCompound, vLen)
	for _, mem := range EquipmentCompound_All {
		EquipmentCompoundDic[mem.Id] = mem
	}
	return true
}
