package models

import (
	"time"
)

/**
 * linux版本 本地数据库的 账号密码
 */
var (
	dbuser   string
	dbpasswd string
	dbhost   string
	dbport   string
	dbname   string
	poolsize int
)

//var LocalMySql string = "root:oneinstack@tcp(192.168.20.47:3306)";
var LocalMySql string

var strallcenter string
var AllCenterConf []string

//var all_center []string = strings.Split(strallcenter, ",")

//---- 和数据库相关的结构 和 数据库操作封装在这里 ----

//---- 数据库里面的 表结构放在这里 ----

type Tb_server_version_log struct {
	Id            int       `orm:"column(id);pk"`
	Serverid      int       `orm:"column(serverid)"`
	Serverversion string    `orm:"size(200);column(serverversion)"`
	Versiondate   time.Time `orm:"type(datetime);column(versiondate)"`
}

type Tb_openserver_log struct {
	Id             int       `orm:"column(id)"`
	Serverid       int       `orm:"column(serverid)"`
	Platform       string    `orm:"size(200);column(platform)"`
	Serverip       string    `orm:"size(200);column(serverip)"`
	Serveropentime time.Time `orm:"type(datetime);column(serveropentime)"`
}

type Tb_mutex struct {
	Id          int    `orm:"column(id)"`
	Mutex_id    string `orm:"size(4);column(mutex_id)"`
	Mutex_gifts string `orm:"size(200);column(mutex_gifts)"`
}

type Tb_gift struct {
	Id        int    `orm:"column(id)"`
	Gift_id   string `orm:"size(4);column(gift_id)"`
	Gift_name string `orm:"size(20);column(gift_name)"`
	Gift_item int    `orm:"column(gift_item)"`
}

type Tb_channel struct {
	Id           int    `orm:"column(id)"`
	Channel_id   string `orm:"size(4);column(channel_id)"`
	Channel_name string `orm:"size(30);column(channel_name)"`
}

type Noticelist struct {
	Id      int    `orm:"column(id)"`
	Title   string `orm:"size(50);column(title)"`
	Descstr string `orm:"size(200);column(descstr)"`
	P1      string `orm:"size(50);column(p1)"`
	P2      string `orm:"size(50);column(p2)"`
	P3      string `orm:"size(50);column(p3)"`
	P4      string `orm:"size(50);column(p4)"`
}

type Noticeread struct {
	Id       int    `orm:"column(id)"`
	Noticeid int    `orm:"column(noticeid)"`
	Pid      int    `orm:"column(pid)"`
	Op       string `orm:"size(20);column(op)"`
}
