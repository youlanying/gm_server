package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Mallboss 原表为商城表.xlsx 的子表 BOSS商店
type Mallboss struct {
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

type MallbossMgr struct {
}

var (
	Mallboss_Model MallbossMgr
	mallbossDic    map[int32]*Mallboss
	// Mallboss_All 商城表.xlsx (BOSS商店)
	Mallboss_All []*Mallboss
)

// Mallboss_Get 商城表.xlsx (BOSS商店)
func Mallboss_Get(Id int32) (*Mallboss, bool) {
	data, ok := mallbossDic[Id]
	if !ok {
		PROTO_ERROR_ID = "商城表.xlsx\nmallboss not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *MallbossMgr) PrintArr() {
	vLen := len(Mallboss_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*MallbossMgr) PrintArrOne(index int) {
	logger.Logf("==Mallboss==:%+v", Mallboss_All[index])
}

func (*MallbossMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := mallbossDic[int32Key]
		if !ok {
			logger.LogWarn("Mallboss PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Mallboss==:%+v", vData)
	}
}

func (*MallbossMgr) Load(buffer []byte) bool {
	Mallboss_All = make([]*Mallboss, 0)
	err := json.Unmarshal(buffer, &Mallboss_All)
	if err != nil {
		logger.LogErr("Mallboss JsonFailed:", err)
		return false
	}
	vLen := len(Mallboss_All)
	mallbossDic = make(map[int32]*Mallboss, vLen)
	for _, mem := range Mallboss_All {
		mallbossDic[mem.Id] = mem
	}
	return true
}
