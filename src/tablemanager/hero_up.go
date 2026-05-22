package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Hero_up 原表为经验表.xlsx 的子表 经验表
type Hero_up struct {
	Id                int32 `json:"id"`                //等级/星级/
	Role_up_exp       int32 `json:"role_up_exp"`       //消耗经验
	Role_add_physical int32 `json:"role_add_physical"` //玩家升级奖励体力
	Role_physical_max int32 `json:"role_physical_max"` //玩家升级体力上限
	Hero_level_limit  int32 `json:"hero_level_limit"`  //玩家升级角色等级上限
	Up_exp            int32 `json:"up_exp"`            //角色升级经验
	Star_item         int32 `json:"star_item"`         //角色升星消耗碎片
	Star_money        int32 `json:"star_money"`        //角色升星消耗金币
	Intimacy_exp      int32 `json:"intimacy_exp"`      //角色亲密度升级经验
	Equip_b_exp       int32 `json:"equip_b_exp"`       //B级装备升级经验
	Equip_a_exp       int32 `json:"equip_a_exp"`       //A级装备升级经验
	Equip_s_exp       int32 `json:"equip_s_exp"`       //S级装备升级经验
	Exchange_item     int32 `json:"exchange_item"`     //代币兑换碎片价格
	Rmb_physical      int32 `json:"rmb_physical"`      //钻石购买体力价格
	A_equip_stone     int32 `json:"a_equip_stone"`     //A级装备升级强化石
	B_equip_stone     int32 `json:"b_equip_stone"`     //B级装备升级强化石
	S_equip_stone     int32 `json:"s_equip_stone"`     //S级装备升级强化石
	A_equip_gold      int32 `json:"a_equip_gold"`      //A级装备升级金币数量
	B_equip_gold      int32 `json:"b_equip_gold"`      //B级装备升级金币数量
	S_equip_gold      int32 `json:"s_equip_gold"`      //S级装备升级金币数量
	Equip_skilllevel  int32 `json:"equip_skilllevel"`  //装备技能等级
}

type Hero_upMgr struct {
}

var (
	Hero_up_Model Hero_upMgr
	hero_upDic    map[int32]*Hero_up
	// Hero_up_All 经验表.xlsx (经验表)
	Hero_up_All []*Hero_up
)

// Hero_up_Get 经验表.xlsx (经验表)
func Hero_up_Get(Id int32) (*Hero_up, bool) {
	data, ok := hero_upDic[Id]
	if !ok {
		PROTO_ERROR_ID = "经验表.xlsx\nhero_up not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Hero_upMgr) PrintArr() {
	vLen := len(Hero_up_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Hero_upMgr) PrintArrOne(index int) {
	logger.Logf("==Hero_up==:%+v", Hero_up_All[index])
}

func (*Hero_upMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := hero_upDic[int32Key]
		if !ok {
			logger.LogWarn("Hero_up PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Hero_up==:%+v", vData)
	}
}

func (*Hero_upMgr) Load(buffer []byte) bool {
	Hero_up_All = make([]*Hero_up, 0)
	err := json.Unmarshal(buffer, &Hero_up_All)
	if err != nil {
		logger.LogErr("Hero_up JsonFailed:", err)
		return false
	}
	vLen := len(Hero_up_All)
	hero_upDic = make(map[int32]*Hero_up, vLen)
	for _, mem := range Hero_up_All {
		hero_upDic[mem.Id] = mem
	}
	return true
}
