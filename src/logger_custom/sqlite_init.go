package logger_custom

import (
	"database/sql"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"
)

const loggerDbPath = "./data/logger.db"

var loggerCreateSQL = `
DROP TABLE IF EXISTS "tb_log_create_role";
CREATE TABLE IF NOT EXISTS "tb_log_create_role" (
	"logid" TEXT PRIMARY KEY,
	"roleid" INTEGER NOT NULL DEFAULT 0,
	"account" TEXT NOT NULL DEFAULT '',
	"serverid" INTEGER NOT NULL DEFAULT 0,
	"ip" TEXT NOT NULL DEFAULT '',
	"mac" TEXT NOT NULL DEFAULT '',
	"devicetype" TEXT NOT NULL DEFAULT '',
	"versionno" TEXT NOT NULL DEFAULT '',
	"cid" TEXT NOT NULL DEFAULT '',
	"createtime" INTEGER NOT NULL DEFAULT 0
);

DROP TABLE IF EXISTS "tb_log_instance_join";
CREATE TABLE IF NOT EXISTS "tb_log_instance_join" (
	"logid" TEXT PRIMARY KEY,
	"roleid" INTEGER NOT NULL DEFAULT 0,
	"account" TEXT NOT NULL DEFAULT '',
	"serverid" INTEGER NOT NULL DEFAULT 0,
	"instanceid" INTEGER NOT NULL DEFAULT 0,
	"activityid" INTEGER NOT NULL DEFAULT 0,
	"daytimes" INTEGER NOT NULL DEFAULT 0,
	"star" TEXT NOT NULL DEFAULT '',
	"createtime" INTEGER NOT NULL DEFAULT 0
);

DROP TABLE IF EXISTS "tb_log_instance_out";
CREATE TABLE IF NOT EXISTS "tb_log_instance_out" (
	"logid" TEXT PRIMARY KEY,
	"instanceid" INTEGER NOT NULL DEFAULT 0,
	"serverid" INTEGER NOT NULL DEFAULT 0,
	"roleid" INTEGER NOT NULL DEFAULT 0,
	"account" TEXT NOT NULL DEFAULT '',
	"rolelevel" INTEGER NOT NULL DEFAULT 0,
	"state" INTEGER NOT NULL DEFAULT 0,
	"daytimes" INTEGER NOT NULL DEFAULT 0,
	"endgold" INTEGER NOT NULL DEFAULT 0,
	"roleexp" INTEGER NOT NULL DEFAULT 0,
	"heroexp" INTEGER NOT NULL DEFAULT 0,
	"objectitems" TEXT NOT NULL DEFAULT '',
	"star" TEXT NOT NULL DEFAULT '',
	"deadhero" TEXT NOT NULL DEFAULT '',
	"battleherolist" TEXT NOT NULL DEFAULT '',
	"timeconsum" INTEGER NOT NULL DEFAULT 0,
	"createtime" INTEGER NOT NULL DEFAULT 0
);

DROP TABLE IF EXISTS "tb_log_role_login";
CREATE TABLE IF NOT EXISTS "tb_log_role_login" (
	"logid" TEXT PRIMARY KEY,
	"roleid" INTEGER NOT NULL DEFAULT 0,
	"account" TEXT NOT NULL DEFAULT '',
	"name" TEXT NOT NULL DEFAULT '',
	"serverid" INTEGER NOT NULL DEFAULT 0,
	"ip" TEXT NOT NULL DEFAULT '',
	"mac" TEXT NOT NULL DEFAULT '',
	"devicetype" TEXT NOT NULL DEFAULT '',
	"versionno" TEXT NOT NULL DEFAULT '',
	"cid" TEXT NOT NULL DEFAULT '',
	"createtime" INTEGER NOT NULL DEFAULT 0
);

DROP TABLE IF EXISTS "tb_log_terminate";
CREATE TABLE IF NOT EXISTS "tb_log_terminate" (
	"logid" TEXT PRIMARY KEY,
	"roleid" INTEGER NOT NULL DEFAULT 0,
	"account" TEXT NOT NULL DEFAULT '',
	"name" TEXT NOT NULL DEFAULT '',
	"serverid" INTEGER NOT NULL DEFAULT 0,
	"errorproto" TEXT NOT NULL DEFAULT '',
	"reason" TEXT NOT NULL DEFAULT '',
	"createtime" INTEGER NOT NULL DEFAULT 0
);
`

func initLoggerSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", loggerDbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite3 fail: %v", err)
	}

	statements := strings.Split(loggerCreateSQL, ";")
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
