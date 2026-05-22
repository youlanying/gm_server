package table

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Hero_model 原表为主人公—英雄角色.xlsx 的子表 英雄模型表
type Hero_model struct {
	Id                   int32     `json:"id"`                   //模板id
	Name                 string    `json:"name"`                 //英雄名
	Hero_class           int32     `json:"hero_class"`           //攻击方式
	Level                int32     `json:"level"`                //出生等级
	Skill_ids            []int32   `json:"skill_ids"`            //技能ID
	Skill_far            []int32   `json:"skill_far"`            //远程普攻技能
	Fight_btn_ids        []int32   `json:"fight_btn_ids"`        //按键对应
	IdeTime              float32   `json:"ideTime"`              //角色进入待机动画时间
	Regular_atlas        string    `json:"regular_atlas"`        //立绘图集
	Regular_icon         string    `json:"regular_icon"`         //立绘图片名称
	Regular_iconpath     string    `json:"regular_iconpath"`     //立绘图片的路径
	Pos_rot              []float32 `json:"pos_rot"`              //主页面3D角色位置（位置xyz旋转xyz）
	Pos4_3               []float32 `json:"pos4_3"`               //立绘位置4：3
	Scale16_9            []float32 `json:"scale16_9"`            //立绘缩放16：9
	Scale4_3             []float32 `json:"scale4_3"`             //立绘缩放4：3
	Atlas                string    `json:"atlas"`                //角色卡牌图集名称
	Icon                 string    `json:"icon"`                 //主界面卡牌头像
	Story_id             int32     `json:"story_id"`             //英雄背景故事（文字表ID）
	Resident_effect      []int32   `json:"resident_effect"`      //常驻特效列表
	HitSound             []string  `json:"hitSound"`             //玩家受击音效列表
	HitSoundProbability  int32     `json:"hitSoundProbability"`  //玩家播放受击音效的播放概率
	DeadSound            []string  `json:"deadSound"`            //玩家死亡音效列表
	DeadSoundProbability int32     `json:"deadSoundProbability"` //玩家播放死亡音效的播放概率
	Special_skill_sound  []string  `json:"special_skill_sound"`  //奥义技玩家音效(三段奥义音效)
	IsPictureRot         int32     `json:"IsPictureRot"`         //立绘是否翻转(1翻转，0不翻转)
	Timeequip            string    `json:"timeequip"`            //角色模型路径
	Timeequippvp         string    `json:"timeequippvp"`         //PVP角色模型路径
	Conditions           []int32   `json:"conditions"`           //角色功能开启条件id
	AutoFightChase       float32   `json:"autoFightChase"`       //追击距离
	AutoFightGuard       float32   `json:"autoFightGuard"`       //警戒范围
	AutoFightMaxDistance float32   `json:"autoFightMaxDistance"` //最大距离检测（超出这个距离会重新检测）
	Order_id             int32     `json:"order_id"`             //开放次序
	Small_sprite         string    `json:"small_sprite"`         //小头像
	Big_sprite           string    `json:"big_sprite"`           //左上血量头像
	Dead_sprite          string    `json:"dead_sprite"`          //死亡头像
	Dead_effect          int32     `json:"dead_effect"`          //死亡特效
	Property             int32     `json:"property"`             //属性
	Gift_item            [][]int32 `json:"gift_item"`            //分解所获物品
	Role_image           string    `json:"role_image"`           //角色详细信息界面英雄立绘
	Role_name_image      string    `json:"role_name_image"`      //角色详细信息界面英雄名字图片名（默认都在Main图集
	Infodes              []int32   `json:"infodes"`              //角色资料分页英雄故事id，声优id，美术id
	Roleimage_position   []float32 `json:"roleimage_position"`   //英雄详细信息界面人物立绘位置
	Formation_atlas      string    `json:"formation_atlas"`      //编队角色卡牌图集名称
	Formation_icon       string    `json:"formation_icon"`       //编队卡牌头像
	Property_limte       []int32   `json:"property_limte"`       //属性上限
	ChangeHeroIcon       []string  `json:"changeHeroIcon"`       //换角色效果图头像
	WarnRange            float32   `json:"warnRange"`            //角色的警戒范围
	AniRoleName          string    `json:"AniRoleName"`          //动态立绘名
	AniRoleSpeak         [][]int32 `json:"AniRoleSpeak"`         //动态立绘对话
	RunAudio             string    `json:"runAudio"`             //走路音效
	Hero_startaudio      int32     `json:"hero_startaudio"`      //战斗开始音效
	Hero_selectaudio     int32     `json:"hero_selectaudio"`     //战斗换人音效
	Hero_loseaudio       int32     `json:"hero_loseaudio"`       //战斗失败音效
	Hero_winaudio        int32     `json:"hero_winaudio"`        //战斗获胜音效
	ProfessionId         int32     `json:"professionId"`         //职能ID
	Prioritylock         float32   `json:"prioritylock"`         //角色优先锁敌
	FightSuccess_Icon    []string  `json:"FightSuccess_Icon"`    //结算胜利界面英雄头像
	Over_closeup         []string  `json:"over_closeup"`         //结算特写动画名称
	AttackDistance       float32   `json:"attackDistance"`       //职业参考攻击距离（攻击自动拉开距离）
	Shortname            string    `json:"shortname"`            //英雄简称
	Atlasintensify       string    `json:"atlasintensify"`       //强化界面头像图集名称
	Intensify_selected   string    `json:"intensify_selected"`   //强化界面头像选中名称
	Intensify_unselected string    `json:"intensify_unselected"` //强化界面头像未选中名称
	Star                 int32     `json:"star"`                 //初始星级
	Seat_property        int32     `json:"seat_property"`        //选人界面英雄位置优先级
	Choosehero_atlas     string    `json:"choosehero_atlas"`     //选人界面角色卡牌图集名称
	Choosehero_icon      string    `json:"choosehero_icon"`      //选人界面卡牌头像名称
	PlayerGravityf       float32   `json:"playerGravityf"`       //角色重力大小
	HintAddHeight        float32   `json:"hintAddHeight"`        //英雄掉血跳字的附加高度
	PositionType         int32     `json:"positionType"`         //站位类型（前、中、后排）
	PositionType2        int32     `json:"positionType2"`        //站位类型2（中坚、两翼）
	PositionParam        int32     `json:"positionParam"`        //站位参数
	Camp                 int32     `json:"camp"`                 //所属阵营
	Becrush              int32     `json:"becrush"`              //是否开始碰撞等级【0否；1是】
	IsDisableHero        int32     `json:"isDisableHero"`        //是否在英雄列表中屏蔽此英雄
	CameraR              float32   `json:"cameraR"`              //摄像机半径增减
	AimCameraPosition    []float32 `json:"aimCameraPosition"`    //远程瞄准镜头位置
	AimCameraRotation    []float32 `json:"aimCameraRotation"`    //远程瞄准镜头角度
	BTTemplates          []string  `json:"BTTemplates"`          //行为树模板组
	Model_path           string    `json:"model_path"`           //人物模型路径（功能不明确）
	Modelfbx_name        string    `json:"modelfbx_name"`        //人物动画路径
}

type Hero_modelMgr struct {
}

var (
	Hero_model_Model Hero_modelMgr
	hero_modelDic    map[int32]*Hero_model
	// Hero_model_All 主人公—英雄角色.xlsx (英雄模型表)
	Hero_model_All []*Hero_model
)

// Hero_model_Get 主人公—英雄角色.xlsx (英雄模型表)
func Hero_model_Get(Id int32) (*Hero_model, bool) {
	data, ok := hero_modelDic[Id]
	if !ok {
		PROTO_ERROR_ID = "主人公—英雄角色.xlsx\nhero_model not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Hero_modelMgr) PrintArr() {
	vLen := len(Hero_model_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Hero_modelMgr) PrintArrOne(index int) {
	logger.Logf("==Hero_model==:%+v", Hero_model_All[index])
}

func (*Hero_modelMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := hero_modelDic[int32Key]
		if !ok {
			logger.LogWarn("Hero_model PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Hero_model==:%+v", vData)
	}
}

func (*Hero_modelMgr) Load(buffer []byte) bool {
	Hero_model_All = make([]*Hero_model, 0)
	err := json.Unmarshal(buffer, &Hero_model_All)
	if err != nil {
		logger.LogErr("Hero_model JsonFailed:", err)
		return false
	}
	vLen := len(Hero_model_All)
	hero_modelDic = make(map[int32]*Hero_model, vLen)
	for _, mem := range Hero_model_All {
		hero_modelDic[mem.Id] = mem
	}
	return true
}
