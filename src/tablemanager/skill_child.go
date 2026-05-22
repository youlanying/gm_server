package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Skill_child 原表为技能子攻击表.xlsm 的子表 子攻击表
type Skill_child struct {
	Id                      int32     `json:"id"`                      //子攻击ID
	Skill_name              string    `json:"skill_name"`              //技能临时名称
	ProcessType             int32     `json:"ProcessType"`             //流程类型
	ProcessParams           []float32 `json:"ProcessParams"`           //流程参数
	ScreenId                int32     `json:"ScreenId"`                //筛选id
	HitRecoverID            int32     `json:"HitRecoverID"`            //受击动作表现id
	ImmuneAbilityInvincible int32     `json:"ImmuneAbilityInvincible"` //免疫闪避无敌
	MainBuffID              [][]int32 `json:"MainBuffID"`              //主干BuffID
	Branch1BuffID           [][]int32 `json:"Branch1BuffID"`           //分支1BuffID
	Branch2BuffID           [][]int32 `json:"Branch2BuffID"`           //分支2BuffID
	Effectresultid          [][]int32 `json:"effectresultid"`          //主干效果ID
	Effectresultid1         [][]int32 `json:"effectresultid1"`         //分支1效果ID
	Effectresultid2         [][]int32 `json:"effectresultid2"`         //分支2效果ID
	SpResValue              int32     `json:"spResValue"`              //击中目标恢复的Sp值
	Energy_capture          int32     `json:"energy_capture"`          //能量获取
	SlowId                  int32     `json:"SlowId"`                  //慢放id
}

type Skill_childMgr struct {
}

var (
	Skill_child_Model Skill_childMgr
	skill_childDic    map[int32]*Skill_child
	// Skill_child_All 技能子攻击表.xlsm (子攻击表)
	Skill_child_All []*Skill_child
)

// Skill_child_Get 技能子攻击表.xlsm (子攻击表)
func Skill_child_Get(Id int32) (*Skill_child, bool) {
	data, ok := skill_childDic[Id]
	if !ok {
		PROTO_ERROR_ID = "技能子攻击表.xlsm\nskill_child not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Skill_childMgr) PrintArr() {
	vLen := len(Skill_child_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Skill_childMgr) PrintArrOne(index int) {
	logger.Logf("==Skill_child==:%+v", Skill_child_All[index])
}

func (*Skill_childMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := skill_childDic[int32Key]
		if !ok {
			logger.LogWarn("Skill_child PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Skill_child==:%+v", vData)
	}
}

func (*Skill_childMgr) Load(buffer []byte) bool {
	Skill_child_All = make([]*Skill_child, 0)
	err := json.Unmarshal(buffer, &Skill_child_All)
	if err != nil {
		logger.LogErr("Skill_child JsonFailed:", err)
		return false
	}
	vLen := len(Skill_child_All)
	skill_childDic = make(map[int32]*Skill_child, vLen)
	for _, mem := range Skill_child_All {
		skill_childDic[mem.Id] = mem
	}
	return true
}
