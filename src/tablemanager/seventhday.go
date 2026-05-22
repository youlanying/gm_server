package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Seventhday 原表为签到与开服七日奖励表.xlsx 的子表 开服七日礼
type Seventhday struct {
	Id     int32   `json:"id"`     //表主键:int
	Reward []int32 `json:"reward"` //奖励(道具ID和数量):list_int
}

type SeventhdayMgr struct {
}

var (
	Seventhday_Model SeventhdayMgr
	seventhdayDic    map[int32]*Seventhday
	// Seventhday_All 签到与开服七日奖励表.xlsx (开服七日礼)
	Seventhday_All []*Seventhday
)

// Seventhday_Get 签到与开服七日奖励表.xlsx (开服七日礼)
func Seventhday_Get(Id int32) (*Seventhday, bool) {
	data, ok := seventhdayDic[Id]
	if !ok {
		PROTO_ERROR_ID = "签到与开服七日奖励表.xlsx\nseventhday not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *SeventhdayMgr) PrintArr() {
	vLen := len(Seventhday_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*SeventhdayMgr) PrintArrOne(index int) {
	logger.Logf("==Seventhday==:%+v", Seventhday_All[index])
}

func (*SeventhdayMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := seventhdayDic[int32Key]
		if !ok {
			logger.LogWarn("Seventhday PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Seventhday==:%+v", vData)
	}
}

func (*SeventhdayMgr) Load(buffer []byte) bool {
	Seventhday_All = make([]*Seventhday, 0)
	err := json.Unmarshal(buffer, &Seventhday_All)
	if err != nil {
		logger.LogErr("Seventhday JsonFailed:", err)
		return false
	}
	vLen := len(Seventhday_All)
	seventhdayDic = make(map[int32]*Seventhday, vLen)
	for _, mem := range Seventhday_All {
		seventhdayDic[mem.Id] = mem
	}
	return true
}
