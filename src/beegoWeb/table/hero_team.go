package table

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Hero_team 原表为主人公—英雄角色.xlsx 的子表 英雄团体表
type Hero_team struct {
	Id        int32   `json:"id"`        //团队id
	Team_name string  `json:"team_name"` //团队名字
	Team_list []int32 `json:"team_list"` //团队英雄列表
}

type Hero_teamMgr struct {
}

var (
	Hero_team_Model Hero_teamMgr
	hero_teamDic    map[int32]*Hero_team
	// Hero_team_All 主人公—英雄角色.xlsx (英雄团体表)
	Hero_team_All []*Hero_team
)

// Hero_team_Get 主人公—英雄角色.xlsx (英雄团体表)
func Hero_team_Get(Id int32) (*Hero_team, bool) {
	data, ok := hero_teamDic[Id]
	if !ok {
		PROTO_ERROR_ID = "主人公—英雄角色.xlsx\nhero_team not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Hero_teamMgr) PrintArr() {
	vLen := len(Hero_team_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Hero_teamMgr) PrintArrOne(index int) {
	logger.Logf("==Hero_team==:%+v", Hero_team_All[index])
}

func (*Hero_teamMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := hero_teamDic[int32Key]
		if !ok {
			logger.LogWarn("Hero_team PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Hero_team==:%+v", vData)
	}
}

func (*Hero_teamMgr) Load(buffer []byte) bool {
	Hero_team_All = make([]*Hero_team, 0)
	err := json.Unmarshal(buffer, &Hero_team_All)
	if err != nil {
		logger.LogErr("Hero_team JsonFailed:", err)
		return false
	}
	vLen := len(Hero_team_All)
	hero_teamDic = make(map[int32]*Hero_team, vLen)
	for _, mem := range Hero_team_All {
		hero_teamDic[mem.Id] = mem
	}
	return true
}
