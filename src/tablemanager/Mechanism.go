package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Mechanism 原表为机关模板表.xlsx 的子表 Sheet1
type Mechanism struct {
	Id         int32 `json:"id"`         //机关模板ID
	TimesLimit int32 `json:"TimesLimit"` //次数限制(0为不限)
	DropId     int32 `json:"DropId"`     //掉落ID
}

type MechanismMgr struct {
}

var (
	Mechanism_Model MechanismMgr
	MechanismDic    map[int32]*Mechanism
	// Mechanism_All 机关模板表.xlsx (Sheet1)
	Mechanism_All []*Mechanism
)

// Mechanism_Get 机关模板表.xlsx (Sheet1)
func Mechanism_Get(Id int32) (*Mechanism, bool) {
	data, ok := MechanismDic[Id]
	if !ok {
		PROTO_ERROR_ID = "机关模板表.xlsx\nMechanism not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MechanismMgr) PrintArr() {
	vLen := len(Mechanism_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MechanismMgr) PrintArrOne(index int) {
	logger.Logf("==Mechanism==:%+v", Mechanism_All[index])
}

func (*MechanismMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := MechanismDic[int32Key]
		if !ok {
			logger.LogWarn("Mechanism PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Mechanism==:%+v", vData)
	}
}

func (*MechanismMgr) Load(buffer []byte) bool {
	Mechanism_All = make([]*Mechanism, 0)
	err := json.Unmarshal(buffer, &Mechanism_All)
	if err != nil {
		logger.LogErr("Mechanism JsonFailed:", err)
		return false
	}
	vLen := len(Mechanism_All)
	MechanismDic = make(map[int32]*Mechanism, vLen)
	for _, mem := range Mechanism_All {
		MechanismDic[mem.Id] = mem
	}
	return true
}
