package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Daysign 原表为签到与开服七日奖励表.xlsx 的子表 日常签到
type Daysign struct {
	Id             int32   `json:"id"`             //表主键:int
	Is_loop        int32   `json:"is_loop"`        //本组奖励是否参与循环（0不参与，1参与）:int
	Sign_1_reward  []int32 `json:"sign_1_reward"`  //签到奖励1(道具ID和数量):list_int
	Sign_2_reward  []int32 `json:"sign_2_reward"`  //签到奖励2:list_int
	Sign_3_reward  []int32 `json:"sign_3_reward"`  //签到奖励3:list_int
	Sign_4_reward  []int32 `json:"sign_4_reward"`  //签到奖励4:list_int
	Sign_5_reward  []int32 `json:"sign_5_reward"`  //签到奖励5:list_int
	Sign_6_reward  []int32 `json:"sign_6_reward"`  //签到奖励6:list_int
	Sign_7_reward  []int32 `json:"sign_7_reward"`  //签到奖励7:list_int
	Sign_8_reward  []int32 `json:"sign_8_reward"`  //签到奖励8:list_int
	Sign_9_reward  []int32 `json:"sign_9_reward"`  //签到奖励9:list_int
	Sign_10_reward []int32 `json:"sign_10_reward"` //签到奖励10:list_int
	Sign_11_reward []int32 `json:"sign_11_reward"` //签到奖励11:list_int
	Sign_12_reward []int32 `json:"sign_12_reward"` //签到奖励12:list_int
	Sign_13_reward []int32 `json:"sign_13_reward"` //签到奖励13:list_int
	Sign_14_reward []int32 `json:"sign_14_reward"` //签到奖励14:list_int
}

type DaysignMgr struct {
}

var (
	Daysign_Model DaysignMgr
	daysignDic    map[int32]*Daysign
	// Daysign_All 签到与开服七日奖励表.xlsx (日常签到)
	Daysign_All []*Daysign
)

// Daysign_Get 签到与开服七日奖励表.xlsx (日常签到)
func Daysign_Get(Id int32) (*Daysign, bool) {
	data, ok := daysignDic[Id]
	if !ok {
		PROTO_ERROR_ID = "签到与开服七日奖励表.xlsx\ndaysign not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *DaysignMgr) PrintArr() {
	vLen := len(Daysign_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*DaysignMgr) PrintArrOne(index int) {
	logger.Logf("==Daysign==:%+v", Daysign_All[index])
}

func (*DaysignMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := daysignDic[int32Key]
		if !ok {
			logger.LogWarn("Daysign PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Daysign==:%+v", vData)
	}
}

func (*DaysignMgr) Load(buffer []byte) bool {
	Daysign_All = make([]*Daysign, 0)
	err := json.Unmarshal(buffer, &Daysign_All)
	if err != nil {
		logger.LogErr("Daysign JsonFailed:", err)
		return false
	}
	vLen := len(Daysign_All)
	daysignDic = make(map[int32]*Daysign, vLen)
	for _, mem := range Daysign_All {
		daysignDic[mem.Id] = mem
	}
	return true
}
