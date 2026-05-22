package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Newbuff 原表为新BUFF表.xlsx 的子表 BUFF
type Newbuff struct {
	Id             int32     `json:"id"`             //id
	Name           string    `json:"name"`           //名称
	Tag            []int32   `json:"tag"`            //标签
	IfUniq         int32     `json:"ifUniq"`         //叠加互斥类型
	IfUniqParam1   int32     `json:"ifUniqParam1"`   //叠加互斥参数1
	IfUniqParam2   []int32   `json:"ifUniqParam2"`   //叠加互斥参数2
	TriggerID      int32     `json:"triggerID"`      //触发器
	TriggerParam   [][]int32 `json:"triggerParam"`   //触发条件
	DeleterID      []int32   `json:"DeleterID"`      //移除器
	EffectTimes    int32     `json:"effectTimes"`    //触发次数
	Duration       int32     `json:"duration"`       //持续时间（毫秒）
	StepTime       int32     `json:"stepTime"`       //触发间隔（毫秒）
	ScreenId       int32     `json:"ScreenId"`       //筛选id
	Effectresultid []int32   `json:"effectresultid"` //效果id
	HitRecoverID   int32     `json:"HitRecoverID"`   //受击动作表现id
}

type NewbuffMgr struct {
}

var (
	Newbuff_Model NewbuffMgr
	newbuffDic    map[int32]*Newbuff
	// Newbuff_All 新BUFF表.xlsx (BUFF)
	Newbuff_All []*Newbuff
)

// Newbuff_Get 新BUFF表.xlsx (BUFF)
func Newbuff_Get(Id int32) (*Newbuff, bool) {
	data, ok := newbuffDic[Id]
	if !ok {
		PROTO_ERROR_ID = "新BUFF表.xlsx\nnewbuff not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *NewbuffMgr) PrintArr() {
	vLen := len(Newbuff_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*NewbuffMgr) PrintArrOne(index int) {
	logger.Logf("==Newbuff==:%+v", Newbuff_All[index])
}

func (*NewbuffMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := newbuffDic[int32Key]
		if !ok {
			logger.LogWarn("Newbuff PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Newbuff==:%+v", vData)
	}
}

func (*NewbuffMgr) Load(buffer []byte) bool {
	Newbuff_All = make([]*Newbuff, 0)
	err := json.Unmarshal(buffer, &Newbuff_All)
	if err != nil {
		logger.LogErr("Newbuff JsonFailed:", err)
		return false
	}
	vLen := len(Newbuff_All)
	newbuffDic = make(map[int32]*Newbuff, vLen)
	for _, mem := range Newbuff_All {
		newbuffDic[mem.Id] = mem
	}
	return true
}
