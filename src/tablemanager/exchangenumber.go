package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
)

// Exchangenumber 原表为兑换码礼包表.xlsx 的子表 Sheet1
type Exchangenumber struct {
	Id    string `json:"id"`    //礼包编码
	Name  string `json:"name"`  //礼包名称
	Item  int32  `json:"item"`  //礼包道具
	Mutex string `json:"mutex"` //互斥组
}

type ExchangenumberMgr struct {
}

var (
	Exchangenumber_Model ExchangenumberMgr
	exchangenumberDic    map[string]*Exchangenumber
	// Exchangenumber_All 兑换码礼包表.xlsx (Sheet1)
	Exchangenumber_All []*Exchangenumber
)

// Exchangenumber_Get 兑换码礼包表.xlsx (Sheet1)
func Exchangenumber_Get(Id string) (*Exchangenumber, bool) {
	data, ok := exchangenumberDic[Id]
	if !ok {
		PROTO_ERROR_ID = "兑换码礼包表.xlsx\nexchangenumber not Id：" + Id
		return nil, false
	}
	return data, true
}
func (this *ExchangenumberMgr) PrintArr() {
	vLen := len(Exchangenumber_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*ExchangenumberMgr) PrintArrOne(index int) {
	logger.Logf("==Exchangenumber==:%+v", Exchangenumber_All[index])
}

func (*ExchangenumberMgr) PrintMapByKey(key interface{}) {
	if strKey, ok := key.(string); ok {
		vData, ok := exchangenumberDic[strKey]
		if !ok {
			logger.LogWarn("Exchangenumber PrintMapByKey Not Find Key:", key)
			return
		}
		logger.Logf("==PrintMapByKey==Exchangenumber==:%+v", vData)
	}
}

func (*ExchangenumberMgr) Load(buffer []byte) bool {
	Exchangenumber_All = make([]*Exchangenumber, 0)
	err := json.Unmarshal(buffer, &Exchangenumber_All)
	if err != nil {
		logger.LogErr("Exchangenumber JsonFailed:", err)
		return false
	}
	vLen := len(Exchangenumber_All)
	exchangenumberDic = make(map[string]*Exchangenumber, vLen)
	for _, mem := range Exchangenumber_All {
		exchangenumberDic[mem.Id] = mem
	}
	return true
}
