package tablemanager

import (
	"encoding/json"
	"gm_server/src/logger"
	"strconv"
)

// CommonEntity 原表为关卡表-通用实体表.xlsx 的子表 通用实体表
type CommonEntity struct {
	Id          int32 `json:"id"`          //实体ID
	Type        int32 `json:"Type"`        //实体类型
	DataTableID int32 `json:"DataTableID"` //实体数据表ID
}

type CommonEntityMgr struct {
}

var (
	CommonEntity_Model CommonEntityMgr
	CommonEntityDic    map[int32]*CommonEntity
	// CommonEntity_All 关卡表-通用实体表.xlsx (通用实体表)
	CommonEntity_All []*CommonEntity
)

// CommonEntity_Get 关卡表-通用实体表.xlsx (通用实体表)
func CommonEntity_Get(Id int32) (*CommonEntity, bool) {
	data, ok := CommonEntityDic[Id]
	if !ok {
		PROTO_ERROR_ID = "关卡表-通用实体表.xlsx\nCommonEntity not Id：" + strconv.FormatInt(int64(Id), 10)
		return nil, false
	}
	return data, true
}
func (this *CommonEntityMgr) PrintArr() {
	vLen := len(CommonEntity_All)
	for i := 0; i < vLen; i++ {
		this.PrintArrOne(i)
	}
}

func (*CommonEntityMgr) PrintArrOne(index int) {
	logger.Logf("==CommonEntity==:%+v", CommonEntity_All[index])
}

func (*CommonEntityMgr) PrintMapByKey(key interface{}) {
	if int32Key, ok := key.(int32); ok {
		vData, ok := CommonEntityDic[int32Key]
		if !ok {
			logger.LogWarn("CommonEntity PrintMapByKey Not Find Key", key)
			return
		}
		logger.Logf("==PrintMapByKey==CommonEntity==:%+v", vData)
	}
}

func (*CommonEntityMgr) Load(buffer []byte) bool {
	CommonEntity_All = make([]*CommonEntity, 0)
	err := json.Unmarshal(buffer, &CommonEntity_All)
	if err != nil {
		logger.LogErr("CommonEntity JsonFailed:", err)
		return false
	}
	vLen := len(CommonEntity_All)
	CommonEntityDic = make(map[int32]*CommonEntity, vLen)
	for _, mem := range CommonEntity_All {
		CommonEntityDic[mem.Id] = mem
	}
	return true
}
