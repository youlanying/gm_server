package cfg

import (
	"flag"
	"fmt"
	"gm_server/src/tools/goini"
	"os"
	"path"
	"sync"
)

var (
	_lock sync.RWMutex
	_file goini.File

	_defServerId = 1 //TODO 0
	_basePath    = path.Join(path.Dir(os.Args[0]), "../")
	//_basePath    = "/data/app/"
	//_defaultFile = path.Join(_basePath, "configs/config.ini")
	_defaultFile = _basePath + "configs/config.ini"
	_configFile  = flag.String("config", _defaultFile, "config filename")
	_serverFlag  = flag.Int("s", _defServerId, "server flag")
)

func init() {
	flag.Parse()
	Reload()

}

func GetBasePath() string {
	return _basePath
}

//	func main() {
//		config := getSection("GLOBAL")
//		_serverId := config["server_id"]
//		_serverName := config["server_name"]
//		if _serverId == "" {
//			//cfglog
//			fmt.Println("server_id not find.")
//		}
//		fmt.Println(" serverId : " + _serverId + " serverName : " + _serverName)
//	}
func Reload() {
	if isFileExist(_defaultFile) == false {
		//_basePath = path.Join(path.Dir(os.Args[0]), "../../")
		//_basePath = path.Join(path.Dir(os.Args[0]))
		if *_configFile == _defaultFile {
			_defaultFile = path.Join(_basePath, "configs/config.ini")
			_configFile = &_defaultFile
		} else {
			_defaultFile = path.Join(_basePath, "configs/config.ini")
		}
	}

	tempPath := *_configFile
	//fmt.Printf("Loading Config: %v.\n", tempPath)
	_lock.Lock()
	_file = loadConfig(tempPath)
	_lock.Unlock()
}

func GetSection(sectionName string) map[string]string {
	_lock.RLock()
	defer _lock.RUnlock()
	return _file.Section(sectionName)
}

func GetGsSectionValue(key string) string {
	if *_serverFlag <= 0 {
		panic("Invalid gameserver -s !")
	}

	tempSid := fmt.Sprintf("GS%d", *_serverFlag)
	conf := GetSection(tempSid)
	str := conf[key]
	if str != "" {
		fmt.Println("GetGsSectionValue[", tempSid, "] [", str, "] = ", str)
		return str
	}

	// 没有指定配置，则直接使用模板 GS 配置。
	configTemp := GetSection("GS_TEMP")
	str = configTemp[key]
	strVec := []rune(str)
	var tmpStr string
	startIndex := 0
	for runeIdx, c := range strVec {
		if c == '%' {
			tmpStr = tmpStr + string(strVec[startIndex:runeIdx])
			startIndex = runeIdx + 1
			tmpStr = fmt.Sprintf("%s%d", tmpStr, *_serverFlag)
			break
		}
	}
	tmpStr = tmpStr + string(strVec[startIndex:])

	str = tmpStr
	return str
}

func isFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func loadConfig(tpath string) goini.File {
	var err error
	//fmt.Println("Loading Config.")//必须注释
	file, err := goini.LoadFile(tpath)
	if err != nil {
		file, err = goini.LoadFile("configs/config.ini")
		if err != nil {
			//fmt.Printf("load [%s] error: %s", tpath, err)
			os.Exit(-1)
		}

	}
	return file
}

func GetGmCmdCfg(key string) (str string) {

	configTemp := GetSection("GM_CMD")
	str = configTemp[key]
	strVec := []rune(str)
	var tmpStr string
	startIndex := 0
	for runeIdx, c := range strVec {
		if c == '%' {
			tmpStr = tmpStr + string(strVec[startIndex:runeIdx])
			startIndex = runeIdx + 1
			tmpStr = fmt.Sprintf("%s%d", tmpStr, *_serverFlag)
			break
		}
	}
	tmpStr = tmpStr + string(strVec[startIndex:])

	str = tmpStr
	return str
}
