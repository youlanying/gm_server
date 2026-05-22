// Package services implements boilerplate code for all services.
package services

import (
	"gm_server/src/cfg"
)

//** TYPES

var defaultDbName string
var URL_MODE string

// Service contains common properties for all services.
//type Service struct {
//	MongoSession *mgo.Session
//	UserID       string
//}

func init() {
	config := cfg.GetSection("BEEGOWEB")
	defaultDbName = config["mongo_dbname"]
	URL_MODE = config["url_mode"]
}

//** PUBLIC FUNCTIONS

// Prepare is called before any controller.
//func (service *Service) Prepare() (err error) {
//	logger.Log("func (service *Service) Prepare() (err error) {")
//	service.MongoSession, err = mongo.CopyMonotonicSession(service.UserID)
//	if err != nil {
//		return err
//	}
//
//	return err
//}
//
//// Finish is called after the controller.
//func (service *Service) Finish() (err error) {
//	defer helper.CatchPanic(&err, service.UserID, "Service.Finish")
//
//	if service.MongoSession != nil {
//		mongo.CloseSession(service.UserID, service.MongoSession)
//		service.MongoSession = nil
//	}
//
//	return err
//}
//
//// DBAction executes the MongoDB literal function
//func (service *Service) DBAction(databaseName string, collectionName string, dbCall mongo.DBCall) (err error) {
//	if databaseName == "" {
//		defaultDbName = defaultDbName
//	}
//	return mongo.Execute(service.UserID, service.MongoSession, databaseName, collectionName, dbCall)
//}
