package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Bossrush_rounds 原表为BossRush表.xlsx 的子表 BOSS RUSH轮次
type Bossrush_rounds struct {
	Id    int32   `json:"id"`    //轮次
	Week1 []int32 `json:"week1"` //第一周（组ID）
	Week2 []int32 `json:"week2"` //第二周
	Week3 []int32 `json:"week3"` //第三周
}

type Bossrush_roundsMgr struct {
}

var (
	Bossrush_rounds_Model Bossrush_roundsMgr
	bossrush_roundsDic    map[int32]*Bossrush_rounds
	// Bossrush_rounds_All BossRush表.xlsx (BOSS RUSH轮次)
	Bossrush_rounds_All []*Bossrush_rounds
)

// Bossrush_rounds_Get BossRush表.xlsx (BOSS RUSH轮次)
func Bossrush_rounds_Get(Id int32) (*Bossrush_rounds, bool) {
	data, ok := bossrush_roundsDic[Id]
	if !ok {
		PROTO_ERROR_ID = "BossRush表.xlsx\nbossrush_rounds not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Bossrush_roundsMgr) PrintArr() {
	vLen := len(Bossrush_rounds_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Bossrush_roundsMgr) PrintArrOne(index int) {
	logger.Logf("==Bossrush_rounds==:%+v", Bossrush_rounds_All[index])
}

func (*Bossrush_roundsMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := bossrush_roundsDic[int32Key]
		if !ok {
			logger.LogWarn("Bossrush_rounds PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Bossrush_rounds==:%+v", vData)
	}
}

func (*Bossrush_roundsMgr) Load(buffer []byte) bool {
	Bossrush_rounds_All = make([]*Bossrush_rounds, 0)
	err := json.Unmarshal(buffer, &Bossrush_rounds_All)
	if err != nil {
		logger.LogErr("Bossrush_rounds JsonFailed:", err)
		return false
	}
	vLen := len(Bossrush_rounds_All)
	bossrush_roundsDic = make(map[int32]*Bossrush_rounds, vLen)
	for _, mem := range Bossrush_rounds_All {
		bossrush_roundsDic[mem.Id] = mem
	}
	return true
}
