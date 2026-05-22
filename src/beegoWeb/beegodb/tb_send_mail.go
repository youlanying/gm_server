package beegodb

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_send_mail struct {
	mailid     int32  ""
	title      string ""
	sendor     string ""
	recvname   string ""
	content    string ""
	itemlist   string ""
	isall      int32  ""
	createtime string ""
	audittime  string ""
	status     int32  ""
	serverurl  string ""
	sendid     string ""
	auditid    string ""
	timetype   int32  ""
	starttime  string ""
	endtime    string ""
	lvstart    int32  ""
	lvend      int32  ""
	sex        int32  ""
	reason     string ""
}

type TB_SEND_MAIL struct {
	MAILID     int32  ""
	TITLE      string ""
	SENDOR     string ""
	RECVNAME   string ""
	CONTENT    string ""
	ITEMLIST   string ""
	ISALL      int32  ""
	CREATETIME string ""
	AUDITTIME  string ""
	STATUS     int32  ""
	SERVERURL  string ""
	SENDID     string ""
	AUDITID    string ""
	TIMETYPE   int32  ""
	STARTTIME  string ""
	ENDTIME    string ""
	LVSTART    int32  ""
	LVEND      int32  ""
	SEX        int32  ""
	REASON     string ""
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_SEND_MAILReadBymailid(conditions string) (map[int32]*TB_SEND_MAIL, bool) {
	selectSql := "select * from `tb_send_mail` where " + conditions
	mapDbData := make(map[int32]*TB_SEND_MAIL)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_send_mail ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_send_mail{}
		rows.Columns()
		err := rows.Scan(&dbData.mailid, &dbData.title, &dbData.sendor, &dbData.recvname, &dbData.content, &dbData.itemlist, &dbData.isall, &dbData.createtime, &dbData.audittime, &dbData.status, &dbData.serverurl, &dbData.sendid, &dbData.auditid, &dbData.timetype, &dbData.starttime, &dbData.endtime, &dbData.lvstart, &dbData.lvend, &dbData.sex, &dbData.reason)
		if err != nil {
			logger.LogErr("get tb_send_mail error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_send_mail(dbData)
		mapDbData[dbData.mailid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_SEND_MAILReadBySQL(conditions string) (map[int32]*TB_SEND_MAIL, bool) {
	selectSql := "select * from `tb_send_mail` " + conditions
	mapDbData := make(map[int32]*TB_SEND_MAIL)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_send_mail ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_send_mail{}
		rows.Columns()
		err := rows.Scan(&dbData.mailid, &dbData.title, &dbData.sendor, &dbData.recvname, &dbData.content, &dbData.itemlist, &dbData.isall, &dbData.createtime, &dbData.audittime, &dbData.status, &dbData.serverurl, &dbData.sendid, &dbData.auditid, &dbData.timetype, &dbData.starttime, &dbData.endtime, &dbData.lvstart, &dbData.lvend, &dbData.sex, &dbData.reason)
		if err != nil {
			logger.LogErr("get sql tb_send_mail error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_send_mail(dbData)
		mapDbData[dbData.mailid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_SEND_MAIL_ReplaceInto(tb_send_mail *TB_SEND_MAIL) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_send_mail`(mailid,title,sendor,recvname,content,itemlist,isall,createtime,audittime,status,serverurl,sendid,auditid,timetype,starttime,endtime,lvstart,lvend,sex,reason) values"
	values := "(" + strconv.FormatInt(int64(tb_send_mail.MAILID), 10) + "," + "'" + tb_send_mail.TITLE + "'" + "," + "'" + tb_send_mail.SENDOR + "'" + "," + "'" + tb_send_mail.RECVNAME + "'" + "," + "'" + tb_send_mail.CONTENT + "'" + "," + "'" + tb_send_mail.ITEMLIST + "'" + "," + strconv.FormatInt(int64(tb_send_mail.ISALL), 10) + "," + "'" + tb_send_mail.CREATETIME + "'" + "," + "'" + tb_send_mail.AUDITTIME + "'" + "," + strconv.FormatInt(int64(tb_send_mail.STATUS), 10) + "," + "'" + tb_send_mail.SERVERURL + "'" + "," + "'" + tb_send_mail.SENDID + "'" + "," + "'" + tb_send_mail.AUDITID + "'" + "," + strconv.FormatInt(int64(tb_send_mail.TIMETYPE), 10) + "," + "'" + tb_send_mail.STARTTIME + "'" + "," + "'" + tb_send_mail.ENDTIME + "'" + "," + strconv.FormatInt(int64(tb_send_mail.LVSTART), 10) + "," + strconv.FormatInt(int64(tb_send_mail.LVEND), 10) + "," + strconv.FormatInt(int64(tb_send_mail.SEX), 10) + "," + "'" + tb_send_mail.REASON + "'" + ")"
	result, err := beegoDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_send_mail error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_send_mail failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_SEND_MAILUpdateBy(tb_send_mail *TB_SEND_MAIL) (int64, bool) {
	kvData, data := tb_send_mailToStr(tb_send_mail)
	selectSql := "UPDATE `tb_send_mail` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_send_mail ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_send_mail RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_SEND_MAILUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_send_mail` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_send_mail ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_send_mail RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//InsertAuto
func TB_SEND_MAILInsertAuto(title string, sendor string, recvname string, content string, itemlist string, isall int32, createtime string, audittime string, status int32, serverurl string, sendid string, auditid string, timetype int32, starttime string, endtime string, lvstart int32, lvend int32, sex int32, reason string) (int64, bool) {
	insertSql := "INSERT INTO `tb_send_mail`(title,sendor,recvname,content,itemlist,isall,createtime,audittime,status,serverurl,sendid,auditid,timetype,starttime,endtime,lvstart,lvend,sex,reason) VALUES"
	values := "(" + "'" + title + "'" + "," + "'" + sendor + "'" + "," + "'" + recvname + "'" + "," + "'" + content + "'" + "," + "'" + itemlist + "'" + "," + strconv.FormatInt(int64(isall), 10) + "," + "'" + createtime + "'" + "," + "'" + audittime + "'" + "," + strconv.FormatInt(int64(status), 10) + "," + "'" + serverurl + "'" + "," + "'" + sendid + "'" + "," + "'" + auditid + "'" + "," + strconv.FormatInt(int64(timetype), 10) + "," + "'" + starttime + "'" + "," + "'" + endtime + "'" + "," + strconv.FormatInt(int64(lvstart), 10) + "," + strconv.FormatInt(int64(lvend), 10) + "," + strconv.FormatInt(int64(sex), 10) + "," + "'" + reason + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_send_mail ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_send_mail RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_SEND_MAIL_Insert(tb_send_mail *TB_SEND_MAIL) (int64, bool) {
	insertSql := "INSERT INTO `tb_send_mail`(mailid,title,sendor,recvname,content,itemlist,isall,createtime,audittime,status,serverurl,sendid,auditid,timetype,starttime,endtime,lvstart,lvend,sex,reason) VALUES"
	values := "(" + strconv.FormatInt(int64(tb_send_mail.MAILID), 10) + "," + "'" + tb_send_mail.TITLE + "'" + "," + "'" + tb_send_mail.SENDOR + "'" + "," + "'" + tb_send_mail.RECVNAME + "'" + "," + "'" + tb_send_mail.CONTENT + "'" + "," + "'" + tb_send_mail.ITEMLIST + "'" + "," + strconv.FormatInt(int64(tb_send_mail.ISALL), 10) + "," + "'" + tb_send_mail.CREATETIME + "'" + "," + "'" + tb_send_mail.AUDITTIME + "'" + "," + strconv.FormatInt(int64(tb_send_mail.STATUS), 10) + "," + "'" + tb_send_mail.SERVERURL + "'" + "," + "'" + tb_send_mail.SENDID + "'" + "," + "'" + tb_send_mail.AUDITID + "'" + "," + strconv.FormatInt(int64(tb_send_mail.TIMETYPE), 10) + "," + "'" + tb_send_mail.STARTTIME + "'" + "," + "'" + tb_send_mail.ENDTIME + "'" + "," + strconv.FormatInt(int64(tb_send_mail.LVSTART), 10) + "," + strconv.FormatInt(int64(tb_send_mail.LVEND), 10) + "," + strconv.FormatInt(int64(tb_send_mail.SEX), 10) + "," + "'" + tb_send_mail.REASON + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_send_mail ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_send_mail RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_SEND_MAILdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_send_mail` where " + conditions
	result, err := beegoDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_send_mail failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_send_mail RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_SEND_MAIL(mailid int32, title string, sendor string, recvname string, content string, itemlist string, isall int32, createtime string, audittime string, status int32, serverurl string, sendid string, auditid string, timetype int32, starttime string, endtime string, lvstart int32, lvend int32, sex int32, reason string) *TB_SEND_MAIL {
	dbData := &TB_SEND_MAIL{
		MAILID:     mailid,
		TITLE:      title,
		SENDOR:     sendor,
		RECVNAME:   recvname,
		CONTENT:    content,
		ITEMLIST:   itemlist,
		ISALL:      isall,
		CREATETIME: createtime,
		AUDITTIME:  audittime,
		STATUS:     status,
		SERVERURL:  serverurl,
		SENDID:     sendid,
		AUDITID:    auditid,
		TIMETYPE:   timetype,
		STARTTIME:  starttime,
		ENDTIME:    endtime,
		LVSTART:    lvstart,
		LVEND:      lvend,
		SEX:        sex,
		REASON:     reason,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_send_mailToStr(dbData *TB_SEND_MAIL) (key string, data string) {
	key = "mailid=" + strconv.FormatInt(int64(dbData.MAILID), 10)
	data = "title=" + "'" + dbData.TITLE + "'" + "," +
		"sendor=" + "'" + dbData.SENDOR + "'" + "," +
		"recvname=" + "'" + dbData.RECVNAME + "'" + "," +
		"content=" + "'" + dbData.CONTENT + "'" + "," +
		"itemlist=" + "'" + dbData.ITEMLIST + "'" + "," +
		"isall=" + strconv.FormatInt(int64(dbData.ISALL), 10) + "," +
		"createtime=" + "'" + dbData.CREATETIME + "'" + "," +
		"audittime=" + "'" + dbData.AUDITTIME + "'" + "," +
		"status=" + strconv.FormatInt(int64(dbData.STATUS), 10) + "," +
		"serverurl=" + "'" + dbData.SERVERURL + "'" + "," +
		"sendid=" + "'" + dbData.SENDID + "'" + "," +
		"auditid=" + "'" + dbData.AUDITID + "'" + "," +
		"timetype=" + strconv.FormatInt(int64(dbData.TIMETYPE), 10) + "," +
		"starttime=" + "'" + dbData.STARTTIME + "'" + "," +
		"endtime=" + "'" + dbData.ENDTIME + "'" + "," +
		"lvstart=" + strconv.FormatInt(int64(dbData.LVSTART), 10) + "," +
		"lvend=" + strconv.FormatInt(int64(dbData.LVEND), 10) + "," +
		"sex=" + strconv.FormatInt(int64(dbData.SEX), 10) + "," +
		"reason=" + "'" + dbData.REASON + "'"
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_send_mail(dbData *tb_send_mail) *TB_SEND_MAIL {
	tbData := &TB_SEND_MAIL{
		MAILID:     dbData.mailid,
		TITLE:      dbData.title,
		SENDOR:     dbData.sendor,
		RECVNAME:   dbData.recvname,
		CONTENT:    dbData.content,
		ITEMLIST:   dbData.itemlist,
		ISALL:      dbData.isall,
		CREATETIME: dbData.createtime,
		AUDITTIME:  dbData.audittime,
		STATUS:     dbData.status,
		SERVERURL:  dbData.serverurl,
		SENDID:     dbData.sendid,
		AUDITID:    dbData.auditid,
		TIMETYPE:   dbData.timetype,
		STARTTIME:  dbData.starttime,
		ENDTIME:    dbData.endtime,
		LVSTART:    dbData.lvstart,
		LVEND:      dbData.lvend,
		SEX:        dbData.sex,
		REASON:     dbData.reason,
	}
	return tbData
}
