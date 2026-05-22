package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Layer 原表为rogue随机迷宫表格.xlsx 的子表 迷宫层数表
type Layer struct {
	Id         int32   `json:"id"`         //层数ID
	Weeks_id   int32   `json:"weeks_id"`   //周目ID
	Room_type  []int32 `json:"room_type"`  //当前层数可随机房间类型
	Next_layer int32   `json:"next_layer"` //后置层数ID
}

type LayerMgr struct {
}

var (
	Layer_Model LayerMgr
	layerDic    map[int32]*Layer
	// Layer_All rogue随机迷宫表格.xlsx (迷宫层数表)
	Layer_All []*Layer
)

// Layer_Get rogue随机迷宫表格.xlsx (迷宫层数表)
func Layer_Get(Id int32) (*Layer, bool) {
	data, ok := layerDic[Id]
	if !ok {
		PROTO_ERROR_ID = "rogue随机迷宫表格.xlsx\nlayer not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *LayerMgr) PrintArr() {
	vLen := len(Layer_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*LayerMgr) PrintArrOne(index int) {
	logger.Logf("==Layer==:%+v", Layer_All[index])
}

func (*LayerMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := layerDic[int32Key]
		if !ok {
			logger.LogWarn("Layer PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Layer==:%+v", vData)
	}
}

func (*LayerMgr) Load(buffer []byte) bool {
	Layer_All = make([]*Layer, 0)
	err := json.Unmarshal(buffer, &Layer_All)
	if err != nil {
		logger.LogErr("Layer JsonFailed:", err)
		return false
	}
	vLen := len(Layer_All)
	layerDic = make(map[int32]*Layer, vLen)
	for _, mem := range Layer_All {
		layerDic[mem.Id] = mem
	}
	return true
}
