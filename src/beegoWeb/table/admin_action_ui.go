package table

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Admin_action_ui 原表为AdminTooles管理权限表.xlsx 的子表 侧边栏UI
type Admin_action_ui struct {
	Id     int32   `json:"id"`     //父li的ID
	Idlist []int32 `json:"idlist"` //折叠栏内子li的ID
	Isfold string  `json:"isfold"` //是否折叠
}

type Admin_action_uiMgr struct {
}

var (
	Admin_action_ui_Model Admin_action_uiMgr
	admin_action_uiDic    map[int32]*Admin_action_ui
	// Admin_action_ui_All AdminTooles管理权限表.xlsx (侧边栏UI)
	Admin_action_ui_All []*Admin_action_ui
)

// Admin_action_ui_Get AdminTooles管理权限表.xlsx (侧边栏UI)
func Admin_action_ui_Get(Id int32) (*Admin_action_ui, bool) {
	data, ok := admin_action_uiDic[Id]
	if !ok {
		PROTO_ERROR_ID = "AdminTooles管理权限表.xlsx\nadmin_action_ui not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Admin_action_uiMgr) PrintArr() {
	vLen := len(Admin_action_ui_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Admin_action_uiMgr) PrintArrOne(index int) {
	logger.Logf("==Admin_action_ui==:%+v", Admin_action_ui_All[index])
}

func (*Admin_action_uiMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := admin_action_uiDic[int32Key]
		if !ok {
			logger.LogWarn("Admin_action_ui PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Admin_action_ui==:%+v", vData)
	}
}

func (*Admin_action_uiMgr) Load(buffer []byte) bool {
	Admin_action_ui_All = make([]*Admin_action_ui, 0)
	err := json.Unmarshal(buffer, &Admin_action_ui_All)
	if err != nil {
		logger.LogErr("Admin_action_ui JsonFailed:", err)
		return false
	}
	vLen := len(Admin_action_ui_All)
	admin_action_uiDic = make(map[int32]*Admin_action_ui, vLen)
	for _, mem := range Admin_action_ui_All {
		admin_action_uiDic[mem.Id] = mem
	}
	return true
}
