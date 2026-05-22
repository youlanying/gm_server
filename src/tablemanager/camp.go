package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Camp 原表为阵营表.xlsx 的子表 Sheet1
type Camp struct {
	Id     int32 `json:"id"`     //编号
	Camp1  int32 `json:"camp1"`  //阵营1
	Camp2  int32 `json:"camp2"`  //阵营2
	Camp3  int32 `json:"camp3"`  //阵营3
	Camp4  int32 `json:"camp4"`  //阵营4
	Camp5  int32 `json:"camp5"`  //阵营5
	Camp6  int32 `json:"camp6"`  //阵营6
	Camp7  int32 `json:"camp7"`  //阵营7
	Camp8  int32 `json:"camp8"`  //阵营8
	Camp9  int32 `json:"camp9"`  //阵营9
	Camp10 int32 `json:"camp10"` //阵营10
}

type CampMgr struct {
}

var (
	Camp_Model CampMgr
	campDic    map[int32]*Camp
	// Camp_All 阵营表.xlsx (Sheet1)
	Camp_All []*Camp
)

// Camp_Get 阵营表.xlsx (Sheet1)
func Camp_Get(Id int32) (*Camp, bool) {
	data, ok := campDic[Id]
	if !ok {
		PROTO_ERROR_ID = "阵营表.xlsx\ncamp not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *CampMgr) PrintArr() {
	vLen := len(Camp_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*CampMgr) PrintArrOne(index int) {
	logger.Logf("==Camp==:%+v", Camp_All[index])
}

func (*CampMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := campDic[int32Key]
		if !ok {
			logger.LogWarn("Camp PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Camp==:%+v", vData)
	}
}

func (*CampMgr) Load(buffer []byte) bool {
	Camp_All = make([]*Camp, 0)
	err := json.Unmarshal(buffer, &Camp_All)
	if err != nil {
		logger.LogErr("Camp JsonFailed:", err)
		return false
	}
	vLen := len(Camp_All)
	campDic = make(map[int32]*Camp, vLen)
	for _, mem := range Camp_All {
		campDic[mem.Id] = mem
	}
	return true
}
