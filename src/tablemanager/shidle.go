package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Shidle 原表为屏蔽字.xlsx 的子表 Sheet1
type Shidle struct {
	Id        int32  `json:"id"`        //ID
	Character string `json:"character"` //名称
}

type ShidleMgr struct {
}

var (
	Shidle_Model ShidleMgr
	shidleDic    map[int32]*Shidle
	// Shidle_All 屏蔽字.xlsx (Sheet1)
	Shidle_All []*Shidle
)

// Shidle_Get 屏蔽字.xlsx (Sheet1)
func Shidle_Get(Id int32) (*Shidle, bool) {
	data, ok := shidleDic[Id]
	if !ok {
		PROTO_ERROR_ID = "屏蔽字.xlsx\nshidle not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *ShidleMgr) PrintArr() {
	vLen := len(Shidle_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*ShidleMgr) PrintArrOne(index int) {
	logger.Logf("==Shidle==:%+v", Shidle_All[index])
}

func (*ShidleMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := shidleDic[int32Key]
		if !ok {
			logger.LogWarn("Shidle PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Shidle==:%+v", vData)
	}
}

func (*ShidleMgr) Load(buffer []byte) bool {
	Shidle_All = make([]*Shidle, 0)
	err := json.Unmarshal(buffer, &Shidle_All)
	if err != nil {
		logger.LogErr("Shidle JsonFailed:", err)
		return false
	}
	vLen := len(Shidle_All)
	shidleDic = make(map[int32]*Shidle, vLen)
	for _, mem := range Shidle_All {
		shidleDic[mem.Id] = mem
	}
	return true
}
