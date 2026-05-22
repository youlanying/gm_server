package netmsg

import (
	"fmt"
	"gm_server/src/ctimer"
	"gm_server/src/helper"
	"gm_server/src/logger"
	"gm_server/src/network"
	network_message "gm_server/src/network/message"
)

/*
 *	处理各个服务器连接及信息分发
 */

const (
	HEART_BEAT_INTV = 5000
	MAX_RETRY_TIMES = 1024
)

type centerGSSession struct {
	serverId   int32
	netSession *network.Session
}

var (
	_isShutDown bool
	//_centerServerSessionMap map[int32]*centerGSSession // server ID

)

func InitServerManager() {
	_isShutDown = false
	//_centerServerSessionMap = make(map[int32]*centerGSSession)
}

//关闭连接
func CloseCenterLink(num int32) {
	conn, ok := GMConnector[num]
	//fmt.Printf("=====conn:%v=\n", conn)
	if ok && conn != nil {
		delete(ReconnectGMEnable, num)
		conn.Stop()
		Isok, ok1 := ReconnectGMEnable[num]
		fmt.Printf("==CloseCenterLink==Isok:%v,ok1:%v\n", Isok, ok1)
		delete(GMConnector, num)
	}
}

func CheckCenterLink(num int32, addrs string) string {
	_, ok := GMConnector[num]
	if !ok {
		fmt.Println("===CheckCenterLink==Create=")
		ReconnectGMEnable[num] = true
		go CreateCenterLink(num, addrs)
		return "create"
	}
	fmt.Println("===CheckCenterLink==in=")
	return "in"
}

//创建center连接
func CreateCenterLink(num int32, addrs string) {
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("beegoWeb to center server Start  %v  %v>>>>>>>>>\n", num, addrs)
	fmt.Println("--------------------------------------------------------")

	DummyConnector(num)
	//TODO 心跳包
	setHeartBeatTimer(num)
	for retryCount := 0; retryCount < MAX_RETRY_TIMES; retryCount++ {
		err := loop(num, addrs)
		if err == nil {
			break
		}
		logger.LogErr("recovered ", err)
		logger.LogWarnf("restart beego server retry times: %d ...", retryCount)
	}

	logger.LogWarn("Center Server Exit!")
	shutDown()
}

func shutDown() {
	if !_isShutDown {
		_isShutDown = true
		logger.Log("shutDown...")
	}
}

func loop(num int32, addrs string) (er interface{}) {
	er = true
	defer func() {
		logger.Logf("beegoWeb server exit.%v", num)
		if er = recover(); er != nil {
			helper.BackTrace("beegoWeb loop()")
		}
	}()
L:
	for {
		//连接center服
		checkConnect(num, addrs)
		select {
		case _, err := <-GMConnector[num].ExitChannel:
			logger.LogErrf("Center chan closed:%v,num:%v", err, num)
			_, ok := ReconnectGMEnable[num]
			if ok {
				ReconnectGMEnable[num] = true
			} else {
				er = nil
				break L
			}
		case msg, err := <-GMConnector[num].MessageChannel:
			if !err {
				logger.LogErrf("net msg err:%v,num:%v", err, num)
				continue
			}
			if gameServerAvailable(num) {
				onCenterMessageHandler(msg)
			}
		case id, err := <-ctimer.C_TIMER:
			timerChanHandler(id, err)
		}
	}
	return
}

func setHeartBeatTimer(num int32) {
	ctimer.CreateTimer(0, HEART_BEAT_INTV, ctimer.CTData{Action: ctimer.ACTION_HEARTBEAT, Data: num})
}

func doHeartBeatRequest(data interface{}) {
	num := data.(int32)
	hbRequest := &network_message.HeartBeatRequest{}
	SendMsgToGMServer(num, hbRequest)
	setHeartBeatTimer(num)
}

func checkConnect(num int32, connectAddr string) bool {
	bol, ok := ReconnectGMEnable[num]
	if ok && bol {
		//logger.Logf("=========checkConnect====ReconnectCenterEnable：%v,%v", bol, ok)
		err, conn := network.ConnectServer(connectAddr, 32767)
		//logger.Logf("=========checkConnect====err:%v, conn：%v", err, conn)
		if err == nil {
			GMConnector[num] = conn
			ReconnectGMEnable[num] = false
			//logger.Logf("===checkConnect==conn:%v", conn)
			isTrue := connectedCenterServer(num)
			return isTrue
		}
		return false
	}
	return true
}

func gameServerAvailable(num int32) bool {
	if ReconnectGMEnable[num] || GMConnector[num] == nil {
		return false
	}
	return true
}
