package netmsg

import (
	"gm_server/src/ctimer"
	"gm_server/src/logger"
)

func timerChanHandler(id int64, ok bool) {
	if !ok {
		logger.LogErrf("GameTimer error: %v", ok)
		return
	}
	ts, has := ctimer.TIMER_MAP[id]
	if has {
		//先 Free timer 确保不会重复触发
		ctimer.CompleteTimer(id)
		timerHandler(ts)
	} else {
		ctimer.FreeTimer(id)
		logger.LogWarnf("Not found timer: %d.", id)
	}
}

//来自定时器的消息处理
func timerHandler(ts *ctimer.CTimer) {
	//TODO

	switch ts.CTData.Action {
	case ctimer.ACTION_SHUTDOWN:
		shutDown()
		STOP_SIG <- 1
		return

	case ctimer.ACTION_HEARTBEAT:
		doHeartBeatRequest(ts.CTData.Data)

	}

}

//func tHandler(user *gsdb.DbUser, msg ctimer.CTData) {

//	switch msg.Action {

//	case ctimer.ACTION_TROOP_TRAIN_COMPLETE:
//		troopTrainTCB(user, msg.Data.(int16))

//	case ctimer.ACTION_BUILDING_COMPLETE:
//		buildingCompletedTCB(user, msg.Data)

//	case ctimer.ACTION_POWER_UPDATE:
//		powerUpdateCB(user, msg.Data)

//	case ctimer.ACTION_RESEARCH_COMPLETE:
//		researchCompletedCB(user, msg.Data)

//	case ctimer.ACTION_OFFICER_EXPIRE:
//		recruitHallProgTCB(user, msg.Data.(int32))

//	case ctimer.ACTION_POWER_RANK_UPDATE:
//		PowerRankUpdateCB(user, msg.Data)

//	case ctimer.ACTION_REPAIR_COMPLETE:
//		RepairCompletedTCB(user, msg.Data)

//	case ctimer.ACTION_GARRISON_TRAIN_COMPLETE:
//		GarrisonCompletedTCB(user, msg.Data)

//	case ctimer.ACTION_VIP_COMPLETE:
//		vipFreeEnd(user)

//	case ctimer.ACTION_DEL_ACTIVITY_REWARD_HISTORY:
//		id := msg.Data.(int64)
//		deleteRewardHistory(id)
//	//------------------------------------ map -------------------------------------
//	case ctimer.ACTION_MARCH_COMPLETE:
//		//占领行军完毕
//		marchCompleteTCB(user, msg.Data.(int32))

//	case ctimer.ACTION_GATHER_COMPLETE:
//		//资源采集完毕
//		gatherCompleteTCB(user, msg.Data.(const_map.MapPst))

//	case ctimer.ACTION_RALLY_COMPLETE:
//		//集结时间到达
//		rallyCompleteTCB(user)

//	case ctimer.ACTION_USER_MAP_INFO_UPDATE:
//		//地图玩家更新
//		userMapInfoTCB(user)

//	case ctimer.ACTION_USER_ALERT_UPDATE:
//		//玩家警告回调
//		userAlertTCB(user)

//	case ctimer.ACTION_FORT_TO_LAND:
//		//占领领地完成
//		fortToLandCompletedTCB(user, msg.Data.(const_map.MapPst))

//	case ctimer.ACTION_LAND_CHANGE_ALLIANCE:
//		// 领土改变了主人
//		landChangeAllianceCompletedTCB(user, msg.Data.(const_map.MapPst))

//	case ctimer.ACTION_MARCH_FIGHT_COMPLETE:
//		//行军战斗
//		MarchFightTCB(user, msg.Data.(int32))

//	case ctimer.ACTION_USE_REBEL_POINT:
//		useRebelPointTCB(user, msg.Data.(int32))

//	default:
//		logger.LogErrf("Unkown timer msg: %v.", msg)
//	}
//}

//---------------------------------- Creat Timer -------------------------------------
