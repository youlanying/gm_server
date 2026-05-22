package usersService

//
//import (
//	"beegoWeb/models/userModels"
//	"beegoWeb/services"
//	"fmt"
//
//	"gopkg.in/mgo.v2/bson"
//)
//
//const (
//	USERS_DB    = "account"
//	COLLECTION  = "accounts"
//	PAGE_LIMIT  = 100
//	MAXN_SERVER = 9
//)
//
//func FindUserList(service *services.Service, pageid int16) (*userModels.UserListInfo, error) {
//
//	var list userModels.UserListInfo
//	list.Info = make([]userModels.UserInfo, 0)
//	if err := service.DBAction(USERS_DB, COLLECTION,
//		func(collection *mgo.Collection) error {
//			queryMap := bson.M{}
//			if pageid > 1 {
//				return collection.Find(queryMap).Sort("id").Skip(int(pageid-1) * PAGE_LIMIT).Limit(PAGE_LIMIT).All(&list.Info)
//			} else {
//				return collection.Find(queryMap).Sort("id").Limit(PAGE_LIMIT).All(&list.Info)
//			}
//		}); err != nil {
//		return nil, err
//	}
//	for i, info := range list.Info {
//		name := "noname"
//		if str, err := get_username(service, info.Id, info.Server_id); err == nil {
//			name = str
//		}
//		list.Info[i].Name = name
//	}
//	return &list, nil
//
//}
//
//func get_username(service *services.Service, uid int64, gsid int16) (string, error) {
//	username := make(map[string]string)
//	var database string
//	database = fmt.Sprintf("gs%d", gsid)
//	if err := service.DBAction(database, "users",
//		func(collection *mgo.Collection) error {
//			queryMap := bson.M{"id": uid}
//			return collection.Find(queryMap).One(username)
//		}); err != nil {
//		return "", err
//	}
//	return username["username"], nil
//}
//
//type ServerId struct {
//	GsId int16 `bson:"server_id"`
//}
//
//func SearchUserByName(service *services.Service, username string) userModels.SnapUserList {
//	var list userModels.SnapUserList
//	var gsInfo ServerId
//	list.Info = make([]userModels.SnapUser, 0)
//	serverNum, _ := getGsNum(service)
//	count := int16(0)
//	for i := 1; i <= serverNum; i++ {
//		if q, err := searchGSByName(service, username, int16(i)); err == nil {
//			for _, info := range q {
//				user := new(userModels.SnapUser)
//				user.Id = info["id"].(int64)
//				user.Name = info["username"].(string)
//				list.Info = append(list.Info, *user)
//				count++
//				if count >= PAGE_LIMIT {
//					break
//				}
//			}
//
//		}
//	}
//
//	for i, info := range list.Info {
//		service.DBAction(USERS_DB, COLLECTION,
//			func(collection *mgo.Collection) error {
//				queryMap := bson.M{"id": info.Id}
//				return collection.Find(queryMap).One(&gsInfo)
//			})
//		list.Info[i].ServerId = fmt.Sprintf("gs%d", gsInfo.GsId)
//	}
//	return list
//}
//
//func searchGSByName(service *services.Service, username string, gsid int16) ([]bson.M, error) {
//	var database string
//	var q []bson.M
//	database = fmt.Sprintf("gs%d", gsid)
//	if err := service.DBAction(database, "users",
//		func(collection *mgo.Collection) error {
//			queryMap := bson.M{"username": bson.M{"$regex": username}}
//			return collection.Find(queryMap).Select(bson.M{"id": 1, "username": 1}).All(&q)
//		}); err != nil {
//
//		return nil, err
//	}
//	return q, nil
//
//}
//func getGsNum(service *services.Service) (int, error) {
//	var q []bson.M
//	if err := service.DBAction(USERS_DB, "servers",
//		func(collection *mgo.Collection) error {
//			queryMap := bson.M{}
//			return collection.Find(queryMap).All(&q)
//		}); err != nil {
//
//		return 0, err
//	}
//	return len(q), nil
//
//}
//
//func SearchUserByUserId(service *services.Service, userid int64) userModels.SnapUserList {
//	var list userModels.SnapUserList
//	var gsInfo ServerId
//	list.Info = make([]userModels.SnapUser, 0)
//	serverNum, _ := getGsNum(service)
//	count := int16(0)
//	for i := 1; i <= serverNum; i++ {
//		if q, err := searchGSByUserId(service, userid, int16(i)); err == nil {
//			for _, info := range q {
//				user := new(userModels.SnapUser)
//				user.Id = info["id"].(int64)
//				user.Name = info["username"].(string)
//				list.Info = append(list.Info, *user)
//				count++
//				if count >= PAGE_LIMIT {
//					break
//				}
//			}
//
//		}
//	}
//
//	for i, info := range list.Info {
//		service.DBAction(USERS_DB, COLLECTION,
//			func(collection *mgo.Collection) error {
//				queryMap := bson.M{"id": info.Id}
//				return collection.Find(queryMap).One(&gsInfo)
//			})
//		list.Info[i].ServerId = fmt.Sprintf("gs%d", gsInfo.GsId)
//	}
//	return list
//}
//
//func searchGSByUserId(service *services.Service, userid int64, gsid int16) ([]bson.M, error) {
//	var database string
//	var q []bson.M
//	database = fmt.Sprintf("gs%d", gsid)
//	if err := service.DBAction(database, "users",
//		func(collection *mgo.Collection) error {
//			queryMap := bson.M{"id": userid}
//			return collection.Find(queryMap).Select(bson.M{"id": 1, "username": 1}).All(&q)
//		}); err != nil {
//
//		return nil, err
//	}
//	return q, nil
//
//}
//
//func FindUserAccountInfo(service *services.Service, userid int64) *userModels.AccountInfo {
//	info := userModels.AccountInfo{}
//	if err := service.DBAction(USERS_DB, COLLECTION,
//		func(collection *mgo.Collection) error {
//			queryMap := bson.M{"id": userid}
//			return collection.Find(queryMap).One(&info)
//		}); err != nil {
//		return nil
//	}
//	return &info
//}
