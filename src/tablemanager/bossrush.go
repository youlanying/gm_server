package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Bossrush 原表为BossRush表.xlsx 的子表 BOSS关数据表
type Bossrush struct {
	Id              int32     `json:"id"`              //id
	Group_id        int32     `json:"group_id"`        //组（同1只boss的三个难度归为1组）
	Boss_id         int32     `json:"boss_id"`         //bossId
	Instance_id     int32     `json:"instance_id"`     //sceneId
	Rewards         [][]int32 `json:"rewards"`         //普通奖励
	Special_rewards [][]int32 `json:"special_rewards"` //特殊奖励
	Difficulty      int32     `json:"difficulty"`      //难度
	Achievement     []int32   `json:"achievement"`     //成就条件
}

type BossrushMgr struct {
}

var (
	Bossrush_Model BossrushMgr
	bossrushDic    map[int32]*Bossrush
	// Bossrush_All BossRush表.xlsx (BOSS关数据表)
	Bossrush_All []*Bossrush
)

// Bossrush_Get BossRush表.xlsx (BOSS关数据表)
func Bossrush_Get(Id int32) (*Bossrush, bool) {
	data, ok := bossrushDic[Id]
	if !ok {
		PROTO_ERROR_ID = "BossRush表.xlsx\nbossrush not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *BossrushMgr) PrintArr() {
	vLen := len(Bossrush_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*BossrushMgr) PrintArrOne(index int) {
	logger.Logf("==Bossrush==:%+v", Bossrush_All[index])
}

func (*BossrushMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := bossrushDic[int32Key]
		if !ok {
			logger.LogWarn("Bossrush PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Bossrush==:%+v", vData)
	}
}

func (*BossrushMgr) Load(buffer []byte) bool {
	Bossrush_All = make([]*Bossrush, 0)
	err := json.Unmarshal(buffer, &Bossrush_All)
	if err != nil {
		logger.LogErr("Bossrush JsonFailed:", err)
		return false
	}
	vLen := len(Bossrush_All)
	bossrushDic = make(map[int32]*Bossrush, vLen)
	for _, mem := range Bossrush_All {
		bossrushDic[mem.Id] = mem
	}
	return true
}
