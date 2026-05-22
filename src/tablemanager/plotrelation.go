package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Plotrelation 原表为剧情总表.xlsx 的子表 剧情章节关联表
type Plotrelation struct {
	Id      int32   `json:"id"`      //章节id
	Type    int32   `json:"type"`    //剧情类型
	Section []int32 `json:"section"` //剧情小节
}

type PlotrelationMgr struct {
}

var (
	Plotrelation_Model PlotrelationMgr
	plotrelationDic    map[int32]*Plotrelation
	// Plotrelation_All 剧情总表.xlsx (剧情章节关联表)
	Plotrelation_All []*Plotrelation
)

// Plotrelation_Get 剧情总表.xlsx (剧情章节关联表)
func Plotrelation_Get(Id int32) (*Plotrelation, bool) {
	data, ok := plotrelationDic[Id]
	if !ok {
		PROTO_ERROR_ID = "剧情总表.xlsx\nplotrelation not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *PlotrelationMgr) PrintArr() {
	vLen := len(Plotrelation_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*PlotrelationMgr) PrintArrOne(index int) {
	logger.Logf("==Plotrelation==:%+v", Plotrelation_All[index])
}

func (*PlotrelationMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := plotrelationDic[int32Key]
		if !ok {
			logger.LogWarn("Plotrelation PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Plotrelation==:%+v", vData)
	}
}

func (*PlotrelationMgr) Load(buffer []byte) bool {
	Plotrelation_All = make([]*Plotrelation, 0)
	err := json.Unmarshal(buffer, &Plotrelation_All)
	if err != nil {
		logger.LogErr("Plotrelation JsonFailed:", err)
		return false
	}
	vLen := len(Plotrelation_All)
	plotrelationDic = make(map[int32]*Plotrelation, vLen)
	for _, mem := range Plotrelation_All {
		plotrelationDic[mem.Id] = mem
	}
	return true
}
