package beegodb

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_admin_user struct {
	id       int32  "ID"
	userid   string "用户名"
	password string "密码"
	groupid  int32  "权限组ID"
	remarks  string "备注"
}

type TB_ADMIN_USER struct {
	ID       int32  "ID"
	USERID   string "用户名"
	PASSWORD string "密码"
	GROUPID  int32  "权限组ID"
	REMARKS  string "备注"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_ADMIN_USERReadByid(conditions string) (map[int32]*TB_ADMIN_USER, bool) {
	selectSql := "select * from `tb_admin_user` where " + conditions
	mapDbData := make(map[int32]*TB_ADMIN_USER)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_admin_user ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_admin_user{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.userid, &dbData.password, &dbData.groupid, &dbData.remarks)
		if err != nil {
			logger.LogErr("get tb_admin_user error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_admin_user(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_ADMIN_USERReadBySQL(conditions string) (map[int32]*TB_ADMIN_USER, bool) {
	selectSql := "select * from `tb_admin_user` " + conditions
	mapDbData := make(map[int32]*TB_ADMIN_USER)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_admin_user ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_admin_user{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.userid, &dbData.password, &dbData.groupid, &dbData.remarks)
		if err != nil {
			logger.LogErr("get sql tb_admin_user error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_admin_user(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_ADMIN_USER_ReplaceInto(tb_admin_user *TB_ADMIN_USER) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_admin_user`(id,userid,password,groupid,remarks) values"
	values := "(" + strconv.FormatInt(int64(tb_admin_user.ID), 10) + "," + "'" + tb_admin_user.USERID + "'" + "," + "'" + tb_admin_user.PASSWORD + "'" + "," + strconv.FormatInt(int64(tb_admin_user.GROUPID), 10) + "," + "'" + tb_admin_user.REMARKS + "'" + ")"
	result, err := beegoDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_admin_user error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_admin_user failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_ADMIN_USERUpdateBy(tb_admin_user *TB_ADMIN_USER) (int64, bool) {
	kvData, data := tb_admin_userToStr(tb_admin_user)
	selectSql := "UPDATE `tb_admin_user` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_admin_user ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_admin_user RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_ADMIN_USERUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_admin_user` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_admin_user ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_admin_user RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//InsertAuto
func TB_ADMIN_USERInsertAuto(userid string, password string, groupid int32, remarks string) (int64, bool) {
	insertSql := "INSERT INTO `tb_admin_user`(userid,password,groupid,remarks) VALUES"
	values := "(" + "'" + userid + "'" + "," + "'" + password + "'" + "," + strconv.FormatInt(int64(groupid), 10) + "," + "'" + remarks + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_admin_user ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_admin_user RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_ADMIN_USER_Insert(tb_admin_user *TB_ADMIN_USER) (int64, bool) {
	insertSql := "INSERT INTO `tb_admin_user`(id,userid,password,groupid,remarks) VALUES"
	values := "(" + strconv.FormatInt(int64(tb_admin_user.ID), 10) + "," + "'" + tb_admin_user.USERID + "'" + "," + "'" + tb_admin_user.PASSWORD + "'" + "," + strconv.FormatInt(int64(tb_admin_user.GROUPID), 10) + "," + "'" + tb_admin_user.REMARKS + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_admin_user ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_admin_user RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_ADMIN_USERdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_admin_user` where " + conditions
	result, err := beegoDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_admin_user failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_admin_user RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_ADMIN_USER(id int32, userid string, password string, groupid int32, remarks string) *TB_ADMIN_USER {
	dbData := &TB_ADMIN_USER{
		ID:       id,
		USERID:   userid,
		PASSWORD: password,
		GROUPID:  groupid,
		REMARKS:  remarks,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_admin_userToStr(dbData *TB_ADMIN_USER) (key string, data string) {
	key = "id=" + strconv.FormatInt(int64(dbData.ID), 10)
	data = "userid=" + "'" + dbData.USERID + "'" + "," +
		"password=" + "'" + dbData.PASSWORD + "'" + "," +
		"groupid=" + strconv.FormatInt(int64(dbData.GROUPID), 10) + "," +
		"remarks=" + "'" + dbData.REMARKS + "'"
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_admin_user(dbData *tb_admin_user) *TB_ADMIN_USER {
	tbData := &TB_ADMIN_USER{
		ID:       dbData.id,
		USERID:   dbData.userid,
		PASSWORD: dbData.password,
		GROUPID:  dbData.groupid,
		REMARKS:  dbData.remarks,
	}
	return tbData
}
