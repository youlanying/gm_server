package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Level_son_auto 原表为关卡-子场景数据表.xlsx 的子表 子场景工具生成表
type Level_son_auto struct {
	Id                 int32   `json:"id"`                 //关卡子场景ID
	Area_id            int32   `json:"area_id"`            //区块ID
	BattleTriggerGroup []int32 `json:"BattleTriggerGroup"` //明雷组ID序列
	MechanismGroup     []int32 `json:"MechanismGroup"`     //机关组刷新表
}

type Level_son_autoMgr struct {
}

var (
	Level_son_auto_Model Level_son_autoMgr
	level_son_autoDic    map[int32]*Level_son_auto
	// Level_son_auto_All 关卡-子场景数据表.xlsx (子场景工具生成表)
	Level_son_auto_All []*Level_son_auto
)

// Level_son_auto_Get 关卡-子场景数据表.xlsx (子场景工具生成表)
func Level_son_auto_Get(Id int32) (*Level_son_auto, bool) {
	data, ok := level_son_autoDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡-子场景数据表.xlsx\nlevel_son_auto not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Level_son_autoMgr) PrintArr() {
	vLen := len(Level_son_auto_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Level_son_autoMgr) PrintArrOne(index int) {
	logger.Logf("==Level_son_auto==:%+v", Level_son_auto_All[index])
}

func (*Level_son_autoMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := level_son_autoDic[int32Key]
		if !ok {
			logger.LogWarn("Level_son_auto PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Level_son_auto==:%+v", vData)
	}
}

func (*Level_son_autoMgr) Load(buffer []byte) bool {
	Level_son_auto_All = make([]*Level_son_auto, 0)
	err := json.Unmarshal(buffer, &Level_son_auto_All)
	if err != nil {
		logger.LogErr("Level_son_auto JsonFailed:", err)
		return false
	}
	vLen := len(Level_son_auto_All)
	level_son_autoDic = make(map[int32]*Level_son_auto, vLen)
	for _, mem := range Level_son_auto_All {
		level_son_autoDic[mem.Id] = mem
	}
	return true
}
