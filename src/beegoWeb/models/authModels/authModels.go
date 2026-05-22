package authModels

import "gm_server/src/beegoWeb/models/netModels"

type KV struct {
	Key   int    `bson:"config_type" json:"config_type"`
	Value string `bson:"value" json:"value"`
}
type AuthUrlInfo struct {
	AuthUrl string `json:"AuthAddr"`
}

type UdidWriteback struct {
	Username string `json:"user_name"`
	OpenUdid string `json:"udid"`
}

type RoleToken struct {
	TokenStr string `json:"token"`
}

type RoleList struct {
	Code  int16      `json:"code"`
	Roles []RoleInfo `json:"roles"`
}

type AllianceBrief struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Abbr  string `json:"abbr"`
	Rank  int16  `json:"rank"`
	Color int16  `json:"color"`
	Title int16  `json:"title"`
}

type RoleInfo struct {
	UserId    int64         `json:"user_id"`
	ServerId  int16         `json:"server_id"`
	Username  string        `json:"user_name"`
	Icon      int16         `json:"icon"`
	Avatar    string        `json:"avatar"`
	Power     int64         `json:"power"`
	Alliance  AllianceBrief `json:"alliance"`
	Forbidden bool          `json:"Forbidden"`
}

type BindInfo struct {
	TokenStr  string `json:"token"`
	Username  string `json:"username"`
	Pwd       string `json:"pwd"`
	AppId     string `json:"appid"`
	Link      int    `json:"link"`
	UserId    int64  `json:"user_id"`
	EnterType int8   `json:"enter_type"`
}

type BindRet struct {
	Code  int16  `json:"code"`
	Token string `json:"token"`
}

type TilePos struct {
	X int16 `json:"x"`
	Y int16 `json:"y"`
}

type ServerList struct {
	Code    int16        `json:"code"`
	Regions []RegionInfo `json:"region_list"`
	Servers []ServerInfo `json:"list"`
}

type ServerInfo struct {
	ServerId   int16   `json:"id"`
	Stat       int16   `json:"stat"`
	Opentime   int64   `json:"open_time"`
	Population int16   `json:"population"`
	Min        TilePos `json:"min"`
	Max        TilePos `json:"max"`
	Flag       int8    `json:"flag"`
}

type RegionInfo struct {
	RegionId                  int     `json:"region_id"`
	GateIp                    string  `json:"gate_ip"`
	GatePort                  string  `json:"gate_port"`
	Status                    int8    `json:"status"`
	ServerProtectPeriod       int64   `json:"server_protect_period"`
	UserRelocateProtectPeriod int64   `json:"user_relocate_protect_period"`
	ServerIds                 []int16 `json:"server_list"`
}

type GuestRoleInfo struct {
	Username string `json:"user_name"`
}

type ForbidCmdReq struct {
	Typ      int8   `json:"typ"`
	RegionId int16  `json:"region_id"`
	Api      int16  `json:"api"`
	Extra    string `json:"extra"`
}

type ForbidCmdItem struct {
	RegionId int16 `json:"region_id"`
	Api      int16 `json:"api"`
}

type ForbidCmdRet struct {
	Code      int16           `json:"code"`
	ForbidCmd []ForbidCmdItem `json:"forbid_list"`
	ForbidUrl []string        `json:"forbid_url"`
}

type UdidWritebackReq struct {
	Typ int8   `json:"typ"`
	Val string `json:"val"`
}

type UdidWritebackRet struct {
	Code int16  `json:"code"`
	Val  string `json:"val"`
}

type AuthReqInfo struct {
	ActionType int    `json:"action_type"`
	Username   string `json:"username"`
	Pwd        string `json:"pwd"`
	DeviceId   string `json:"device_id"`
	OpenUdid   string `json:"open_udid"`
	PlatformId int16  `json:"platform_id"`
	Locale     string `json:"locale"`
	ServerId   int16  `json:"server_id"`
	UserId     int64  `json:"user_id"`
	Mode       int8   `json:"mode"`
	AppId      string `json:"appid"`
}

type AuthRespInfo struct {
	Code       int16                `json:"code"`
	Token      string               `json:"token"`
	ExtraData  string               `json:"extra_data"`
	GateIp     string               `json:"gate_ip"`
	GatePort   string               `json:"gate_port"`
	UserId     int64                `json:"user_id"`
	ServerId   int16                `json:"server_id"`
	UdidWBInfo UdidWriteback        `json:"udid_writeback"`
	IsoCode    string               `json:"iso_code"`
	NetProxy   []netModels.NetProxy `json:"net_proxy_list"`
}

type ForbidCmdAck struct {
	Code    int16
	Typ     int8
	CmdList []int16
}
