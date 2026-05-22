package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Sceneitemdata 原表为关卡表-交互物件数据表.xlsx 的子表 Data
type Sceneitemdata struct {
	Id              int32   `json:"id"`              //交互物件模型id
	Drop            int32   `json:"drop"`            //掉落ID
	DropType        int32   `json:"dropType"`        //掉落类型
	DropParm        []int32 `json:"dropParm"`        //掉落参数
	Drop_coineffect int32   `json:"drop_coineffect"` //掉落金币特效/血球特效
	Drop_geteffect  int32   `json:"drop_geteffect"`  //获得掉落物品特效
}

type SceneitemdataMgr struct {
}

var (
	Sceneitemdata_Model SceneitemdataMgr
	sceneitemdataDic    map[int32]*Sceneitemdata
	// Sceneitemdata_All 关卡表-交互物件数据表.xlsx (Data)
	Sceneitemdata_All []*Sceneitemdata
)

// Sceneitemdata_Get 关卡表-交互物件数据表.xlsx (Data)
func Sceneitemdata_Get(Id int32) (*Sceneitemdata, bool) {
	data, ok := sceneitemdataDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡表-交互物件数据表.xlsx\nsceneitemdata not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *SceneitemdataMgr) PrintArr() {
	vLen := len(Sceneitemdata_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*SceneitemdataMgr) PrintArrOne(index int) {
	logger.Logf("==Sceneitemdata==:%+v", Sceneitemdata_All[index])
}

func (*SceneitemdataMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := sceneitemdataDic[int32Key]
		if !ok {
			logger.LogWarn("Sceneitemdata PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Sceneitemdata==:%+v", vData)
	}
}

func (*SceneitemdataMgr) Load(buffer []byte) bool {
	Sceneitemdata_All = make([]*Sceneitemdata, 0)
	err := json.Unmarshal(buffer, &Sceneitemdata_All)
	if err != nil {
		logger.LogErr("Sceneitemdata JsonFailed:", err)
		return false
	}
	vLen := len(Sceneitemdata_All)
	sceneitemdataDic = make(map[int32]*Sceneitemdata, vLen)
	for _, mem := range Sceneitemdata_All {
		sceneitemdataDic[mem.Id] = mem
	}
	return true
}
