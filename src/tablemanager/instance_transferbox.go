package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Instance_transferbox 原表为场景传送盒子表.xlsx 的子表 场景传送盒子表
type Instance_transferbox struct {
	Id      int32 `json:"id"`      //传送盒子ID
	SceneId int32 `json:"SceneId"` //盒子所在关卡子场景ID
}

type Instance_transferboxMgr struct {
}

var (
	Instance_transferbox_Model Instance_transferboxMgr
	instance_transferboxDic    map[int32]*Instance_transferbox
	// Instance_transferbox_All 场景传送盒子表.xlsx (场景传送盒子表)
	Instance_transferbox_All []*Instance_transferbox
)

// Instance_transferbox_Get 场景传送盒子表.xlsx (场景传送盒子表)
func Instance_transferbox_Get(Id int32) (*Instance_transferbox, bool) {
	data, ok := instance_transferboxDic[Id]
	if !ok {
		PROTO_ERROR_ID = "场景传送盒子表.xlsx\ninstance_transferbox not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Instance_transferboxMgr) PrintArr() {
	vLen := len(Instance_transferbox_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Instance_transferboxMgr) PrintArrOne(index int) {
	logger.Logf("==Instance_transferbox==:%+v", Instance_transferbox_All[index])
}

func (*Instance_transferboxMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := instance_transferboxDic[int32Key]
		if !ok {
			logger.LogWarn("Instance_transferbox PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Instance_transferbox==:%+v", vData)
	}
}

func (*Instance_transferboxMgr) Load(buffer []byte) bool {
	Instance_transferbox_All = make([]*Instance_transferbox, 0)
	err := json.Unmarshal(buffer, &Instance_transferbox_All)
	if err != nil {
		logger.LogErr("Instance_transferbox JsonFailed:", err)
		return false
	}
	vLen := len(Instance_transferbox_All)
	instance_transferboxDic = make(map[int32]*Instance_transferbox, vLen)
	for _, mem := range Instance_transferbox_All {
		instance_transferboxDic[mem.Id] = mem
	}
	return true
}
