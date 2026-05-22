package ctimer

const (
	ACTION_MAIL_SAVE = 30 // 邮件定时保存

	//	// 全局消息
	ACTION_SHUTDOWN = 10000

	ACTION_USER_AUTO_SAVE = 10003 // 玩家定时存盘
	ACTION_HEARTBEAT      = 10004 // 心跳包定时器

	ACTION_WORLD_TICK_UPDATE = 10013 // 游戏帧

	ACTION_GATE_ONLINE_STAT = 10015 //

	ACTION_DEL_ACTIVITY_REWARD_HISTORY = 10035 // 删除活动奖励历史记录

	ACTION_GS_SNAPSHOT = 10039 // GS存盘Hash定时

	//================================================================================
	// 以下为新项目所用，鉴于定时器最大长度为65535个，但ACTION为int16目前考虑id从30001开始
	ACTION_HOOK_ON_UPDATE_10S    = 30001 // 10s 定时
	ACTION_HOOK_ON_UPDATE_1M     = 30002 // 1min 定时
	ACTION_HOOK_ON_UPDATE_5M     = 30003 // 5min 定时
	ACTION_HOOK_ON_UPDATE_10M    = 30004 // 10min 定时
	ACTION_HOOK_ON_UPDATE_DAY    = 30005 // 每日（0点） 定时
	ACTION_HOOK_ON_UPDATE_CLOCK4 = 30006 // 每日整点（4点）刷新 定时

)

type CTData struct {
	Action int16
	Data   interface{}
}

type CTimer struct {
	Receiver int64
	Expired  int64 // expired time, timestamp(ms)
	Index    int64 // slot index of timerwheel
	CTData   CTData
}
