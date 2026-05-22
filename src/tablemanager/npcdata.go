package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Npcdata 原表为关卡表-场景npc数据表.xlsx 的子表 场景npc数据表
type Npcdata struct {
	Id                  int32     `json:"id"`                  //NPC模板ID
	Type_detail         int32     `json:"type_detail"`         //怪物子类型
	Drop                int32     `json:"drop"`                //掉落ID
	DropType            int32     `json:"dropType"`            //掉落类型
	DropParm            []int32   `json:"dropParm"`            //掉落参数
	Drop_coineffect     int32     `json:"drop_coineffect"`     //掉落金币特效/血球特效
	Drop_itemeffect     int32     `json:"drop_itemeffect"`     //掉落物品特效
	Drop_geteffect      int32     `json:"drop_geteffect"`      //获得掉落物品特效
	Level               int32     `json:"level"`               //怪物等级
	Attr_init           [][]int32 `json:"attr_init"`           //怪物属性
	Miss_skyhurt        int32     `json:"miss_skyhurt"`        //是否免疫浮空
	Miss_flyhurt        int32     `json:"miss_flyhurt"`        //是否免疫击飞
	Miss_normalhurt     int32     `json:"miss_normalhurt"`     //是否免疫硬直
	Miss_bemove         int32     `json:"miss_bemove"`         //是否免疫被动位移
	Becrush             int32     `json:"becrush"`             //是否可以被挤动
	Crush_other         int32     `json:"crush_other"`         //是否可以挤动玩家
	Gravityf            float32   `json:"gravityf"`            //怪物重力
	DeadTime            float32   `json:"deadTime"`            //尸体存留时间
	Skill_list          []int32   `json:"skill_list"`          //技能列表
	Npc_rotationSpeed   float32   `json:"npc_rotationSpeed"`   //旋转速度（冒险中）
	HpPercent           []float32 `json:"hpPercent"`           //血量到达百分比后播放timeline
	TimeLinePath        []int32   `json:"timeLinePath"`        //timeLine路径
	Hp_num              int32     `json:"hp_num"`              //Boss血条数量
	BrokenSuperArmor    int32     `json:"BrokenSuperArmor"`    //是否会被奥义受击
	SuperArmorBuffs     []int32   `json:"superArmorBuffs"`     //被动BUFF（霸体护甲）
	WeaknessPart        []int32   `json:"weaknessPart"`        //NPC弱点ID
	Flytype             int32     `json:"flytype"`             //飞行标记
	Defense             []float32 `json:"defense"`             //防御属性[减伤角度，减伤倍率，是否开启霸体，默认开关]
	Start_buff          []int32   `json:"start_buff"`          //初始Buff
	Locked_priority     int32     `json:"locked_priority"`     //锁定优先级
	BtTreeSharePathList []string  `json:"btTreeSharePathList"` //行为树模板路径
	Camp                int32     `json:"camp"`                //所属阵营
}

type NpcdataMgr struct {
}

var (
	Npcdata_Model NpcdataMgr
	npcdataDic    map[int32]*Npcdata
	// Npcdata_All 关卡表-场景npc数据表.xlsx (场景npc数据表)
	Npcdata_All []*Npcdata
)

// Npcdata_Get 关卡表-场景npc数据表.xlsx (场景npc数据表)
func Npcdata_Get(Id int32) (*Npcdata, bool) {
	data, ok := npcdataDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡表-场景npc数据表.xlsx\nnpcdata not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *NpcdataMgr) PrintArr() {
	vLen := len(Npcdata_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*NpcdataMgr) PrintArrOne(index int) {
	logger.Logf("==Npcdata==:%+v", Npcdata_All[index])
}

func (*NpcdataMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := npcdataDic[int32Key]
		if !ok {
			logger.LogWarn("Npcdata PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Npcdata==:%+v", vData)
	}
}

func (*NpcdataMgr) Load(buffer []byte) bool {
	Npcdata_All = make([]*Npcdata, 0)
	err := json.Unmarshal(buffer, &Npcdata_All)
	if err != nil {
		logger.LogErr("Npcdata JsonFailed:", err)
		return false
	}
	vLen := len(Npcdata_All)
	npcdataDic = make(map[int32]*Npcdata, vLen)
	for _, mem := range Npcdata_All {
		npcdataDic[mem.Id] = mem
	}
	return true
}
