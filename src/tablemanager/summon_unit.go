package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Summon_unit 原表为召唤生物表.xlsx 的子表 召唤生物
type Summon_unit struct {
	Id                  int32     `json:"id"`                  //召唤物id
	Protoid             int32     `json:"protoid"`             //npc模型id
	Npcdataid           int32     `json:"npcdataid"`           //npc数据id
	Npc_modelName       string    `json:"npc_modelName"`       //prefab路径
	AniBornId           int32     `json:"aniBornId"`           //怪物出生动画id
	BtTreeSharePathList []string  `json:"btTreeSharePathList"` //行为树模板路径
	SonBtTreeIndex      []int32   `json:"sonBtTreeIndex"`      //行为树子树索引
	Growth_attr         [][]int32 `json:"growth_attr"`         //属性成长
	Collision_radius    float32   `json:"collision_radius"`    //检测碰撞半径
}

type Summon_unitMgr struct {
}

var (
	Summon_unit_Model Summon_unitMgr
	summon_unitDic    map[int32]*Summon_unit
	// Summon_unit_All 召唤生物表.xlsx (召唤生物)
	Summon_unit_All []*Summon_unit
)

// Summon_unit_Get 召唤生物表.xlsx (召唤生物)
func Summon_unit_Get(Id int32) (*Summon_unit, bool) {
	data, ok := summon_unitDic[Id]
	if !ok {
		PROTO_ERROR_ID = "召唤生物表.xlsx\nsummon_unit not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Summon_unitMgr) PrintArr() {
	vLen := len(Summon_unit_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Summon_unitMgr) PrintArrOne(index int) {
	logger.Logf("==Summon_unit==:%+v", Summon_unit_All[index])
}

func (*Summon_unitMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := summon_unitDic[int32Key]
		if !ok {
			logger.LogWarn("Summon_unit PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Summon_unit==:%+v", vData)
	}
}

func (*Summon_unitMgr) Load(buffer []byte) bool {
	Summon_unit_All = make([]*Summon_unit, 0)
	err := json.Unmarshal(buffer, &Summon_unit_All)
	if err != nil {
		logger.LogErr("Summon_unit JsonFailed:", err)
		return false
	}
	vLen := len(Summon_unit_All)
	summon_unitDic = make(map[int32]*Summon_unit, vLen)
	for _, mem := range Summon_unit_All {
		summon_unitDic[mem.Id] = mem
	}
	return true
}
