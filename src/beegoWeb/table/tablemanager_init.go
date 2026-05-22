package table

func init() {

	Dic = make(map[string]TableMgrInterface, 10+OtherTableLen)

	//AdminTooles管理权限表.xlsx
	Register("admin_action", &Admin_action_Model)
	//AdminTooles管理权限表.xlsx
	Register("admin_action_ui", &Admin_action_ui_Model)
	//主人公—英雄角色.xlsx
	Register("consume_material", &Consume_material_Model)
	//主人公—英雄角色.xlsx
	Register("consume_material_ratio", &Consume_material_ratio_Model)
	//主人公—英雄角色.xlsx
	Register("hero_break", &Hero_break_Model)
	//主人公—英雄角色.xlsx
	Register("hero_information", &Hero_information_Model)
	//主人公—英雄角色.xlsx
	Register("hero_model", &Hero_model_Model)
	//主人公—英雄角色.xlsx
	Register("hero_star", &Hero_star_Model)
	//主人公—英雄角色.xlsx
	Register("hero_team", &Hero_team_Model)
	//道具表.xlsx
	Register("itemInfo", &ItemInfo_Model)
	initOther()
}
