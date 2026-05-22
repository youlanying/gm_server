package beegodb

import (
	"database/sql"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"
)

var sqliteBeegoDB *sql.DB

const beegoDbPath = "./data/beego.db"

var beegoCreateSQL = `
DROP TABLE IF EXISTS "tb_admin_user";
CREATE TABLE IF NOT EXISTS "tb_admin_user" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"userid" TEXT NOT NULL DEFAULT '0',
	"password" TEXT NOT NULL DEFAULT '0',
	"groupid" INTEGER NOT NULL DEFAULT 0,
	"remarks" TEXT NOT NULL DEFAULT '0'
);

INSERT INTO "tb_admin_user" ("id", "userid", "password", "groupid", "remarks") VALUES (1, 'admin', '21232f297a57a5a743894a0e4a801fc3', 1, '');

DROP TABLE IF EXISTS "tb_admin_group";
CREATE TABLE IF NOT EXISTS "tb_admin_group" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" TEXT NOT NULL DEFAULT '0',
	"actionlist" TEXT NOT NULL DEFAULT '[]'
);

INSERT INTO "tb_admin_group" ("id", "name", "actionlist") VALUES (1, '超级管理员', '[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,95,96,97,98,99,104,105,1000]');

DROP TABLE IF EXISTS "tb_center_list";
CREATE TABLE IF NOT EXISTS "tb_center_list" (
	"id" INTEGER NOT NULL DEFAULT 0,
	"name" TEXT NOT NULL DEFAULT '0',
	"ip" TEXT NOT NULL DEFAULT '0',
	"port" TEXT NOT NULL DEFAULT '0',
	"serverpath" TEXT NOT NULL DEFAULT '0',
	PRIMARY KEY ("id")
);

INSERT INTO "tb_center_list" ("id", "name", "ip", "port", "serverpath") VALUES (1000, '57开发服', '192.168.20.57', '2023', '');

DROP TABLE IF EXISTS "tb_send_mail";
CREATE TABLE IF NOT EXISTS "tb_send_mail" (
	"mailid" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" TEXT NOT NULL DEFAULT '0',
	"sendor" TEXT NOT NULL DEFAULT '0',
	"recvname" TEXT NOT NULL DEFAULT '0',
	"content" TEXT NOT NULL DEFAULT '0',
	"itemlist" TEXT NOT NULL DEFAULT '0',
	"isall" INTEGER NOT NULL DEFAULT 0,
	"createtime" TEXT NOT NULL DEFAULT '0',
	"audittime" TEXT NOT NULL DEFAULT '0',
	"status" INTEGER NOT NULL DEFAULT 0,
	"serverurl" TEXT NOT NULL DEFAULT '0',
	"sendid" TEXT NOT NULL DEFAULT '0',
	"auditid" TEXT NOT NULL DEFAULT '0',
	"timetype" INTEGER NOT NULL DEFAULT 0,
	"starttime" TEXT NOT NULL DEFAULT '0',
	"endtime" TEXT NOT NULL DEFAULT '0',
	"lvstart" INTEGER NOT NULL DEFAULT 0,
	"lvend" INTEGER NOT NULL DEFAULT 0,
	"sex" INTEGER NOT NULL DEFAULT 0,
	"reason" TEXT NOT NULL DEFAULT '0'
);

DROP TABLE IF EXISTS "tb_publicnotice";
CREATE TABLE IF NOT EXISTS "tb_publicnotice" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"platformid" INTEGER NOT NULL DEFAULT 0,
	"serverid" INTEGER NOT NULL DEFAULT 0,
	"noticetype" INTEGER NOT NULL DEFAULT 0,
	"label" INTEGER NOT NULL DEFAULT 0,
	"priority" INTEGER NOT NULL DEFAULT 0,
	"titleshort" TEXT NOT NULL DEFAULT '0',
	"title" TEXT NOT NULL DEFAULT '0',
	"content" TEXT NOT NULL DEFAULT '0',
	"starttime" INTEGER NOT NULL DEFAULT 0,
	"endtime" INTEGER NOT NULL DEFAULT 0,
	"createtime" INTEGER NOT NULL DEFAULT 0,
	"audittime" INTEGER NOT NULL DEFAULT 0,
	"sendid" TEXT NOT NULL DEFAULT '0',
	"auditid" TEXT NOT NULL DEFAULT '0'
);

DROP TABLE IF EXISTS "tb_send_marquee";
CREATE TABLE IF NOT EXISTS "tb_send_marquee" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"mtype" INTEGER NOT NULL DEFAULT 0,
	"num" INTEGER NOT NULL DEFAULT 0,
	"content" TEXT NOT NULL DEFAULT '0',
	"createtime" TEXT NOT NULL DEFAULT '0',
	"audittime" TEXT NOT NULL DEFAULT '0',
	"endtime" TEXT NOT NULL DEFAULT '0',
	"sendok" TEXT NOT NULL DEFAULT '0',
	"serverurl" TEXT NOT NULL DEFAULT '0',
	"sendid" TEXT NOT NULL DEFAULT '0',
	"auditid" TEXT NOT NULL DEFAULT '0'
);
`

func initBeegoSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", beegoDbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite3 fail: %v", err)
	}

	statements := strings.Split(beegoCreateSQL, ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db.Exec(stmt)
		if err != nil {
			sqlPreview := stmt
			if len(sqlPreview) > 100 {
				sqlPreview = sqlPreview[:100]
			}
			return nil, fmt.Errorf("exec sql fail: %v, sql: %s", err, sqlPreview)
		}
	}

	return db, nil
}
