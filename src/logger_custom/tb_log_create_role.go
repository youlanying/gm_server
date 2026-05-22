package logger_custom

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_log_create_role struct {
	logid      string "id"
	roleid     uint64 "角色id"
	account    string "账号"
	serverid   int32  "服务器id"
	ip         string "IP地址"
	mac        string "mac地址"
	devicetype string "设备类型"
	versionno  string "版本"
	cid        string "渠道ID"
	createtime int64  "创建时间"
}

type TB_LOG_CREATE_ROLE struct {
	LOGID      string "id"
	ROLEID     uint64 "角色id"
	ACCOUNT    string "账号"
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
func TB_LOG_CREATE_ROLEReadBylogid(conditions string) (map[string]*TB_LOG_CREATE_ROLE, bool) {
	selectSql := "select * from `tb_log_create_role` where " + conditions
	mapDbData := make(map[string]*TB_LOG_CREATE_ROLE)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_log_create_role ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_create_role{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.roleid, &dbData.account, &dbData.serverid, &dbData.ip, &dbData.mac, &dbData.devicetype, &dbData.versionno, &dbData.cid, &dbData.createtime)
		if err != nil {
			logger.LogErr("get tb_log_create_role error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_create_role(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_CREATE_ROLEReadBySQL(conditions string) (map[string]*TB_LOG_CREATE_ROLE, bool) {
	selectSql := "select * from `tb_log_create_role` " + conditions
	mapDbData := make(map[string]*TB_LOG_CREATE_ROLE)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_log_create_role ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_create_role{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.roleid, &dbData.account, &dbData.serverid, &dbData.ip, &dbData.mac, &dbData.devicetype, &dbData.versionno, &dbData.cid, &dbData.createtime)
		if err != nil {
			logger.LogErr("get sql tb_log_create_role error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_create_role(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_LOG_CREATE_ROLE_ReplaceInto(tb_log_create_role *TB_LOG_CREATE_ROLE) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_log_create_role`(logid,roleid,account,serverid,ip,mac,devicetype,versionno,cid,createtime) values"
	values := "(" + "'" + tb_log_create_role.LOGID + "'" + "," + strconv.FormatUint(tb_log_create_role.ROLEID, 10) + "," + "'" + tb_log_create_role.ACCOUNT + "'" + "," + strconv.FormatInt(int64(tb_log_create_role.SERVERID), 10) + "," + "'" + tb_log_create_role.IP + "'" + "," + "'" + tb_log_create_role.MAC + "'" + "," + "'" + tb_log_create_role.DEVICETYPE + "'" + "," + "'" + tb_log_create_role.VERSIONNO + "'" + "," + "'" + tb_log_create_role.CID + "'" + "," + strconv.FormatInt(tb_log_create_role.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_log_create_role error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_log_create_role failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_CREATE_ROLEUpdateBy(tb_log_create_role *TB_LOG_CREATE_ROLE) (int64, bool) {
	kvData, data := tb_log_create_roleToStr(tb_log_create_role)
	selectSql := "UPDATE `tb_log_create_role` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_log_create_role ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_log_create_role RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_CREATE_ROLEUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_log_create_role` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_log_create_role ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_log_create_role RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Insert
func TB_LOG_CREATE_ROLEInsert(logid string, roleid uint64, account string, serverid int32, ip string, mac string, devicetype string, versionno string, cid string, createtime int64) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_create_role`(logid,roleid,account,serverid,ip,mac,devicetype,versionno,cid,createtime) VALUES"
	values := "(" + "'" + logid + "'" + "," + strconv.FormatUint(roleid, 10) + "," + "'" + account + "'" + "," + strconv.FormatInt(int64(serverid), 10) + "," + "'" + ip + "'" + "," + "'" + mac + "'" + "," + "'" + devicetype + "'" + "," + "'" + versionno + "'" + "," + "'" + cid + "'" + "," + strconv.FormatInt(createtime, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_create_role ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_create_role RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_LOG_CREATE_ROLE_Insert(tb_log_create_role *TB_LOG_CREATE_ROLE) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_create_role`(logid,roleid,account,serverid,ip,mac,devicetype,versionno,cid,createtime) VALUES"
	values := "(" + "'" + tb_log_create_role.LOGID + "'" + "," + strconv.FormatUint(tb_log_create_role.ROLEID, 10) + "," + "'" + tb_log_create_role.ACCOUNT + "'" + "," + strconv.FormatInt(int64(tb_log_create_role.SERVERID), 10) + "," + "'" + tb_log_create_role.IP + "'" + "," + "'" + tb_log_create_role.MAC + "'" + "," + "'" + tb_log_create_role.DEVICETYPE + "'" + "," + "'" + tb_log_create_role.VERSIONNO + "'" + "," + "'" + tb_log_create_role.CID + "'" + "," + strconv.FormatInt(tb_log_create_role.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_create_role ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_create_role RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_LOG_CREATE_ROLEdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_log_create_role` where " + conditions
	result, err := loggerDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_log_create_role failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_log_create_role RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_LOG_CREATE_ROLE(logid string, roleid uint64, account string, serverid int32, ip string, mac string, devicetype string, versionno string, cid string, createtime int64) *TB_LOG_CREATE_ROLE {
	dbData := &TB_LOG_CREATE_ROLE{
		LOGID:      logid,
		ROLEID:     roleid,
		ACCOUNT:    account,
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
func tb_log_create_roleToStr(dbData *TB_LOG_CREATE_ROLE) (key string, data string) {
	key = "logid=" + "'" + dbData.LOGID + "'"
	data = "roleid=" + strconv.FormatUint(dbData.ROLEID, 10) + "," +
		"account=" + "'" + dbData.ACCOUNT + "'" + "," +
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
func strTotb_log_create_role(dbData *tb_log_create_role) *TB_LOG_CREATE_ROLE {
	tbData := &TB_LOG_CREATE_ROLE{
		LOGID:      dbData.logid,
		ROLEID:     dbData.roleid,
		ACCOUNT:    dbData.account,
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
