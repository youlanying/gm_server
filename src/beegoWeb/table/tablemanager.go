package table

import (
	"gm_server/src/logger"
	"io/ioutil"
	"os"
	"path/filepath"
)

type TableMgrInterface interface {
	PrintArr()
	PrintArrOne(index int)
	PrintMapByKey(key interface{})
	Load(vBuffer []byte) bool //加载
}

const OtherTableLen int = 0

const configFileSizeLimit = 10 << 20

var Dic map[string]TableMgrInterface

var PROTO_ERROR_ID string //记录错误表信息

func Register(key string, val TableMgrInterface) {
	Dic[key] = val
}

var ConfigFile string

func LoadAllTable() bool {

	//读取目录下文件
	var err error

	_basePath, _ := os.Getwd()
	_basePath = filepath.Join(_basePath, "beegoWeb", "config", "json")
	_, err = ioutil.ReadDir(_basePath)

	if err != nil {
		str, _ := os.Getwd()
		logger.LogErrf("read Dir error:%v,nowDir:%v ", err.Error(), str)
		return false
	}
	for key, mem := range Dic {
		jsonPath := filepath.Join(_basePath, key+".json")
		_, err := os.Stat(jsonPath)
		if os.IsNotExist(err) {
			logger.LogWarnf("LoadTable %v config file not found, skipping", key)
			continue
		}
		result := LoadTable(key, jsonPath, mem)
		if result == false {
			return false
		}
	}
	logger.LogWarn("---- ReLoadAllTable Success ----")
	return true
}

func LoadTable(name string, path string, mgr TableMgrInterface) bool {

	configFile, err := os.Open(path)
	if err != nil {
		logger.LogErrf("LoadTable %v Failed to open config file '%s': %s", name, path, err)
		return false
	}

	fi, _ := configFile.Stat()
	if size := fi.Size(); size > (configFileSizeLimit) {
		logger.LogErrf("LoadTable %v config file (%q) size exceeds reasonable limit (%d) - aborting", name, path, size)
		return false // REVUE: shouldn't this return an error, then?
	}

	if fi.Size() == 0 {
		logger.LogErrf("LoadTable %v config file (%q) is empty, skipping", name, path)
		return false
	}

	buffer := make([]byte, fi.Size())
	_, err = configFile.Read(buffer)
	if err != nil {
		logger.LogErrf("LoadTable %v configFile.Read err: %s", name, err)
		return false
	}
	buffer = []byte(os.ExpandEnv(string(buffer))) //特殊

	//解析json格式数据
	result := mgr.Load(buffer)

	return result
}

func initOther() {
	//这里新增要修改 上面的 OtherTableLen
	//Register("action",&ActionModel)
	//Register("collider",&ColliderInfoModel)
}
