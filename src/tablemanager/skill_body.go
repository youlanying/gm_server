package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Skill_body 原表为技能体表.xlsx 的子表 技能体表
type Skill_body struct {
	Id                int32   `json:"id"`                //技能体ID
	SourceOffsetId    int32   `json:"SourceOffsetId"`    //无目标的出生偏移id
	TargetOffsetId    int32   `json:"TargetOffsetId"`    //有目标的出生偏移
	Type              int32   `json:"type"`              //技能体类型
	BulletEffectId    int32   `json:"bulletEffectId"`    //子弹特效名称
	MoveType          int32   `json:"MoveType"`          //移动类型
	Length            float32 `json:"Length"`            //移动距离或激光长度
	CreateOnGround    float32 `json:"CreateOnGround"`    //是否创建到地面
	LifeTime          float32 `json:"LifeTime"`          //存活时间
	DisappearEffectId int32   `json:"DisappearEffectId"` //子弹消失特效
	BulletCanBlocked  int32   `json:"bulletCanBlocked"`  //子弹是否能被子弹挡墙挡住碰到子弹墙直接消失
	Ispenetrate       int32   `json:"ispenetrate"`       //子弹是否能穿透目标
	TriggerCamp       int32   `json:"TriggerCamp"`       //命中触发阵营关系类型
	Trigger           int32   `json:"trigger"`           //技能触发时机
	TriggerInterval   float32 `json:"TriggerInterval"`   //触发间隔
	Skillid           int32   `json:"skillid"`           //技能组ID只有创建技能体效果出来的技能体会填这个（走技能逻辑）
}

type Skill_bodyMgr struct {
}

var (
	Skill_body_Model Skill_bodyMgr
	skill_bodyDic    map[int32]*Skill_body
	// Skill_body_All 技能体表.xlsx (技能体表)
	Skill_body_All []*Skill_body
)

// Skill_body_Get 技能体表.xlsx (技能体表)
func Skill_body_Get(Id int32) (*Skill_body, bool) {
	data, ok := skill_bodyDic[Id]
	if !ok {
		PROTO_ERROR_ID = "技能体表.xlsx\nskill_body not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Skill_bodyMgr) PrintArr() {
	vLen := len(Skill_body_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Skill_bodyMgr) PrintArrOne(index int) {
	logger.Logf("==Skill_body==:%+v", Skill_body_All[index])
}

func (*Skill_bodyMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := skill_bodyDic[int32Key]
		if !ok {
			logger.LogWarn("Skill_body PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Skill_body==:%+v", vData)
	}
}

func (*Skill_bodyMgr) Load(buffer []byte) bool {
	Skill_body_All = make([]*Skill_body, 0)
	err := json.Unmarshal(buffer, &Skill_body_All)
	if err != nil {
		logger.LogErr("Skill_body JsonFailed:", err)
		return false
	}
	vLen := len(Skill_body_All)
	skill_bodyDic = make(map[int32]*Skill_body, vLen)
	for _, mem := range Skill_body_All {
		skill_bodyDic[mem.Id] = mem
	}
	return true
}
