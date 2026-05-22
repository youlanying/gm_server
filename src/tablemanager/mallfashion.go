package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Mallfashion 原表为商城表.xlsx 的子表 时装商店
type Mallfashion struct {
	Id           int32  `json:"id"`           //商品id
	Item_protoid int32  `json:"item_protoid"` //道具ID道具ID，读取《道具表》中的道具ID
	Item_num     int32  `json:"item_num"`     //数量
	Price_pre    int32  `json:"price_pre"`    //原价不填原价或者原价=折扣价时不显示折扣
	Price_then   int32  `json:"price_then"`   //折扣价
	Price_str    string `json:"price_str"`    //折扣比例
	Cost_type    int32  `json:"cost_type"`    //购买货币类型/物品ID
	Cost_num     int32  `json:"cost_num"`     //购买货币/物品数量
	Ishot        int32  `json:"ishot"`        //是否热卖
}

type MallfashionMgr struct {
}

var (
	Mallfashion_Model MallfashionMgr
	mallfashionDic    map[int32]*Mallfashion
	// Mallfashion_All 商城表.xlsx (时装商店)
	Mallfashion_All []*Mallfashion
)

// Mallfashion_Get 商城表.xlsx (时装商店)
func Mallfashion_Get(Id int32) (*Mallfashion, bool) {
	data, ok := mallfashionDic[Id]
	if !ok {
		PROTO_ERROR_ID = "商城表.xlsx\nmallfashion not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MallfashionMgr) PrintArr() {
	vLen := len(Mallfashion_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MallfashionMgr) PrintArrOne(index int) {
	logger.Logf("==Mallfashion==:%+v", Mallfashion_All[index])
}

func (*MallfashionMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := mallfashionDic[int32Key]
		if !ok {
			logger.LogWarn("Mallfashion PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Mallfashion==:%+v", vData)
	}
}

func (*MallfashionMgr) Load(buffer []byte) bool {
	Mallfashion_All = make([]*Mallfashion, 0)
	err := json.Unmarshal(buffer, &Mallfashion_All)
	if err != nil {
		logger.LogErr("Mallfashion JsonFailed:", err)
		return false
	}
	vLen := len(Mallfashion_All)
	mallfashionDic = make(map[int32]*Mallfashion, vLen)
	for _, mem := range Mallfashion_All {
		mallfashionDic[mem.Id] = mem
	}
	return true
}
