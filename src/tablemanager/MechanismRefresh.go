package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// MechanismRefresh 原表为机关刷新表.xlsx 的子表 机关刷新表
type MechanismRefresh struct {
	Id               int32   `json:"id"`               //机关刷新id
	MechanismObjects []int32 `json:"MechanismObjects"` //机关实体
}

type MechanismRefreshMgr struct {
}

var (
	MechanismRefresh_Model MechanismRefreshMgr
	MechanismRefreshDic    map[int32]*MechanismRefresh
	// MechanismRefresh_All 机关刷新表.xlsx (机关刷新表)
	MechanismRefresh_All []*MechanismRefresh
)

// MechanismRefresh_Get 机关刷新表.xlsx (机关刷新表)
func MechanismRefresh_Get(Id int32) (*MechanismRefresh, bool) {
	data, ok := MechanismRefreshDic[Id]
	if !ok {
		PROTO_ERROR_ID = "机关刷新表.xlsx\nMechanismRefresh not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MechanismRefreshMgr) PrintArr() {
	vLen := len(MechanismRefresh_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MechanismRefreshMgr) PrintArrOne(index int) {
	logger.Logf("==MechanismRefresh==:%+v", MechanismRefresh_All[index])
}

func (*MechanismRefreshMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := MechanismRefreshDic[int32Key]
		if !ok {
			logger.LogWarn("MechanismRefresh PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==MechanismRefresh==:%+v", vData)
	}
}

func (*MechanismRefreshMgr) Load(buffer []byte) bool {
	MechanismRefresh_All = make([]*MechanismRefresh, 0)
	err := json.Unmarshal(buffer, &MechanismRefresh_All)
	if err != nil {
		logger.LogErr("MechanismRefresh JsonFailed:", err)
		return false
	}
	vLen := len(MechanismRefresh_All)
	MechanismRefreshDic = make(map[int32]*MechanismRefresh, vLen)
	for _, mem := range MechanismRefresh_All {
		MechanismRefreshDic[mem.Id] = mem
	}
	return true
}
