package table

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Admin_action 原表为AdminTooles管理权限表.xlsx 的子表 权限ID
type Admin_action struct {
	Id   int32  `json:"id"`   //权限ID
	Name string `json:"name"` //权限名称
	Href string `json:"href"` //链接路径
	Icon string `json:"icon"` //图标
}

type Admin_actionMgr struct {
}

var (
	Admin_action_Model Admin_actionMgr
	admin_actionDic    map[int32]*Admin_action
	// Admin_action_All AdminTooles管理权限表.xlsx (权限ID)
	Admin_action_All []*Admin_action
)

// Admin_action_Get AdminTooles管理权限表.xlsx (权限ID)
func Admin_action_Get(Id int32) (*Admin_action, bool) {
	data, ok := admin_actionDic[Id]
	if !ok {
		PROTO_ERROR_ID = "AdminTooles管理权限表.xlsx\nadmin_action not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Admin_actionMgr) PrintArr() {
	vLen := len(Admin_action_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Admin_actionMgr) PrintArrOne(index int) {
	logger.Logf("==Admin_action==:%+v", Admin_action_All[index])
}

func (*Admin_actionMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := admin_actionDic[int32Key]
		if !ok {
			logger.LogWarn("Admin_action PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Admin_action==:%+v", vData)
	}
}

func (*Admin_actionMgr) Load(buffer []byte) bool {
	Admin_action_All = make([]*Admin_action, 0)
	err := json.Unmarshal(buffer, &Admin_action_All)
	if err != nil {
		logger.LogErr("Admin_action JsonFailed:", err)
		return false
	}
	vLen := len(Admin_action_All)
	admin_actionDic = make(map[int32]*Admin_action, vLen)
	for _, mem := range Admin_action_All {
		admin_actionDic[mem.Id] = mem
	}
	return true
}
