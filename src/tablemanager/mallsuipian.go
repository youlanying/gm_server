package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Mallsuipian 原表为商城表.xlsx 的子表 碎片商店
type Mallsuipian struct {
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

type MallsuipianMgr struct {
}

var (
	Mallsuipian_Model MallsuipianMgr
	mallsuipianDic    map[int32]*Mallsuipian
	// Mallsuipian_All 商城表.xlsx (碎片商店)
	Mallsuipian_All []*Mallsuipian
)

// Mallsuipian_Get 商城表.xlsx (碎片商店)
func Mallsuipian_Get(Id int32) (*Mallsuipian, bool) {
	data, ok := mallsuipianDic[Id]
	if !ok {
		PROTO_ERROR_ID = "商城表.xlsx\nmallsuipian not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MallsuipianMgr) PrintArr() {
	vLen := len(Mallsuipian_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MallsuipianMgr) PrintArrOne(index int) {
	logger.Logf("==Mallsuipian==:%+v", Mallsuipian_All[index])
}

func (*MallsuipianMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := mallsuipianDic[int32Key]
		if !ok {
			logger.LogWarn("Mallsuipian PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Mallsuipian==:%+v", vData)
	}
}

func (*MallsuipianMgr) Load(buffer []byte) bool {
	Mallsuipian_All = make([]*Mallsuipian, 0)
	err := json.Unmarshal(buffer, &Mallsuipian_All)
	if err != nil {
		logger.LogErr("Mallsuipian JsonFailed:", err)
		return false
	}
	vLen := len(Mallsuipian_All)
	mallsuipianDic = make(map[int32]*Mallsuipian, vLen)
	for _, mem := range Mallsuipian_All {
		mallsuipianDic[mem.Id] = mem
	}
	return true
}
