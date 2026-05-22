package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Skill 原表为技能表.xlsx 的子表 技能表
type Skill struct {
	Id              int32       `json:"id"`              //技能ID
	Skill_name      string      `json:"skill_name"`      //技能备注名称
	SkillCostType   int32       `json:"SkillCostType"`   //消耗类型
	SkillCostValue  float32     `json:"SkillCostValue"`  //消耗值类型是2那么填能量
	AttackPointTime [][]float32 `json:"attackPointTime"` //子攻击序列
	Priority        int32       `json:"priority"`        //技能优先级
	CanCutByHurt    int32       `json:"canCutByHurt"`    //释放技能的时候是否霸体
	CutHurtTime     []float32   `json:"cutHurtTime"`     //霸体时间区间
}

type SkillMgr struct {
}

var (
	Skill_Model SkillMgr
	skillDic    map[int32]*Skill
	// Skill_All 技能表.xlsx (技能表)
	Skill_All []*Skill
)

// Skill_Get 技能表.xlsx (技能表)
func Skill_Get(Id int32) (*Skill, bool) {
	data, ok := skillDic[Id]
	if !ok {
		PROTO_ERROR_ID = "技能表.xlsx\nskill not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *SkillMgr) PrintArr() {
	vLen := len(Skill_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*SkillMgr) PrintArrOne(index int) {
	logger.Logf("==Skill==:%+v", Skill_All[index])
}

func (*SkillMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := skillDic[int32Key]
		if !ok {
			logger.LogWarn("Skill PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Skill==:%+v", vData)
	}
}

func (*SkillMgr) Load(buffer []byte) bool {
	Skill_All = make([]*Skill, 0)
	err := json.Unmarshal(buffer, &Skill_All)
	if err != nil {
		logger.LogErr("Skill JsonFailed:", err)
		return false
	}
	vLen := len(Skill_All)
	skillDic = make(map[int32]*Skill, vLen)
	for _, mem := range Skill_All {
		skillDic[mem.Id] = mem
	}
	return true
}
