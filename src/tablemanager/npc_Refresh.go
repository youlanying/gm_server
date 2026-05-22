package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Npc_Refresh 原表为怪物刷新表.xlsx 的子表 怪物刷新表
type Npc_Refresh struct {
	RefreshID             int32      `json:"RefreshID"`             //刷怪ID
	SceneID               int32      `json:"sceneID"`               //区域ID
	BoxPosition           []float32  `json:"boxPosition"`           //位置
	BoxRotation           []float32  `json:"boxRotation"`           //旋转
	BoxScale              []float32  `json:"boxScale"`              //缩放
	Npc_entity            [][]string `json:"npc_entity"`            //NPC实体ID
	NpcRefreshProbability [][]int32  `json:"npcRefreshProbability"` //每只怪刷新概率
	NpcWhether            []string   `json:"npcWhether"`            //每波NPC击杀个数
	TimelineIds           []int32    `json:"timelineIds"`           //timeline列表id
	DelyRefreshNpcs       []float32  `json:"delyRefreshNpcs"`       //每波怪物刷新延迟时间
	DelyRefreshNpc_random []float32  `json:"delyRefreshNpc_random"` //每只怪物刷新延迟时间
	RefrashType           int32      `json:"refrashType"`           //刷新方式
	NextBoxID             int32      `json:"nextBoxID"`             //下一个盒子ID（对应刷新方式3）
	WarningTime           float32    `json:"warningTime"`           //每波怪物的预警时间（刷新间隔）
	Refresh_points        []int32    `json:"refresh_points"`        //随机刷怪点（此列如果有数据则按照随机刷怪走，没有则按照正常的走）
}

type Npc_RefreshMgr struct {
}

var (
	Npc_Refresh_Model Npc_RefreshMgr
	npc_RefreshDic    map[int32]*Npc_Refresh
	// Npc_Refresh_All 怪物刷新表.xlsx (怪物刷新表)
	Npc_Refresh_All []*Npc_Refresh
)

// Npc_Refresh_Get 怪物刷新表.xlsx (怪物刷新表)
func Npc_Refresh_Get(Id int32) (*Npc_Refresh, bool) {
	data, ok := npc_RefreshDic[Id]
	if !ok {
		PROTO_ERROR_ID = "怪物刷新表.xlsx\nnpc_Refresh not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Npc_RefreshMgr) PrintArr() {
	vLen := len(Npc_Refresh_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Npc_RefreshMgr) PrintArrOne(index int) {
	logger.Logf("==Npc_Refresh==:%+v", Npc_Refresh_All[index])
}

func (*Npc_RefreshMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := npc_RefreshDic[int32Key]
		if !ok {
			logger.LogWarn("Npc_Refresh PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Npc_Refresh==:%+v", vData)
	}
}

func (*Npc_RefreshMgr) Load(buffer []byte) bool {
	Npc_Refresh_All = make([]*Npc_Refresh, 0)
	err := json.Unmarshal(buffer, &Npc_Refresh_All)
	if err != nil {
		logger.LogErr("Npc_Refresh JsonFailed:", err)
		return false
	}
	vLen := len(Npc_Refresh_All)
	npc_RefreshDic = make(map[int32]*Npc_Refresh, vLen)
	for _, mem := range Npc_Refresh_All {
		npc_RefreshDic[mem.RefreshID] = mem
	}
	return true
}
