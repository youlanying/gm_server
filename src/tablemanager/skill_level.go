package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Skill_level 原表为技能升级表.xlsx 的子表 技能升级
type Skill_level struct {
	Id         int32 `json:"id"`         //等级
	Cost_money int32 `json:"cost_money"` //消耗金钱
}

type Skill_levelMgr struct {
}

var (
	Skill_level_Model Skill_levelMgr
	skill_levelDic    map[int32]*Skill_level
	// Skill_level_All 技能升级表.xlsx (技能升级)
	Skill_level_All []*Skill_level
)

// Skill_level_Get 技能升级表.xlsx (技能升级)
func Skill_level_Get(Id int32) (*Skill_level, bool) {
	data, ok := skill_levelDic[Id]
	if !ok {
		PROTO_ERROR_ID = "技能升级表.xlsx\nskill_level not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Skill_levelMgr) PrintArr() {
	vLen := len(Skill_level_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Skill_levelMgr) PrintArrOne(index int) {
	logger.Logf("==Skill_level==:%+v", Skill_level_All[index])
}

func (*Skill_levelMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := skill_levelDic[int32Key]
		if !ok {
			logger.LogWarn("Skill_level PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Skill_level==:%+v", vData)
	}
}

func (*Skill_levelMgr) Load(buffer []byte) bool {
	Skill_level_All = make([]*Skill_level, 0)
	err := json.Unmarshal(buffer, &Skill_level_All)
	if err != nil {
		logger.LogErr("Skill_level JsonFailed:", err)
		return false
	}
	vLen := len(Skill_level_All)
	skill_levelDic = make(map[int32]*Skill_level, vLen)
	for _, mem := range Skill_level_All {
		skill_levelDic[mem.Id] = mem
	}
	return true
}
