package beegodb

import (
	"database/sql"
	"encoding/json"
	"gm_server/src/beegoWeb/table"
	"gm_server/src/logger"
	"gm_server/src/tools"

	_ "modernc.org/sqlite"
)

var (
	beegoDB *sql.DB
)

var (
	//缓存用户数据
	TbAdminUserAllData map[string]*TB_ADMIN_USER
	TbAdminUserIntMap  map[int32]*TB_ADMIN_USER
	//缓存分组数据库
	TbAdminGroupAllData map[int32]*TB_ADMIN_GROUP
	//缓存中心服(平台)列表
	TbCenterListAllData map[int32]*TB_CENTER_LIST
)

// 初始化数据库连接
func OpenDB() {
	var err error
	beegoDB, err = initBeegoSQLiteDB()
	if err != nil {
		logger.LogErr("connect sqlite3 fail ! [%s]", err)
	} else {
		logger.Log("connect to sqlite3 success")
		InitTbCenterList()
		InitTbAdminUser()
		InitTbAdminGroup()
	}
}

// 关闭数据库连接
func CloseDB() {
	if beegoDB != nil {
		beegoDB.Close()
		beegoDB = nil
		logger.Logf("Disconnect %v gameDB...")
	}
}

// InitTbAdminUser 初始化用户数据库表
func InitTbAdminUser() {
	TbAdminUserAllData = make(map[string]*TB_ADMIN_USER)
	TbAdminUserIntMap = make(map[int32]*TB_ADMIN_USER)
	all, ok := TB_ADMIN_USERReadBySQL("")
	if len(all) == 0 && ok {
		userData := MakeTB_ADMIN_USER(1, "admin", "21232f297a57a5a743894a0e4a801fc3", 1, "")
		TB_ADMIN_USER_Insert(userData)
		TbAdminUserAllData["admin"] = userData
		return
	}
	if !ok {
		return
	}
	TbAdminUserIntMap = all
	for _, i2 := range all {
		TbAdminUserAllData[i2.USERID] = i2
	}
}

// InitTbAdminGroup 初始化权限组数据库表
func InitTbAdminGroup() {
	TbAdminGroupAllData = make(map[int32]*TB_ADMIN_GROUP)
	all, ok := TB_ADMIN_GROUPReadBySQL("")
	if len(all) == 0 && ok {
		actionList := make([]int32, 0)
		for _, i2 := range table.Admin_action_All {
			actionList = append(actionList, i2.Id)
		}
		for i, _ := range TbCenterListAllData {
			actionList = append(actionList, i)
		}
		userData := MakeTB_ADMIN_GROUP(1, "超级管理员", actionList)
		TB_ADMIN_GROUP_Insert(userData)
		TbAdminGroupAllData[1] = userData
		return
	}
	if !ok {
		return
	}
	TbAdminGroupAllData = all
}

// InitTbCenterList 中心服务器列表 平台列表
func InitTbCenterList() {
	all, ok := TB_CENTER_LISTReadBySQL("")
	TbCenterListAllData = make(map[int32]*TB_CENTER_LIST)
	if len(all) == 0 && ok {
		centerData := MakeTB_CENTER_LIST(1000, "57开发服", "192.168.20.57", "2023", "")
		TB_CENTER_LIST_Insert(centerData)
		TbCenterListAllData[1000] = centerData
		return
	}
	if !ok {
		return
	}
	TbCenterListAllData = all
}

// 获取 中心服务器列表 平台列表 最大ID
func GetMaxCenterId() int32 {
	n := int32(1000)
	for i, _ := range TbCenterListAllData {
		if i > n {
			n = i
		}
	}
	return n
}

// UpdateSuperAdminAuthority 更新超级管理员权限
func UpdateSuperAdminAuthority(groupId int32) {
	Group, ok := TbAdminGroupAllData[groupId]
	if ok {
		aCTIONLIST := Group.ACTIONLIST
		for _, centerData := range TbCenterListAllData {
			isMb := tools.ListsMember(centerData.ID, aCTIONLIST)
			if !isMb {
				aCTIONLIST = append(aCTIONLIST, centerData.ID)
			}
		}
		Group.ACTIONLIST = aCTIONLIST
		TbAdminGroupAllData[groupId] = Group
		TB_ADMIN_GROUPUpdateBy(Group)
	}
}

// UpdateTbAdminGroup 更新权限组数据库表
func UpdateTbAdminGroup(oneGroup *TB_ADMIN_GROUP) {
	row, ok := TB_ADMIN_GROUPUpdateBy(oneGroup)
	logger.Logf("===UpdateTbAdminGroup==err:%v==row:%v", ok, row)
	if !ok {
		logger.LogErrf("===UpdateTbAdminGroup==err:%v", ok)
		return
	}
	InitTbAdminGroup()
}

func DeleteGroupId(gId int32) {
	for _, group := range TbAdminGroupAllData {
		acList := group.ACTIONLIST
		isMb := tools.ListsMember(gId, acList)
		if isMb {
			newAcList := tools.ListsDelete(gId, acList)
			group.ACTIONLIST = newAcList
			TB_ADMIN_GROUPUpdateBy(group)
		}
	}
	InitTbAdminGroup()
}

// interface ToJson
func ToJson(v interface{}) string {
	byteJson, errs := json.Marshal(v) //转换成JSON返回的是byte[]
	if errs != nil {
		logger.LogErr(" SaveJson ERROR", errs.Error())
		return ""
	} else {
		strJson := "'" + string(byteJson) + "'"
		return strJson
	}
}

// string Json To []int32
func JsonToIntList(v string) []int32 {
	var intList []int32
	if v == "" || v == "[]" || v == "{}" {
		return intList
	}
	buffer := []byte(v)
	err := json.Unmarshal(buffer, &intList)
	if err != nil {
		logger.LogErr("JsonFailed :", err)
		return intList
	}
	return intList
}

// string Json To map[int32]int32
func JsonToIntIntMap(v string) map[int32]int32 {
	mapData := make(map[int32]int32)
	if v == "" || v == "[]" || v == "{}" {
		return mapData
	}
	buffer := []byte(v)
	err := json.Unmarshal(buffer, &mapData)
	if err != nil {
		logger.LogErr("JsonFailed", err)
		return mapData
	}
	return mapData
}

// string Json To map[string]int32
func JsonToStringIntMap(v string) map[string]int32 {
	mapData := make(map[string]int32)
	if v == "" || v == "[]" || v == "{}" {
		return mapData
	}
	buffer := []byte(v)
	err := json.Unmarshal(buffer, &mapData)
	if err != nil {
		logger.LogErr("JsonFailed", err)
		return mapData
	}
	return mapData
}

// string Json To map[int32]string
func JsonToIntStringMap(v string) map[int32]string {
	mapData := make(map[int32]string)
	if v == "" || v == "[]" || v == "{}" {
		return mapData
	}
	buffer := []byte(v)
	err := json.Unmarshal(buffer, &mapData)
	if err != nil {
		logger.LogErr("JsonFailed", err)
		return mapData
	}
	return mapData
}

// string Json To map[string]string
func JsonToStringStringMap(v string) map[string]string {
	mapData := make(map[string]string)
	if v == "" || v == "[]" || v == "{}" {
		return mapData
	}
	buffer := []byte(v)
	err := json.Unmarshal(buffer, &mapData)
	if err != nil {
		logger.LogErr("JsonFailed", err)
		return mapData
	}
	return mapData
}
