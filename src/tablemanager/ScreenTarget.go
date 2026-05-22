package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// ScreenTarget 原表为筛选目标表.xlsx 的子表 筛选目标表
type ScreenTarget struct {
	Id             int32     `json:"Id"`             //筛选目标id
	TargetType     int32     `json:"TargetType"`     //筛选目标类型
	TargetParams   []int32   `json:"TargetParams"`   //筛选目标参数
	RangeType      int32     `json:"RangeType"`      //筛选目标范围类型
	RangeParams    []float32 `json:"RangeParams"`    //筛选目标范围参数
	RangePosOffset []float32 `json:"RangePosOffset"` //范围位置偏移
	RangeAngle     float32   `json:"RangeAngle"`     //范围角度
}

type ScreenTargetMgr struct {
}

var (
	ScreenTarget_Model ScreenTargetMgr
	ScreenTargetDic    map[int32]*ScreenTarget
	// ScreenTarget_All 筛选目标表.xlsx (筛选目标表)
	ScreenTarget_All []*ScreenTarget
)

// ScreenTarget_Get 筛选目标表.xlsx (筛选目标表)
func ScreenTarget_Get(Id int32) (*ScreenTarget, bool) {
	data, ok := ScreenTargetDic[Id]
	if !ok {
		PROTO_ERROR_ID = "筛选目标表.xlsx\nScreenTarget not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *ScreenTargetMgr) PrintArr() {
	vLen := len(ScreenTarget_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*ScreenTargetMgr) PrintArrOne(index int) {
	logger.Logf("==ScreenTarget==:%+v", ScreenTarget_All[index])
}

func (*ScreenTargetMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := ScreenTargetDic[int32Key]
		if !ok {
			logger.LogWarn("ScreenTarget PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==ScreenTarget==:%+v", vData)
	}
}

func (*ScreenTargetMgr) Load(buffer []byte) bool {
	ScreenTarget_All = make([]*ScreenTarget, 0)
	err := json.Unmarshal(buffer, &ScreenTarget_All)
	if err != nil {
		logger.LogErr("ScreenTarget JsonFailed:", err)
		return false
	}
	vLen := len(ScreenTarget_All)
	ScreenTargetDic = make(map[int32]*ScreenTarget, vLen)
	for _, mem := range ScreenTarget_All {
		ScreenTargetDic[mem.Id] = mem
	}
	return true
}
