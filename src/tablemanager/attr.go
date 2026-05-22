package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Attr 原表为属性统计.xlsx 的子表 Sheet1
type Attr struct {
	Id              int32   `json:"id"`              //类型ID
	AttrCoefficient float32 `json:"attrCoefficient"` //战力系数
}

type AttrMgr struct {
}

var (
	Attr_Model AttrMgr
	attrDic    map[int32]*Attr
	// Attr_All 属性统计.xlsx (Sheet1)
	Attr_All []*Attr
)

// Attr_Get 属性统计.xlsx (Sheet1)
func Attr_Get(Id int32) (*Attr, bool) {
	data, ok := attrDic[Id]
	if !ok {
		PROTO_ERROR_ID = "属性统计.xlsx\nattr not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *AttrMgr) PrintArr() {
	vLen := len(Attr_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*AttrMgr) PrintArrOne(index int) {
	logger.Logf("==Attr==:%+v", Attr_All[index])
}

func (*AttrMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := attrDic[int32Key]
		if !ok {
			logger.LogWarn("Attr PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Attr==:%+v", vData)
	}
}

func (*AttrMgr) Load(buffer []byte) bool {
	Attr_All = make([]*Attr, 0)
	err := json.Unmarshal(buffer, &Attr_All)
	if err != nil {
		logger.LogErr("Attr JsonFailed:", err)
		return false
	}
	vLen := len(Attr_All)
	attrDic = make(map[int32]*Attr, vLen)
	for _, mem := range Attr_All {
		attrDic[mem.Id] = mem
	}
	return true
}
