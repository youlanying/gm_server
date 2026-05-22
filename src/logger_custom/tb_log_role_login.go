package logger_custom

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_log_role_login struct {
	logid      string "id"
	roleid     uint64 "角色id"
	account    string "账号"
	name       string "名字"
	serverid   int32  "服务器id"
	ip         string "IP地址"
	mac        string "mac地址"
	devicetype string "设备类型"
	versionno  string "版本"
	cid        string "渠道ID"
	createtime int64  "创建时间"
}

type TB_LOG_ROLE_LOGIN struct {
	LOGID      string "id"
	ROLEID     uint64 "角色id"
	ACCOUNT    string "账号"
	NAME       string "名字"
	SERVERID   int32  "服务器id"
	IP         string "IP地址"
	MAC        string "mac地址"
	DEVICETYPE string "设备类型"
	VERSIONNO  string "版本"
	CID        string "渠道ID"
	CREATETIME int64  "创建时间"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_ROLE_LOGINReadBylogid(conditions string) (map[string]*TB_LOG_ROLE_LOGIN, bool) {
	selectSql := "select * from `tb_log_role_login` where " + conditions
	mapDbData := make(map[string]*TB_LOG_ROLE_LOGIN)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_log_role_login ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_role_login{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.roleid, &dbData.account, &dbData.name, &dbData.serverid, &dbData.ip, &dbData.mac, &dbData.devicetype, &dbData.versionno, &dbData.cid, &dbData.createtime)
		if err != nil {
			logger.LogErr("get tb_log_role_login error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_role_login(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_ROLE_LOGINReadBySQL(conditions string) (map[string]*TB_LOG_ROLE_LOGIN, bool) {
	selectSql := "select * from `tb_log_role_login` " + conditions
	mapDbData := make(map[string]*TB_LOG_ROLE_LOGIN)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_log_role_login ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_role_login{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.roleid, &dbData.account, &dbData.name, &dbData.serverid, &dbData.ip, &dbData.mac, &dbData.devicetype, &dbData.versionno, &dbData.cid, &dbData.createtime)
		if err != nil {
			logger.LogErr("get sql tb_log_role_login error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_role_login(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_LOG_ROLE_LOGIN_ReplaceInto(tb_log_role_login *TB_LOG_ROLE_LOGIN) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_log_role_login`(logid,roleid,account,name,serverid,ip,mac,devicetype,versionno,cid,createtime) values"
	values := "(" + "'" + tb_log_role_login.LOGID + "'" + "," + strconv.FormatUint(tb_log_role_login.ROLEID, 10) + "," + "'" + tb_log_role_login.ACCOUNT + "'" + "," + "'" + tb_log_role_login.NAME + "'" + "," + strconv.FormatInt(int64(tb_log_role_login.SERVERID), 10) + "," + "'" + tb_log_role_login.IP + "'" + "," + "'" + tb_log_role_login.MAC + "'" + "," + "'" + tb_log_role_login.DEVICETYPE + "'" + "," + "'" + tb_log_role_login.VERSIONNO + "'" + "," + "'" + tb_log_role_login.CID + "'" + "," + strconv.FormatInt(tb_log_role_login.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_log_role_login error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_log_role_login failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_ROLE_LOGINUpdateBy(tb_log_role_login *TB_LOG_ROLE_LOGIN) (int64, bool) {
	kvData, data := tb_log_role_loginToStr(tb_log_role_login)
	selectSql := "UPDATE `tb_log_role_login` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_log_role_login ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_log_role_login RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_ROLE_LOGINUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_log_role_login` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_log_role_login ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_log_role_login RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Insert
func TB_LOG_ROLE_LOGINInsert(logid string, roleid uint64, account string, name string, serverid int32, ip string, mac string, devicetype string, versionno string, cid string, createtime int64) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_role_login`(logid,roleid,account,name,serverid,ip,mac,devicetype,versionno,cid,createtime) VALUES"
	values := "(" + "'" + logid + "'" + "," + strconv.FormatUint(roleid, 10) + "," + "'" + account + "'" + "," + "'" + name + "'" + "," + strconv.FormatInt(int64(serverid), 10) + "," + "'" + ip + "'" + "," + "'" + mac + "'" + "," + "'" + devicetype + "'" + "," + "'" + versionno + "'" + "," + "'" + cid + "'" + "," + strconv.FormatInt(createtime, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_role_login ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_role_login RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_LOG_ROLE_LOGIN_Insert(tb_log_role_login *TB_LOG_ROLE_LOGIN) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_role_login`(logid,roleid,account,name,serverid,ip,mac,devicetype,versionno,cid,createtime) VALUES"
	values := "(" + "'" + tb_log_role_login.LOGID + "'" + "," + strconv.FormatUint(tb_log_role_login.ROLEID, 10) + "," + "'" + tb_log_role_login.ACCOUNT + "'" + "," + "'" + tb_log_role_login.NAME + "'" + "," + strconv.FormatInt(int64(tb_log_role_login.SERVERID), 10) + "," + "'" + tb_log_role_login.IP + "'" + "," + "'" + tb_log_role_login.MAC + "'" + "," + "'" + tb_log_role_login.DEVICETYPE + "'" + "," + "'" + tb_log_role_login.VERSIONNO + "'" + "," + "'" + tb_log_role_login.CID + "'" + "," + strconv.FormatInt(tb_log_role_login.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_role_login ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_role_login RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_LOG_ROLE_LOGINdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_log_role_login` where " + conditions
	result, err := loggerDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_log_role_login failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_log_role_login RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_LOG_ROLE_LOGIN(logid string, roleid uint64, account string, name string, serverid int32, ip string, mac string, devicetype string, versionno string, cid string, createtime int64) *TB_LOG_ROLE_LOGIN {
	dbData := &TB_LOG_ROLE_LOGIN{
		LOGID:      logid,
		ROLEID:     roleid,
		ACCOUNT:    account,
		NAME:       name,
		SERVERID:   serverid,
		IP:         ip,
		MAC:        mac,
		DEVICETYPE: devicetype,
		VERSIONNO:  versionno,
		CID:        cid,
		CREATETIME: createtime,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_log_role_loginToStr(dbData *TB_LOG_ROLE_LOGIN) (key string, data string) {
	key = "logid=" + "'" + dbData.LOGID + "'"
	data = "roleid=" + strconv.FormatUint(dbData.ROLEID, 10) + "," +
		"account=" + "'" + dbData.ACCOUNT + "'" + "," +
		"name=" + "'" + dbData.NAME + "'" + "," +
		"serverid=" + strconv.FormatInt(int64(dbData.SERVERID), 10) + "," +
		"ip=" + "'" + dbData.IP + "'" + "," +
		"mac=" + "'" + dbData.MAC + "'" + "," +
		"devicetype=" + "'" + dbData.DEVICETYPE + "'" + "," +
		"versionno=" + "'" + dbData.VERSIONNO + "'" + "," +
		"cid=" + "'" + dbData.CID + "'" + "," +
		"createtime=" + strconv.FormatInt(dbData.CREATETIME, 10)
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_log_role_login(dbData *tb_log_role_login) *TB_LOG_ROLE_LOGIN {
	tbData := &TB_LOG_ROLE_LOGIN{
		LOGID:      dbData.logid,
		ROLEID:     dbData.roleid,
		ACCOUNT:    dbData.account,
		NAME:       dbData.name,
		SERVERID:   dbData.serverid,
		IP:         dbData.ip,
		MAC:        dbData.mac,
		DEVICETYPE: dbData.devicetype,
		VERSIONNO:  dbData.versionno,
		CID:        dbData.cid,
		CREATETIME: dbData.createtime,
	}
	return tbData
}
