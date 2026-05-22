package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
)

// Npc_object 原表为关卡表-场景npc实体表.xlsx 的子表 NPC实体表
type Npc_object struct {
	Id                string    `json:"id"`                //实体ID[关卡ID+随机套数+NPC类型+编号]
	Instance_level_id int32     `json:"instance_level_id"` //区域ID
	Type              int32     `json:"type"`              //NPC类型
	Protoid           int32     `json:"protoid"`           //NPC模版模型
	Pos               []float32 `json:"pos"`               //坐标
	Forward           []float32 `json:"forward"`           //朝向
	Npcdataid         int32     `json:"npcdataid"`         //NPC数据ID
	IsRefresh         int32     `json:"isRefresh"`         //是否刷新
	AniBornId         int32     `json:"aniBornId"`         //怪物出生动画id
	HideBtnType       int32     `json:"hideBtnType"`       //隐藏按钮类型（0无,1圆形，2矩形,3受击必死）
	HideBtnParms      []float32 `json:"hideBtnParms"`      //隐藏按钮参数
}

type Npc_objectMgr struct {
}

var (
	Npc_object_Model Npc_objectMgr
	npc_objectDic    map[string]*Npc_object
	// Npc_object_All 关卡表-场景npc实体表.xlsx (NPC实体表)
	Npc_object_All []*Npc_object
)

// Npc_object_Get 关卡表-场景npc实体表.xlsx (NPC实体表)
func Npc_object_Get(Id string) (*Npc_object, bool) {
	data, ok := npc_objectDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡表-场景npc实体表.xlsx\nnpc_object not Id：" + Id
		return nil, false
	}
	return data, true
}
func (this *Npc_objectMgr) PrintArr() {
	vLen := len(Npc_object_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Npc_objectMgr) PrintArrOne(index int) {
	logger.Logf("==Npc_object==:%+v", Npc_object_All[index])
}

func (*Npc_objectMgr) PrintMapByKey(key interface{}) {
	if strKey, ok := key.(string); ok {
		vData, ok := npc_objectDic[strKey]
		if !ok {
			logger.LogWarn("Npc_object PrintMapByKey Not Find Key:", key)
			return
		}
		logger.Logf("==PrintMapByKey==Npc_object==:%+v", vData)
	}
}

func (*Npc_objectMgr) Load(buffer []byte) bool {
	Npc_object_All = make([]*Npc_object, 0)
	err := json.Unmarshal(buffer, &Npc_object_All)
	if err != nil {
		logger.LogErr("Npc_object JsonFailed:", err)
		return false
	}
	vLen := len(Npc_object_All)
	npc_objectDic = make(map[string]*Npc_object, vLen)
	for _, mem := range Npc_object_All {
		npc_objectDic[mem.Id] = mem
	}
	return true
}
