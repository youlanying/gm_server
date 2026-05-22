package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Hero_talent 原表为角色天赋表.xlsx 的子表 Sheet1
type Hero_talent struct {
	Id       int32 `json:"id"`       //ID
	Talent_1 int32 `json:"talent_1"` //技能1
	Talent_2 int32 `json:"talent_2"` //技能2
	Talent_3 int32 `json:"talent_3"` //闪避
	Talent_4 int32 `json:"talent_4"` //奥义
}

type Hero_talentMgr struct {
}

var (
	Hero_talent_Model Hero_talentMgr
	hero_talentDic    map[int32]*Hero_talent
	// Hero_talent_All 角色天赋表.xlsx (Sheet1)
	Hero_talent_All []*Hero_talent
)

// Hero_talent_Get 角色天赋表.xlsx (Sheet1)
func Hero_talent_Get(Id int32) (*Hero_talent, bool) {
	data, ok := hero_talentDic[Id]
	if !ok {
		PROTO_ERROR_ID = "角色天赋表.xlsx\nhero_talent not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Hero_talentMgr) PrintArr() {
	vLen := len(Hero_talent_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Hero_talentMgr) PrintArrOne(index int) {
	logger.Logf("==Hero_talent==:%+v", Hero_talent_All[index])
}

func (*Hero_talentMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := hero_talentDic[int32Key]
		if !ok {
			logger.LogWarn("Hero_talent PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Hero_talent==:%+v", vData)
	}
}

func (*Hero_talentMgr) Load(buffer []byte) bool {
	Hero_talent_All = make([]*Hero_talent, 0)
	err := json.Unmarshal(buffer, &Hero_talent_All)
	if err != nil {
		logger.LogErr("Hero_talent JsonFailed:", err)
		return false
	}
	vLen := len(Hero_talent_All)
	hero_talentDic = make(map[int32]*Hero_talent, vLen)
	for _, mem := range Hero_talent_All {
		hero_talentDic[mem.Id] = mem
	}
	return true
}
