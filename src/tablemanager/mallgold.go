package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Mallgold 原表为商城表.xlsx 的子表 金币商店
type Mallgold struct {
	Id           int32 `json:"id"`           //商品id
	Item_protoid int32 `json:"item_protoid"` //道具ID道具ID，读取《道具表》中的道具ID
	Item_num     int32 `json:"item_num"`     //数量
	Buytimes     int32 `json:"buytimes"`     //限购次数0不限购
	Price_pre    int32 `json:"price_pre"`    //原价
	Price_then   int32 `json:"price_then"`   //折扣价
	Price_str    int32 `json:"price_str"`    //折扣比例
	Cost_type    int32 `json:"cost_type"`    //购买货币类型
	Cost_num     int32 `json:"cost_num"`     //购买货币/物品数量
	Ishot        int32 `json:"ishot"`        //是否热卖
}

type MallgoldMgr struct {
}

var (
	Mallgold_Model MallgoldMgr
	mallgoldDic    map[int32]*Mallgold
	// Mallgold_All 商城表.xlsx (金币商店)
	Mallgold_All []*Mallgold
)

// Mallgold_Get 商城表.xlsx (金币商店)
func Mallgold_Get(Id int32) (*Mallgold, bool) {
	data, ok := mallgoldDic[Id]
	if !ok {
		PROTO_ERROR_ID = "商城表.xlsx\nmallgold not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MallgoldMgr) PrintArr() {
	vLen := len(Mallgold_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MallgoldMgr) PrintArrOne(index int) {
	logger.Logf("==Mallgold==:%+v", Mallgold_All[index])
}

func (*MallgoldMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := mallgoldDic[int32Key]
		if !ok {
			logger.LogWarn("Mallgold PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Mallgold==:%+v", vData)
	}
}

func (*MallgoldMgr) Load(buffer []byte) bool {
	Mallgold_All = make([]*Mallgold, 0)
	err := json.Unmarshal(buffer, &Mallgold_All)
	if err != nil {
		logger.LogErr("Mallgold JsonFailed:", err)
		return false
	}
	vLen := len(Mallgold_All)
	mallgoldDic = make(map[int32]*Mallgold, vLen)
	for _, mem := range Mallgold_All {
		mallgoldDic[mem.Id] = mem
	}
	return true
}
