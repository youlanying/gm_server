package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Defaultdata 原表为默认数据表.xlsx 的子表 Sheet1
type Defaultdata struct {
	Id            int32       `json:"id"`            //key
	Value         []int32     `json:"value"`         //value
	String_value  [][]string  `json:"string_value"`  //string二维列表
	Vector3_value [][]float32 `json:"vector3_value"` //Vector3列表
	ValueString   []string    `json:"valueString"`   //string列表
	ValueFloat    []float32   `json:"valueFloat"`    //float列表
}

type DefaultdataMgr struct {
}

var (
	Defaultdata_Model DefaultdataMgr
	defaultdataDic    map[int32]*Defaultdata
	// Defaultdata_All 默认数据表.xlsx (Sheet1)
	Defaultdata_All []*Defaultdata
)

// Defaultdata_Get 默认数据表.xlsx (Sheet1)
func Defaultdata_Get(Id int32) (*Defaultdata, bool) {
	data, ok := defaultdataDic[Id]
	if !ok {
		PROTO_ERROR_ID = "默认数据表.xlsx\ndefaultdata not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *DefaultdataMgr) PrintArr() {
	vLen := len(Defaultdata_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*DefaultdataMgr) PrintArrOne(index int) {
	logger.Logf("==Defaultdata==:%+v", Defaultdata_All[index])
}

func (*DefaultdataMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := defaultdataDic[int32Key]
		if !ok {
			logger.LogWarn("Defaultdata PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Defaultdata==:%+v", vData)
	}
}

func (*DefaultdataMgr) Load(buffer []byte) bool {
	Defaultdata_All = make([]*Defaultdata, 0)
	err := json.Unmarshal(buffer, &Defaultdata_All)
	if err != nil {
		logger.LogErr("Defaultdata JsonFailed:", err)
		return false
	}
	vLen := len(Defaultdata_All)
	defaultdataDic = make(map[int32]*Defaultdata, vLen)
	for _, mem := range Defaultdata_All {
		defaultdataDic[mem.Id] = mem
	}
	return true
}
