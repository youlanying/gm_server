package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Talentinfo 原表为技能树模板表.xlsx 的子表 Sheet1
type Talentinfo struct {
	Id              int32   `json:"id"`              //ID
	Slot            int32   `json:"slot"`            //槽位
	Level           int32   `json:"level"`           //等级
	Before_talent   int32   `json:"before_talent"`   //前置
	Next_talent     int32   `json:"next_talent"`     //后置
	Consumeitem     []int32 `json:"consumeitem"`     //消耗道具
	Consumemoney    int32   `json:"consumemoney"`    //消耗金币
	Condition_type  int32   `json:"condition_type"`  //条件类型
	Condition_level int32   `json:"condition_level"` //当前条件等级
}

type TalentinfoMgr struct {
}

var (
	Talentinfo_Model TalentinfoMgr
	talentinfoDic    map[int32]*Talentinfo
	// Talentinfo_All 技能树模板表.xlsx (Sheet1)
	Talentinfo_All []*Talentinfo
)

// Talentinfo_Get 技能树模板表.xlsx (Sheet1)
func Talentinfo_Get(Id int32) (*Talentinfo, bool) {
	data, ok := talentinfoDic[Id]
	if !ok {
		PROTO_ERROR_ID = "技能树模板表.xlsx\ntalentinfo not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *TalentinfoMgr) PrintArr() {
	vLen := len(Talentinfo_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*TalentinfoMgr) PrintArrOne(index int) {
	logger.Logf("==Talentinfo==:%+v", Talentinfo_All[index])
}

func (*TalentinfoMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := talentinfoDic[int32Key]
		if !ok {
			logger.LogWarn("Talentinfo PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Talentinfo==:%+v", vData)
	}
}

func (*TalentinfoMgr) Load(buffer []byte) bool {
	Talentinfo_All = make([]*Talentinfo, 0)
	err := json.Unmarshal(buffer, &Talentinfo_All)
	if err != nil {
		logger.LogErr("Talentinfo JsonFailed:", err)
		return false
	}
	vLen := len(Talentinfo_All)
	talentinfoDic = make(map[int32]*Talentinfo, vLen)
	for _, mem := range Talentinfo_All {
		talentinfoDic[mem.Id] = mem
	}
	return true
}
