package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Equipinfo 原表为装备表.xlsx 的子表 装备表
type Equipinfo struct {
	Id          int32     `json:"id"`          //装备ID
	Column      int32     `json:"column"`      //装备栏位
	Init_attr   [][]int32 `json:"init_attr"`   //初始属性
	Growth_attr [][]int32 `json:"growth_attr"` //成长属性
	Skillid     int32     `json:"skillid"`     //装备技能
}

type EquipinfoMgr struct {
}

var (
	Equipinfo_Model EquipinfoMgr
	equipinfoDic    map[int32]*Equipinfo
	// Equipinfo_All 装备表.xlsx (装备表)
	Equipinfo_All []*Equipinfo
)

// Equipinfo_Get 装备表.xlsx (装备表)
func Equipinfo_Get(Id int32) (*Equipinfo, bool) {
	data, ok := equipinfoDic[Id]
	if !ok {
		PROTO_ERROR_ID = "装备表.xlsx\nequipinfo not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *EquipinfoMgr) PrintArr() {
	vLen := len(Equipinfo_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*EquipinfoMgr) PrintArrOne(index int) {
	logger.Logf("==Equipinfo==:%+v", Equipinfo_All[index])
}

func (*EquipinfoMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := equipinfoDic[int32Key]
		if !ok {
			logger.LogWarn("Equipinfo PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Equipinfo==:%+v", vData)
	}
}

func (*EquipinfoMgr) Load(buffer []byte) bool {
	Equipinfo_All = make([]*Equipinfo, 0)
	err := json.Unmarshal(buffer, &Equipinfo_All)
	if err != nil {
		logger.LogErr("Equipinfo JsonFailed:", err)
		return false
	}
	vLen := len(Equipinfo_All)
	equipinfoDic = make(map[int32]*Equipinfo, vLen)
	for _, mem := range Equipinfo_All {
		equipinfoDic[mem.Id] = mem
	}
	return true
}
