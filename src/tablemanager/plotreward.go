package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Plotreward 原表为剧情总表.xlsx 的子表 剧情奖励及资源
type Plotreward struct {
	Id           int32   `json:"id"`           //剧情小节id
	Plottype     int32   `json:"plottype"`     //剧情类型
	Owner        int32   `json:"owner"`        //所属
	Goods        []int32 `json:"goods"`        //奖励内容：[道具id，数量]
	Proper_value []int32 `json:"proper_value"` //属性奖励：[属性id，值]
}

type PlotrewardMgr struct {
}

var (
	Plotreward_Model PlotrewardMgr
	plotrewardDic    map[int32]*Plotreward
	// Plotreward_All 剧情总表.xlsx (剧情奖励及资源)
	Plotreward_All []*Plotreward
)

// Plotreward_Get 剧情总表.xlsx (剧情奖励及资源)
func Plotreward_Get(Id int32) (*Plotreward, bool) {
	data, ok := plotrewardDic[Id]
	if !ok {
		PROTO_ERROR_ID = "剧情总表.xlsx\nplotreward not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *PlotrewardMgr) PrintArr() {
	vLen := len(Plotreward_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*PlotrewardMgr) PrintArrOne(index int) {
	logger.Logf("==Plotreward==:%+v", Plotreward_All[index])
}

func (*PlotrewardMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := plotrewardDic[int32Key]
		if !ok {
			logger.LogWarn("Plotreward PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Plotreward==:%+v", vData)
	}
}

func (*PlotrewardMgr) Load(buffer []byte) bool {
	Plotreward_All = make([]*Plotreward, 0)
	err := json.Unmarshal(buffer, &Plotreward_All)
	if err != nil {
		logger.LogErr("Plotreward JsonFailed:", err)
		return false
	}
	vLen := len(Plotreward_All)
	plotrewardDic = make(map[int32]*Plotreward, vLen)
	for _, mem := range Plotreward_All {
		plotrewardDic[mem.Id] = mem
	}
	return true
}
