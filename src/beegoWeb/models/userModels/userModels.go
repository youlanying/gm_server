package userModels

type UserInfo struct {
	Id                 int64  `bson:"id" json:"id"`
	Udid               string `bson:"udid" json:"udid"`
	Account            string `bson:"account" json:"accounts"`
	Name               string `json:"account"`
	Register_server_id int16  `bson:"register_server_id",json:"register_server_id"`
	Server_id          int16  `bson:"server_id" ,json:"server_id"`
}

type UserListInfo struct {
	Info []UserInfo
}

// 运营平台

// 用户查询列表数据
type SnapUser struct {
	Id       int64  `bson:"id" json:"id"`
	Name     string `json:"name"`
	ServerId string `bson:"server_id"  json:"server"`
}

// 用户列表
type SnapUserList struct {
	Info []SnapUser `json:"info"`
}

// 玩家详细数据
type UserDetail struct {
	Id              int64  `json:"id"`
	ServerId        string `json:"server"`
	PosX            int16  `json:"pos_x"`
	PosY            int16  `json:"pos_y"`
	AsName          string `json:"as_name"`
	AsId            int32  `json:"as_id"`
	Name            string `json:"name"`
	Jid             string `json:"chat_account"`
	HQlevel         int16  `json:"hq_level"`
	HeroLevel       int16  `json:"hero_level"`
	Money           int64  `json:"money"`
	VipPoints       int32  `json:"vip_point"`
	VipLevel        int16  `json:"vip_level"`
	MonthCardExpire int64  `json:"monthcard_expire"`
	Consumption     int64  `json:"consumption"`
	Power           int64  `json:"power"`
	RegisteTime     string `json:"registe_time"`
	RegisteUdid     string `json:"registe_udid"`
	LastLoginTime   string `json:"last_login_time"`
	AccountType     string `json:"account_type"`
	Lang            int8   `json:"lang"`

	Shield        int64  `json:"shield"`
	Ip            string `json:"ip"`
	EndpointType  int8   `json:"endpoint_type"`
	EndpointToken string `json:"endpoint_token"`
	PushSetting   int32  `json:"push_setting"`
	Domino        int64  `json:"domino"`
	Udid          string `json:"udid"`
}

// 玩家封装信息
type UserDetailInfo struct {
	Info UserDetail `json:"info"`
}

// 账户信息
type AccountInfo struct {
	Id      int64  `json:"id" bson:"id"`
	Account string `json:"account" bson:"account"`
}

// 道具信息
type ItemInfo struct {
	Id       int16  `json:"code"`
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
	SubType  int8   `json:"subtype"`
	Price    int64  `json:"price"`
	ItemId   string `json:"itemid"`
}
type UserItemsInfo struct {
	Info []ItemInfo `json:"info"`
}

// 修改资源道具返回
type ResRsp struct {
	Message  int16 `json:"code"`     //消息结果 0成功 不等于0是失败
	Quantity int64 `json:"quantity"` // 当前数量
}

// 修改部队返回
type AarmRsp struct {
	Message      int16 `json:"code"`     //消息结果
	Quantity     int64 `json:"quantity"` // 当前数量
	CityQuantity int64 `json:"city_quantity"`
}

// 邮件返回接口
type MailRsp struct {
	Message int16 `json:"code"`
}

type StrRsp struct {
	Message string `json:"msg"`
}

// 内部充值接口返回
type IapRspInfo struct {
	Message   int16 `json:"code"`
	Complete  int16 `json:"finished_number"`
	Money     int64 `json:"money"`
	VipPoints int64 `json:"vip_point"`
}

// 资源信息
type ResInfo struct {
	Id       int16  `json:"code"`
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
}

// 用户玩家资源的信息
type UserRes struct {
	Info []ResInfo `json:"info"`
}

// 部队信息
type ArmyInfo struct {
	Type         int16  `json:"code"`
	Name         string `json:"name"`
	Quantity     int64  `json:"quantity"`
	CityQuantity int64  `json:"city_quantity"`
}

// 玩家部队信息
type UserTroop struct {
	Armies []ArmyInfo `json:"info"` //玩家部队
}

type Svr struct {
	SvrId int16 `json:"svr"`
}

type SvrInfo struct {
	SvrLst []Svr `json:"all_svr"`
}

type RewardItem struct {
	ItemId  int16 `json:"item_id"`
	ItemCnt int32 `json:"item_cnt"`
}

type AllRewardItem struct {
	Rewards []RewardItem `json:"info"`
}

type UserFieldInfo struct {
	Field []FieldInfo `json:"field_info"`
	March []MarchInfo `json:"march_info"`
}

type MarchInfo struct {
	X      int16      `json:"pos_x"`
	Y      int16      `json:"pos_y"`
	Retime int64      `json:"remaining_time"`
	Armies []ArmyInfo `json:"armies"`
}
type FieldInfo struct {
	X      int16      `json:"pos_x"`
	Y      int16      `json:"pos_y"`
	Gtype  int16      `json:"group_type"`
	Armies []ArmyInfo `json:"armies"`
}

type BatchIapPkgInfo struct {
	Users []int64 `json:"users"`
	PkgId int32   `json:"pkg_id"`
}
