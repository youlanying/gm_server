package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Skillgroups 原表为技能组表.xlsx 的子表 技能组表
type Skillgroups struct {
	Id                int32     `json:"id"`                //技能组ID
	SkillIdList       [][]int32 `json:"SkillIdList"`       //技能ID列表
	TypeList          []int32   `json:"TypeList"`          //技能机制类型
	SkillCostTimes    int32     `json:"SkillCostTimes"`    //CD积攒次数
	InnerCD           float32   `json:"InnerCD"`           //内置CD
	SkillLocationType int32     `json:"SkillLocationType"` //技能位置类型
	SkillDescId       []int32   `json:"SkillDescId"`       //技能描述
	SkillIcon         []string  `json:"skillIcon"`         //技能图标
}

type SkillgroupsMgr struct {
}

var (
	Skillgroups_Model SkillgroupsMgr
	skillgroupsDic    map[int32]*Skillgroups
	// Skillgroups_All 技能组表.xlsx (技能组表)
	Skillgroups_All []*Skillgroups
)

// Skillgroups_Get 技能组表.xlsx (技能组表)
func Skillgroups_Get(Id int32) (*Skillgroups, bool) {
	data, ok := skillgroupsDic[Id]
	if !ok {
		PROTO_ERROR_ID = "技能组表.xlsx\nskillgroups not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *SkillgroupsMgr) PrintArr() {
	vLen := len(Skillgroups_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*SkillgroupsMgr) PrintArrOne(index int) {
	logger.Logf("==Skillgroups==:%+v", Skillgroups_All[index])
}

func (*SkillgroupsMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := skillgroupsDic[int32Key]
		if !ok {
			logger.LogWarn("Skillgroups PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Skillgroups==:%+v", vData)
	}
}

func (*SkillgroupsMgr) Load(buffer []byte) bool {
	Skillgroups_All = make([]*Skillgroups, 0)
	err := json.Unmarshal(buffer, &Skillgroups_All)
	if err != nil {
		logger.LogErr("Skillgroups JsonFailed:", err)
		return false
	}
	vLen := len(Skillgroups_All)
	skillgroupsDic = make(map[int32]*Skillgroups, vLen)
	for _, mem := range Skillgroups_All {
		skillgroupsDic[mem.Id] = mem
	}
	return true
}
