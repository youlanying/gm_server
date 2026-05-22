package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Roomtype 原表为rogue随机迷宫表格.xlsx 的子表 房间类型表
type Roomtype struct {
	Id         int32   `json:"id"`         //房间类型ID
	Sceneslist []int32 `json:"sceneslist"` //当前房间类型可随机场景
}

type RoomtypeMgr struct {
}

var (
	Roomtype_Model RoomtypeMgr
	roomtypeDic    map[int32]*Roomtype
	// Roomtype_All rogue随机迷宫表格.xlsx (房间类型表)
	Roomtype_All []*Roomtype
)

// Roomtype_Get rogue随机迷宫表格.xlsx (房间类型表)
func Roomtype_Get(Id int32) (*Roomtype, bool) {
	data, ok := roomtypeDic[Id]
	if !ok {
		PROTO_ERROR_ID = "rogue随机迷宫表格.xlsx\nroomtype not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *RoomtypeMgr) PrintArr() {
	vLen := len(Roomtype_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*RoomtypeMgr) PrintArrOne(index int) {
	logger.Logf("==Roomtype==:%+v", Roomtype_All[index])
}

func (*RoomtypeMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := roomtypeDic[int32Key]
		if !ok {
			logger.LogWarn("Roomtype PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Roomtype==:%+v", vData)
	}
}

func (*RoomtypeMgr) Load(buffer []byte) bool {
	Roomtype_All = make([]*Roomtype, 0)
	err := json.Unmarshal(buffer, &Roomtype_All)
	if err != nil {
		logger.LogErr("Roomtype JsonFailed:", err)
		return false
	}
	vLen := len(Roomtype_All)
	roomtypeDic = make(map[int32]*Roomtype, vLen)
	for _, mem := range Roomtype_All {
		roomtypeDic[mem.Id] = mem
	}
	return true
}
