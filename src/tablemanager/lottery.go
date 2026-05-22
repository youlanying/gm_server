package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Lottery 原表为抽卡数据表.xlsx 的子表 抽卡数据表
type Lottery struct {
	Id            int32     `json:"id"`            //表ID
	Type          int32     `json:"type"`          //卡池类型：0道具池,1常驻池，2FES池，3活动角色池；4,UP活动池
	Activityid    int32     `json:"activityid"`    //活动ID
	Costtype      int32     `json:"costtype"`      //抽取消耗类型0金币，1抽卡券
	Dailyoncefree int32     `json:"dailyoncefree"` //是否每日免费1次（0不免费，1免费）
	Onceprice     int32     `json:"onceprice"`     //单抽价格
	Dailytenfree  int32     `json:"dailytenfree"`  //是否每日免费10次（0不免费，1免费）
	Tentimesprice int32     `json:"tentimesprice"` //十抽价格
	Exteralgift   [][]int32 `json:"exteralgift"`   //每次抽卡额外赠送道具[{}]
	Keepcount     int32     `json:"keepcount"`     //保底抽取次数
	Limitcount    int32     `json:"limitcount"`    //必得次数
	Poolid        int32     `json:"poolid"`        //基础奖池组int
	Poolidkeep    int32     `json:"poolidkeep"`    //保底奖池组
	Poolidlimit   int32     `json:"poolidlimit"`   //必得奖池组
}

type LotteryMgr struct {
}

var (
	Lottery_Model LotteryMgr
	lotteryDic    map[int32]*Lottery
	// Lottery_All 抽卡数据表.xlsx (抽卡数据表)
	Lottery_All []*Lottery
)

// Lottery_Get 抽卡数据表.xlsx (抽卡数据表)
func Lottery_Get(Id int32) (*Lottery, bool) {
	data, ok := lotteryDic[Id]
	if !ok {
		PROTO_ERROR_ID = "抽卡数据表.xlsx\nlottery not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *LotteryMgr) PrintArr() {
	vLen := len(Lottery_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*LotteryMgr) PrintArrOne(index int) {
	logger.Logf("==Lottery==:%+v", Lottery_All[index])
}

func (*LotteryMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := lotteryDic[int32Key]
		if !ok {
			logger.LogWarn("Lottery PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Lottery==:%+v", vData)
	}
}

func (*LotteryMgr) Load(buffer []byte) bool {
	Lottery_All = make([]*Lottery, 0)
	err := json.Unmarshal(buffer, &Lottery_All)
	if err != nil {
		logger.LogErr("Lottery JsonFailed:", err)
		return false
	}
	vLen := len(Lottery_All)
	lotteryDic = make(map[int32]*Lottery, vLen)
	for _, mem := range Lottery_All {
		lotteryDic[mem.Id] = mem
	}
	return true
}
