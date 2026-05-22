package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Condition_data 原表为条件表.xlsx 的子表 Sheet2
type Condition_data struct {
	Id              int32   `json:"id"`              //条件ID
	Data_type       int32   `json:"data_type"`       //数据类型
	Judge_type      int32   `json:"judge_type"`      //条件判断类型
	Match_condition []int32 `json:"match_condition"` //匹配条件:[]
	Finish_value    int32   `json:"finish_value"`    //条件达成数据
}

type Condition_dataMgr struct {
}

var (
	Condition_data_Model Condition_dataMgr
	condition_dataDic    map[int32]*Condition_data
	// Condition_data_All 条件表.xlsx (Sheet2)
	Condition_data_All []*Condition_data
)

// Condition_data_Get 条件表.xlsx (Sheet2)
func Condition_data_Get(Id int32) (*Condition_data, bool) {
	data, ok := condition_dataDic[Id]
	if !ok {
		PROTO_ERROR_ID = "条件表.xlsx\ncondition_data not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Condition_dataMgr) PrintArr() {
	vLen := len(Condition_data_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Condition_dataMgr) PrintArrOne(index int) {
	logger.Logf("==Condition_data==:%+v", Condition_data_All[index])
}

func (*Condition_dataMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := condition_dataDic[int32Key]
		if !ok {
			logger.LogWarn("Condition_data PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Condition_data==:%+v", vData)
	}
}

func (*Condition_dataMgr) Load(buffer []byte) bool {
	Condition_data_All = make([]*Condition_data, 0)
	err := json.Unmarshal(buffer, &Condition_data_All)
	if err != nil {
		logger.LogErr("Condition_data JsonFailed:", err)
		return false
	}
	vLen := len(Condition_data_All)
	condition_dataDic = make(map[int32]*Condition_data, vLen)
	for _, mem := range Condition_data_All {
		condition_dataDic[mem.Id] = mem
	}
	return true
}
