package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Fight_area 原表为战斗区域表.xlsx 的子表 战斗区域表
type Fight_area struct {
	Id                         int32       `json:"id"`                         //战斗区域ID
	FightWall_Type             []int32     `json:"fightWall_Type"`             //战斗区域墙类型（1圆形，2矩形）
	FightWall_Position         [][]float32 `json:"fightWall_Position"`         //战斗区域墙位置
	FightWall_Rotation         [][]float32 `json:"fightWall_Rotation"`         //战斗区域墙旋转角度
	FightWall_Scale            [][]float32 `json:"fightWall_Scale"`            //战斗区域墙缩放
	WarnTriggerEntities        []string    `json:"WarnTriggerEntities"`        //警戒触发战斗区域的实体idList
	KillTriggerEntities        []string    `json:"KillTriggerEntities"`        //击杀触发战斗区域的实体idList
	Generating_Condition       []int32     `json:"generating_Condition"`       //战斗区域空气墙触发条件
	GeneratingTrigger          []float32   `json:"generatingTrigger"`          //生成触发器位置
	GeneratingTrigger_Rotation []float32   `json:"generatingTrigger_Rotation"` //生成触发器旋转
	GeneratingTrigger_Scale    []float32   `json:"generatingTrigger_Scale"`    //生成触发器缩放
	WarnCloseEntities          []string    `json:"WarnCloseEntities"`          //警戒关闭战斗区域实体idList
	KillCloseEntities          []string    `json:"KillCloseEntities"`          //击杀关闭战斗区域实体idList
	Disappear_conditions       []int32     `json:"disappear_conditions"`       //战斗空气墙解除条件
	DisappearTrigger           []float32   `json:"disappearTrigger"`           //消失触发器位置
	DisappearTrigger_Rotation  []float32   `json:"disappearTrigger_Rotation"`  //消失触发器旋转
	DisappearTrigger_Scale     []float32   `json:"disappearTrigger_Scale"`     //消失触发器缩放
	MoveArea                   [][]float32 `json:"moveArea"`                   //怪物活动区域位置
	MoveArea_Scale             []float32   `json:"moveArea_Scale"`             //怪物活动区域
	LookBoxPos                 []float32   `json:"lookBoxPos"`                 //战斗墙消失的时候看向的盒子位置
	LookBoxAngle               []float32   `json:"lookBoxAngle"`               //战斗墙消失的时候看向的盒子旋转信息
	LookBoxScale               []float32   `json:"lookBoxScale"`               //战斗墙消失的时候看向的盒子缩放信息
}

type Fight_areaMgr struct {
}

var (
	Fight_area_Model Fight_areaMgr
	fight_areaDic    map[int32]*Fight_area
	// Fight_area_All 战斗区域表.xlsx (战斗区域表)
	Fight_area_All []*Fight_area
)

// Fight_area_Get 战斗区域表.xlsx (战斗区域表)
func Fight_area_Get(Id int32) (*Fight_area, bool) {
	data, ok := fight_areaDic[Id]
	if !ok {
		PROTO_ERROR_ID = "战斗区域表.xlsx\nfight_area not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *Fight_areaMgr) PrintArr() {
	vLen := len(Fight_area_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*Fight_areaMgr) PrintArrOne(index int) {
	logger.Logf("==Fight_area==:%+v", Fight_area_All[index])
}

func (*Fight_areaMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := fight_areaDic[int32Key]
		if !ok {
			logger.LogWarn("Fight_area PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Fight_area==:%+v", vData)
	}
}

func (*Fight_areaMgr) Load(buffer []byte) bool {
	Fight_area_All = make([]*Fight_area, 0)
	err := json.Unmarshal(buffer, &Fight_area_All)
	if err != nil {
		logger.LogErr("Fight_area JsonFailed:", err)
		return false
	}
	vLen := len(Fight_area_All)
	fight_areaDic = make(map[int32]*Fight_area, vLen)
	for _, mem := range Fight_area_All {
		fight_areaDic[mem.Id] = mem
	}
	return true
}
