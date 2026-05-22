package logger_custom

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_log_instance_join struct {
	logid      string "id"
	roleid     uint64 "角色id"
	account    string "账号"
	serverid   int32  "服务器id"
	instanceid int32  "关卡ID"
	activityid int32  "活动ID"
	daytimes   int32  "次数"
	star       string "星级"
	createtime int64  "创建时间"
}

type TB_LOG_INSTANCE_JOIN struct {
	LOGID      string  "id"
	ROLEID     uint64  "角色id"
	ACCOUNT    string  "账号"
	SERVERID   int32   "服务器id"
	INSTANCEID int32   "关卡ID"
	ACTIVITYID int32   "活动ID"
	DAYTIMES   int32   "次数"
	STAR       []int32 "星级"
	CREATETIME int64   "创建时间"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_INSTANCE_JOINReadBylogid(conditions string) (map[string]*TB_LOG_INSTANCE_JOIN, bool) {
	selectSql := "select * from `tb_log_instance_join` where " + conditions
	mapDbData := make(map[string]*TB_LOG_INSTANCE_JOIN)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_log_instance_join ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_instance_join{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.roleid, &dbData.account, &dbData.serverid, &dbData.instanceid, &dbData.activityid, &dbData.daytimes, &dbData.star, &dbData.createtime)
		if err != nil {
			logger.LogErr("get tb_log_instance_join error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_instance_join(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_INSTANCE_JOINReadBySQL(conditions string) (map[string]*TB_LOG_INSTANCE_JOIN, bool) {
	selectSql := "select * from `tb_log_instance_join` " + conditions
	mapDbData := make(map[string]*TB_LOG_INSTANCE_JOIN)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_log_instance_join ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_instance_join{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.roleid, &dbData.account, &dbData.serverid, &dbData.instanceid, &dbData.activityid, &dbData.daytimes, &dbData.star, &dbData.createtime)
		if err != nil {
			logger.LogErr("get sql tb_log_instance_join error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_instance_join(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_LOG_INSTANCE_JOIN_ReplaceInto(tb_log_instance_join *TB_LOG_INSTANCE_JOIN) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_log_instance_join`(logid,roleid,account,serverid,instanceid,activityid,daytimes,star,createtime) values"
	values := "(" + "'" + tb_log_instance_join.LOGID + "'" + "," + strconv.FormatUint(tb_log_instance_join.ROLEID, 10) + "," + "'" + tb_log_instance_join.ACCOUNT + "'" + "," + strconv.FormatInt(int64(tb_log_instance_join.SERVERID), 10) + "," + strconv.FormatInt(int64(tb_log_instance_join.INSTANCEID), 10) + "," + strconv.FormatInt(int64(tb_log_instance_join.ACTIVITYID), 10) + "," + strconv.FormatInt(int64(tb_log_instance_join.DAYTIMES), 10) + "," + ToJson(tb_log_instance_join.STAR) + "," + strconv.FormatInt(tb_log_instance_join.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_log_instance_join error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_log_instance_join failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_INSTANCE_JOINUpdateBy(tb_log_instance_join *TB_LOG_INSTANCE_JOIN) (int64, bool) {
	kvData, data := tb_log_instance_joinToStr(tb_log_instance_join)
	selectSql := "UPDATE `tb_log_instance_join` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_log_instance_join ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_log_instance_join RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_INSTANCE_JOINUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_log_instance_join` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_log_instance_join ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_log_instance_join RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Insert
func TB_LOG_INSTANCE_JOINInsert(logid string, roleid uint64, account string, serverid int32, instanceid int32, activityid int32, daytimes int32, star []int32, createtime int64) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_instance_join`(logid,roleid,account,serverid,instanceid,activityid,daytimes,star,createtime) VALUES"
	values := "(" + "'" + logid + "'" + "," + strconv.FormatUint(roleid, 10) + "," + "'" + account + "'" + "," + strconv.FormatInt(int64(serverid), 10) + "," + strconv.FormatInt(int64(instanceid), 10) + "," + strconv.FormatInt(int64(activityid), 10) + "," + strconv.FormatInt(int64(daytimes), 10) + "," + ToJson(star) + "," + strconv.FormatInt(createtime, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_instance_join ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_instance_join RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_LOG_INSTANCE_JOIN_Insert(tb_log_instance_join *TB_LOG_INSTANCE_JOIN) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_instance_join`(logid,roleid,account,serverid,instanceid,activityid,daytimes,star,createtime) VALUES"
	values := "(" + "'" + tb_log_instance_join.LOGID + "'" + "," + strconv.FormatUint(tb_log_instance_join.ROLEID, 10) + "," + "'" + tb_log_instance_join.ACCOUNT + "'" + "," + strconv.FormatInt(int64(tb_log_instance_join.SERVERID), 10) + "," + strconv.FormatInt(int64(tb_log_instance_join.INSTANCEID), 10) + "," + strconv.FormatInt(int64(tb_log_instance_join.ACTIVITYID), 10) + "," + strconv.FormatInt(int64(tb_log_instance_join.DAYTIMES), 10) + "," + ToJson(tb_log_instance_join.STAR) + "," + strconv.FormatInt(tb_log_instance_join.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_instance_join ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_instance_join RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_LOG_INSTANCE_JOINdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_log_instance_join` where " + conditions
	result, err := loggerDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_log_instance_join failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_log_instance_join RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_LOG_INSTANCE_JOIN(logid string, roleid uint64, account string, serverid int32, instanceid int32, activityid int32, daytimes int32, star []int32, createtime int64) *TB_LOG_INSTANCE_JOIN {
	dbData := &TB_LOG_INSTANCE_JOIN{
		LOGID:      logid,
		ROLEID:     roleid,
		ACCOUNT:    account,
		SERVERID:   serverid,
		INSTANCEID: instanceid,
		ACTIVITYID: activityid,
		DAYTIMES:   daytimes,
		STAR:       star,
		CREATETIME: createtime,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_log_instance_joinToStr(dbData *TB_LOG_INSTANCE_JOIN) (key string, data string) {
	key = "logid=" + "'" + dbData.LOGID + "'"
	data = "roleid=" + strconv.FormatUint(dbData.ROLEID, 10) + "," +
		"account=" + "'" + dbData.ACCOUNT + "'" + "," +
		"serverid=" + strconv.FormatInt(int64(dbData.SERVERID), 10) + "," +
		"instanceid=" + strconv.FormatInt(int64(dbData.INSTANCEID), 10) + "," +
		"activityid=" + strconv.FormatInt(int64(dbData.ACTIVITYID), 10) + "," +
		"daytimes=" + strconv.FormatInt(int64(dbData.DAYTIMES), 10) + "," +
		"star=" + ToJson(dbData.STAR) + "," +
		"createtime=" + strconv.FormatInt(dbData.CREATETIME, 10)
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_log_instance_join(dbData *tb_log_instance_join) *TB_LOG_INSTANCE_JOIN {
	tbData := &TB_LOG_INSTANCE_JOIN{
		LOGID:      dbData.logid,
		ROLEID:     dbData.roleid,
		ACCOUNT:    dbData.account,
		SERVERID:   dbData.serverid,
		INSTANCEID: dbData.instanceid,
		ACTIVITYID: dbData.activityid,
		DAYTIMES:   dbData.daytimes,
		STAR:       JsonToIntList(dbData.star),
		CREATETIME: dbData.createtime,
	}
	return tbData
}
