package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Instance_level_result 原表为关卡_场景表.xlsx 的子表 消耗倍数对应产出倍数（待优化）
type Instance_level_result struct {
	Id            int32 `json:"id"`            //key
	Cost_itemid   int32 `json:"cost_itemid"`   //消耗物品id
	Cost_itemnum  int32 `json:"cost_itemnum"`  //消耗物品id数量
	Result_double int32 `json:"result_double"` //产出倍数
}

type Instance_level_resultMgr struct {
}

var (
	Instance_level_result_Model Instance_level_resultMgr
	instance_level_resultDic    map[int32]*Instance_level_result
	// Instance_level_result_All 关卡_场景表.xlsx (消耗倍数对应产出倍数（待优化）)
	Instance_level_result_All []*Instance_level_result
)

// Instance_level_result_Get 关卡_场景表.xlsx (消耗倍数对应产出倍数（待优化）)
func Instance_level_result_Get(Id int32) (*Instance_level_result, bool) {
	data, ok := instance_level_resultDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡_场景表.xlsx\ninstance_level_result not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Instance_level_resultMgr) PrintArr() {
	vLen := len(Instance_level_result_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Instance_level_resultMgr) PrintArrOne(index int) {
	logger.Logf("==Instance_level_result==:%+v", Instance_level_result_All[index])
}

func (*Instance_level_resultMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := instance_level_resultDic[int32Key]
		if !ok {
			logger.LogWarn("Instance_level_result PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Instance_level_result==:%+v", vData)
	}
}

func (*Instance_level_resultMgr) Load(buffer []byte) bool {
	Instance_level_result_All = make([]*Instance_level_result, 0)
	err := json.Unmarshal(buffer, &Instance_level_result_All)
	if err != nil {
		logger.LogErr("Instance_level_result JsonFailed:", err)
		return false
	}
	vLen := len(Instance_level_result_All)
	instance_level_resultDic = make(map[int32]*Instance_level_result, vLen)
	for _, mem := range Instance_level_result_All {
		instance_level_resultDic[mem.Id] = mem
	}
	return true
}
