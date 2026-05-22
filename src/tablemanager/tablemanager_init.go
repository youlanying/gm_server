package tablemanager

func init() {

	Dic = make(map[string]TableMgrInterface, 74+OtherTableLen)

	//BossRush表.xlsx
	Register("bossrush", &Bossrush_Model)
	//BossRush表.xlsx
	Register("bossrush_rounds", &Bossrush_rounds_Model)
	//rank等级表.xlsx
	Register("ranklevel", &Ranklevel_Model)
	//rogue随机迷宫表格.xlsx
	Register("layer", &Layer_Model)
	//rogue随机迷宫表格.xlsx
	Register("moze_scenes", &Moze_scenes_Model)
	//rogue随机迷宫表格.xlsx
	Register("roomtype", &Roomtype_Model)
	//主人公—团体属性.xlsx
	Register("role_team", &Role_team_Model)
	//主人公—英雄角色.xlsx
	Register("consume_material", &Consume_material_Model)
	//主人公—英雄角色.xlsx
	Register("consume_material_ratio", &Consume_material_ratio_Model)
	//主人公—英雄角色.xlsx
	Register("hero_break", &Hero_break_Model)
	//主人公—英雄角色.xlsx
	Register("hero_information", &Hero_information_Model)
	//主人公—英雄角色.xlsx
	Register("hero_star", &Hero_star_Model)
	//主人公—英雄角色.xlsx
	Register("hero_team", &Hero_team_Model)
	//任务表.xlsx
	Register("quest", &Quest_Model)
	//任务表.xlsx
	Register("quest_active", &Quest_active_Model)
	//兑换码礼包表.xlsx
	Register("exchangenumber", &Exchangenumber_Model)
	//关卡-子场景数据表.xlsx
	Register("level_son_auto", &Level_son_auto_Model)
	//关卡-明雷数据表.xlsx
	Register("BattleTriggerGroup", &BattleTriggerGroup_Model)
	//关卡_场景表.xlsx
	Register("instance_level", &Instance_level_Model)
	//关卡_场景表.xlsx
	Register("instance_level_result", &Instance_level_result_Model)
	//关卡_子场景表.xlsx
	Register("instance_level_son", &Instance_level_son_Model)
	//关卡条件表.xlsx
	Register("condition_level", &Condition_level_Model)
	//关卡表-交互物件数据表.xlsx
	Register("sceneitemdata", &Sceneitemdata_Model)
	//关卡表-关联场景盒子数据表.xlsx
	Register("AssociatedSceneBoxData", &AssociatedSceneBoxData_Model)
	//关卡表-场景npc实体表.xlsx
	Register("npc_object", &Npc_object_Model)
	//关卡表-场景npc数据表.xlsx
	Register("npcdata", &Npcdata_Model)
	//关卡表-场景npc模型表.xlsx
	Register("npc", &Npc_Model)
	//关卡表-通用实体表.xlsx
	Register("CommonEntity", &CommonEntity_Model)
	//剧情总表.xlsx
	Register("plotrelation", &Plotrelation_Model)
	//剧情总表.xlsx
	Register("plotreward", &Plotreward_Model)
	//功能开启表.xlsx
	Register("function", &Function_Model)
	//召唤生物表.xlsx
	Register("summon_unit", &Summon_unit_Model)
	//合成表.xlsx
	Register("EquipmentCompound", &EquipmentCompound_Model)
	//商城表.xlsx
	Register("mall_paging", &Mall_paging_Model)
	//商城表.xlsx
	Register("mallboss", &Mallboss_Model)
	//商城表.xlsx
	Register("malldiamond", &Malldiamond_Model)
	//商城表.xlsx
	Register("mallfashion", &Mallfashion_Model)
	//商城表.xlsx
	Register("mallgold", &Mallgold_Model)
	//商城表.xlsx
	Register("mallmaze", &Mallmaze_Model)
	//商城表.xlsx
	Register("mallsuipian", &Mallsuipian_Model)
	//场景传送盒子映射表.xlsx
	Register("instance_transfer_map", &Instance_transfer_map_Model)
	//场景传送盒子表.xlsx
	Register("instance_transferbox", &Instance_transferbox_Model)
	//屏蔽字.xlsx
	Register("shidle", &Shidle_Model)
	//属性统计.xlsx
	Register("attr", &Attr_Model)
	//怪物刷新表.xlsx
	Register("npc_Refresh", &Npc_Refresh_Model)
	//战斗区域表.xlsx
	Register("fight_area", &Fight_area_Model)
	//扫荡掉落表.xlsx
	Register("sweep_drop", &Sweep_drop_Model)
	//扫荡表.xlsx
	Register("sweep_level", &Sweep_level_Model)
	//技能体表.xlsx
	Register("skill_body", &Skill_body_Model)
	//技能升级表.xlsx
	Register("skill_level", &Skill_level_Model)
	//技能子攻击表.xlsm
	Register("skill_child", &Skill_child_Model)
	//技能树模板表.xlsx
	Register("talentinfo", &Talentinfo_Model)
	//技能组表.xlsx
	Register("skillgroups", &Skillgroups_Model)
	//技能表.xlsx
	Register("skill", &Skill_Model)
	//抽卡数据表.xlsx
	Register("lottery", &Lottery_Model)
	//抽卡数据表.xlsx
	Register("lotterypool", &Lotterypool_Model)
	//掉落表.xlsx
	Register("drop", &Drop_Model)
	//新BUFF表.xlsx
	Register("newbuff", &Newbuff_Model)
	//机关刷新表.xlsx
	Register("MechanismRefresh", &MechanismRefresh_Model)
	//机关实体表.xlsx
	Register("MechanismObj", &MechanismObj_Model)
	//机关模板表.xlsx
	Register("Mechanism", &Mechanism_Model)
	//条件表.xlsx
	Register("condition_data", &Condition_data_Model)
	//章节星数奖励.xlsx
	Register("chapter_star_reward", &Chapter_star_reward_Model)
	//筛选目标表.xlsx
	Register("ScreenTarget", &ScreenTarget_Model)
	//签到与开服七日奖励表.xlsx
	Register("daysign", &Daysign_Model)
	//签到与开服七日奖励表.xlsx
	Register("seventhday", &Seventhday_Model)
	//经验表.xlsx
	Register("hero_up", &Hero_up_Model)
	//装备表.xlsx
	Register("equipinfo", &Equipinfo_Model)
	//角色天赋表.xlsx
	Register("hero_talent", &Hero_talent_Model)
	//运营活动.xlsx
	Register("activity", &Activity_Model)
	//运营活动.xlsx
	Register("time", &Time_Model)
	//道具表.xlsx
	Register("itemInfo", &ItemInfo_Model)
	//阵营表.xlsx
	Register("camp", &Camp_Model)
	//默认数据表.xlsx
	Register("defaultdata", &Defaultdata_Model)
	initOther()
}
