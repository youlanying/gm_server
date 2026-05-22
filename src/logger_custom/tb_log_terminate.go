package logger_custom

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_log_terminate struct {
	logid      string "id"
	roleid     uint64 "角色id"
	account    string "账号"
	name       string "名字"
	serverid   int32  "服务器id"
	errorproto string "可能引起报错的数据表及Id"
	reason     string "原因"
	createtime int64  "创建时间"
}

type TB_LOG_TERMINATE struct {
	LOGID      string "id"
	ROLEID     uint64 "角色id"
	ACCOUNT    string "账号"
	NAME       string "名字"
	SERVERID   int32  "服务器id"
	ERRORPROTO string "可能引起报错的数据表及Id"
	REASON     string "原因"
	CREATETIME int64  "创建时间"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_TERMINATEReadBylogid(conditions string) (map[string]*TB_LOG_TERMINATE, bool) {
	selectSql := "select * from `tb_log_terminate` where " + conditions
	mapDbData := make(map[string]*TB_LOG_TERMINATE)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_log_terminate ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_terminate{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.roleid, &dbData.account, &dbData.name, &dbData.serverid, &dbData.errorproto, &dbData.reason, &dbData.createtime)
		if err != nil {
			logger.LogErr("get tb_log_terminate error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_terminate(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_TERMINATEReadBySQL(conditions string) (map[string]*TB_LOG_TERMINATE, bool) {
	selectSql := "select * from `tb_log_terminate` " + conditions
	mapDbData := make(map[string]*TB_LOG_TERMINATE)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_log_terminate ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_terminate{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.roleid, &dbData.account, &dbData.name, &dbData.serverid, &dbData.errorproto, &dbData.reason, &dbData.createtime)
		if err != nil {
			logger.LogErr("get sql tb_log_terminate error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_terminate(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_LOG_TERMINATE_ReplaceInto(tb_log_terminate *TB_LOG_TERMINATE) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_log_terminate`(logid,roleid,account,name,serverid,errorproto,reason,createtime) values"
	values := "(" + "'" + tb_log_terminate.LOGID + "'" + "," + strconv.FormatUint(tb_log_terminate.ROLEID, 10) + "," + "'" + tb_log_terminate.ACCOUNT + "'" + "," + "'" + tb_log_terminate.NAME + "'" + "," + strconv.FormatInt(int64(tb_log_terminate.SERVERID), 10) + "," + "'" + tb_log_terminate.ERRORPROTO + "'" + "," + "'" + tb_log_terminate.REASON + "'" + "," + strconv.FormatInt(tb_log_terminate.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_log_terminate error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_log_terminate failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_TERMINATEUpdateBy(tb_log_terminate *TB_LOG_TERMINATE) (int64, bool) {
	kvData, data := tb_log_terminateToStr(tb_log_terminate)
	selectSql := "UPDATE `tb_log_terminate` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_log_terminate ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_log_terminate RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_TERMINATEUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_log_terminate` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_log_terminate ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_log_terminate RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Insert
func TB_LOG_TERMINATEInsert(logid string, roleid uint64, account string, name string, serverid int32, errorproto string, reason string, createtime int64) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_terminate`(logid,roleid,account,name,serverid,errorproto,reason,createtime) VALUES"
	values := "(" + "'" + logid + "'" + "," + strconv.FormatUint(roleid, 10) + "," + "'" + account + "'" + "," + "'" + name + "'" + "," + strconv.FormatInt(int64(serverid), 10) + "," + "'" + errorproto + "'" + "," + "'" + reason + "'" + "," + strconv.FormatInt(createtime, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_terminate ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_terminate RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_LOG_TERMINATE_Insert(tb_log_terminate *TB_LOG_TERMINATE) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_terminate`(logid,roleid,account,name,serverid,errorproto,reason,createtime) VALUES"
	values := "(" + "'" + tb_log_terminate.LOGID + "'" + "," + strconv.FormatUint(tb_log_terminate.ROLEID, 10) + "," + "'" + tb_log_terminate.ACCOUNT + "'" + "," + "'" + tb_log_terminate.NAME + "'" + "," + strconv.FormatInt(int64(tb_log_terminate.SERVERID), 10) + "," + "'" + tb_log_terminate.ERRORPROTO + "'" + "," + "'" + tb_log_terminate.REASON + "'" + "," + strconv.FormatInt(tb_log_terminate.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_terminate ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_terminate RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_LOG_TERMINATEdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_log_terminate` where " + conditions
	result, err := loggerDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_log_terminate failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_log_terminate RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_LOG_TERMINATE(logid string, roleid uint64, account string, name string, serverid int32, errorproto string, reason string, createtime int64) *TB_LOG_TERMINATE {
	dbData := &TB_LOG_TERMINATE{
		LOGID:      logid,
		ROLEID:     roleid,
		ACCOUNT:    account,
		NAME:       name,
		SERVERID:   serverid,
		ERRORPROTO: errorproto,
		REASON:     reason,
		CREATETIME: createtime,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_log_terminateToStr(dbData *TB_LOG_TERMINATE) (key string, data string) {
	key = "logid=" + "'" + dbData.LOGID + "'"
	data = "roleid=" + strconv.FormatUint(dbData.ROLEID, 10) + "," +
		"account=" + "'" + dbData.ACCOUNT + "'" + "," +
		"name=" + "'" + dbData.NAME + "'" + "," +
		"serverid=" + strconv.FormatInt(int64(dbData.SERVERID), 10) + "," +
		"errorproto=" + "'" + dbData.ERRORPROTO + "'" + "," +
		"reason=" + "'" + dbData.REASON + "'" + "," +
		"createtime=" + strconv.FormatInt(dbData.CREATETIME, 10)
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_log_terminate(dbData *tb_log_terminate) *TB_LOG_TERMINATE {
	tbData := &TB_LOG_TERMINATE{
		LOGID:      dbData.logid,
		ROLEID:     dbData.roleid,
		ACCOUNT:    dbData.account,
		NAME:       dbData.name,
		SERVERID:   dbData.serverid,
		ERRORPROTO: dbData.errorproto,
		REASON:     dbData.reason,
		CREATETIME: dbData.createtime,
	}
	return tbData
}
