package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Hero_star 原表为主人公—英雄角色.xlsx 的子表 英雄升星表
type Hero_star struct {
	Id             int32     `json:"id"`             //英雄模板Id*100+星级
	Proto_id       int32     `json:"proto_id"`       //英雄模板id
	Level          int32     `json:"level"`          //星级
	Attr_init      [][]int32 `json:"attr_init"`      //基础属性
	Growth_attr    [][]int32 `json:"growth_attr"`    //属性成长
	Talent_skillid int32     `json:"talent_skillid"` //天赋技能id（技能组表ID）
}

type Hero_starMgr struct {
}

var (
	Hero_star_Model Hero_starMgr
	hero_starDic    map[int32]*Hero_star
	// Hero_star_All 主人公—英雄角色.xlsx (英雄升星表)
	Hero_star_All []*Hero_star
)

// Hero_star_Get 主人公—英雄角色.xlsx (英雄升星表)
func Hero_star_Get(Id int32) (*Hero_star, bool) {
	data, ok := hero_starDic[Id]
	if !ok {
		PROTO_ERROR_ID = "主人公—英雄角色.xlsx\nhero_star not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Hero_starMgr) PrintArr() {
	vLen := len(Hero_star_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Hero_starMgr) PrintArrOne(index int) {
	logger.Logf("==Hero_star==:%+v", Hero_star_All[index])
}

func (*Hero_starMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := hero_starDic[int32Key]
		if !ok {
			logger.LogWarn("Hero_star PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Hero_star==:%+v", vData)
	}
}

func (*Hero_starMgr) Load(buffer []byte) bool {
	Hero_star_All = make([]*Hero_star, 0)
	err := json.Unmarshal(buffer, &Hero_star_All)
	if err != nil {
		logger.LogErr("Hero_star JsonFailed:", err)
		return false
	}
	vLen := len(Hero_star_All)
	hero_starDic = make(map[int32]*Hero_star, vLen)
	for _, mem := range Hero_star_All {
		hero_starDic[mem.Id] = mem
	}
	return true
}
