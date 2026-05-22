package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// BattleTriggerGroup 原表为关卡-明雷数据表.xlsx 的子表 明雷组表
type BattleTriggerGroup struct {
	Id           int32   `json:"id"`           //明雷组ID
	BattleId     int32   `json:"BattleId"`     //战斗场景的子场景ID
	RefreshBoxId []int32 `json:"refreshBoxId"` //刷怪盒子ID
	HeroExp      int32   `json:"HeroExp"`      //通关角色经验奖励
	TeamExp      int32   `json:"TeamExp"`      //通关团队经验奖励
}

type BattleTriggerGroupMgr struct {
}

var (
	BattleTriggerGroup_Model BattleTriggerGroupMgr
	BattleTriggerGroupDic    map[int32]*BattleTriggerGroup
	// BattleTriggerGroup_All 关卡-明雷数据表.xlsx (明雷组表)
	BattleTriggerGroup_All []*BattleTriggerGroup
)

// BattleTriggerGroup_Get 关卡-明雷数据表.xlsx (明雷组表)
func BattleTriggerGroup_Get(Id int32) (*BattleTriggerGroup, bool) {
	data, ok := BattleTriggerGroupDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡-明雷数据表.xlsx\nBattleTriggerGroup not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *BattleTriggerGroupMgr) PrintArr() {
	vLen := len(BattleTriggerGroup_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*BattleTriggerGroupMgr) PrintArrOne(index int) {
	logger.Logf("==BattleTriggerGroup==:%+v", BattleTriggerGroup_All[index])
}

func (*BattleTriggerGroupMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := BattleTriggerGroupDic[int32Key]
		if !ok {
			logger.LogWarn("BattleTriggerGroup PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==BattleTriggerGroup==:%+v", vData)
	}
}

func (*BattleTriggerGroupMgr) Load(buffer []byte) bool {
	BattleTriggerGroup_All = make([]*BattleTriggerGroup, 0)
	err := json.Unmarshal(buffer, &BattleTriggerGroup_All)
	if err != nil {
		logger.LogErr("BattleTriggerGroup JsonFailed:", err)
		return false
	}
	vLen := len(BattleTriggerGroup_All)
	BattleTriggerGroupDic = make(map[int32]*BattleTriggerGroup, vLen)
	for _, mem := range BattleTriggerGroup_All {
		BattleTriggerGroupDic[mem.Id] = mem
	}
	return true
}
