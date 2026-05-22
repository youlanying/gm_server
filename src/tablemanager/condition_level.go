package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Condition_level 原表为关卡条件表.xlsx 的子表 关卡条件表
type Condition_level struct {
	Id              int32   `json:"id"`              //ID
	Data_type       int32   `json:"data_type"`       //条件类型
	Judge_type      int32   `json:"judge_type"`      //判断类型
	Match_condition []int32 `json:"match_condition"` //条件参数
}

type Condition_levelMgr struct {
}

var (
	Condition_level_Model Condition_levelMgr
	condition_levelDic    map[int32]*Condition_level
	// Condition_level_All 关卡条件表.xlsx (关卡条件表)
	Condition_level_All []*Condition_level
)

// Condition_level_Get 关卡条件表.xlsx (关卡条件表)
func Condition_level_Get(Id int32) (*Condition_level, bool) {
	data, ok := condition_levelDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡条件表.xlsx\ncondition_level not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Condition_levelMgr) PrintArr() {
	vLen := len(Condition_level_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Condition_levelMgr) PrintArrOne(index int) {
	logger.Logf("==Condition_level==:%+v", Condition_level_All[index])
}

func (*Condition_levelMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := condition_levelDic[int32Key]
		if !ok {
			logger.LogWarn("Condition_level PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Condition_level==:%+v", vData)
	}
}

func (*Condition_levelMgr) Load(buffer []byte) bool {
	Condition_level_All = make([]*Condition_level, 0)
	err := json.Unmarshal(buffer, &Condition_level_All)
	if err != nil {
		logger.LogErr("Condition_level JsonFailed:", err)
		return false
	}
	vLen := len(Condition_level_All)
	condition_levelDic = make(map[int32]*Condition_level, vLen)
	for _, mem := range Condition_level_All {
		condition_levelDic[mem.Id] = mem
	}
	return true
}
