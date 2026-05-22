package logger_custom

import (
	"database/sql"
	"encoding/json"
	"gm_server/src/logger"

	_ "modernc.org/sqlite"
)

var (
	loggerDB      *sql.DB
	DB_FRAME_TIME int64
)

// OpenDB 初始化数据库连接
func OpenDB() {
	var err error
	loggerDB, err = initLoggerSQLiteDB()
	if err != nil {
		logger.LogErr("connect sqlite3 fail ! [%s]", err)
	} else {
		logger.Log("connect to sqlite3 success")
	}
}

//关闭数据库连接
func CloseDB() {

	if loggerDB != nil {
		loggerDB.Close()
		loggerDB = nil
		logger.Logf("Disconnect %v gameDB...")
	}
}

//测试数据库连接
func Testdb() {
	rows, err := loggerDB.Query("select account,roleid from tb_account where account = 10002")
	if err != nil {
		logger.Log("select fail [%s]", err)
	}

	var mapUser map[int]string
	mapUser = make(map[int]string)

	for rows.Next() {
		var id int
		var username string
		rows.Columns()
		err := rows.Scan(&id, &username)
		if err != nil {
			logger.Log("get user info error [%s]", err)
		}
		mapUser[id] = username
	}

	for k, v := range mapUser {
		logger.Log(k, v)
	}
}

func SetDbFrame(t int64) {
	DB_FRAME_TIME = t
}

//// Get_DB_MaxItemId 获取最大的道具ID
//func Get_DB_MaxItemId() uint64 {
//	var maxitemid uint64
//	rows := gameDB.QueryRow("SELECT MAX(itemid) AS maxitemid FROM `tb_role_item`")
//	err1 := rows.Scan(&maxitemid)
//	if err1 != nil {
//		logger.LogErr("get user info error [%s]", err1)
//		return 0
//	} else {
//		return maxitemid
//	}
//}

//interface ToJson
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

//string Json To []int32
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

//string Json To map[int32]int32
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

//string Json To map[string]int32
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

//string Json To map[int32]string
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

//string Json To map[string]string
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
