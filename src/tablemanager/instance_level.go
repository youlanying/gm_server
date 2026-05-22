package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Instance_level 原表为关卡_场景表.xlsx 的子表 场景表
type Instance_level struct {
	Id                    int32       `json:"id"`                    //关卡ID
	Before_level          []int32     `json:"before_level"`          //前置关卡id
	Next_level            []int32     `json:"next_level"`            //后置关卡id
	IsBeg                 int32       `json:"isBeg"`                 //是否是大关
	Level_nameImg         []string    `json:"level_nameImg"`         //关卡显示信息
	MonsterList           []int32     `json:"monsterList"`           //读取NPC模型表ID
	MonsterListData       []int32     `json:"monsterListData"`       //读取NPC数据表ID
	Level_son_id          []int32     `json:"level_son_id"`          //子场景序列
	CreatePlayerPoint     [][]float32 `json:"createPlayerPoint"`     //角色出生位置与转向
	Camera                string      `json:"Camera"`                //关卡对应的镜头
	Camerapos             []float32   `json:"Camerapos"`             //关卡对应镜头位置
	DelayChapterName      float32     `json:"delayChapterName"`      //延迟展示章节名时间
	Chapter_id            int32       `json:"chapter_id"`            //章节ID
	Chapter_name          int32       `json:"chapter_name"`          //章节名称
	Chapter_namepic       []string    `json:"chapter_namepic"`       //章节图集名称,图片名称
	Name                  int32       `json:"name"`                  //关卡名称
	Level_type            int32       `json:"level_type"`            //关卡玩法类型
	Level_anchor_coord    []float32   `json:"level_anchor_coord"`    //关卡节点坐标
	Better_award          []int32     `json:"better_award"`          //显示常规通关奖励物品ID
	Better_award_First    []int32     `json:"better_award_First"`    //显示首次通关奖励
	Better_award_FullStar []int32     `json:"better_award_FullStar"` //显示三星通关奖励
	Hero_likability       int32       `json:"hero_likability"`       //通关角色好感度奖励
	Hero_exp              int32       `json:"hero_exp"`              //通关角色经验奖励
	Team_exp              int32       `json:"team_exp"`              //通关团队经验奖励
	Rank_type             int32       `json:"rank_type"`             //榜单类型
	Score_type            []int32     `json:"score_type"`            //关卡积分类型
	Max_money             int32       `json:"max_money"`             //金钱预警
	Role_max_exp          int32       `json:"role_max_exp"`          //战队经验预警
	Hero_max_exp          int32       `json:"hero_max_exp"`          //角色英雄经验预警
	Level_of_difficulty   []float32   `json:"level_of_difficulty"`   //关卡难度系数
	Scene_hurt_param      int32       `json:"scene_hurt_param"`      //关卡伤害公式调整
	HeroList              []int32     `json:"heroList"`              //关卡指定的角色ID
	Ban_heros             []int32     `json:"ban_heros"`             //关卡禁用的角色ID
	Type                  int32       `json:"type"`                  //关卡类型
	Level_limit           int32       `json:"level_limit"`           //战队等级限制
	Physical_lost         int32       `json:"physical_lost"`         //体力消耗
	Day_times             int32       `json:"day_times"`             //每日限制次数
	IsCameraPath          int32       `json:"isCameraPath"`          //是否开启走镜
	WalkView_ani          string      `json:"walkView_ani"`          //走镜动画
	StartLightMapId       int32       `json:"startLightMapId"`       //初始LightMap
	Fightsceneeffect      []string    `json:"fightsceneeffect"`      //预加载的战斗场景特效
	Fightsceneatlas       []string    `json:"fightsceneatlas"`       //预加载的战斗界面图集
	Star_id               []int32     `json:"star_id"`               //星级条件ID
	NpcNumLimit           int32       `json:"npcNumLimit"`           //场中拥有的怪物上限
	Story_trigger_id      int32       `json:"story_trigger_id"`      //关卡开场剧情
	StartTimelineId       int32       `json:"startTimelineId"`       //关卡开场timeline
	ContinueTime          []float32   `json:"continueTime"`          //开场剧情时间和显示boss头像时间
	Distance              float32     `json:"distance"`              //BOSS区域距离（单位/米）
	Over_time             int32       `json:"over_time"`             //通关时间(用作ui显示倒计时)
	IsBossLevel           int32       `json:"isBossLevel"`           //是否为BOSS关
	Success_condition     []int32     `json:"success_condition"`     //通关条件
	Fail_condition        []int32     `json:"fail_condition"`        //失败条件
}

type Instance_levelMgr struct {
}

var (
	Instance_level_Model Instance_levelMgr
	instance_levelDic    map[int32]*Instance_level
	// Instance_level_All 关卡_场景表.xlsx (场景表)
	Instance_level_All []*Instance_level
)

// Instance_level_Get 关卡_场景表.xlsx (场景表)
func Instance_level_Get(Id int32) (*Instance_level, bool) {
	data, ok := instance_levelDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡_场景表.xlsx\ninstance_level not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Instance_levelMgr) PrintArr() {
	vLen := len(Instance_level_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Instance_levelMgr) PrintArrOne(index int) {
	logger.Logf("==Instance_level==:%+v", Instance_level_All[index])
}

func (*Instance_levelMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := instance_levelDic[int32Key]
		if !ok {
			logger.LogWarn("Instance_level PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Instance_level==:%+v", vData)
	}
}

func (*Instance_levelMgr) Load(buffer []byte) bool {
	Instance_level_All = make([]*Instance_level, 0)
	err := json.Unmarshal(buffer, &Instance_level_All)
	if err != nil {
		logger.LogErr("Instance_level JsonFailed:", err)
		return false
	}
	vLen := len(Instance_level_All)
	instance_levelDic = make(map[int32]*Instance_level, vLen)
	for _, mem := range Instance_level_All {
		instance_levelDic[mem.Id] = mem
	}
	return true
}
