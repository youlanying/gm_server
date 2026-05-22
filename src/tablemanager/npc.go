package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Npc 原表为关卡表-场景npc模型表.xlsx 的子表 NPC模型表
type Npc struct {
	Id                   int32       `json:"id"`                   //NPC模板ID
	Npc_modelName        string      `json:"npc_modelName"`        //NPC模型
	Name                 string      `json:"name"`                 //名称
	Dead_fade_effect     int32       `json:"dead_fade_effect"`     //尸体消失特效
	Dead_effect          int32       `json:"dead_effect"`          //死亡特效
	BirthEffect          int32       `json:"birthEffect"`          //出生特效ID
	Scale                float32     `json:"scale"`                //缩放
	HitSound             []string    `json:"hitSound"`             //NPC受击音效
	HitSoundProbability  int32       `json:"hitSoundProbability"`  //受击音效播放概率
	DeadSound            []string    `json:"deadSound"`            //NPC死亡音效
	DeadSoundProbability int32       `json:"deadSoundProbability"` //死亡的时候音效播放概率
	HurtEffect           int32       `json:"hurtEffect"`           //受击特效
	ResidentEffect       []int32     `json:"residentEffect"`       //NPC全局特效
	Billboard_height     float32     `json:"billboard_height"`     //姓名版高度
	HintAddHeight        float32     `json:"hintAddHeight"`        //BOSS掉血跳字的附加高度
	BornAction_ID        int32       `json:"bornAction_ID"`        //出生动作组编号
	Boss_icon            string      `json:"boss_icon"`            //BOSS头像（战斗UI用）
	Buffeffect_scale     []float32   `json:"buffeffect_scale"`     //BUFF资源特效缩放大小
	IsCloseCollision     int32       `json:"isCloseCollision"`     //(死亡)是否关闭碰撞
	LockEnmyEffectScale  float32     `json:"lockEnmyEffectScale"`  //锁敌特效缩放值
	MagicElementId       int32       `json:"MagicElementId"`       //元素属性ID
	CameraDistanceVal    [][]float32 `json:"CameraDistanceVal"`    //镜头的最大最小距离
	HeadIcon             []string    `json:"headIcon"`             //NPC头像
	IsTurnBack           int32       `json:"isTurnBack"`           //是否有转身
	Collider_grade       int32       `json:"collider_grade"`       //碰撞等级
}

type NpcMgr struct {
}

var (
	Npc_Model NpcMgr
	npcDic    map[int32]*Npc
	// Npc_All 关卡表-场景npc模型表.xlsx (NPC模型表)
	Npc_All []*Npc
)

// Npc_Get 关卡表-场景npc模型表.xlsx (NPC模型表)
func Npc_Get(Id int32) (*Npc, bool) {
	data, ok := npcDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡表-场景npc模型表.xlsx\nnpc not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *NpcMgr) PrintArr() {
	vLen := len(Npc_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*NpcMgr) PrintArrOne(index int) {
	logger.Logf("==Npc==:%+v", Npc_All[index])
}

func (*NpcMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := npcDic[int32Key]
		if !ok {
			logger.LogWarn("Npc PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Npc==:%+v", vData)
	}
}

func (*NpcMgr) Load(buffer []byte) bool {
	Npc_All = make([]*Npc, 0)
	err := json.Unmarshal(buffer, &Npc_All)
	if err != nil {
		logger.LogErr("Npc JsonFailed:", err)
		return false
	}
	vLen := len(Npc_All)
	npcDic = make(map[int32]*Npc, vLen)
	for _, mem := range Npc_All {
		npcDic[mem.Id] = mem
	}
	return true
}
