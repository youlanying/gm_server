package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Hero_information 原表为主人公—英雄角色.xlsx 的子表 角色信息表
type Hero_information struct {
	Id            int32   `json:"id"`            //角色id
	Name          string  `json:"name"`          //角色名称
	Hero_class    int32   `json:"hero_class"`    //攻击类型
	Hero_element  int32   `json:"hero_element"`  //元素类型
	Shardid       int32   `json:"shardid"`       //碎片id
	Shardnum      int32   `json:"shardnum"`      //合成所需数量
	Star          int32   `json:"star"`          //初始星级
	Skill_ids     []int32 `json:"skill_ids"`     //初始技能组
	Advanced_type int32   `json:"advanced_type"` //进阶类型
}

type Hero_informationMgr struct {
}

var (
	Hero_information_Model Hero_informationMgr
	hero_informationDic    map[int32]*Hero_information
	// Hero_information_All 主人公—英雄角色.xlsx (角色信息表)
	Hero_information_All []*Hero_information
)

// Hero_information_Get 主人公—英雄角色.xlsx (角色信息表)
func Hero_information_Get(Id int32) (*Hero_information, bool) {
	data, ok := hero_informationDic[Id]
	if !ok {
		PROTO_ERROR_ID = "主人公—英雄角色.xlsx\nhero_information not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Hero_informationMgr) PrintArr() {
	vLen := len(Hero_information_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Hero_informationMgr) PrintArrOne(index int) {
	logger.Logf("==Hero_information==:%+v", Hero_information_All[index])
}

func (*Hero_informationMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := hero_informationDic[int32Key]
		if !ok {
			logger.LogWarn("Hero_information PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Hero_information==:%+v", vData)
	}
}

func (*Hero_informationMgr) Load(buffer []byte) bool {
	Hero_information_All = make([]*Hero_information, 0)
	err := json.Unmarshal(buffer, &Hero_information_All)
	if err != nil {
		logger.LogErr("Hero_information JsonFailed:", err)
		return false
	}
	vLen := len(Hero_information_All)
	hero_informationDic = make(map[int32]*Hero_information, vLen)
	for _, mem := range Hero_information_All {
		hero_informationDic[mem.Id] = mem
	}
	return true
}
