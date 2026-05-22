package netmsg

import (
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"gm_server/src/logger"
	"gm_server/src/network"
	network_message "gm_server/src/network/message"
)

//------------------------------------- 消息注册函数 ------------------------------------------------
// Center心跳包
func doCNHeartBeatRequest(netSession *network.Session, tempMsg proto.Message) {
	logger.Log("HeartBeat   HeartBeat")
}

func connectedCenterServer(num int32) bool {
	connSync := &network_message.GM_ConnectedNtf{}
	isTrue := SendMsgToGMServer(num, connSync)
	return isTrue
}

func onCenterMessageHandler(msg *network.NetMessage) {
	if _gmMsgRegister.OnMsgReceiveMessageHandler(msg) {
		return
	}
}

func doGM_UpdateVersionResponse(netSession *network.Session, tempMsg proto.Message) {
	uvr := tempMsg.(*network_message.GM_UpdateVersionResponse)
	logger.Log("--------------->>> version update over! state:", uvr.State, ", sids:", uvr.Sids)
	s := GetSession(uvr.SessionId)
	if s != nil {
		s.Write(*uvr)
		DelSession(uvr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_UserDataResponse(netSession *network.Session, tempMsg proto.Message) {
	udr := tempMsg.(*network_message.GM_UserDataResponse)
	s := GetSession(udr.SessionId)
	if s != nil {
		s.Write(*udr)
		DelSession(udr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_KickRoleResponse(netSession *network.Session, tempMsg proto.Message) {
	kr := tempMsg.(*network_message.GM_KickRoleResponse)
	s := GetSession(kr.SessionId)
	if s != nil {
		s.Write(*kr)
		DelSession(kr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_NoTalkResponse(netSession *network.Session, tempMsg proto.Message) {
	nt := tempMsg.(*network_message.GM_NoTalkResponse)
	s := GetSession(nt.SessionId)
	if s != nil {
		s.Write(*nt)
		DelSession(nt.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_PassRookieResponse(netSession *network.Session, tempMsg proto.Message) {
	pr := tempMsg.(*network_message.GM_PassRookieResponse)
	s := GetSession(pr.SessionId)
	if s != nil {
		s.Write(*pr)
		DelSession(pr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_UpdateUserDataResponse(netSession *network.Session, tempMsg proto.Message) {
	uudr := tempMsg.(*network_message.GM_UpdateUserDataResponse)
	s := GetSession(uudr.SessionId)
	if s != nil {
		s.Write(*uudr)
		DelSession(uudr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_QuestDataResponse(netSession *network.Session, tempMsg proto.Message) {
	qdr := tempMsg.(*network_message.GM_QuestDataResponse)
	s := GetSession(qdr.SessionId)
	if s != nil {
		s.Write(*qdr)
		DelSession(qdr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_ItemDataResponse(netSession *network.Session, tempMsg proto.Message) {
	idr := tempMsg.(*network_message.GM_ItemDataResponse)
	s := GetSession(idr.SessionId)
	if s != nil {
		s.Write(*idr)
		DelSession(idr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_UpdateItemDataResponse(netSession *network.Session, tempMsg proto.Message) {
	uidr := tempMsg.(*network_message.GM_UpdateItemDataResponse)
	s := GetSession(uidr.SessionId)
	if s != nil {
		s.Write(*uidr)
		DelSession(uidr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_StrengthenResponse(netSession *network.Session, tempMsg proto.Message) {
	sr := tempMsg.(*network_message.GM_StrengthenResponse)
	s := GetSession(sr.SessionId)
	if s != nil {
		s.Write(*sr)
		DelSession(sr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_UpdateStrengthenResponse(netSession *network.Session, tempMsg proto.Message) {
	usr := tempMsg.(*network_message.GM_UpdateStrengthenResponse)
	s := GetSession(usr.SessionId)
	if s != nil {
		s.Write(*usr)
		DelSession(usr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_HeroProgressResponse(netSession *network.Session, tempMsg proto.Message) {
	usr := tempMsg.(*network_message.GM_HeroProgressResponse)
	s := GetSession(usr.SessionId)
	if s != nil {
		s.Write(*usr)
		DelSession(usr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_GetTerminateResponse(netSession *network.Session, tempMsg proto.Message) {
	dmr := tempMsg.(*network_message.GM_GetTerminateResponse)
	s := GetSession(dmr.SessionId)
	if s != nil {
		s.Write(*dmr)
		DelSession(dmr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

//---- 邮件相关 Response ----

func doGM_SendRoleMailResponse(netSession *network.Session, tempMsg proto.Message) {
	mr := tempMsg.(*network_message.GM_GetAllMailRequest)
	s := GetSession(mr.SessionID)
	if s != nil {
		s.Write(*mr)
		DelSession(mr.SessionID)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_GetAllMailResponse(netSession *network.Session, tempMsg proto.Message) {
	dmr := tempMsg.(*network_message.GM_GetAllMailResponse)
	s := GetSession(dmr.SessionID)
	if s != nil {
		s.Write(*dmr)
		DelSession(dmr.SessionID)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_DelMailResponse(netSession *network.Session, tempMsg proto.Message) {
	dmr := tempMsg.(*network_message.GM_DelMailResponse)
	s := GetSession(dmr.SessionID)
	if s != nil {
		s.Write(*dmr)
		DelSession(dmr.SessionID)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_CreateNumberResponse(netSession *network.Session, tempMsg proto.Message) {
	dmr := tempMsg.(*network_message.GM_CreateNumberResponse)
	s := GetSession(dmr.SessionId)
	if s != nil {
		s.Write(*dmr)
		DelSession(dmr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}

func doGM_AddNoticeResponse(netSession *network.Session, tempMsg proto.Message) {
	dmr := tempMsg.(*network_message.GM_AddNoticeResponse)
	s := GetSession(dmr.SessionId)
	if s != nil {
		s.Write(*dmr)
		DelSession(dmr.SessionId)
	} else {
		beego.Warn("can not find session")
	}
}
