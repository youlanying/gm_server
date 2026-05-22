package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Instance_transfer_map 原表为场景传送盒子映射表.xlsx 的子表 场景传送盒子映射表
type Instance_transfer_map struct {
	Id     int32   `json:"id"`     //传送盒子ID
	Target []int32 `json:"target"` //传送目标
}

type Instance_transfer_mapMgr struct {
}

var (
	Instance_transfer_map_Model Instance_transfer_mapMgr
	instance_transfer_mapDic    map[int32]*Instance_transfer_map
	// Instance_transfer_map_All 场景传送盒子映射表.xlsx (场景传送盒子映射表)
	Instance_transfer_map_All []*Instance_transfer_map
)

// Instance_transfer_map_Get 场景传送盒子映射表.xlsx (场景传送盒子映射表)
func Instance_transfer_map_Get(Id int32) (*Instance_transfer_map, bool) {
	data, ok := instance_transfer_mapDic[Id]
	if !ok {
		PROTO_ERROR_ID = "场景传送盒子映射表.xlsx\ninstance_transfer_map not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Instance_transfer_mapMgr) PrintArr() {
	vLen := len(Instance_transfer_map_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Instance_transfer_mapMgr) PrintArrOne(index int) {
	logger.Logf("==Instance_transfer_map==:%+v", Instance_transfer_map_All[index])
}

func (*Instance_transfer_mapMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := instance_transfer_mapDic[int32Key]
		if !ok {
			logger.LogWarn("Instance_transfer_map PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Instance_transfer_map==:%+v", vData)
	}
}

func (*Instance_transfer_mapMgr) Load(buffer []byte) bool {
	Instance_transfer_map_All = make([]*Instance_transfer_map, 0)
	err := json.Unmarshal(buffer, &Instance_transfer_map_All)
	if err != nil {
		logger.LogErr("Instance_transfer_map JsonFailed:", err)
		return false
	}
	vLen := len(Instance_transfer_map_All)
	instance_transfer_mapDic = make(map[int32]*Instance_transfer_map, vLen)
	for _, mem := range Instance_transfer_map_All {
		instance_transfer_mapDic[mem.Id] = mem
	}
	return true
}
