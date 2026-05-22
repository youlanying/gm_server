package beegodb

//%% !!! 此代码为自动生成 !!!
import (
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_center_list struct {
	id         int32  "平台Id"
	name       string "平台名字"
	ip         string "IP地址"
	port       string "端口"
	serverpath string "表存放路径"
}

type TB_CENTER_LIST struct {
	ID         int32  "平台Id"
	NAME       string "平台名字"
	IP         string "IP地址"
	PORT       string "端口"
	SERVERPATH string "表存放路径"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_CENTER_LISTReadByid(conditions string) (map[int32]*TB_CENTER_LIST, bool) {
	selectSql := "select * from `tb_center_list` where " + conditions
	mapDbData := make(map[int32]*TB_CENTER_LIST)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_center_list ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_center_list{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.name, &dbData.ip, &dbData.port, &dbData.serverpath)
		if err != nil {
			logger.LogErr("get tb_center_list error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_center_list(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_CENTER_LISTReadBySQL(conditions string) (map[int32]*TB_CENTER_LIST, bool) {
	selectSql := "select * from `tb_center_list` " + conditions
	mapDbData := make(map[int32]*TB_CENTER_LIST)
	rows, err := beegoDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_center_list ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_center_list{}
		rows.Columns()
		err := rows.Scan(&dbData.id, &dbData.name, &dbData.ip, &dbData.port, &dbData.serverpath)
		if err != nil {
			logger.LogErr("get sql tb_center_list error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_center_list(dbData)
		mapDbData[dbData.id] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_CENTER_LIST_ReplaceInto(tb_center_list *TB_CENTER_LIST) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_center_list`(id,name,ip,port,serverpath) values"
	values := "(" + strconv.FormatInt(int64(tb_center_list.ID), 10) + "," + "'" + tb_center_list.NAME + "'" + "," + "'" + tb_center_list.IP + "'" + "," + "'" + tb_center_list.PORT + "'" + "," + "'" + tb_center_list.SERVERPATH + "'" + ")"
	result, err := beegoDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_center_list error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_center_list failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_CENTER_LISTUpdateBy(tb_center_list *TB_CENTER_LIST) (int64, bool) {
	kvData, data := tb_center_listToStr(tb_center_list)
	selectSql := "UPDATE `tb_center_list` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_center_list ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_center_list RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_CENTER_LISTUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_center_list` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := beegoDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_center_list ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_center_list RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Insert
func TB_CENTER_LISTInsert(id int32, name string, ip string, port string, serverpath string) (int64, bool) {
	insertSql := "INSERT INTO `tb_center_list`(id,name,ip,port,serverpath) VALUES"
	values := "(" + strconv.FormatInt(int64(id), 10) + "," + "'" + name + "'" + "," + "'" + ip + "'" + "," + "'" + port + "'" + "," + "'" + serverpath + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_center_list ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_center_list RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_CENTER_LIST_Insert(tb_center_list *TB_CENTER_LIST) (int64, bool) {
	insertSql := "INSERT INTO `tb_center_list`(id,name,ip,port,serverpath) VALUES"
	values := "(" + strconv.FormatInt(int64(tb_center_list.ID), 10) + "," + "'" + tb_center_list.NAME + "'" + "," + "'" + tb_center_list.IP + "'" + "," + "'" + tb_center_list.PORT + "'" + "," + "'" + tb_center_list.SERVERPATH + "'" + ")"
	result, err := beegoDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_center_list ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_center_list RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_CENTER_LISTdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_center_list` where " + conditions
	result, err := beegoDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_center_list failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_center_list RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_CENTER_LIST(id int32, name string, ip string, port string, serverpath string) *TB_CENTER_LIST {
	dbData := &TB_CENTER_LIST{
		ID:         id,
		NAME:       name,
		IP:         ip,
		PORT:       port,
		SERVERPATH: serverpath,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_center_listToStr(dbData *TB_CENTER_LIST) (key string, data string) {
	key = "id=" + strconv.FormatInt(int64(dbData.ID), 10)
	data = "name=" + "'" + dbData.NAME + "'" + "," +
		"ip=" + "'" + dbData.IP + "'" + "," +
		"port=" + "'" + dbData.PORT + "'" + "," +
		"serverpath=" + "'" + dbData.SERVERPATH + "'"
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_center_list(dbData *tb_center_list) *TB_CENTER_LIST {
	tbData := &TB_CENTER_LIST{
		ID:         dbData.id,
		NAME:       dbData.name,
		IP:         dbData.ip,
		PORT:       dbData.port,
		SERVERPATH: dbData.serverpath,
	}
	return tbData
}
