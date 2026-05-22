package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// AssociatedSceneBoxData 原表为关卡表-关联场景盒子数据表.xlsx 的子表 数据表
type AssociatedSceneBoxData struct {
	Id             int32 `json:"id"`             //数据表ID
	EventType      int32 `json:"EventType"`      //事件类型
	IntEventParams int32 `json:"IntEventParams"` //事件参数
}

type AssociatedSceneBoxDataMgr struct {
}

var (
	AssociatedSceneBoxData_Model AssociatedSceneBoxDataMgr
	AssociatedSceneBoxDataDic    map[int32]*AssociatedSceneBoxData
	// AssociatedSceneBoxData_All 关卡表-关联场景盒子数据表.xlsx (数据表)
	AssociatedSceneBoxData_All []*AssociatedSceneBoxData
)

// AssociatedSceneBoxData_Get 关卡表-关联场景盒子数据表.xlsx (数据表)
func AssociatedSceneBoxData_Get(Id int32) (*AssociatedSceneBoxData, bool) {
	data, ok := AssociatedSceneBoxDataDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡表-关联场景盒子数据表.xlsx\nAssociatedSceneBoxData not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *AssociatedSceneBoxDataMgr) PrintArr() {
	vLen := len(AssociatedSceneBoxData_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*AssociatedSceneBoxDataMgr) PrintArrOne(index int) {
	logger.Logf("==AssociatedSceneBoxData==:%+v", AssociatedSceneBoxData_All[index])
}

func (*AssociatedSceneBoxDataMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := AssociatedSceneBoxDataDic[int32Key]
		if !ok {
			logger.LogWarn("AssociatedSceneBoxData PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==AssociatedSceneBoxData==:%+v", vData)
	}
}

func (*AssociatedSceneBoxDataMgr) Load(buffer []byte) bool {
	AssociatedSceneBoxData_All = make([]*AssociatedSceneBoxData, 0)
	err := json.Unmarshal(buffer, &AssociatedSceneBoxData_All)
	if err != nil {
		logger.LogErr("AssociatedSceneBoxData JsonFailed:", err)
		return false
	}
	vLen := len(AssociatedSceneBoxData_All)
	AssociatedSceneBoxDataDic = make(map[int32]*AssociatedSceneBoxData, vLen)
	for _, mem := range AssociatedSceneBoxData_All {
		AssociatedSceneBoxDataDic[mem.Id] = mem
	}
	return true
}
