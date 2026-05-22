package authService

//import (
//	"beegoWeb/models/authModels"
//	"beegoWeb/services"
//	"beegoWeb/utilities/mongo"
//
//	"gopkg.in/mgo.v2/bson"
//)
//
//const (
//	WEB_DB                   = ""
//	FORBIDED_CMD_COLLECTION  = "forbided_cmd"
//	CONFIG_COLLECTION        = "gate_config"
//	CONFIG_TYPE_UDID1_SWITCH = 2
//)
//
//var (
//	gateConfig     map[int]string //gate config
//	_udidWriteBack map[string]string
//	_forbidedCmds  map[string]bool //key:url
//)
//
//func init() {
//	gateConfig = make(map[int]string)
//	_udidWriteBack = make(map[string]string)
//	_forbidedCmds = make(map[string]bool)
//}

//func LoadUdidCfg(dbname string) error {
//	logger.Log("func LoadUdidCfg(dbname string) error {")
//	session, err := mongo.CopyMonotonicSession("init")
//	if err != nil {
//		return err
//	}
//	defer mongo.CloseSession("init", session)
//
//	var q []bson.M
//	coll := session.DB(dbname).C(CONFIG_COLLECTION)
//	coll.Find(nil).All(&q)
//	for _, info := range q {
//		kv := *m2Info(info)
//		gateConfig[kv.Key] = kv.Value
//	}
//
//	return nil
//}
//
//func m2Info(doc bson.M) *authModels.KV {
//	info := new(authModels.KV)
//	info.Key = doc["config_type"].(int)
//	info.Value = doc["value"].(string)
//	return info
//}
//
//func GetGateConfig(typ int) string {
//	if s, exist := gateConfig[typ]; exist {
//		return s
//	}
//
//	return ""
//}
//
//func AddUdidWriteback(username string, udid string) {
//	_udidWriteBack[username] = udid
//}
//
//func GetUdidWriteback(username string) (string, bool) {
//	udid1, exist := _udidWriteBack[username]
//	return udid1, exist
//}
//
//func LoadForbidCmd(dbname string) error {
//	logger.Log("func LoadForbidCmd(dbname string) error {")
//	session, err := mongo.CopyMonotonicSession("init")
//	if err != nil {
//		return err
//	}
//	defer mongo.CloseSession("init", session)
//
//	var q []bson.M
//	coll := session.DB(dbname).C(FORBIDED_CMD_COLLECTION)
//	coll.Find(nil).All(&q)
//	for _, info := range q {
//		s := info["cmd"].(string)
//		_forbidedCmds[s] = true
//	}
//	return nil
//}
//
//func DelForbidCmd(service *services.Service, cmd string) (bool, error) {
//	if err := service.DBAction(WEB_DB, FORBIDED_CMD_COLLECTION,
//		func(collection *mgo.Collection) error {
//			cond := bson.M{"cmd": cmd}
//			err := collection.Remove(cond)
//			return err
//		}); err != nil {
//		return false, err
//	}
//	delete(_forbidedCmds, cmd)
//	return true, nil
//}
//
//func AddForbidCmd(service *services.Service, cmd string) (bool, error) {
//	if err := service.DBAction(WEB_DB, FORBIDED_CMD_COLLECTION,
//		func(collection *mgo.Collection) error {
//			cond := bson.M{"cmd": cmd}
//			_, err := collection.Upsert(cond, bson.M{"$set": cond})
//			return err
//		}); err != nil {
//		return false, err
//	}
//	// 更新cache
//	_forbidedCmds[cmd] = true
//
//	return true, nil
//}
//
//func GetForbidCmd() []string {
//	ret := make([]string, 0)
//	for url, _ := range _forbidedCmds {
//		ret = append(ret, url)
//	}
//	return ret
//}
//
//func IsForbidCmd(cmd string) bool {
//	_, exist := _forbidedCmds[cmd]
//	return exist
//}
//
//func SetUdidWritebackVal(service *services.Service, val string) error {
//	if err := service.DBAction(WEB_DB, CONFIG_COLLECTION,
//		func(collection *mgo.Collection) error {
//			cond := bson.M{"config_type": CONFIG_TYPE_UDID1_SWITCH}
//			doc := bson.M{
//				"config_type": CONFIG_TYPE_UDID1_SWITCH,
//				"value":       val,
//			}
//			_, err := collection.Upsert(cond, doc)
//			return err
//		}); err != nil {
//		return err
//	}
//
//	// 更新cache.
//	gateConfig[CONFIG_TYPE_UDID1_SWITCH] = val
//	return nil
//}
//
//func GetUdidWritebackVal() string {
//	v, p := gateConfig[CONFIG_TYPE_UDID1_SWITCH]
//	if p {
//		return v
//	}
//	return "false"
//}
