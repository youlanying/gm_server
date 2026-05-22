package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Drop 原表为掉落表.xlsx 的子表 Sheet1
type Drop struct {
	Id            int32     `json:"id"`            //掉落id
	Drop_mode     int32     `json:"drop_mode"`     //掉落模式
	Drop_count    int32     `json:"drop_count"`    //掉落数量
	Drop_contents [][]int32 `json:"drop_contents"` //掉落内容
}

type DropMgr struct {
}

var (
	Drop_Model DropMgr
	dropDic    map[int32]*Drop
	// Drop_All 掉落表.xlsx (Sheet1)
	Drop_All []*Drop
)

// Drop_Get 掉落表.xlsx (Sheet1)
func Drop_Get(Id int32) (*Drop, bool) {
	data, ok := dropDic[Id]
	if !ok {
		PROTO_ERROR_ID = "掉落表.xlsx\ndrop not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *DropMgr) PrintArr() {
	vLen := len(Drop_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*DropMgr) PrintArrOne(index int) {
	logger.Logf("==Drop==:%+v", Drop_All[index])
}

func (*DropMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := dropDic[int32Key]
		if !ok {
			logger.LogWarn("Drop PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Drop==:%+v", vData)
	}
}

func (*DropMgr) Load(buffer []byte) bool {
	Drop_All = make([]*Drop, 0)
	err := json.Unmarshal(buffer, &Drop_All)
	if err != nil {
		logger.LogErr("Drop JsonFailed:", err)
		return false
	}
	vLen := len(Drop_All)
	dropDic = make(map[int32]*Drop, vLen)
	for _, mem := range Drop_All {
		dropDic[mem.Id] = mem
	}
	return true
}
