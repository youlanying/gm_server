package logger_custom

import (
	"gm_server/src/logger"
	"log"
	"strconv"
	"time"
)

var (
	tbLogRoleLoginList    []*TB_LOG_ROLE_LOGIN
	tbLogCreateRoleList   []*TB_LOG_CREATE_ROLE
	tbLogInstanceJoinList []*TB_LOG_INSTANCE_JOIN
	tbLogInstanceOutList  []*TB_LOG_INSTANCE_OUT
	tbLogTerminateList    []*TB_LOG_TERMINATE

	tbLogRoleLoginChan    chan *TB_LOG_ROLE_LOGIN
	tbLogCreateRoleChan   chan *TB_LOG_CREATE_ROLE
	tbLogInstanceJoinChan chan *TB_LOG_INSTANCE_JOIN
	tbLogInstanceOutChan  chan *TB_LOG_INSTANCE_OUT
	tbLogTerminateChan    chan *TB_LOG_TERMINATE

	//numTbLogRoleLogin    int
	//numTbLogCreateRole   int
	//numTbLogInstanceJoin int
	//numTbLogInstanceOut  int
	//numTbLogTerminate    int
)

func init() {
	//fmt.Println("=====gmLog====init==ok===")
	//Log("=====gmLog====init==ok===")
	tbLogRoleLoginChan = make(chan *TB_LOG_ROLE_LOGIN, 10000)
	tbLogCreateRoleChan = make(chan *TB_LOG_CREATE_ROLE, 10000)
	tbLogInstanceJoinChan = make(chan *TB_LOG_INSTANCE_JOIN, 10000)
	tbLogInstanceOutChan = make(chan *TB_LOG_INSTANCE_OUT, 10000)
	tbLogTerminateChan = make(chan *TB_LOG_TERMINATE, 10000)

	tbLogRoleLoginList = make([]*TB_LOG_ROLE_LOGIN, 0)
	tbLogCreateRoleList = make([]*TB_LOG_CREATE_ROLE, 0)
	tbLogInstanceJoinList = make([]*TB_LOG_INSTANCE_JOIN, 0)
	tbLogInstanceOutList = make([]*TB_LOG_INSTANCE_OUT, 0)
	tbLogTerminateList = make([]*TB_LOG_TERMINATE, 0)

	go SaveLogDb()
}

// SaveLogDb 存储LOG
func SaveLogDb() {
	for {
		select {
		case tmpTbLogRoleLogin := <-tbLogRoleLoginChan:
			TB_LOG_ROLE_LOGIN_Insert(tmpTbLogRoleLogin)
		case tmpTbLogCreateRole := <-tbLogCreateRoleChan:
			TB_LOG_CREATE_ROLE_Insert(tmpTbLogCreateRole)
		case tmpTbLogInstanceJoin := <-tbLogInstanceJoinChan:
			TB_LOG_INSTANCE_JOIN_Insert(tmpTbLogInstanceJoin)
		case tmpTbLogInstanceOut := <-tbLogInstanceOutChan:
			TB_LOG_INSTANCE_OUT_Insert(tmpTbLogInstanceOut)
		case tmpTbLogTerminate := <-tbLogTerminateChan:
			TB_LOG_TERMINATE_Insert(tmpTbLogTerminate)
			//default:

		}
	}
}

// logWrite gmLog用Logger, 如需在其他目录或有不同规则,可考虑移植logger中的init相关到此
func logWrite(format string, v ...interface{}) {
	str := logger.FormatLogf(logger.LOG_INFO, logger.GetDefaultCallDepth(), format, v...)
	if logger := logger.GetLogger("gmLog"); logger != nil {
		logger.Println(str)
	}
	if (logger.Getlogfilemaxsize() & logger.LerrorExit) != 0 {
		log.Fatalln(str)
	} else {
		log.Println(str)
	}
}

func GmLog_Login(RoleId uint64, Account, RoleName string, ServerId int32, Ip, Mac, Model, VersionNo, CId, LoginType, PhoneNumber string) {
	nowS := CurrentMS()
	logWrite("cmd: login_log, roleid: %v, account: %v, rolename: %v, serverid: %v, ip: %v, mac: %v, devicetype: %v, versionno: %v, channelid: %v, time: %v",
		RoleId, Account, RoleName, ServerId, Ip, Mac, Model, VersionNo, CId, nowS)
	// todo DBOperation?
	tbLogRoleLogin := &TB_LOG_ROLE_LOGIN{
		LOGID:      strconv.FormatInt(nowS, 10) + "_" + strconv.FormatUint(RoleId, 10),
		ROLEID:     RoleId,
		ACCOUNT:    Account,
		NAME:       RoleName,
		SERVERID:   ServerId,
		IP:         Ip,
		MAC:        Mac,
		DEVICETYPE: Model,
		VERSIONNO:  VersionNo,
		CID:        CId,
		CREATETIME: nowS,
	}
	tbLogRoleLoginChan <- tbLogRoleLogin
	//tbLogRoleLoginList = append(tbLogRoleLoginList, tbLogRoleLogin)
	//TB_LOG_ROLE_LOGIN_Insert(tbLogRoleLogin)
}

func GmLog_Logout(RoleId uint64, Account, RoleName string, ServerId int32, MapNode, Reason, LoginTime, LogoutTime, CId, LoginType, NickName, HeadImgUrl, PhoneNumber, Ip, DeviceType, VersionNo string) {
	nowS := CurrentS()
	logWrite("cmd: logout_log, roleid: %v, account: %v, rolename: %v, serverid: %v, mapnode: %v, ip: %v, devicetype: %v, versionno: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, ServerId, MapNode, Ip, DeviceType, VersionNo, Reason, nowS)
	// todo DBOperation?
}

func GmLog_CreateRoleId(RoleId uint64, Account, IpAddress string, ServerId int32, Mac, CId string) {
	nowS := CurrentS()
	logWrite("cmd: create_roleid_log, roleid: %v, account: %v, ipaddress: %v, mac: %v, serverid: %v, channelid: %v, time: %v",
		RoleId, Account, IpAddress, Mac, ServerId, CId, nowS)
	// todo DBOperation?
	tbLogCreateRole := &TB_LOG_CREATE_ROLE{
		LOGID:      strconv.FormatInt(nowS, 10) + "_" + strconv.FormatUint(RoleId, 10),
		ROLEID:     RoleId,
		ACCOUNT:    Account,
		SERVERID:   ServerId,
		IP:         IpAddress,
		MAC:        Mac,
		DEVICETYPE: "",
		VERSIONNO:  "",
		CID:        CId,
		CREATETIME: nowS,
	}
	tbLogCreateRoleChan <- tbLogCreateRole
	//tbLogCreateRoleList = append(tbLogCreateRoleList, tbLogCreateRole)
	//TB_LOG_CREATE_ROLE_Insert(tbLogCreateRole)
}

func GmLog_CreateRole(RoleId uint64, Account, RoleName, IpAddress string, ServerId int32, Mac, CId, DeviceType, VersionNo string) {
	nowS := CurrentS()
	logWrite("cmd: create_role_log, roleid: %v, account: %v, rolename: %v, ipaddress: %v, mac: %v, serverid: %v, channelid: %v, time: %v",
		RoleId, Account, RoleName, IpAddress, Mac, ServerId, CId, nowS)
	// todo DBOperation?
}

func GmLog_ChangeRoleName(RoleId uint64, Account, RoleName string) {
	nowS := CurrentS()
	logWrite("cmd: change_rolename, roleid: %v, account: %v, rolename: %v, time: %v",
		RoleId, Account, RoleName, nowS)
}

func GmLog_AddBlackLog(Account string, Time int64) {
	nowS := CurrentS()
	logWrite("cmd: add_black_log, account: %v, black_time: %v, time: %v",
		Account, Time, nowS)
}

func GmLog_RoleLevelUp(RoleId uint64, Account, RoleName string, Level, Exp int32) {
	nowS := CurrentS()
	logWrite("cmd: role_uplevel_log, roleid: %v, account: %v, rolename: %v, orilevel: %v, curlevel: %v, curexp: %v, time: %v",
		RoleId, Account, RoleName, Level-1, Level, Exp, nowS)
}

func GmLog_HeroUpLevel(RoleId uint64, Account, RoleName string, HeroId uint64, HeroProtoId, Level, Exp int32) {
	nowS := CurrentS()
	logWrite("cmd: hero_uplevel_log, roleid: %v, account: %v, rolename: %v, heroid: %v, heroproto: %v, orilevel: %v, curlevel: %v, curexp: %v, time: %v",
		RoleId, Account, RoleName, HeroId, HeroProtoId, Level-1, Level, Exp, nowS)
}

func GmLog_HeroUpGrade(RoleId uint64, Account, RoleName string, HeroId uint64, HeroProtoId, Grade int32) {
	nowS := CurrentS()
	logWrite("cmd: hero_upgrade_log, roleid: %v, account: %v, rolename: %v, heroid: %v, heroproto: %v, grade: %v, time: %v",
		RoleId, Account, RoleName, HeroId, HeroProtoId, Grade, nowS)
}

func GmLog_AddMoney(RoleId uint64, Account, RoleName string, OriMoney, Money, CurMoney int32, Reason string) {
	nowS := CurrentS()
	logWrite("cmd: add_money_log, roleid: %v, account: %v, rolename: %v, orimoney: %v, money: %v, curmoney: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, OriMoney, Money, CurMoney, Reason, nowS)
}

func GmLog_LoseMoney(RoleId uint64, Account, RoleName string, OriMoney, Money, CurMoney int32, Reason string) {
	nowS := CurrentS()
	logWrite("cmd: lose_money_log, roleid: %v, account: %v, rolename: %v, orimoney: %v, money: %v, curmoney: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, OriMoney, Money, CurMoney, Reason, nowS)
}

func GmLog_DeleteItem(RoleId uint64, Account, RoleName string, SLot int32, ItemId uint64, ProtoId, ItemType, Count int32, Reason string) {
	nowS := CurrentS()
	logWrite("cmd: delete_item_log, roleid: %v, account: %v, rolename: %v, slot: %v, itemid: %v, protoid: %v, itemtype: %v, count: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, SLot, ItemId, ProtoId, ItemType, Count, Reason, nowS)
}

func GmLog_ConsumeItem(RoleId uint64, Account, RoleName string, SLot int32, ItemId uint64, ProtoId, ItemType, ConsumeCount, LeftCount, CurLevel int32, Reason string) {
	nowS := CurrentS()
	logWrite("cmd: consume_item_log, roleid: %v, account: %v, rolename: %v, slot: %v, itemid: %v, protoid: %v, itemtype: %v, consumecount: %v, leftcount: %v, curlevel: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, SLot, ItemId, ProtoId, ItemType, ConsumeCount, LeftCount, CurLevel, Reason, nowS)
}

func GmLog_AddItem(RoleId uint64, Account, RoleName string, SLot int32, ItemId uint64, ProtoId, ItemType, ConsumeCount, LeftCount, CurLevel int32, Reason string) {
	nowS := CurrentS()
	logWrite("cmd: add_item_log, roleid: %v, account: %v, rolename: %v, slot: %v, itemid: %v, protoid: %v, itemtype: %v, consumecount: %v, leftcount: %v, curlevel: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, SLot, ItemId, ProtoId, ItemType, ConsumeCount, LeftCount, CurLevel, Reason, nowS)
}

func GmLog_GoldObtain(RoleId uint64, Account, RoleName string, OriGold, Gold, CurGold int32, Reason, Charge string) {
	nowS := CurrentS()
	logWrite("cmd: gold_obtain, roleid: %v, account: %v, rolename: %v, origold: %v, count: %v, curgold: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, OriGold, Gold, CurGold, Reason, nowS)
}

func GmLog_GoldLose(RoleId uint64, Account, RoleName string, OriGold, Gold, CurGold int32, Reason string) {
	nowS := CurrentS()
	logWrite("cmd: gold_lose, roleid: %v, account: %v, rolename: %v, origold: %v, count: %v, curgold: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, OriGold, Gold, CurGold, Reason, nowS)
}

func GmLog_ObtainPhysical(RoleId uint64, Account, RoleName string, OriPhysical, AddPhysical, CurPhysical int32, Reason string) {
	nowS := CurrentS()
	logWrite("cmd: obtain_physical_log, roleid: %v, account: %v, rolename: %v, oriphysical: %v, addphysical: %v, curphysical: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, OriPhysical, AddPhysical, CurPhysical, Reason, nowS)
}

func GmLog_LosePhysical(RoleId uint64, Account, RoleName string, OriPhysical, RemPhysical, CurPhysical int32, Reason string) {
	nowS := CurrentS()
	logWrite("cmd: lose_physical_log, roleid: %v, account: %v, rolename: %v, oriphysical: %v, remphysical: %v, curphysical: %v, reason: %v, time: %v",
		RoleId, Account, RoleName, OriPhysical, RemPhysical, CurPhysical, Reason, nowS)
}

func GmLog_QuestFinish(RoleId uint64, Account, RoleName string, QuestId, QuestType int32) {
	nowS := CurrentS()
	logWrite("cmd: quest_finish_log, roleid: %v, account: %v, rolename: %v, questid: %v, questtype: %v, time: %v",
		RoleId, Account, RoleName, QuestId, QuestType, nowS)
}

func GmLog_QuestConditionFinish(RoleId uint64, Account, RoleName string, QuestId, QuestType int32) {
	nowS := CurrentS()
	logWrite("cmd: quest_condition_finish_log, roleid: %v, account: %v, rolename: %v, questid: %v, questtype: %v, time: %v",
		RoleId, Account, RoleName, QuestId, QuestType, nowS)
}

func GmLog_AddQuestIdList(RoleId uint64, Account, RoleName string, OriginalQIdList, RemRepeatQuest, ActualQIdList []int32) {
	nowS := CurrentS()
	logWrite("cmd: quest_condition_finish_log, roleid: %v, account: %v, rolename: %v, addquestid: %v, removalrepeatqid: %v, actualquestid: %v, time: %v",
		RoleId, Account, RoleName, OriginalQIdList, RemRepeatQuest, ActualQIdList, nowS)
}

func GmLog_InstanceJoin(RoleId uint64, Account string, ServerId, InstanceId, CostDouble, DayTimes, activityId int32, Star []int32) {
	nowS := CurrentS()
	logWrite("cmd: instance_Join_log, roleid: %v, account: %v, instanceid: %v, costdouble: %v, DayTimes: %v, time: %v",
		RoleId, Account, InstanceId, CostDouble, DayTimes, nowS)
	tbLogInstanceJoin := &TB_LOG_INSTANCE_JOIN{
		LOGID:      strconv.FormatInt(nowS, 10) + "_" + strconv.FormatUint(RoleId, 10),
		ROLEID:     RoleId,
		ACCOUNT:    Account,
		SERVERID:   ServerId,
		INSTANCEID: InstanceId,
		ACTIVITYID: activityId,
		DAYTIMES:   DayTimes,
		STAR:       Star,
		CREATETIME: nowS,
	}
	tbLogInstanceJoinChan <- tbLogInstanceJoin
	//tbLogInstanceJoinList = append(tbLogInstanceJoinList, tbLogInstanceJoin)
	//TB_LOG_INSTANCE_JOIN_Insert(tbLogInstanceJoin)
}

func GmLogInstanceOut(ServerId int32, RoleId uint64, Account string, roleLevel, InstanceId, result, DayTimes, endGold, RoleExp, HeroExp int32, objectItemList [][]int32, deadHero []int32, battleHeroDataList [][]int32, timeConsume int32) {
	nowS := CurrentS()
	//logWrite("cmd: instance_out_log, roleid: %v, account: %v, instanceid: %v, result: %v, usedtimes: %v, gold: %v, roleexp: %v, heroexp: %v, score: %v, add_items: %v, star: %v, time: %v",
	//	RoleId, Account, InstanceId, result, timeConsume, endGold, RoleExp, HeroExp, Score, objectItemList, Star, nowS)
	tbLogInstanceOut := &TB_LOG_INSTANCE_OUT{
		LOGID:          strconv.FormatInt(nowS, 10) + "_" + strconv.FormatUint(RoleId, 10),
		INSTANCEID:     InstanceId,
		SERVERID:       ServerId,
		ROLEID:         RoleId,
		ACCOUNT:        Account,
		ROLELEVEL:      roleLevel,
		STATE:          result,
		DAYTIMES:       DayTimes,
		ENDGOLD:        endGold,
		ROLEEXP:        RoleExp,
		HEROEXP:        HeroExp,
		OBJECTITEMS:    objectItemList,
		STAR:           []int32{},
		DEADHERO:       deadHero,
		BATTLEHEROLIST: battleHeroDataList,
		TIMECONSUM:     timeConsume,
		CREATETIME:     nowS,
	}
	tbLogInstanceOutChan <- tbLogInstanceOut
	//tbLogInstanceOutList = append(tbLogInstanceOutList, tbLogInstanceOut)
	//TB_LOG_INSTANCE_OUT_Insert(tbLogInstanceOut)
}

func GmLog_InstanceCheckError(RoleId uint64, Account, RoleName string, InstanceId, Money, MaxMoney, RoleExp, MaxRoleExp, HeroExp, MaxHeroExp, Score, MaxScore, CMaxHp, MaxHp, CMaxAttack, MaxAttack, CMaxDefence, MaxDefence int32) {
	nowS := CurrentS()
	logWrite("cmd: instance_check_error, roleid: %v, account: %v, rolename: %v, instanceid: %v, money: %v, maxmoney: %v, roleexp: %v, maxroleexp: %v, heroexp: %v, maxheroexp: %v, score: %v, maxscore: %v, cmaxhp: %v, maxhp: %v, cmaxattack: %v, maxattack: %v, cmaxdefence: %v, maxdefence: %v, time: %v",
		RoleId, Account, RoleName, InstanceId, Money, MaxMoney, RoleExp, MaxRoleExp, HeroExp, MaxHeroExp, Score, MaxScore, CMaxHp, MaxHp, CMaxAttack, MaxAttack, CMaxDefence, MaxDefence, nowS)
}

func GmLog_PrayCalcResultLog(RoleId uint64, Account, RoleName string, PrayType, IsSeries, Money int32, ItemList []int32) {
	nowS := CurrentS()
	logWrite("cmd: pray_calcresult_log, roleid: %v, account: %v, rolename: %v, praytype: %v, isseries: %v, money: %v, itemlist: %v, time: %v",
		RoleId, Account, RoleName, PrayType, IsSeries, Money, ItemList, nowS)
}

// todo vip据说取消了,所以没写

func GmLog_ChargeFaild(RoleId uint64, Account, RoleName string, Rmb, ChargeId, ChargeGold, ProtoRmb int32, StrPlat_GoodsId, StrPlat_OrderId, StrPlat_PayTime, StrPlat_CId string) {
	nowS := CurrentS()
	logWrite("cmd: charge_failed_log, roleid: %v, account: %v, rolename: %v, rmb: %v, chargeid: %v, chargegold: %v, protormb: %v, platgoodid: %v, platoid: %v, platcid: %v, paytime: %v, time: %v",
		RoleId, Account, RoleName, Rmb, ChargeId, ChargeGold, ProtoRmb, StrPlat_GoodsId, StrPlat_OrderId, StrPlat_PayTime, StrPlat_CId, nowS)
}

func GmLog_ChargeAbnormal(RoleId uint64, Account, RoleName string, Rmb, ChargeId, ChargeGold, ProtoRmb int32, StrPlat_GoodsId, StrPlat_OrderId, StrPlat_PayTime, StrPlat_CId string, BuyOrTime int32) {
	nowS := CurrentS()
	logWrite("cmd: charge_abnormal_log, roleid: %v, account: %v, rolename: %v, rmb: %v, chargeid: %v, chargegold: %v, protormb: %v, platgoodid: %v, platoid: %v, platcid: %v, paytime: %v, BuyOrTime: %v, time: %v",
		RoleId, Account, RoleName, Rmb, ChargeId, ChargeGold, ProtoRmb, StrPlat_GoodsId, StrPlat_OrderId, StrPlat_PayTime, StrPlat_CId, BuyOrTime, nowS)
}

func GmLog_ChargeOk(RoleId uint64, Account, RoleName string, Rmb, ChargeId, ChargeGold, GiftGold, OriGold int32, StrPlat_GoodsId, StrPlat_OrderId, StrPlat_PayTime, StrPlat_CId string, BuyOrTime int32) {
	nowS := CurrentS()
	logWrite("cmd: charge_ok_log, roleid: %v, account: %v, rolename: %v, rmb: %v, chargeid: %v, chargegold: %v, giftgold: %v, origold: %v, platgoodid: %v, platoid: %v, platcid: %v, paytime: %v, BuyOrTime: %v, time: %v",
		RoleId, Account, RoleName, Rmb, ChargeId, ChargeGold, GiftGold, OriGold, StrPlat_GoodsId, StrPlat_OrderId, StrPlat_PayTime, StrPlat_CId, BuyOrTime, nowS)
}

func GmLog_ObtainWealthLog(RoleId uint64, Account, RoleName string, UsedTimes, RandomTimes, Money int32) {
	nowS := CurrentS()
	logWrite("cmd: obtainwealth_log, roleid: %v, account: %v, rolename: %v, usedtimes: %v, randomtimes: %v, money: %v, time: %v",
		RoleId, Account, RoleName, UsedTimes, RandomTimes, Money, nowS)
}

func GmLog_HeroSkillLevelUpLog(RoleId uint64, Account, RoleName string, HeroId uint64, HeroProtoId, SkillId int32, SkillType []int32, CostMoney, CurSkillLev int32) {
	nowS := CurrentS()
	logWrite("cmd: hero_skill_levelup_log, roleid: %v, account: %v, rolename: %v, heroid: %v, heroproto: %v, skillid: %v, skilltype: %v, costmoney: %v, curskillLev: %v, time: %v",
		RoleId, Account, RoleName, HeroId, HeroProtoId, SkillId, SkillType, CostMoney, CurSkillLev, nowS)
}

func GmLog_EquipLevelUpLog(RoleId uint64, Account, RoleName string, ItemId uint64, ItemProtoId, CurEquipLevel, NewExp, CurEquipStar, CostMoney int32) {
	nowS := CurrentS()
	logWrite("cmd: equip_levelup_log, roleid: %v, account: %v, rolename: %v, itemid: %v, itemprotoid: %v, curequiplevel: %v, newexp: %v, curequipstar: %v, costmoney: %v, time: %v",
		RoleId, Account, RoleName, ItemId, ItemProtoId, CurEquipLevel, NewExp, CurEquipStar, CostMoney, nowS)
}

func GmLog_EquipStarUpLog(RoleId uint64, Account, RoleName string, ItemId uint64, ItemProtoId, CurEquipStar, CostMoney int32) {
	nowS := CurrentS()
	logWrite("cmd: equip_starup_log, roleid: %v, account: %v, rolename: %v, itemid: %v, itemprotoid: %v, curequipstar: %v, costmoney: %v, time: %v",
		RoleId, Account, RoleName, ItemId, ItemProtoId, CurEquipStar, CostMoney, nowS)
}

func GmLog_SendMailLog(RecvName, Title string, ItemList []int32, Sendor, IsOnline, SendOk string) {
	nowS := CurrentS()
	logWrite("cmd: gm_send_mail_log, recvname: %v, title: %v, itemlist: %v, sendor: %v, isonline: %v, sendok: %v, time: %v",
		RecvName, Title, ItemList, Sendor, IsOnline, SendOk, nowS)
}

func GmLog_GmSendCmdLog(SendStr string) {
	nowS := CurrentS()
	logWrite("cmd: gm_send_cmd_log, sendstr: %v, time: %v", SendStr, nowS)
}

func GmLog_GmRankLog(TypeStr, RankStr string) {
	nowS := CurrentS()
	logWrite("cmd: gm_rank_log, type: %v, rankstr: %v, time: %v", TypeStr, RankStr, nowS)
}

func GmLog_ChangeNameLog(RoleId uint64, Account, OriName, NewName string, CostGold int32) {
	nowS := CurrentS()
	logWrite("cmd: change_name_log, type: %v, roleid: %v, account: %v, oriname: %v, newname: %v, costgold: %v, time: %v", RoleId, Account, OriName, NewName, CostGold, nowS)
}

func GmLog_ChangeIconLog(RoleId uint64, Account string, OldIcon, NewIcon int32) {
	nowS := CurrentS()
	logWrite("cmd: change_icon_log, roleid: %v, account: %v, oldicon: %v, newicon: %v, time: %v", RoleId, Account, OldIcon, NewIcon, nowS)
}

func GmLog_ChangeSexLog(RoleId uint64, Account string, OldSex, NewSex int32) {
	nowS := CurrentS()
	logWrite("cmd: change_sex_log, roleid: %v, account: %v, oldsex: %v, newsex: %v, time: %v", RoleId, Account, OldSex, NewSex, nowS)
}

func GmLog_ChangeAutographLog(RoleId uint64, Account, OldAutograph, NewAutograph string) {
	nowS := CurrentS()
	logWrite("cmd: change_sex_log, roleid: %v, account: %v, oldautograph: %v, newautograph: %v, time: %v", RoleId, Account, OldAutograph, NewAutograph, nowS)
}

// 这个传参顺序有修改,使用时请注意!
func GmLog_UseExchangeNumber(RoleId uint64, RoleName, Number string) {
	nowS := CurrentS()
	logWrite("cmd: use_exchange_number, roleid: %v, rolename: %v, number: %v, time: %v",
		RoleId, RoleName, Number, nowS)
	// todo DBOperation?
}

// todo redis据说不用了,所以没写

func GmLog_BuyMallItem(RoleId uint64, Account, RoleName string, MallId, ItemProtoId, ItemNum, ItemPrice, MallDiscount, BuyPrice, MoneyType int32) {
	nowS := CurrentS()
	logWrite("cmd: buy_mall_item, roleid: %v, account: %v, rolename: %v, mallid: %v, itemprotoid: %v, itemnum: %v, mallprice: %v, malldiscount: %v, buyprice: %v, type: %v, time: %v",
		RoleId, Account, RoleName, MallId, ItemProtoId, ItemNum, ItemPrice, MallDiscount, BuyPrice, MoneyType, nowS)
}

func GmLog_OnWeapon(RoleId uint64, Account, RoleName string, Slot, ItemId uint64, ProtoId, ItemType, CurLevel int32) {
	nowS := CurrentS()
	logWrite("cmd: on_weanpon_log, roleid: %v, account: %v, rolename: %v, slot: %v, itemid: %v, protoid: %v, itemtype: %v, curlevel: %v, time: %v",
		RoleId, Account, RoleName, Slot, ItemId, ProtoId, ItemType, CurLevel, nowS)
}

func GmLog_OpMail(RoleId uint64, Account, RoleName string, Flag int32, MailData string, CurLevel int32) {
	nowS := CurrentS()
	logWrite("cmd: op_mail_log, roleid: %v, account: %v, rolename: %v, flag: %v, maildata: %v, curlevel: %v, time: %v",
		RoleId, Account, RoleName, Flag, MailData, CurLevel, nowS)
}

// todo 装备升星貌似没了,所以没写

func GmLog_BuyFightShortCutLog(RoleId uint64, Account, RoleName string, Levelid, Buyid, FlagNum, Use_money int32) {
	nowS := CurrentS()
	logWrite("cmd: buy_fight_shortcut_log, roleid: %v, account: %v, rolename: %v, levelid: %v, buyid: %v, flagnum: %v, usemoney: %v, time: %v",
		RoleId, Account, RoleName, Levelid, Buyid, FlagNum, Use_money, nowS)
}

func GmLog_QuickChallengeLog(RoleId uint64, Account, RoleName string, LevelId, Useitemid, Usenum int32, NewDropItemlist string, NewMoney, NewHeroEndExp, RoleEndExp int32) {
	nowS := CurrentS()
	logWrite("cmd: quick_challenge_log, roleid: %v, account: %v, rolename: %v, levelid: %v, Useitemid: %v, Usenum: %v, dropitem: %v, obtainMoney: %v, obtainheroexp: %v, obtainexp: %v, time: %v",
		RoleId, Account, RoleName, LevelId, Useitemid, Usenum, NewDropItemlist, NewMoney, NewHeroEndExp, RoleEndExp, nowS)
}

// todo jpush不确定还用不用,装备类升星貌似被取缔了,所以没写

func GmLog_UseGoldAwardLog(RoleId uint64, Account, RoleName string, Awardid, Use_gold int32, GiftItemList string, ActId int32) {
	nowS := CurrentS()
	logWrite("cmd: use_gold_award_log, roleid: %v, account: %v, rolename: %v, awardid: %v, usegold: %v, giftitemlist: %v, actid: %v, time: %v",
		RoleId, Account, RoleName, Awardid, Use_gold, GiftItemList, ActId, nowS)
}

func GmLog_RandomItemListLog(RoleId uint64, Account, RoleName string, Protoid, UsedItemNum int32, AllItemList string, Type, Price int32) {
	nowS := CurrentS()
	logWrite("cmd: random_itemlist_log, roleid: %v, account: %v, rolename: %v, itemclsid: %v, usenum: %v, itemlist: %v, moneytype: %v, price: %v, time: %v",
		RoleId, Account, RoleName, Protoid, UsedItemNum, AllItemList, Type, Price, nowS)
}

func GmLog_AddTerminate(reason string, serverId int32, roleId uint64, account, roleName, errorProto string) {
	nowMS := CurrentMS()
	tbLogTerminate := &TB_LOG_TERMINATE{
		LOGID:      strconv.FormatInt(int64(serverId), 10) + "_" + strconv.FormatInt(nowMS, 10),
		ROLEID:     roleId,
		ACCOUNT:    account,
		NAME:       roleName,
		SERVERID:   serverId,
		ERRORPROTO: errorProto,
		REASON:     reason,
		CREATETIME: CurrentS(),
	}
	tbLogTerminateChan <- tbLogTerminate
	//tbLogTerminateList = append(tbLogTerminateList, tbLogTerminate)
	//TB_LOG_TERMINATE_Insert(tbLogTerminate)
}

// 返回unix时间戳秒
func CurrentS() int64 {
	return int64(time.Now().UnixNano() / 1000000 / 1000)
}

// 返回unix时间戳毫秒
func CurrentMS() int64 {
	return int64(time.Now().UnixNano() / 1000000)
}
