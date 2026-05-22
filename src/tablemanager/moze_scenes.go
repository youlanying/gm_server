package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Moze_scenes 原表为rogue随机迷宫表格.xlsx 的子表 迷宫场景表
type Moze_scenes struct {
	Id                  int32     `json:"id"`                  //迷宫场景ID
	Room_id             int32     `json:"room_id"`             //所属房间类型ID
	Crossing_num        int32     `json:"crossing_num"`        //路口数量
	Hidebox_probability int32     `json:"hidebox_probability"` //隐藏房间出现几率
	Over_id             int32     `json:"over_id"`             //场景通关条件
	Refreshboxid        [][]int32 `json:"refreshboxid"`        //刷怪盒子ID
}

type Moze_scenesMgr struct {
}

var (
	Moze_scenes_Model Moze_scenesMgr
	moze_scenesDic    map[int32]*Moze_scenes
	// Moze_scenes_All rogue随机迷宫表格.xlsx (迷宫场景表)
	Moze_scenes_All []*Moze_scenes
)

// Moze_scenes_Get rogue随机迷宫表格.xlsx (迷宫场景表)
func Moze_scenes_Get(Id int32) (*Moze_scenes, bool) {
	data, ok := moze_scenesDic[Id]
	if !ok {
		PROTO_ERROR_ID = "rogue随机迷宫表格.xlsx\nmoze_scenes not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Moze_scenesMgr) PrintArr() {
	vLen := len(Moze_scenes_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Moze_scenesMgr) PrintArrOne(index int) {
	logger.Logf("==Moze_scenes==:%+v", Moze_scenes_All[index])
}

func (*Moze_scenesMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := moze_scenesDic[int32Key]
		if !ok {
			logger.LogWarn("Moze_scenes PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Moze_scenes==:%+v", vData)
	}
}

func (*Moze_scenesMgr) Load(buffer []byte) bool {
	Moze_scenes_All = make([]*Moze_scenes, 0)
	err := json.Unmarshal(buffer, &Moze_scenes_All)
	if err != nil {
		logger.LogErr("Moze_scenes JsonFailed:", err)
		return false
	}
	vLen := len(Moze_scenes_All)
	moze_scenesDic = make(map[int32]*Moze_scenes, vLen)
	for _, mem := range Moze_scenes_All {
		moze_scenesDic[mem.Id] = mem
	}
	return true
}
