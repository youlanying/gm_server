package tablemanager

import (
	"gm_server/src/logger"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type TableMgrInterface interface {
	PrintArr()
	PrintArrOne(index int)
	PrintMapByKey(key interface{})
	Load(vBuffer []byte) bool //加载
}

const OtherTableLen int = 2

const configFileSizeLimit = 10 << 20

var Dic map[string]TableMgrInterface

var PROTO_ERROR_ID string //记录错误表信息

func Register(key string, val TableMgrInterface) {
	Dic[key] = val
}

/**
 * 这个方法 需要主动去调用
 * 如果 服务器读取配置路径失败 那么服务器不该运行
 */
// LoadAllTable 加载表
func LoadAllTable() bool {
	//记录错误表信息
	PROTO_ERROR_ID = "null"

	//读取目录下文件
	_basePath := path.Join(path.Dir(os.Args[0]), "../")
	pathPre := _basePath + "/configs/tableJson/"
	//battleFile := _basePath + "/configs/battleJson/"

	_, err := ioutil.ReadDir(pathPre)
	if err != nil {
		str, _ := os.Getwd()
		logger.LogErrf("read Dir error:%v,nowDir:%v \n", err.Error(), str)
		return false
	}
	logger.LogWarn("======load==========open====")
	for key, mem := range Dic {
		jsonPath := filepath.Join(pathPre, key+".json")
		result := LoadTable(key, jsonPath, mem)
		if result == false {
			return false
		}
	}
	logger.LogWarn("======load===Dic=======over====")
	//不启用battle服 暂时注释掉
	//actionPath := battleFile + "action.json"
	//result := LoadTable("action", actionPath, &ActionModel)
	//if result == false {
	//	return false
	//}
	//logger.LogWarn("======load===action=======over====")
	//
	//colliderPath := battleFile + "collider.json"
	//result = LoadTable("collider", colliderPath, &ColliderInfoModel)
	//if result == false {
	//	return false
	//}
	//logger.LogWarn("======load===collider=======over====")

	logger.LogWarn("---- ReLoadAllTable Success ----")
	return true
}

func LoadTable(name string, path string, mgr TableMgrInterface) bool {
	result := false
	configFile, err := os.Open(path)
	if err != nil {
		logger.LogErrf("LoadTable Failed to open config %s file '%s': %s", name, path, err)
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
		logger.LogErrf("LoadTable configFile.Read %v error: %v", name, err)
		return false
	}
	buffer = []byte(os.ExpandEnv(string(buffer))) //特殊
	logger.LogWarnf("====LoadTable=========Load===name:%v", name)
	//解析json格式数据
	result = mgr.Load(buffer)
	//if name == "npc_object" {
	//	npcObject()
	//}
	return result
}

func initOther() {
	//这里新增要修改 上面的 OtherTableLen
	//Register("action",&ActionModel)
	//Register("collider",&ColliderInfoModel)
}

func toLinux(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

func rootPath() string {
	var fp, _ = filepath.Abs(path.Dir(os.Args[0]))
	return fp
}

//func ToAbsolutePath(p string) string {
//	//if "" == p || "." == p {
//	//	return toLinux(rootPath())
//	//}
//	fmt.Println("1111111111111111111111111111111 rootPath()", rootPath())
//	fmt.Println("1111111111111111111111111111111 p", p)
//	var linuxPath = toLinux(p)
//	var paths = strings.Split(linuxPath, "/")
//
//	var rp string
//	if 0 < len(paths) {
//		switch paths[0] {
//		case "":
//			fallthrough
//		case ".":
//			fallthrough
//		case "..":
//			rp = toLinux(rootPath())
//		}
//	}
//	var realPaths []string
//	if "" != rp {
//		realPaths = strings.Split(rp, "/")
//	}
//	if 0 < len(paths) {
//		realPaths = append(realPaths, paths...)
//	}
//
//	return path.Join(realPaths...)
//}

//func Load(vbuffer *[]byte, mgr *hero_informationMgr) bool {
//	buffer := *vbuffer
//	err := json.Unmarshal(buffer, &Hero_information_All)
//	if err != nil {
//		fmt.Println("JsonFailed", err)
//		return false
//	}
//	vLen := len(Hero_information_All)
//	hero_informationDict = make(map[int32]*Hero_information, vLen)
//	for _, mem := range Hero_information_All {
//		hero_informationDict[mem.Id] = mem
//	}
//	return true
//}

//// NpcObjectLevel 以关卡ID为Key的map
//var _NpcObjectLevel map[int32][]*Npc_object
//
//// NpcObjectLevelMap var NpcObjectMap =  make(map[int32]*Npc_object)
//var _NpcObjectLevelMap map[int32]map[string]*Npc_object
//
//func npcObject() {
//	_NpcObjectLevel = make(map[int32][]*Npc_object)
//	_NpcObjectLevelMap = make(map[int32]map[string]*Npc_object)
//	for _, v := range Npc_object_All {
//		noList, ok := _NpcObjectLevel[v.Instance_level_id]
//		if ok {
//			noList = append(noList, v)
//			_NpcObjectLevel[v.Instance_level_id] = noList
//		} else {
//			var realList []*Npc_object
//			realList = append(realList, v)
//			_NpcObjectLevel[v.Instance_level_id] = realList
//		}
//		noList1, ok1 := _NpcObjectLevelMap[v.Instance_level_id]
//		if ok1 {
//			noList1[v.Id] = v
//			_NpcObjectLevelMap[v.Instance_level_id] = noList1
//		} else {
//			noList1 := make(map[string]*Npc_object)
//			noList1[v.Id] = v
//			_NpcObjectLevelMap[v.Instance_level_id] = noList1
//		}
//	}
//}

//func GetNpcObjectLevel(levelId int32) (allNpcObjectList []*Npc_object, ok bool) {
//	allNpcObjectList, ok = _NpcObjectLevel[levelId]
//	if !ok {
//		PROTO_ERROR_ID = "关卡表-场景npc实体表.xlsx\nnpc_object not instance_level_id：" + strconv.FormatInt(int64(levelId), 10)
//		return
//	}
//	return
//}
