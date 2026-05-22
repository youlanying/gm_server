package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Malldiamond 原表为商城表.xlsx 的子表 钻石商店
type Malldiamond struct {
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

type MalldiamondMgr struct {
}

var (
	Malldiamond_Model MalldiamondMgr
	malldiamondDic    map[int32]*Malldiamond
	// Malldiamond_All 商城表.xlsx (钻石商店)
	Malldiamond_All []*Malldiamond
)

// Malldiamond_Get 商城表.xlsx (钻石商店)
func Malldiamond_Get(Id int32) (*Malldiamond, bool) {
	data, ok := malldiamondDic[Id]
	if !ok {
		PROTO_ERROR_ID = "商城表.xlsx\nmalldiamond not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MalldiamondMgr) PrintArr() {
	vLen := len(Malldiamond_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MalldiamondMgr) PrintArrOne(index int) {
	logger.Logf("==Malldiamond==:%+v", Malldiamond_All[index])
}

func (*MalldiamondMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := malldiamondDic[int32Key]
		if !ok {
			logger.LogWarn("Malldiamond PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Malldiamond==:%+v", vData)
	}
}

func (*MalldiamondMgr) Load(buffer []byte) bool {
	Malldiamond_All = make([]*Malldiamond, 0)
	err := json.Unmarshal(buffer, &Malldiamond_All)
	if err != nil {
		logger.LogErr("Malldiamond JsonFailed:", err)
		return false
	}
	vLen := len(Malldiamond_All)
	malldiamondDic = make(map[int32]*Malldiamond, vLen)
	for _, mem := range Malldiamond_All {
		malldiamondDic[mem.Id] = mem
	}
	return true
}
