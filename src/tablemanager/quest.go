package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Quest 原表为任务表.xlsx 的子表 任务
type Quest struct {
	Id            int32     `json:"id"`            //任务id
	Quest_trigger int32     `json:"quest_trigger"` //任务触发方式
	Tag           int32     `json:"tag"`           //任务类型
	Pre_questid   int32     `json:"pre_questid"`   //前置任务
	After_questid []int32   `json:"after_questid"` //后置任务
	Condition     int32     `json:"condition"`     //条件id
	Reward_items  [][]int32 `json:"reward_items"`  //任务奖励
	Activity      int32     `json:"activity"`      //任务活跃度
}

type QuestMgr struct {
}

var (
	Quest_Model QuestMgr
	questDic    map[int32]*Quest
	// Quest_All 任务表.xlsx (任务)
	Quest_All []*Quest
)

// Quest_Get 任务表.xlsx (任务)
func Quest_Get(Id int32) (*Quest, bool) {
	data, ok := questDic[Id]
	if !ok {
		PROTO_ERROR_ID = "任务表.xlsx\nquest not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *QuestMgr) PrintArr() {
	vLen := len(Quest_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*QuestMgr) PrintArrOne(index int) {
	logger.Logf("==Quest==:%+v", Quest_All[index])
}

func (*QuestMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := questDic[int32Key]
		if !ok {
			logger.LogWarn("Quest PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Quest==:%+v", vData)
	}
}

func (*QuestMgr) Load(buffer []byte) bool {
	Quest_All = make([]*Quest, 0)
	err := json.Unmarshal(buffer, &Quest_All)
	if err != nil {
		logger.LogErr("Quest JsonFailed:", err)
		return false
	}
	vLen := len(Quest_All)
	questDic = make(map[int32]*Quest, vLen)
	for _, mem := range Quest_All {
		questDic[mem.Id] = mem
	}
	return true
}
