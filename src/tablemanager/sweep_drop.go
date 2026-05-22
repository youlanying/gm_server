package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Sweep_drop 原表为扫荡掉落表.xlsx 的子表 Sheet2
type Sweep_drop struct {
	Id      int32     `json:"id"`      //掉落id
	Group   [][]int32 `json:"group"`   //掉落组概率
	Group_1 [][]int32 `json:"group_1"` //掉落物品组_1
	Group_2 [][]int32 `json:"group_2"` //掉落物品组_2
	Group_3 [][]int32 `json:"group_3"` //掉落物品组_3
	Group_4 [][]int32 `json:"group_4"` //掉落物品组_4
}

type Sweep_dropMgr struct {
}

var (
	Sweep_drop_Model Sweep_dropMgr
	sweep_dropDic    map[int32]*Sweep_drop
	// Sweep_drop_All 扫荡掉落表.xlsx (Sheet2)
	Sweep_drop_All []*Sweep_drop
)

// Sweep_drop_Get 扫荡掉落表.xlsx (Sheet2)
func Sweep_drop_Get(Id int32) (*Sweep_drop, bool) {
	data, ok := sweep_dropDic[Id]
	if !ok {
		PROTO_ERROR_ID = "扫荡掉落表.xlsx\nsweep_drop not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Sweep_dropMgr) PrintArr() {
	vLen := len(Sweep_drop_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Sweep_dropMgr) PrintArrOne(index int) {
	logger.Logf("==Sweep_drop==:%+v", Sweep_drop_All[index])
}

func (*Sweep_dropMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := sweep_dropDic[int32Key]
		if !ok {
			logger.LogWarn("Sweep_drop PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Sweep_drop==:%+v", vData)
	}
}

func (*Sweep_dropMgr) Load(buffer []byte) bool {
	Sweep_drop_All = make([]*Sweep_drop, 0)
	err := json.Unmarshal(buffer, &Sweep_drop_All)
	if err != nil {
		logger.LogErr("Sweep_drop JsonFailed:", err)
		return false
	}
	vLen := len(Sweep_drop_All)
	sweep_dropDic = make(map[int32]*Sweep_drop, vLen)
	for _, mem := range Sweep_drop_All {
		sweep_dropDic[mem.Id] = mem
	}
	return true
}
