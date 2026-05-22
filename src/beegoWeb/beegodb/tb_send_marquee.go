package beegodb

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_send_marquee struct {
	id         int32  ""
	mtype      int32  ""
	num        int32  ""
	content    string ""
	createtime string ""
	audittime  string ""
	endtime    string ""
	sendok     string ""
	serverurl  string ""
	sendid     string ""
	auditid    string ""
}

type TB_SEND_MARQUEE struct {
	ID         int32  ""
	MTYPE      int32  ""
	NUM        int32  ""
	CONTENT    string ""
	CREATETIME string ""
	AUDITTIME  string ""
	ENDTIME    string ""
	SENDOK     string ""
	SERVERURL  string ""
	SENDID     string ""
	AUDITID    string ""
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_SEND_MARQUEEReadByid(conditions string) (map[int32]*TB_SEND_MARQUEE, bool) {
	selectSql := "select * from `tb_send_marquee` where " + conditions
	mapDbData := make(map[int32]*TB_SEND_MARQUEE)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_send_marquee ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_send_marquee{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.mtype, &dbData.num, &dbData.content, &dbData.createtime, &dbData.audittime, &dbData.endtime, &dbData.sendok, &dbData.serverurl, &dbData.sendid, &dbData.auditid)
		if err != nil {
			logger.LogErr("get tb_send_marquee error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_send_marquee(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_SEND_MARQUEEReadBySQL(conditions string) (map[int32]*TB_SEND_MARQUEE, bool) {
	selectSql := "select * from `tb_send_marquee` " + conditions
	mapDbData := make(map[int32]*TB_SEND_MARQUEE)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_send_marquee ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_send_marquee{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.mtype, &dbData.num, &dbData.content, &dbData.createtime, &dbData.audittime, &dbData.endtime, &dbData.sendok, &dbData.serverurl, &dbData.sendid, &dbData.auditid)
		if err != nil {
			logger.LogErr("get sql tb_send_marquee error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_send_marquee(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_SEND_MARQUEE_ReplaceInto(tb_send_marquee *TB_SEND_MARQUEE) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_send_marquee`(id,mtype,num,content,createtime,audittime,endtime,sendok,serverurl,sendid,auditid) values"
	values := "(" + strconv.FormatInt(int64(tb_send_marquee.ID), 10) + "," + strconv.FormatInt(int64(tb_send_marquee.MTYPE), 10) + "," + strconv.FormatInt(int64(tb_send_marquee.NUM), 10) + "," + "'" + tb_send_marquee.CONTENT + "'" + "," + "'" + tb_send_marquee.CREATETIME + "'" + "," + "'" + tb_send_marquee.AUDITTIME + "'" + "," + "'" + tb_send_marquee.ENDTIME + "'" + "," + "'" + tb_send_marquee.SENDOK + "'" + "," + "'" + tb_send_marquee.SERVERURL + "'" + "," + "'" + tb_send_marquee.SENDID + "'" + "," + "'" + tb_send_marquee.AUDITID + "'" + ")"
	result, err := beegoDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_send_marquee error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_send_marquee failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_SEND_MARQUEEUpdateBy(tb_send_marquee *TB_SEND_MARQUEE) (int64, bool) {
	kvData, data := tb_send_marqueeToStr(tb_send_marquee)
	selectSql := "UPDATE `tb_send_marquee` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_send_marquee ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_send_marquee RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_SEND_MARQUEEUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_send_marquee` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_send_marquee ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_send_marquee RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//InsertAuto
func TB_SEND_MARQUEEInsertAuto(mtype int32, num int32, content string, createtime string, audittime string, endtime string, sendok string, serverurl string, sendid string, auditid string) (int64, bool) {
	insertSql := "INSERT INTO `tb_send_marquee`(mtype,num,content,createtime,audittime,endtime,sendok,serverurl,sendid,auditid) VALUES"
	values := "(" + strconv.FormatInt(int64(mtype), 10) + "," + strconv.FormatInt(int64(num), 10) + "," + "'" + content + "'" + "," + "'" + createtime + "'" + "," + "'" + audittime + "'" + "," + "'" + endtime + "'" + "," + "'" + sendok + "'" + "," + "'" + serverurl + "'" + "," + "'" + sendid + "'" + "," + "'" + auditid + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_send_marquee ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_send_marquee RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_SEND_MARQUEE_Insert(tb_send_marquee *TB_SEND_MARQUEE) (int64, bool) {
	insertSql := "INSERT INTO `tb_send_marquee`(id,mtype,num,content,createtime,audittime,endtime,sendok,serverurl,sendid,auditid) VALUES"
	values := "(" + strconv.FormatInt(int64(tb_send_marquee.ID), 10) + "," + strconv.FormatInt(int64(tb_send_marquee.MTYPE), 10) + "," + strconv.FormatInt(int64(tb_send_marquee.NUM), 10) + "," + "'" + tb_send_marquee.CONTENT + "'" + "," + "'" + tb_send_marquee.CREATETIME + "'" + "," + "'" + tb_send_marquee.AUDITTIME + "'" + "," + "'" + tb_send_marquee.ENDTIME + "'" + "," + "'" + tb_send_marquee.SENDOK + "'" + "," + "'" + tb_send_marquee.SERVERURL + "'" + "," + "'" + tb_send_marquee.SENDID + "'" + "," + "'" + tb_send_marquee.AUDITID + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_send_marquee ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_send_marquee RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_SEND_MARQUEEdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_send_marquee` where " + conditions
	result, err := beegoDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_send_marquee failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_send_marquee RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_SEND_MARQUEE(id int32, mtype int32, num int32, content string, createtime string, audittime string, endtime string, sendok string, serverurl string, sendid string, auditid string) *TB_SEND_MARQUEE {
	dbData := &TB_SEND_MARQUEE{
		ID:         id,
		MTYPE:      mtype,
		NUM:        num,
		CONTENT:    content,
		CREATETIME: createtime,
		AUDITTIME:  audittime,
		ENDTIME:    endtime,
		SENDOK:     sendok,
		SERVERURL:  serverurl,
		SENDID:     sendid,
		AUDITID:    auditid,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_send_marqueeToStr(dbData *TB_SEND_MARQUEE) (key string, data string) {
	key = "id=" + strconv.FormatInt(int64(dbData.ID), 10)
	data = "mtype=" + strconv.FormatInt(int64(dbData.MTYPE), 10) + "," +
		"num=" + strconv.FormatInt(int64(dbData.NUM), 10) + "," +
		"content=" + "'" + dbData.CONTENT + "'" + "," +
		"createtime=" + "'" + dbData.CREATETIME + "'" + "," +
		"audittime=" + "'" + dbData.AUDITTIME + "'" + "," +
		"endtime=" + "'" + dbData.ENDTIME + "'" + "," +
		"sendok=" + "'" + dbData.SENDOK + "'" + "," +
		"serverurl=" + "'" + dbData.SERVERURL + "'" + "," +
		"sendid=" + "'" + dbData.SENDID + "'" + "," +
		"auditid=" + "'" + dbData.AUDITID + "'"
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_send_marquee(dbData *tb_send_marquee) *TB_SEND_MARQUEE {
	tbData := &TB_SEND_MARQUEE{
		ID:         dbData.id,
		MTYPE:      dbData.mtype,
		NUM:        dbData.num,
		CONTENT:    dbData.content,
		CREATETIME: dbData.createtime,
		AUDITTIME:  dbData.audittime,
		ENDTIME:    dbData.endtime,
		SENDOK:     dbData.sendok,
		SERVERURL:  dbData.serverurl,
		SENDID:     dbData.sendid,
		AUDITID:    dbData.auditid,
	}
	return tbData
}
