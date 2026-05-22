package network

/*
 *	消息注册机
 */
import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"gm_server/src/bridge"
	"gm_server/src/helper"
	"gm_server/src/logger"
	"gm_server/src/stack_helper"
	"reflect"
)

type MessageRegister struct {
	description      string
	msgCMDHandlerMap map[string]*msgCBHander
}

type MsgCBFunc func(netSession *Session, msg proto.Message)
type msgCBHander struct {
	handler MsgCBFunc
	async   bool //TODO 存在阻塞主线程时 异步操作为真（暂定）
	msgType reflect.Type
}

func CreatMessageRegister(des string) *MessageRegister {
	msgRegister := MessageRegister{
		description:      des,
		msgCMDHandlerMap: make(map[string]*msgCBHander),
	}
	//registeCMDHandler("LoginRequest", false, onLoginHandler)
	return &msgRegister
}

func (self *MessageRegister) RegisteCMDHandler(cmd string, handlerFunc MsgCBFunc, async bool) {
	if _, has := self.msgCMDHandlerMap[cmd]; has {
		logger.LogErrf("[" + self.description + "] Repeated CMD !")
		return
	}
	msgType := proto.MessageType(bridge.MSG_PACKAGE_NAME + cmd)
	if msgType == nil {
		logger.LogErrf("[" + self.description + "] CMD NOT FOUND! ====> " + cmd)
		return
	}
	cmdhandler := msgCBHander{
		handler: handlerFunc,
		async:   async,
		msgType: msgType,
	}
	logger.Logf("Registe CMD: [%s], msgT: {%v}.", cmd, msgType)
	self.msgCMDHandlerMap[cmd] = &cmdhandler
}

func (self *MessageRegister) executeCmdFunc(hander *msgCBHander, netSession *Session, data interface{}) {
	defer func() {
		if err := recover(); err != nil {
			helper.BackTrace("executeCmdFunc routinue")
			stack_helper.NetWorkBackTrace("executeCmdFunc routinue==>", reflect.New(hander.msgType.Elem()).Interface().(proto.Message).String())
			logger.LogErrf("[%s] executeCmdFunc err:%v", self.description, err)
		}
	}()
	//reflect.New(t2.Elem()).Interface().(Speaker)
	msgType := hander.msgType
	msg := reflect.New(msgType.Elem()).Interface().(proto.Message)

	err := proto.Unmarshal(data.([]byte), msg)
	if err != nil {
		logger.LogErrf(fmt.Sprintf("[%s] proto[%v].Unmarshal error : %v. ---> msg:%+v.", self.description, msgType, err, msg))
		return
	}
	hander.handler(netSession, msg)
}

func (self *MessageRegister) OnMsgReceiveMessageHandler(msg *NetMessage) bool {
	//logger.LogDebugf("[%s] ReceiveMessage CMD: %v ", self.description, msg.Head)
	hander, has := self.msgCMDHandlerMap[msg.Head]
	if has {
		if hander.async {
			go self.executeCmdFunc(hander, msg.NetSession, msg.Body)
		} else {
			self.executeCmdFunc(hander, msg.NetSession, msg.Body)
		}
		return true
	}
	return false
}
