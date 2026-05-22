package logger_custom

//%% !!! 此代码为自动生成 !!!
import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

///////////////////////////////////////////////////////////////////////
type tb_log_instance_out struct {
	logid          string "id"
	instanceid     int32  "关卡ID"
	serverid       int32  "服务器id"
	roleid         uint64 "角色id"
	account        string "账号"
	rolelevel      int32  "玩家等级"
	state          int32  "通关状态0失败,1通关,2主动退出,3超时,4团队死亡"
	daytimes       int32  "次数"
	endgold        int32  "获得金币"
	roleexp        int32  "获得经验"
	heroexp        int32  "英雄获得经验"
	objectitems    string "获得物品Id, Num"
	star           string "星级"
	deadhero       string "死亡英雄"
	battleherolist string "出战英雄数据"
	timeconsum     int32  "通关耗时秒"
	createtime     int64  "创建时间"
}

type TB_LOG_INSTANCE_OUT struct {
	LOGID          string    "id"
	INSTANCEID     int32     "关卡ID"
	SERVERID       int32     "服务器id"
	ROLEID         uint64    "角色id"
	ACCOUNT        string    "账号"
	ROLELEVEL      int32     "玩家等级"
	STATE          int32     "通关状态0失败,1通关,2主动退出,3超时,4团队死亡"
	DAYTIMES       int32     "次数"
	ENDGOLD        int32     "获得金币"
	ROLEEXP        int32     "获得经验"
	HEROEXP        int32     "英雄获得经验"
	OBJECTITEMS    [][]int32 "获得物品Id, Num"
	STAR           []int32   "星级"
	DEADHERO       []int32   "死亡英雄"
	BATTLEHEROLIST [][]int32 "出战英雄数据"
	TIMECONSUM     int32     "通关耗时秒"
	CREATETIME     int64     "创建时间"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_INSTANCE_OUTReadBylogid(conditions string) (map[string]*TB_LOG_INSTANCE_OUT, bool) {
	selectSql := "select * from `tb_log_instance_out` where " + conditions
	mapDbData := make(map[string]*TB_LOG_INSTANCE_OUT)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail tb_log_instance_out ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_instance_out{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.instanceid, &dbData.serverid, &dbData.roleid, &dbData.account, &dbData.rolelevel, &dbData.state, &dbData.daytimes, &dbData.endgold, &dbData.roleexp, &dbData.heroexp, &dbData.objectitems, &dbData.star, &dbData.deadhero, &dbData.battleherolist, &dbData.timeconsum, &dbData.createtime)
		if err != nil {
			logger.LogErr("get tb_log_instance_out error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_instance_out(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//read_data by primary key
//conditions = "account=Account"
func TB_LOG_INSTANCE_OUTReadBySQL(conditions string) (map[string]*TB_LOG_INSTANCE_OUT, bool) {
	selectSql := "select * from `tb_log_instance_out` " + conditions
	mapDbData := make(map[string]*TB_LOG_INSTANCE_OUT)
	rows, err := loggerDB.Query(selectSql)
	if err != nil {
		logger.LogErr("select fail sql tb_log_instance_out ", err)
		return mapDbData, false
	}
	for rows.Next() {
		dbData := &tb_log_instance_out{}
		rows.Columns()
		err := rows.Scan(&dbData.logid, &dbData.instanceid, &dbData.serverid, &dbData.roleid, &dbData.account, &dbData.rolelevel, &dbData.state, &dbData.daytimes, &dbData.endgold, &dbData.roleexp, &dbData.heroexp, &dbData.objectitems, &dbData.star, &dbData.deadhero, &dbData.battleherolist, &dbData.timeconsum, &dbData.createtime)
		if err != nil {
			logger.LogErr("get sql tb_log_instance_out error: ", err)
			return mapDbData, false
		}
		tbData := strTotb_log_instance_out(dbData)
		mapDbData[dbData.logid] = tbData
	}
	return mapDbData, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//REPLACE INTO 不存在insert 存在update 慎用
func TB_LOG_INSTANCE_OUT_ReplaceInto(tb_log_instance_out *TB_LOG_INSTANCE_OUT) (int64, bool) {
	replaceInfoSql := "REPLACE INTO `tb_log_instance_out`(logid,instanceid,serverid,roleid,account,rolelevel,state,daytimes,endgold,roleexp,heroexp,objectitems,star,deadhero,battleherolist,timeconsum,createtime) values"
	values := "(" + "'" + tb_log_instance_out.LOGID + "'" + "," + strconv.FormatInt(int64(tb_log_instance_out.INSTANCEID), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.SERVERID), 10) + "," + strconv.FormatUint(tb_log_instance_out.ROLEID, 10) + "," + "'" + tb_log_instance_out.ACCOUNT + "'" + "," + strconv.FormatInt(int64(tb_log_instance_out.ROLELEVEL), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.STATE), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.DAYTIMES), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.ENDGOLD), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.ROLEEXP), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.HEROEXP), 10) + "," + ToJson(tb_log_instance_out.OBJECTITEMS) + "," + ToJson(tb_log_instance_out.STAR) + "," + ToJson(tb_log_instance_out.DEADHERO) + "," + ToJson(tb_log_instance_out.BATTLEHEROLIST) + "," + strconv.FormatInt(int64(tb_log_instance_out.TIMECONSUM), 10) + "," + strconv.FormatInt(tb_log_instance_out.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(replaceInfoSql + values)
	if err != nil {
		logger.LogErr("REPLACE INTO tb_log_instance_out error: ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("REPLACE INTO tb_log_instance_out failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_INSTANCE_OUTUpdateBy(tb_log_instance_out *TB_LOG_INSTANCE_OUT) (int64, bool) {
	kvData, data := tb_log_instance_outToStr(tb_log_instance_out)
	selectSql := "UPDATE `tb_log_instance_out` SET " + data + " WHERE " + kvData
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateBy tb_log_instance_out ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateBy tb_log_instance_out RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//key = "account=Account"
func TB_LOG_INSTANCE_OUTUpdateByKey(key string, data string) (int64, bool) {
	selectSql := "UPDATE `tb_log_instance_out` SET " + data + " WHERE " + key
	//logger.LogErr("selectSql [%s]", selectSql)
	result, err := loggerDB.Exec(selectSql)
	if err != nil {
		logger.LogErr("UpdateByKey tb_log_instance_out ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("UpdateByKey tb_log_instance_out RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Insert
func TB_LOG_INSTANCE_OUTInsert(logid string, instanceid int32, serverid int32, roleid uint64, account string, rolelevel int32, state int32, daytimes int32, endgold int32, roleexp int32, heroexp int32, objectitems [][]int32, star []int32, deadhero []int32, battleherolist [][]int32, timeconsum int32, createtime int64) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_instance_out`(logid,instanceid,serverid,roleid,account,rolelevel,state,daytimes,endgold,roleexp,heroexp,objectitems,star,deadhero,battleherolist,timeconsum,createtime) VALUES"
	values := "(" + "'" + logid + "'" + "," + strconv.FormatInt(int64(instanceid), 10) + "," + strconv.FormatInt(int64(serverid), 10) + "," + strconv.FormatUint(roleid, 10) + "," + "'" + account + "'" + "," + strconv.FormatInt(int64(rolelevel), 10) + "," + strconv.FormatInt(int64(state), 10) + "," + strconv.FormatInt(int64(daytimes), 10) + "," + strconv.FormatInt(int64(endgold), 10) + "," + strconv.FormatInt(int64(roleexp), 10) + "," + strconv.FormatInt(int64(heroexp), 10) + "," + ToJson(objectitems) + "," + ToJson(star) + "," + ToJson(deadhero) + "," + ToJson(battleherolist) + "," + strconv.FormatInt(int64(timeconsum), 10) + "," + strconv.FormatInt(createtime, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_instance_out ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_instance_out RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//insert 传结构体指针
func TB_LOG_INSTANCE_OUT_Insert(tb_log_instance_out *TB_LOG_INSTANCE_OUT) (int64, bool) {
	insertSql := "INSERT INTO `tb_log_instance_out`(logid,instanceid,serverid,roleid,account,rolelevel,state,daytimes,endgold,roleexp,heroexp,objectitems,star,deadhero,battleherolist,timeconsum,createtime) VALUES"
	values := "(" + "'" + tb_log_instance_out.LOGID + "'" + "," + strconv.FormatInt(int64(tb_log_instance_out.INSTANCEID), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.SERVERID), 10) + "," + strconv.FormatUint(tb_log_instance_out.ROLEID, 10) + "," + "'" + tb_log_instance_out.ACCOUNT + "'" + "," + strconv.FormatInt(int64(tb_log_instance_out.ROLELEVEL), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.STATE), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.DAYTIMES), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.ENDGOLD), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.ROLEEXP), 10) + "," + strconv.FormatInt(int64(tb_log_instance_out.HEROEXP), 10) + "," + ToJson(tb_log_instance_out.OBJECTITEMS) + "," + ToJson(tb_log_instance_out.STAR) + "," + ToJson(tb_log_instance_out.DEADHERO) + "," + ToJson(tb_log_instance_out.BATTLEHEROLIST) + "," + strconv.FormatInt(int64(tb_log_instance_out.TIMECONSUM), 10) + "," + strconv.FormatInt(tb_log_instance_out.CREATETIME, 10) + ")"
	result, err := loggerDB.Exec(insertSql + values)
	if err != nil {
		logger.LogErr("Insert tb_log_instance_out ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("Insert tb_log_instance_out RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//where conditions = "account=Account"
func TB_LOG_INSTANCE_OUTdeleteBy(conditions string) (int64, bool) {
	delSql := "delete from `tb_log_instance_out` where " + conditions
	result, err := loggerDB.Exec(delSql)
	if err != nil {
		logger.LogErr("delete tb_log_instance_out failed ", err)
		return 0, false
	}
	idAff, er := result.RowsAffected()
	if er != nil {
		logger.LogErr("delete tb_log_instance_out RowsAffected failed:", er)
		return 0, false
	}
	return idAff, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//生成结构体
func MakeTB_LOG_INSTANCE_OUT(logid string, instanceid int32, serverid int32, roleid uint64, account string, rolelevel int32, state int32, daytimes int32, endgold int32, roleexp int32, heroexp int32, objectitems [][]int32, star []int32, deadhero []int32, battleherolist [][]int32, timeconsum int32, createtime int64) *TB_LOG_INSTANCE_OUT {
	dbData := &TB_LOG_INSTANCE_OUT{
		LOGID:          logid,
		INSTANCEID:     instanceid,
		SERVERID:       serverid,
		ROLEID:         roleid,
		ACCOUNT:        account,
		ROLELEVEL:      rolelevel,
		STATE:          state,
		DAYTIMES:       daytimes,
		ENDGOLD:        endgold,
		ROLEEXP:        roleexp,
		HEROEXP:        heroexp,
		OBJECTITEMS:    objectitems,
		STAR:           star,
		DEADHERO:       deadhero,
		BATTLEHEROLIST: battleherolist,
		TIMECONSUM:     timeconsum,
		CREATETIME:     createtime,
	}
	return dbData
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// to string
func tb_log_instance_outToStr(dbData *TB_LOG_INSTANCE_OUT) (key string, data string) {
	key = "logid=" + "'" + dbData.LOGID + "'"
	data = "instanceid=" + strconv.FormatInt(int64(dbData.INSTANCEID), 10) + "," +
		"serverid=" + strconv.FormatInt(int64(dbData.SERVERID), 10) + "," +
		"roleid=" + strconv.FormatUint(dbData.ROLEID, 10) + "," +
		"account=" + "'" + dbData.ACCOUNT + "'" + "," +
		"rolelevel=" + strconv.FormatInt(int64(dbData.ROLELEVEL), 10) + "," +
		"state=" + strconv.FormatInt(int64(dbData.STATE), 10) + "," +
		"daytimes=" + strconv.FormatInt(int64(dbData.DAYTIMES), 10) + "," +
		"endgold=" + strconv.FormatInt(int64(dbData.ENDGOLD), 10) + "," +
		"roleexp=" + strconv.FormatInt(int64(dbData.ROLEEXP), 10) + "," +
		"heroexp=" + strconv.FormatInt(int64(dbData.HEROEXP), 10) + "," +
		"objectitems=" + ToJson(dbData.OBJECTITEMS) + "," +
		"star=" + ToJson(dbData.STAR) + "," +
		"deadhero=" + ToJson(dbData.DEADHERO) + "," +
		"battleherolist=" + ToJson(dbData.BATTLEHEROLIST) + "," +
		"timeconsum=" + strconv.FormatInt(int64(dbData.TIMECONSUM), 10) + "," +
		"createtime=" + strconv.FormatInt(dbData.CREATETIME, 10)
	return key, data
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Json string TO 结构体
func strTotb_log_instance_out(dbData *tb_log_instance_out) *TB_LOG_INSTANCE_OUT {
	tbData := &TB_LOG_INSTANCE_OUT{
		LOGID:          dbData.logid,
		INSTANCEID:     dbData.instanceid,
		SERVERID:       dbData.serverid,
		ROLEID:         dbData.roleid,
		ACCOUNT:        dbData.account,
		ROLELEVEL:      dbData.rolelevel,
		STATE:          dbData.state,
		DAYTIMES:       dbData.daytimes,
		ENDGOLD:        dbData.endgold,
		ROLEEXP:        dbData.roleexp,
		HEROEXP:        dbData.heroexp,
		OBJECTITEMS:    jsonTotb_log_instance_out_objectitems(dbData.objectitems),
		STAR:           JsonToIntList(dbData.star),
		DEADHERO:       JsonToIntList(dbData.deadhero),
		BATTLEHEROLIST: jsonTotb_log_instance_out_battleherolist(dbData.battleherolist),
		TIMECONSUM:     dbData.timeconsum,
		CREATETIME:     dbData.createtime,
	}
	return tbData
}

func jsonTotb_log_instance_out_objectitems(v string) [][]int32 {
	mapData := make([][]int32, 0)
	if v == "" || v == "[]" || v == "{}" {
		return mapData
	}
	buffer := []byte(v)
	err := json.Unmarshal(buffer, &mapData)
	if err != nil {
		logger.LogErr("JsonFailed", err)
		return mapData
	}
	return mapData
}

func jsonTotb_log_instance_out_battleherolist(v string) [][]int32 {
	mapData := make([][]int32, 0)
	if v == "" || v == "[]" || v == "{}" {
		return mapData
	}
	buffer := []byte(v)
	err := json.Unmarshal(buffer, &mapData)
	if err != nil {
		logger.LogErr("JsonFailed", err)
		return mapData
	}
	return mapData
}
