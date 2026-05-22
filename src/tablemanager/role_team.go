package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Role_team 原表为主人公—团体属性.xlsx 的子表 主角团体属性表
type Role_team struct {
	Id           int32 `json:"id"`           //御崎市等级
	Exp          int32 `json:"exp"`          //升级经验
	Physical     int32 `json:"physical"`     //体力上限
	Hero_max     int32 `json:"hero_max"`     //英雄最大等级
	Add_physical int32 `json:"add_physical"` //达到等级奖励体力
	Treasure_max int32 `json:"treasure_max"` //配备最大等级
	Partner_max  int32 `json:"partner_max"`  //战团最大等级
}

type Role_teamMgr struct {
}

var (
	Role_team_Model Role_teamMgr
	role_teamDic    map[int32]*Role_team
	// Role_team_All 主人公—团体属性.xlsx (主角团体属性表)
	Role_team_All []*Role_team
)

// Role_team_Get 主人公—团体属性.xlsx (主角团体属性表)
func Role_team_Get(Id int32) (*Role_team, bool) {
	data, ok := role_teamDic[Id]
	if !ok {
		PROTO_ERROR_ID = "主人公—团体属性.xlsx\nrole_team not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Role_teamMgr) PrintArr() {
	vLen := len(Role_team_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Role_teamMgr) PrintArrOne(index int) {
	logger.Logf("==Role_team==:%+v", Role_team_All[index])
}

func (*Role_teamMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := role_teamDic[int32Key]
		if !ok {
			logger.LogWarn("Role_team PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Role_team==:%+v", vData)
	}
}

func (*Role_teamMgr) Load(buffer []byte) bool {
	Role_team_All = make([]*Role_team, 0)
	err := json.Unmarshal(buffer, &Role_team_All)
	if err != nil {
		logger.LogErr("Role_team JsonFailed:", err)
		return false
	}
	vLen := len(Role_team_All)
	role_teamDic = make(map[int32]*Role_team, vLen)
	for _, mem := range Role_team_All {
		role_teamDic[mem.Id] = mem
	}
	return true
}
