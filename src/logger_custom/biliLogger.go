package logger_custom

import (
	"encoding/json"
	"fmt"
	"gm_server/src/cfg"
	"gm_server/src/tools"
	"io"
	"log"
	"os"
	"path"
	"time"
)

const (
	BILILOGTYPE_OPS = "ops"
	BILILOGTYPE_BG  = "bg"

	BgLogPath  = "/data/log/bgamelog" //${应用}-${日期}.log
	OpsLogPath = "/data/log/opslog/"  //默认目录为 /data/log/opslog/${应用名}/*.log
)

var (
	_APP_NAME               string
	LOGFILE_MAXSIZE_DEFAULT int64 = 50 << 20
	_logfile_maxsize        int64 = LOGFILE_MAXSIZE_DEFAULT
	_ops_logger             *log.Logger
	_bg_logger              *log.Logger
	_ops_logtime            *time.Time
	_bg_logtime             *time.Time
	canLog                  bool
)

type LogFile struct {
	*os.File
}

// InitBiliLogger 初始化BiliLog
func InitBiliLogger(logType, logfile string, maxSize int64) {
	//读取全局配置
	config := cfg.GetSection("GLOBAL")
	canLog = config["usebililog"] == "true"
	_APP_NAME = config["appname"]
	if _APP_NAME == "" {
		log.Println("[BiliLog] appname not find, use default: yznws")
		_APP_NAME = "yznws" // todo 默认app名称
	}
	var filePath string
	switch logType {
	case BILILOGTYPE_OPS:
		filePath = OpsLogPath + _APP_NAME
	case BILILOGTYPE_BG:
		filePath = BgLogPath
	default:
		log.Println("logType error, no type:", logType)
		return
	}
	fullpath := path.Join(filePath, "", logfile)
	if _, loc_stat := os.Stat(filePath); loc_stat != nil {
		err := os.MkdirAll(filePath, 0777)
		if err != nil {
			log.Println("MkdirAll err:", err)
			return
		}
	}
	_logfile_maxsize = int64(maxSize) << 10 //单位是k
	if _logfile_maxsize < (1 << 16) {       //日志文件最小64k
		_logfile_maxsize = LOGFILE_MAXSIZE_DEFAULT
	}
	log.Println("logfile size:", _logfile_maxsize)
	file, err := OpenLogFile(fullpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("error opening file %v\n", err)
		return
	}
	logger := log.New(io.MultiWriter(file, os.Stderr), "", 0)
	t := time.Now()
	switch logType {
	case BILILOGTYPE_OPS:
		_ops_logger = logger
		_ops_logtime = &t
	case BILILOGTYPE_BG:
		_bg_logger = logger
		_bg_logtime = &t
	}
}

func OpenLogFile(name string, flag int, perm os.FileMode) (file *LogFile, err error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	lf := LogFile{}
	lf.File = f
	return &lf, nil
}

// CheckLogTime 检查log时间，隔天则生成新文件
func CheckLogTime(logType string) {
	now := time.Now()
	//strNow := now.Format(time.RFC3339)
	y, m, d := now.Date()
	var filePath string
	if logType == BILILOGTYPE_OPS && _ops_logtime != nil {
		isOtherDay := tools.OtherDay(now.Unix(), _ops_logtime.Unix())
		if isOtherDay {
			filePath = OpsLogPath + _APP_NAME
			fileName := fmt.Sprintf("%d%d%d.log", y, int(m), d)
			fullpath := path.Join(filePath, "", fileName)
			file, err := OpenLogFile(fullpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Printf("error opening file %v\n", err)
				return
			}
			logger := log.New(io.MultiWriter(file, os.Stderr), "", 0)
			_ops_logger = logger
			_ops_logtime = &now
		}
	} else if logType == BILILOGTYPE_BG && _bg_logtime != nil {
		isOtherDay := tools.OtherDay(now.Unix(), _bg_logtime.Unix())
		if isOtherDay {
			filePath = BgLogPath
			fileName := fmt.Sprintf("%s-%d%d%d.log", _APP_NAME, y, int(m), d)
			fullpath := path.Join(filePath, "", fileName)
			file, err := OpenLogFile(fullpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Printf("error opening file %v\n", err)
				return
			}
			logger := log.New(io.MultiWriter(file, os.Stderr), "", 0)
			_bg_logger = logger
			_bg_logtime = &now
		}
	}
}

// BiliLog_BgLog bg类型log的基础写入函数
func BiliLog_BgLog(logData map[string]interface{}) {
	CheckLogTime(BILILOGTYPE_BG)
	b, err := json.Marshal(logData)
	if err != nil {
		log.Printf("json.Marshal failed:%v \n\tlog data:%v\n", err, logData)
		return
	}
	_bg_logger.Println(string(b))
}

// BiliLog_OpsLog ops类型log的基础写入函数
func BiliLog_OpsLog(logData map[string]interface{}) {
	CheckLogTime(BILILOGTYPE_OPS)
	b, err := json.Marshal(logData)
	if err != nil {
		log.Printf("json.Marshal failed:%v \n\tlog data:%v\n", err, logData)
		return
	}
	_ops_logger.Println(string(b))
}

// Bili_GetLogTimeNow 用于统一获取bili要求的log时间格式
func Bili_GetLogTimeNow() string {
	return time.Now().Format(time.RFC3339)
}

// todo 之后可以根据不同log编写专用方法；
// key不要使用message、type这两个保留字段；
// value尽量不要使用数组类型，数组不支持检索；
// ops必须字段：time、level、app_id、instance_id，其他详细请参照文档
// bg必须字段：time、game_name、game_name_type、log_id、env，其他详细请参照文档

//-----------------------------------------------
//--------------      OPS LOG      --------------
//-----------------------------------------------

// BiliOpsLog_GameServerStart 服务器启动
func BiliOpsLog_GameServerStart(serverId int32, version string) {
	if !canLog {
		return
	}
	logData := make(map[string]interface{})
	nowStr := Bili_GetLogTimeNow()
	logData["time"] = nowStr
	logData["level"] = "INFO"
	logData["app_id"] = _APP_NAME
	logData["instance_id"] = _APP_NAME
	logData["log"] = fmt.Sprintf("game server %d start, version: %s", serverId, version)
	BiliLog_OpsLog(logData)
}

// BiliOpsLog_TableError 表单错误
func BiliOpsLog_TableError(serverId int32, roleId uint64, account, roleName, protoErrId string) {
	if !canLog {
		return
	}
	logData := make(map[string]interface{})
	nowStr := Bili_GetLogTimeNow()
	logData["time"] = nowStr
	logData["level"] = "WARNING"
	logData["app_id"] = _APP_NAME
	logData["instance_id"] = _APP_NAME
	logMap := make(map[string]interface{})
	logMap["msg"] = "Table Error"
	logMap["server_id"] = serverId
	logMap["role_id"] = roleId
	logMap["account"] = account
	logMap["role_name"] = roleName
	logMap["error_id"] = protoErrId
	logData["log"] = logMap
	BiliLog_OpsLog(logData)
}

//-----------------------------------------------
//--------------      BG LOG      ---------------
//-----------------------------------------------

// BiliBgLog_JoinLevel 玩家进战斗（PVE）
func BiliBgLog_JoinLevel(roleId uint64, account string, serverId, levelId, oriTimes, activityId int32, oriStar []int32, dropList interface{}, curTime int64) {
	if !canLog {
		return
	}
	nowStr := Bili_GetLogTimeNow()
	logId := fmt.Sprintf("%s%d-%d-%d", _APP_NAME, serverId, roleId, time.Now().Unix()) // todo 这个回头再确定下？诸如限制长度、字母开头之类的
	logData := make(map[string]interface{})
	logData["time"] = nowStr
	logData["game_name"] = _APP_NAME
	logData["game_name_type"] = "kapai"
	logData["log_id"] = logId
	logData["env"] = "qa"
	logData["account"] = account
	logData["level_id"] = levelId
	logData["oritimes"] = oriTimes
	logData["activity_id"] = activityId
	logData["oristar"] = oriStar
	logData["drop_list"] = dropList
	logData["join_time"] = curTime
	BiliLog_BgLog(logData)
}
