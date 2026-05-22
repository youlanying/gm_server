package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Chapter_star_reward 原表为章节星数奖励.xlsx 的子表 Sheet1
type Chapter_star_reward struct {
	Id         int32     `json:"id"`         //id
	Chapter_id int32     `json:"chapter_id"` //章节
	Star_count int32     `json:"star_count"` //档次/星数
	Items      [][]int32 `json:"items"`      //普通本物品ID
}

type Chapter_star_rewardMgr struct {
}

var (
	Chapter_star_reward_Model Chapter_star_rewardMgr
	chapter_star_rewardDic    map[int32]*Chapter_star_reward
	// Chapter_star_reward_All 章节星数奖励.xlsx (Sheet1)
	Chapter_star_reward_All []*Chapter_star_reward
)

// Chapter_star_reward_Get 章节星数奖励.xlsx (Sheet1)
func Chapter_star_reward_Get(Id int32) (*Chapter_star_reward, bool) {
	data, ok := chapter_star_rewardDic[Id]
	if !ok {
		PROTO_ERROR_ID = "章节星数奖励.xlsx\nchapter_star_reward not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Chapter_star_rewardMgr) PrintArr() {
	vLen := len(Chapter_star_reward_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Chapter_star_rewardMgr) PrintArrOne(index int) {
	logger.Logf("==Chapter_star_reward==:%+v", Chapter_star_reward_All[index])
}

func (*Chapter_star_rewardMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := chapter_star_rewardDic[int32Key]
		if !ok {
			logger.LogWarn("Chapter_star_reward PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Chapter_star_reward==:%+v", vData)
	}
}

func (*Chapter_star_rewardMgr) Load(buffer []byte) bool {
	Chapter_star_reward_All = make([]*Chapter_star_reward, 0)
	err := json.Unmarshal(buffer, &Chapter_star_reward_All)
	if err != nil {
		logger.LogErr("Chapter_star_reward JsonFailed:", err)
		return false
	}
	vLen := len(Chapter_star_reward_All)
	chapter_star_rewardDic = make(map[int32]*Chapter_star_reward, vLen)
	for _, mem := range Chapter_star_reward_All {
		chapter_star_rewardDic[mem.Id] = mem
	}
	return true
}
