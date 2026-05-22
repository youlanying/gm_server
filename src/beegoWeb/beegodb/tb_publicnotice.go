package beegodb

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_publicnotice struct {
	id         int32  "ID"
	platformid int32  "平台ID"
	serverid   int32  "服务器ID"
	noticetype int32  "公告类型"
	label      int32  "标签"
	priority   int32  "优先级"
	titleshort string "短标题"
	title      string "标题"
	content    string "内容"
	starttime  int64  "开始时间"
	endtime    int64  "结束时间"
	createtime int64  "创建时间"
	audittime  int64  "审核时间"
	sendid     string "提交ID"
	auditid    string "审核ID"
}

type TB_PUBLICNOTICE struct {
	ID         int32  "ID"
	PLATFORMID int32  "平台ID"
	SERVERID   int32  "服务器ID"
	NOTICETYPE int32  "公告类型"
	LABEL      int32  "标签"
	PRIORITY   int32  "优先级"
	TITLESHORT string "短标题"
	TITLE      string "标题"
	CONTENT    string "内容"
	STARTTIME  int64  "开始时间"
	ENDTIME    int64  "结束时间"
	CREATETIME int64  "创建时间"
	AUDITTIME  int64  "审核时间"
	SENDID     string "提交ID"
	AUDITID    string "审核ID"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_PUBLICNOTICEReadByid(conditions string) (map[int32]*TB_PUBLICNOTICE, bool) {
	selectSql := "select * from `tb_publicnotice` where " + conditions
	mapDbData := make(map[int32]*TB_PUBLICNOTICE)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_publicnotice ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_publicnotice{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.platformid, &dbData.serverid, &dbData.noticetype, &dbData.label, &dbData.priority, &dbData.titleshort, &dbData.title, &dbData.content, &dbData.starttime, &dbData.endtime, &dbData.createtime, &dbData.audittime, &dbData.sendid, &dbData.auditid)
		if err != nil {
			logger.LogErr("get tb_publicnotice error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_publicnotice(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_PUBLICNOTICEReadBySQL(conditions string) (map[int32]*TB_PUBLICNOTICE, bool) {
	selectSql := "select * from `tb_publicnotice` " + conditions
	mapDbData := make(map[int32]*TB_PUBLICNOTICE)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_publicnotice ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_publicnotice{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.platformid, &dbData.serverid, &dbData.noticetype, &dbData.label, &dbData.priority, &dbData.titleshort, &dbData.title, &dbData.content, &dbData.starttime, &dbData.endtime, &dbData.createtime, &dbData.audittime, &dbData.sendid, &dbData.auditid)
		if err != nil {
			logger.LogErr("get sql tb_publicnotice error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_publicnotice(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_PUBLICNOTICE_ReplaceInto(tb_publicnotice *TB_PUBLICNOTICE) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_publicnotice`(id,platformid,serverid,noticetype,label,priority,titleshort,title,content,starttime,endtime,createtime,audittime,sendid,auditid) values"
	values := "(" + strconv.FormatInt(int64(tb_publicnotice.ID), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.PLATFORMID), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.SERVERID), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.NOTICETYPE), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.LABEL), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.PRIORITY), 10) + "," + "'" + tb_publicnotice.TITLESHORT + "'" + "," + "'" + tb_publicnotice.TITLE + "'" + "," + "'" + tb_publicnotice.CONTENT + "'" + "," + strconv.FormatInt(tb_publicnotice.STARTTIME, 10) + "," + strconv.FormatInt(tb_publicnotice.ENDTIME, 10) + "," + strconv.FormatInt(tb_publicnotice.CREATETIME, 10) + "," + strconv.FormatInt(tb_publicnotice.AUDITTIME, 10) + "," + "'" + tb_publicnotice.SENDID + "'" + "," + "'" + tb_publicnotice.AUDITID + "'" + ")"
	result, err := beegoDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_publicnotice error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_publicnotice failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_PUBLICNOTICEUpdateBy(tb_publicnotice *TB_PUBLICNOTICE) (int64, bool) {
	kvData, data := tb_publicnoticeToStr(tb_publicnotice)
	selectSql := "UPDATE `tb_publicnotice` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_publicnotice ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_publicnotice RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_PUBLICNOTICEUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_publicnotice` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_publicnotice ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_publicnotice RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//InsertAuto
func TB_PUBLICNOTICEInsertAuto(platformid int32, serverid int32, noticetype int32, label int32, priority int32, titleshort string, title string, content string, starttime int64, endtime int64, createtime int64, audittime int64, sendid string, auditid string) (int64, bool) {
	insertSql := "INSERT INTO `tb_publicnotice`(platformid,serverid,noticetype,label,priority,titleshort,title,content,starttime,endtime,createtime,audittime,sendid,auditid) VALUES"
	values := "(" + strconv.FormatInt(int64(platformid), 10) + "," + strconv.FormatInt(int64(serverid), 10) + "," + strconv.FormatInt(int64(noticetype), 10) + "," + strconv.FormatInt(int64(label), 10) + "," + strconv.FormatInt(int64(priority), 10) + "," + "'" + titleshort + "'" + "," + "'" + title + "'" + "," + "'" + content + "'" + "," + strconv.FormatInt(starttime, 10) + "," + strconv.FormatInt(endtime, 10) + "," + strconv.FormatInt(createtime, 10) + "," + strconv.FormatInt(audittime, 10) + "," + "'" + sendid + "'" + "," + "'" + auditid + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_publicnotice ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_publicnotice RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_PUBLICNOTICE_Insert(tb_publicnotice *TB_PUBLICNOTICE) (int64, bool) {
	insertSql := "INSERT INTO `tb_publicnotice`(id,platformid,serverid,noticetype,label,priority,titleshort,title,content,starttime,endtime,createtime,audittime,sendid,auditid) VALUES"
	values := "(" + strconv.FormatInt(int64(tb_publicnotice.ID), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.PLATFORMID), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.SERVERID), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.NOTICETYPE), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.LABEL), 10) + "," + strconv.FormatInt(int64(tb_publicnotice.PRIORITY), 10) + "," + "'" + tb_publicnotice.TITLESHORT + "'" + "," + "'" + tb_publicnotice.TITLE + "'" + "," + "'" + tb_publicnotice.CONTENT + "'" + "," + strconv.FormatInt(tb_publicnotice.STARTTIME, 10) + "," + strconv.FormatInt(tb_publicnotice.ENDTIME, 10) + "," + strconv.FormatInt(tb_publicnotice.CREATETIME, 10) + "," + strconv.FormatInt(tb_publicnotice.AUDITTIME, 10) + "," + "'" + tb_publicnotice.SENDID + "'" + "," + "'" + tb_publicnotice.AUDITID + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_publicnotice ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_publicnotice RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_PUBLICNOTICEdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_publicnotice` where " + conditions
	result, err := beegoDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_publicnotice failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_publicnotice RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_PUBLICNOTICE(id int32, platformid int32, serverid int32, noticetype int32, label int32, priority int32, titleshort string, title string, content string, starttime int64, endtime int64, createtime int64, audittime int64, sendid string, auditid string) *TB_PUBLICNOTICE {
	dbData := &TB_PUBLICNOTICE{
		ID:         id,
		PLATFORMID: platformid,
		SERVERID:   serverid,
		NOTICETYPE: noticetype,
		LABEL:      label,
		PRIORITY:   priority,
		TITLESHORT: titleshort,
		TITLE:      title,
		CONTENT:    content,
		STARTTIME:  starttime,
		ENDTIME:    endtime,
		CREATETIME: createtime,
		AUDITTIME:  audittime,
		SENDID:     sendid,
		AUDITID:    auditid,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_publicnoticeToStr(dbData *TB_PUBLICNOTICE) (key string, data string) {
	key = "id=" + strconv.FormatInt(int64(dbData.ID), 10)
	data = "platformid=" + strconv.FormatInt(int64(dbData.PLATFORMID), 10) + "," +
		"serverid=" + strconv.FormatInt(int64(dbData.SERVERID), 10) + "," +
		"noticetype=" + strconv.FormatInt(int64(dbData.NOTICETYPE), 10) + "," +
		"label=" + strconv.FormatInt(int64(dbData.LABEL), 10) + "," +
		"priority=" + strconv.FormatInt(int64(dbData.PRIORITY), 10) + "," +
		"titleshort=" + "'" + dbData.TITLESHORT + "'" + "," +
		"title=" + "'" + dbData.TITLE + "'" + "," +
		"content=" + "'" + dbData.CONTENT + "'" + "," +
		"starttime=" + strconv.FormatInt(dbData.STARTTIME, 10) + "," +
		"endtime=" + strconv.FormatInt(dbData.ENDTIME, 10) + "," +
		"createtime=" + strconv.FormatInt(dbData.CREATETIME, 10) + "," +
		"audittime=" + strconv.FormatInt(dbData.AUDITTIME, 10) + "," +
		"sendid=" + "'" + dbData.SENDID + "'" + "," +
		"auditid=" + "'" + dbData.AUDITID + "'"
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_publicnotice(dbData *tb_publicnotice) *TB_PUBLICNOTICE {
	tbData := &TB_PUBLICNOTICE{
		ID:         dbData.id,
		PLATFORMID: dbData.platformid,
		SERVERID:   dbData.serverid,
		NOTICETYPE: dbData.noticetype,
		LABEL:      dbData.label,
		PRIORITY:   dbData.priority,
		TITLESHORT: dbData.titleshort,
		TITLE:      dbData.title,
		CONTENT:    dbData.content,
		STARTTIME:  dbData.starttime,
		ENDTIME:    dbData.endtime,
		CREATETIME: dbData.createtime,
		AUDITTIME:  dbData.audittime,
		SENDID:     dbData.sendid,
		AUDITID:    dbData.auditid,
	}
	return tbData
}
