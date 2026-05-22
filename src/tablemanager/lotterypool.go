package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Lotterypool 原表为抽卡数据表.xlsx 的子表 卡池数据表
type Lotterypool struct {
	Id   int32     `json:"id"`   //奖池id
	Pool [][]int32 `json:"pool"` //奖池内容[{道具1id，数量，权重};{道具2id，数量，权重};{道具3id，数量，权重};{道具4id，数量，权重};{道具5id，数量，权重};{道具6id，数量，权重}]
}

type LotterypoolMgr struct {
}

var (
	Lotterypool_Model LotterypoolMgr
	lotterypoolDic    map[int32]*Lotterypool
	// Lotterypool_All 抽卡数据表.xlsx (卡池数据表)
	Lotterypool_All []*Lotterypool
)

// Lotterypool_Get 抽卡数据表.xlsx (卡池数据表)
func Lotterypool_Get(Id int32) (*Lotterypool, bool) {
	data, ok := lotterypoolDic[Id]
	if !ok {
		PROTO_ERROR_ID = "抽卡数据表.xlsx\nlotterypool not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *LotterypoolMgr) PrintArr() {
	vLen := len(Lotterypool_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*LotterypoolMgr) PrintArrOne(index int) {
	logger.Logf("==Lotterypool==:%+v", Lotterypool_All[index])
}

func (*LotterypoolMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := lotterypoolDic[int32Key]
		if !ok {
			logger.LogWarn("Lotterypool PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Lotterypool==:%+v", vData)
	}
}

func (*LotterypoolMgr) Load(buffer []byte) bool {
	Lotterypool_All = make([]*Lotterypool, 0)
	err := json.Unmarshal(buffer, &Lotterypool_All)
	if err != nil {
		logger.LogErr("Lotterypool JsonFailed:", err)
		return false
	}
	vLen := len(Lotterypool_All)
	lotterypoolDic = make(map[int32]*Lotterypool, vLen)
	for _, mem := range Lotterypool_All {
		lotterypoolDic[mem.Id] = mem
	}
	return true
}
