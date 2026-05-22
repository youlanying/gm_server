package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Quest_active 原表为任务表.xlsx 的子表 每日活跃度
type Quest_active struct {
	Id           int32   `json:"id"`           //活跃度值
	Reward_items []int32 `json:"reward_items"` //奖励
}

type Quest_activeMgr struct {
}

var (
	Quest_active_Model Quest_activeMgr
	quest_activeDic    map[int32]*Quest_active
	// Quest_active_All 任务表.xlsx (每日活跃度)
	Quest_active_All []*Quest_active
)

// Quest_active_Get 任务表.xlsx (每日活跃度)
func Quest_active_Get(Id int32) (*Quest_active, bool) {
	data, ok := quest_activeDic[Id]
	if !ok {
		PROTO_ERROR_ID = "任务表.xlsx\nquest_active not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Quest_activeMgr) PrintArr() {
	vLen := len(Quest_active_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Quest_activeMgr) PrintArrOne(index int) {
	logger.Logf("==Quest_active==:%+v", Quest_active_All[index])
}

func (*Quest_activeMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := quest_activeDic[int32Key]
		if !ok {
			logger.LogWarn("Quest_active PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Quest_active==:%+v", vData)
	}
}

func (*Quest_activeMgr) Load(buffer []byte) bool {
	Quest_active_All = make([]*Quest_active, 0)
	err := json.Unmarshal(buffer, &Quest_active_All)
	if err != nil {
		logger.LogErr("Quest_active JsonFailed:", err)
		return false
	}
	vLen := len(Quest_active_All)
	quest_activeDic = make(map[int32]*Quest_active, vLen)
	for _, mem := range Quest_active_All {
		quest_activeDic[mem.Id] = mem
	}
	return true
}
