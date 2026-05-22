package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Time 原表为运营活动.xlsx 的子表 时间表
type Time struct {
	Id        int32   `json:"id"`        //时间ID
	Type      int32   `json:"type"`      //时间类型
	Weekday   int32   `json:"weekday"`   //星期几
	Timestart []int32 `json:"timestart"` //开始时间
	Timeend   int32   `json:"timeend"`   //持续时间
	Daydelay  int32   `json:"daydelay"`  //延迟秒数
}

type TimeMgr struct {
}

var (
	Time_Model TimeMgr
	timeDic    map[int32]*Time
	// Time_All 运营活动.xlsx (时间表)
	Time_All []*Time
)

// Time_Get 运营活动.xlsx (时间表)
func Time_Get(Id int32) (*Time, bool) {
	data, ok := timeDic[Id]
	if !ok {
		PROTO_ERROR_ID = "运营活动.xlsx\ntime not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *TimeMgr) PrintArr() {
	vLen := len(Time_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*TimeMgr) PrintArrOne(index int) {
	logger.Logf("==Time==:%+v", Time_All[index])
}

func (*TimeMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := timeDic[int32Key]
		if !ok {
			logger.LogWarn("Time PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Time==:%+v", vData)
	}
}

func (*TimeMgr) Load(buffer []byte) bool {
	Time_All = make([]*Time, 0)
	err := json.Unmarshal(buffer, &Time_All)
	if err != nil {
		logger.LogErr("Time JsonFailed:", err)
		return false
	}
	vLen := len(Time_All)
	timeDic = make(map[int32]*Time, vLen)
	for _, mem := range Time_All {
		timeDic[mem.Id] = mem
	}
	return true
}
