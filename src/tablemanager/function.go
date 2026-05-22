package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Function 原表为功能开启表.xlsx 的子表 Sheet1
type Function struct {
	Id         int32   `json:"id"`         //功能ID
	Conditions []int32 `json:"conditions"` //条件表
}

type FunctionMgr struct {
}

var (
	Function_Model FunctionMgr
	functionDic    map[int32]*Function
	// Function_All 功能开启表.xlsx (Sheet1)
	Function_All []*Function
)

// Function_Get 功能开启表.xlsx (Sheet1)
func Function_Get(Id int32) (*Function, bool) {
	data, ok := functionDic[Id]
	if !ok {
		PROTO_ERROR_ID = "功能开启表.xlsx\nfunction not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *FunctionMgr) PrintArr() {
	vLen := len(Function_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*FunctionMgr) PrintArrOne(index int) {
	logger.Logf("==Function==:%+v", Function_All[index])
}

func (*FunctionMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := functionDic[int32Key]
		if !ok {
			logger.LogWarn("Function PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Function==:%+v", vData)
	}
}

func (*FunctionMgr) Load(buffer []byte) bool {
	Function_All = make([]*Function, 0)
	err := json.Unmarshal(buffer, &Function_All)
	if err != nil {
		logger.LogErr("Function JsonFailed:", err)
		return false
	}
	vLen := len(Function_All)
	functionDic = make(map[int32]*Function, vLen)
	for _, mem := range Function_All {
		functionDic[mem.Id] = mem
	}
	return true
}
