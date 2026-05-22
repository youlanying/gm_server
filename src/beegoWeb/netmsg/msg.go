package netmsg

import (
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"gm_server/src/logger"
	"gm_server/src/network"
)

const (
	QUEUE_SIZE = 128
)

type Msg struct {
	RegionId int
	Api      int16
	Data     interface{}
}

var (
	MsgData chan Msg

	//ReconnectGateEnable bool
	//GateConnector       *network.ConnectInfo
	ReconnectGMEnable map[int32]bool
	GMConnector       map[int32]*network.ConnectInfo
)

func init() {
	ReconnectGMEnable = make(map[int32]bool)
	GMConnector = make(map[int32]*network.ConnectInfo)
}

func SendMsg(regionId int, api int16, data interface{}) {
	msg := Msg{
		RegionId: regionId,
		Api:      api,
		Data:     data,
	}
	MsgData <- msg
}

func RecMsg(sId int64) interface{} {
	if s := GetSession(sId); s != nil {
		return s.Read()
	} else {
		beego.Error("Cannot find Session.")
		return nil
	}
}

// 发送信息至网关服务器
func SendMsgToGate(pbData proto.Message) {
	//if ReconnectGateEnable || GateConnector == nil {
	//	return
	//}
	//GateConnector.Send(pbData)
}

func DummyConnector(num int32) {
	//dummyConn := network.ConnectInfo{
	//	ExitChannel:    make(chan int),
	//	MessageChannel: make(chan *network.NetMessage),
	//}
	//GateConnector = &dummyConn

	dummyConnCn := network.ConnectInfo{
		ExitChannel:    make(chan int),
		MessageChannel: make(chan *network.NetMessage),
	}
	GMConnector[num] = &dummyConnCn
}

//GS 向指定玩家发送信息
//func sendMsgToUser(uid int64, pbData proto.Message) {
//	msg := bridge.SerializeB_PktUserMsgNtf(uid, consts.NIL_SESSION_ID, pbData)
//	logger.LogDebugf("send[%s]MsgToUser<%d> --> %v", bridge.GetCMDbyMessageName(pbData), uid, pbData)
//	sendMsgToGate(msg)
//}

func SendMsgToGMServer(num int32, pbData proto.Message) bool {
	//fmt.Printf("=====SendMsgToGMServer===CenterConnector[num]:%+v==ReconnectCenterEnable[num]:%v\n",CenterConnector[num],ReconnectCenterEnable[num])
	isBool, ok := ReconnectGMEnable[num]
	_, ok1 := GMConnector[num]
	if !ok || isBool || !ok1 {
		logger.Logf("ReconnectCenterEnable = [%v],=CenterConnector = [%v]", ReconnectGMEnable[num], GMConnector[num])
		return false
	}

	err := GMConnector[num].Send(pbData)
	if err != nil {
		logger.LogErrf("========SendMsgToGMServer======err:%+v", err)
		return false
	}
	return true
}
