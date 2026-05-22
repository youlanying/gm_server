package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Mall_paging 原表为商城表.xlsx 的子表 商店分页
type Mall_paging struct {
	Id           int32   `json:"id"`           //商店分页ID商店类型分页ID1：金币分页2：钻石分页3：代币分页根据分页ID顺序，在商店中进行分页排序
	Refresh      int32   `json:"refresh"`      //商店是否可刷新1：为刷新0：为不可刷新
	Function     int32   `json:"function"`     //功能开启ID跟随功能开启从而开启商店0为当前商店无需跟随功能开启填写功能开启表中的功能ID
	Refresh_time []int32 `json:"refresh_time"` //刷新时间[5,8,10,15,19]刷新时间通过数组配置第1数字为当天第1次刷新时间最后数字为当天最后刷新时间-1：为不可刷新
	Drop         []int32 `json:"drop"`         //掉落ID填写《掉落表》中掉落ID在《掉落表》进行商品的随机选择与等级区间掉落ID组成[101,102,103]
	Activity     int32   `json:"activity"`     //活动ID根据运营活动表中的活动ID，判断限时商店的出现与消失当前商店不限时填0
	Consume      []int32 `json:"consume"`      //消耗货币数量当前刷新所消耗的代币数量
	Currency     int32   `json:"currency"`     //刷新货币类型道具ID
}

type Mall_pagingMgr struct {
}

var (
	Mall_paging_Model Mall_pagingMgr
	mall_pagingDic    map[int32]*Mall_paging
	// Mall_paging_All 商城表.xlsx (商店分页)
	Mall_paging_All []*Mall_paging
)

// Mall_paging_Get 商城表.xlsx (商店分页)
func Mall_paging_Get(Id int32) (*Mall_paging, bool) {
	data, ok := mall_pagingDic[Id]
	if !ok {
		PROTO_ERROR_ID = "商城表.xlsx\nmall_paging not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Mall_pagingMgr) PrintArr() {
	vLen := len(Mall_paging_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Mall_pagingMgr) PrintArrOne(index int) {
	logger.Logf("==Mall_paging==:%+v", Mall_paging_All[index])
}

func (*Mall_pagingMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := mall_pagingDic[int32Key]
		if !ok {
			logger.LogWarn("Mall_paging PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Mall_paging==:%+v", vData)
	}
}

func (*Mall_pagingMgr) Load(buffer []byte) bool {
	Mall_paging_All = make([]*Mall_paging, 0)
	err := json.Unmarshal(buffer, &Mall_paging_All)
	if err != nil {
		logger.LogErr("Mall_paging JsonFailed:", err)
		return false
	}
	vLen := len(Mall_paging_All)
	mall_pagingDic = make(map[int32]*Mall_paging, vLen)
	for _, mem := range Mall_paging_All {
		mall_pagingDic[mem.Id] = mem
	}
	return true
}
