package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Ranklevel 原表为rank等级表.xlsx 的子表 工作表1
type Ranklevel struct {
	Id              int32     `json:"id"`              //rank等级
	Level_limit     int32     `json:"level_limit"`     //需求等级
	Attrid          int32     `json:"attrid"`          //属性提升
	Usermoney       int32     `json:"usermoney"`       //进阶金币消耗
	User_item_fire  [][]int32 `json:"user_item_fire"`  //道具消耗1/火属性角色
	User_item_water [][]int32 `json:"user_item_water"` //道具消耗2/水属性角色
	User_item_wind  [][]int32 `json:"user_item_wind"`  //道具消耗3/风属性角色
	User_item_light [][]int32 `json:"user_item_light"` //道具消耗4/光属性角色
	User_item_dark  [][]int32 `json:"user_item_dark"`  //道具消耗5/暗属性角色
}

type RanklevelMgr struct {
}

var (
	Ranklevel_Model RanklevelMgr
	ranklevelDic    map[int32]*Ranklevel
	// Ranklevel_All rank等级表.xlsx (工作表1)
	Ranklevel_All []*Ranklevel
)

// Ranklevel_Get rank等级表.xlsx (工作表1)
func Ranklevel_Get(Id int32) (*Ranklevel, bool) {
	data, ok := ranklevelDic[Id]
	if !ok {
		PROTO_ERROR_ID = "rank等级表.xlsx\nranklevel not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *RanklevelMgr) PrintArr() {
	vLen := len(Ranklevel_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*RanklevelMgr) PrintArrOne(index int) {
	logger.Logf("==Ranklevel==:%+v", Ranklevel_All[index])
}

func (*RanklevelMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := ranklevelDic[int32Key]
		if !ok {
			logger.LogWarn("Ranklevel PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Ranklevel==:%+v", vData)
	}
}

func (*RanklevelMgr) Load(buffer []byte) bool {
	Ranklevel_All = make([]*Ranklevel, 0)
	err := json.Unmarshal(buffer, &Ranklevel_All)
	if err != nil {
		logger.LogErr("Ranklevel JsonFailed:", err)
		return false
	}
	vLen := len(Ranklevel_All)
	ranklevelDic = make(map[int32]*Ranklevel, vLen)
	for _, mem := range Ranklevel_All {
		ranklevelDic[mem.Id] = mem
	}
	return true
}
