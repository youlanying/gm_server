package beegodb

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_admin_group struct {
	id         int32  "角色组id"
	name       string "用户组名字"
	actionlist string "用户组权限"
}

type TB_ADMIN_GROUP struct {
	ID         int32   "角色组id"
	NAME       string  "用户组名字"
	ACTIONLIST []int32 "用户组权限"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_ADMIN_GROUPReadByid(conditions string) (map[int32]*TB_ADMIN_GROUP, bool) {
	selectSql := "select * from `tb_admin_group` where " + conditions
	mapDbData := make(map[int32]*TB_ADMIN_GROUP)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_admin_group ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_admin_group{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.name, &dbData.actionlist)
		if err != nil {
			logger.LogErr("get tb_admin_group error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_admin_group(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_ADMIN_GROUPReadBySQL(conditions string) (map[int32]*TB_ADMIN_GROUP, bool) {
	selectSql := "select * from `tb_admin_group` " + conditions
	mapDbData := make(map[int32]*TB_ADMIN_GROUP)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_admin_group ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_admin_group{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.name, &dbData.actionlist)
		if err != nil {
			logger.LogErr("get sql tb_admin_group error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_admin_group(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_ADMIN_GROUP_ReplaceInto(tb_admin_group *TB_ADMIN_GROUP) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_admin_group`(id,name,actionlist) values"
	values := "(" + strconv.FormatInt(int64(tb_admin_group.ID), 10) + "," + "'" + tb_admin_group.NAME + "'" + "," + ToJson(tb_admin_group.ACTIONLIST) + ")"
	result, err := beegoDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_admin_group error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_admin_group failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_ADMIN_GROUPUpdateBy(tb_admin_group *TB_ADMIN_GROUP) (int64, bool) {
	kvData, data := tb_admin_groupToStr(tb_admin_group)
	selectSql := "UPDATE `tb_admin_group` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_admin_group ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_admin_group RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_ADMIN_GROUPUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_admin_group` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_admin_group ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_admin_group RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//InsertAuto
func TB_ADMIN_GROUPInsertAuto(name string, actionlist []int32) (int64, bool) {
	insertSql := "INSERT INTO `tb_admin_group`(name,actionlist) VALUES"
	values := "(" + "'" + name + "'" + "," + ToJson(actionlist) + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_admin_group ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_admin_group RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_ADMIN_GROUP_Insert(tb_admin_group *TB_ADMIN_GROUP) (int64, bool) {
	insertSql := "INSERT INTO `tb_admin_group`(id,name,actionlist) VALUES"
	values := "(" + strconv.FormatInt(int64(tb_admin_group.ID), 10) + "," + "'" + tb_admin_group.NAME + "'" + "," + ToJson(tb_admin_group.ACTIONLIST) + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_admin_group ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_admin_group RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_ADMIN_GROUPdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_admin_group` where " + conditions
	result, err := beegoDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_admin_group failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_admin_group RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_ADMIN_GROUP(id int32, name string, actionlist []int32) *TB_ADMIN_GROUP {
	dbData := &TB_ADMIN_GROUP{
		ID:         id,
		NAME:       name,
		ACTIONLIST: actionlist,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_admin_groupToStr(dbData *TB_ADMIN_GROUP) (key string, data string) {
	key = "id=" + strconv.FormatInt(int64(dbData.ID), 10)
	data = "name=" + "'" + dbData.NAME + "'" + "," +
		"actionlist=" + ToJson(dbData.ACTIONLIST)
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_admin_group(dbData *tb_admin_group) *TB_ADMIN_GROUP {
	tbData := &TB_ADMIN_GROUP{
		ID:         dbData.id,
		NAME:       dbData.name,
		ACTIONLIST: JsonToIntList(dbData.actionlist),
	}
	return tbData
}
