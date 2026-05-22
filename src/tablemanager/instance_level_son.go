package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Instance_level_son 原表为关卡_子场景表.xlsx 的子表 子场景表
type Instance_level_son struct {
	Id                   int32   `json:"id"`                   //关卡子场景ID
	Area_id              int32   `json:"area_id"`              //区块ID
	AssociatedSceneBoxId []int32 `json:"AssociatedSceneBoxId"` //关联场景盒子id（通用实体id）
	RefreshBoxId         []int32 `json:"refreshBoxId"`         //刷怪盒子序列
	TransferBox          []int32 `json:"TransferBox"`          //传送盒子
	FightAreaId          []int32 `json:"fightAreaId"`          //战斗区域ID
	AreaSize             string  `json:"AreaSize"`             //区域大小
	Success_condition    []int32 `json:"success_condition"`    //通关条件
	Fail_condition       []int32 `json:"fail_condition"`       //失败条件
	Hero_exp             int32   `json:"hero_exp"`             //通关角色经验奖励
	Team_exp             int32   `json:"team_exp"`             //通关团队经验奖励
}

type Instance_level_sonMgr struct {
}

var (
	Instance_level_son_Model Instance_level_sonMgr
	instance_level_sonDic    map[int32]*Instance_level_son
	// Instance_level_son_All 关卡_子场景表.xlsx (子场景表)
	Instance_level_son_All []*Instance_level_son
)

// Instance_level_son_Get 关卡_子场景表.xlsx (子场景表)
func Instance_level_son_Get(Id int32) (*Instance_level_son, bool) {
	data, ok := instance_level_sonDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡_子场景表.xlsx\ninstance_level_son not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Instance_level_sonMgr) PrintArr() {
	vLen := len(Instance_level_son_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Instance_level_sonMgr) PrintArrOne(index int) {
	logger.Logf("==Instance_level_son==:%+v", Instance_level_son_All[index])
}

func (*Instance_level_sonMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := instance_level_sonDic[int32Key]
		if !ok {
			logger.LogWarn("Instance_level_son PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Instance_level_son==:%+v", vData)
	}
}

func (*Instance_level_sonMgr) Load(buffer []byte) bool {
	Instance_level_son_All = make([]*Instance_level_son, 0)
	err := json.Unmarshal(buffer, &Instance_level_son_All)
	if err != nil {
		logger.LogErr("Instance_level_son JsonFailed:", err)
		return false
	}
	vLen := len(Instance_level_son_All)
	instance_level_sonDic = make(map[int32]*Instance_level_son, vLen)
	for _, mem := range Instance_level_son_All {
		instance_level_sonDic[mem.Id] = mem
	}
	return true
}
