package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// Activity 原表为运营活动.xlsx 的子表 运营活动表
type Activity struct {
	Id           int32   `json:"id"`           //活动编号
	Type         int32   `json:"type"`         //活动类型
	Instance     []int32 `json:"instance"`     //副本列表：[]
	Time         int32   `json:"time"`         //活动起止时间:时间表ID
	Frequency    []int32 `json:"frequency"`    //活动开启频率:[时间表ID,]
	Join_time    int32   `json:"join_time"`    //活动副本可参与次数
	Cd_cost      int32   `json:"cd_cost"`      //活动副本清除CD消耗魔晶
	Functionid   int32   `json:"functionid"`   //所属功能条件限制：功能表ID
	Condition    []int32 `json:"condition"`    //其他限制条件:[条件表id,]
	Ispush       int32   `json:"ispush"`       //是否推送
	Push_title   string  `json:"push_title"`   //推送标题
	Push_content string  `json:"push_content"` //推送内容
	Push_time    int32   `json:"push_time"`    //推送提前时间（秒数，但精度为分钟）
}

type ActivityMgr struct {
}

var (
	Activity_Model ActivityMgr
	activityDic    map[int32]*Activity
	// Activity_All 运营活动.xlsx (运营活动表)
	Activity_All []*Activity
)

// Activity_Get 运营活动.xlsx (运营活动表)
func Activity_Get(Id int32) (*Activity, bool) {
	data, ok := activityDic[Id]
	if !ok {
		PROTO_ERROR_ID = "运营活动.xlsx\nactivity not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *ActivityMgr) PrintArr() {
	vLen := len(Activity_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*ActivityMgr) PrintArrOne(index int) {
	logger.Logf("==Activity==:%+v", Activity_All[index])
}

func (*ActivityMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := activityDic[int32Key]
		if !ok {
			logger.LogWarn("Activity PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==Activity==:%+v", vData)
	}
}

func (*ActivityMgr) Load(buffer []byte) bool {
	Activity_All = make([]*Activity, 0)
	err := json.Unmarshal(buffer, &Activity_All)
	if err != nil {
		logger.LogErr("Activity JsonFailed:", err)
		return false
	}
	vLen := len(Activity_All)
	activityDic = make(map[int32]*Activity, vLen)
	for _, mem := range Activity_All {
		activityDic[mem.Id] = mem
	}
	return true
}
