package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Mallmaze 原表为商城表.xlsx 的子表 迷宫商店
type Mallmaze struct {
	Id           int32  `json:"id"`           //商品id
	Item_protoid int32  `json:"item_protoid"` //道具ID道具ID，读取《道具表》中的道具ID
	Item_num     int32  `json:"item_num"`     //数量
	Buytimes     int32  `json:"buytimes"`     //限购次数0不限购
	Price_pre    int32  `json:"price_pre"`    //原价
	Price_then   int32  `json:"price_then"`   //折扣价
	Price_str    string `json:"price_str"`    //折扣比例
	Cost_type    int32  `json:"cost_type"`    //购买货币类型
	Cost_num     int32  `json:"cost_num"`     //购买货币/物品数量
	Ishot        int32  `json:"ishot"`        //是否热卖
}

type MallmazeMgr struct {
}

var (
	Mallmaze_Model MallmazeMgr
	mallmazeDic    map[int32]*Mallmaze
	// Mallmaze_All 商城表.xlsx (迷宫商店)
	Mallmaze_All []*Mallmaze
)

// Mallmaze_Get 商城表.xlsx (迷宫商店)
func Mallmaze_Get(Id int32) (*Mallmaze, bool) {
	data, ok := mallmazeDic[Id]
	if !ok {
		PROTO_ERROR_ID = "商城表.xlsx\nmallmaze not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MallmazeMgr) PrintArr() {
	vLen := len(Mallmaze_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MallmazeMgr) PrintArrOne(index int) {
	logger.Logf("==Mallmaze==:%+v", Mallmaze_All[index])
}

func (*MallmazeMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := mallmazeDic[int32Key]
		if !ok {
			logger.LogWarn("Mallmaze PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Mallmaze==:%+v", vData)
	}
}

func (*MallmazeMgr) Load(buffer []byte) bool {
	Mallmaze_All = make([]*Mallmaze, 0)
	err := json.Unmarshal(buffer, &Mallmaze_All)
	if err != nil {
		logger.LogErr("Mallmaze JsonFailed:", err)
		return false
	}
	vLen := len(Mallmaze_All)
	mallmazeDic = make(map[int32]*Mallmaze, vLen)
	for _, mem := range Mallmaze_All {
		mallmazeDic[mem.Id] = mem
	}
	return true
}
