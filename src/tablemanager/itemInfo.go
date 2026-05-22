package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// ItemInfo 原表为道具表.xlsx 的子表 道具表
type ItemInfo struct {
	Id            int32     `json:"id"`            //物品ID
	Rare_level    int32     `json:"rare_level"`    //稀有程度
	Stacknum      int32     `json:"stacknum"`      //最大堆叠数
	Effectivetime int32     `json:"effectivetime"` //有效期
	Type          int32     `json:"type"`          //物品大类型（背包类型）
	Subtype       int32     `json:"subtype"`       //物品小类型（功能类型）
	Useparam      int32     `json:"useparam"`      //使用参数
	Price         [][]int32 `json:"price"`         //卖出/分解
}

type ItemInfoMgr struct {
}

var (
	ItemInfo_Model ItemInfoMgr
	itemInfoDic    map[int32]*ItemInfo
	// ItemInfo_All 道具表.xlsx (道具表)
	ItemInfo_All []*ItemInfo
)

// ItemInfo_Get 道具表.xlsx (道具表)
func ItemInfo_Get(Id int32) (*ItemInfo, bool) {
	data, ok := itemInfoDic[Id]
	if !ok {
		PROTO_ERROR_ID = "道具表.xlsx\nitemInfo not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *ItemInfoMgr) PrintArr() {
	vLen := len(ItemInfo_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*ItemInfoMgr) PrintArrOne(index int) {
	logger.Logf("==ItemInfo==:%+v", ItemInfo_All[index])
}

func (*ItemInfoMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := itemInfoDic[int32Key]
		if !ok {
			logger.LogWarn("ItemInfo PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==ItemInfo==:%+v", vData)
	}
}

func (*ItemInfoMgr) Load(buffer []byte) bool {
	ItemInfo_All = make([]*ItemInfo, 0)
	err := json.Unmarshal(buffer, &ItemInfo_All)
	if err != nil {
		logger.LogErr("ItemInfo JsonFailed:", err)
		return false
	}
	vLen := len(ItemInfo_All)
	itemInfoDic = make(map[int32]*ItemInfo, vLen)
	for _, mem := range ItemInfo_All {
		itemInfoDic[mem.Id] = mem
	}
	return true
}
