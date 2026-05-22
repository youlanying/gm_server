package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Sweep_level 原表为扫荡表.xlsx 的子表 Sheet1
type Sweep_level struct {
	Id          int32     `json:"id"`          //关卡ID
	Asklevel    int32     `json:"asklevel"`    //要求等级
	Score_on    int32     `json:"score_on"`    //允许扫荡的要求积分
	Score_max   int32     `json:"score_max"`   //最高分数
	Score_gears []int32   `json:"score_gears"` //关卡扫荡分数档位
	Score_pro   []float32 `json:"score_pro"`   //扫荡概率
	Money_max   int32     `json:"money_max"`   //关卡金币数
	Ifsweep     int32     `json:"ifsweep"`     //是否可以扫荡
	Useitemid   int32     `json:"useitemid"`   //扫荡所消耗的道具ID
	Usenum      int32     `json:"usenum"`      //消耗扫荡道具数量
	Droplist    []int32   `json:"droplist"`    //掉落ID
	Droplist2   []int32   `json:"droplist2"`   //集齐眼睛时的掉落ID
	Ifeye       int32     `json:"ifeye"`       //是否判断眼睛
	Eyedrop     []int32   `json:"eyedrop"`     //开启眼睛
}

type Sweep_levelMgr struct {
}

var (
	Sweep_level_Model Sweep_levelMgr
	sweep_levelDic    map[int32]*Sweep_level
	// Sweep_level_All 扫荡表.xlsx (Sheet1)
	Sweep_level_All []*Sweep_level
)

// Sweep_level_Get 扫荡表.xlsx (Sheet1)
func Sweep_level_Get(Id int32) (*Sweep_level, bool) {
	data, ok := sweep_levelDic[Id]
	if !ok {
		PROTO_ERROR_ID = "扫荡表.xlsx\nsweep_level not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Sweep_levelMgr) PrintArr() {
	vLen := len(Sweep_level_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Sweep_levelMgr) PrintArrOne(index int) {
	logger.Logf("==Sweep_level==:%+v", Sweep_level_All[index])
}

func (*Sweep_levelMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := sweep_levelDic[int32Key]
		if !ok {
			logger.LogWarn("Sweep_level PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Sweep_level==:%+v", vData)
	}
}

func (*Sweep_levelMgr) Load(buffer []byte) bool {
	Sweep_level_All = make([]*Sweep_level, 0)
	err := json.Unmarshal(buffer, &Sweep_level_All)
	if err != nil {
		logger.LogErr("Sweep_level JsonFailed:", err)
		return false
	}
	vLen := len(Sweep_level_All)
	sweep_levelDic = make(map[int32]*Sweep_level, vLen)
	for _, mem := range Sweep_level_All {
		sweep_levelDic[mem.Id] = mem
	}
	return true
}
