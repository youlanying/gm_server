package models

import "gm_server/src/beegoWeb/beegodb"

//---- 和客户端通信 或者逻辑模块需要的 自定义结构放在这里 ----

/**
 * 通信时的 逻辑结构
 */
type User struct {
	ID       int    `sql:"id"`
	UserID   string `sql:"userid"`
	Password string `sql:"password"` //是字符串 而 不是  []string
	GroupID  int    `sql:"groupid"`
}

/**
 * 通信时的 逻辑结构
 */
type GroupMem struct {
	Gid      int32
	Gname    string
	Gactlist string //是字符串 而 不是  []string
}

type AllPlantServer struct {
	Platform     string
	PlatformName string
	ServerList   []*OneServer
}

type OneServer struct {
	Sviewid      int
	Sserverid    int
	Sserver_name string
	Sstatus      int
	Sgateips     string
	Smaxnum      int
	Scmdurl      string
	Sdbid        int
}

/**
 * 其实 就是 make_proto_admin_action
 */
type AdminAction struct {
	ID   int
	Name string
}

type Plant struct {
	PlatformName    string
	Dbhost          string
	Dbport          int
	Dbuser          string
	Dbpasswd        string
	Center_node_api string
	Dbname          string
	Purview         int
}

/**
 * 继承的知识
 * 貌似 只能集合进来
 * 看MailMem
 */
type PlantDetail struct {
	Plant      *Plant
	ServerList []OneServer //这个属性是 查询数据库后 获得的
}

/**
 * http result
 */
type HttpResult struct {
	Cmd    string "json:'cmd'" //一定要大写 不然反射不出来 linzi add at 2020 0119
	Result string "json:'result'"
	State  string "json:'state'"
}

/**
 * 真正的继承写法
 * send_mail 循环需要的结构体
 */
type MailMem struct {
	beegodb.TB_SEND_MAIL
	Disabled string
	Style    string
	Mclass   string
	Sendok   string
	Recvname string
}

/**
 * 真正的继承写法
 * send_mail 循环需要的结构体
 */
type PublicNoticeMem struct {
	beegodb.TB_PUBLICNOTICE
	Disabled string
	Style    string
	Mclass   string
	Sendok   string
}

/**
 * 跑马灯信息
 * 客户端用
 */
type MarInfo struct {
	beegodb.TB_SEND_MARQUEE
	Disabled string
	Style    string
	Mclass   string
	TypeStr  string //由 1 2转化成 普通跑马灯 重要中央提示
}

type ProtoData struct {
	//Tb_proto_data
	CreatetimeStr string
	SendtimeStr   string
}

type ProtoTable struct {
	//Tb_proto_table
}

type ServerVersion struct {
	Id        int32
	CName     string
	Dir       string
	VersionId string
}

/**
 * 表头中 一列的表头
 */
type TableTitle struct {
	CS   string
	Key  string
	Type string
}

func StrToTableTitle(str string) []TableTitle {
	return nil
}
