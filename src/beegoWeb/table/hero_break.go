package table

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Hero_break 原表为主人公—英雄角色.xlsx 的子表 英雄突破表
type Hero_break struct {
	Id                  int32     `json:"id"`                  //英雄ID
	Break_list          [][]int32 `json:"break_list"`          //突破等级
	Break_material_list [][]int32 `json:"break_material_list"` //突破材料
	Cost_gold           [][]int32 `json:"cost_gold"`           //突破金币
}

type Hero_breakMgr struct {
}

var (
	Hero_break_Model Hero_breakMgr
	hero_breakDic    map[int32]*Hero_break
	// Hero_break_All 主人公—英雄角色.xlsx (英雄突破表)
	Hero_break_All []*Hero_break
)

// Hero_break_Get 主人公—英雄角色.xlsx (英雄突破表)
func Hero_break_Get(Id int32) (*Hero_break, bool) {
	data, ok := hero_breakDic[Id]
	if !ok {
		PROTO_ERROR_ID = "主人公—英雄角色.xlsx\nhero_break not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Hero_breakMgr) PrintArr() {
	vLen := len(Hero_break_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Hero_breakMgr) PrintArrOne(index int) {
	logger.Logf("==Hero_break==:%+v", Hero_break_All[index])
}

func (*Hero_breakMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := hero_breakDic[int32Key]
		if !ok {
			logger.LogWarn("Hero_break PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Hero_break==:%+v", vData)
	}
}

func (*Hero_breakMgr) Load(buffer []byte) bool {
	Hero_break_All = make([]*Hero_break, 0)
	err := json.Unmarshal(buffer, &Hero_break_All)
	if err != nil {
		logger.LogErr("Hero_break JsonFailed:", err)
		return false
	}
	vLen := len(Hero_break_All)
	hero_breakDic = make(map[int32]*Hero_break, vLen)
	for _, mem := range Hero_break_All {
		hero_breakDic[mem.Id] = mem
	}
	return true
}
