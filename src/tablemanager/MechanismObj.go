package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// MechanismObj 原表为机关实体表.xlsx 的子表 Sheet1
type MechanismObj struct {
	Id        int32 `json:"id"`        //机关实体ID
	Mechanism int32 `json:"Mechanism"` //机关模板表
}

type MechanismObjMgr struct {
}

var (
	MechanismObj_Model MechanismObjMgr
	MechanismObjDic    map[int32]*MechanismObj
	// MechanismObj_All 机关实体表.xlsx (Sheet1)
	MechanismObj_All []*MechanismObj
)

// MechanismObj_Get 机关实体表.xlsx (Sheet1)
func MechanismObj_Get(Id int32) (*MechanismObj, bool) {
	data, ok := MechanismObjDic[Id]
	if !ok {
		PROTO_ERROR_ID = "机关实体表.xlsx\nMechanismObj not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MechanismObjMgr) PrintArr() {
	vLen := len(MechanismObj_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MechanismObjMgr) PrintArrOne(index int) {
	logger.Logf("==MechanismObj==:%+v", MechanismObj_All[index])
}

func (*MechanismObjMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := MechanismObjDic[int32Key]
		if !ok {
			logger.LogWarn("MechanismObj PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==MechanismObj==:%+v", vData)
	}
}

func (*MechanismObjMgr) Load(buffer []byte) bool {
	MechanismObj_All = make([]*MechanismObj, 0)
	err := json.Unmarshal(buffer, &MechanismObj_All)
	if err != nil {
		logger.LogErr("MechanismObj JsonFailed:", err)
		return false
	}
	vLen := len(MechanismObj_All)
	MechanismObjDic = make(map[int32]*MechanismObj, vLen)
	for _, mem := range MechanismObj_All {
		MechanismObjDic[mem.Id] = mem
	}
	return true
}
