package netmsg

import (
	"gm_server/src/logger"
	"gm_server/src/network"
)

/*
 *	beegoWeb 消息注册
 */

/*
	信息Handler注册 初始化
*/
func RegisterInit() {
	logger.Logf("----------------------MSG register Init----------------------------")
	registerGateH()
}

//------------------- GATE 信息注册 --------------------------------------------
var (
	_gateMsgRegister *network.MessageRegister
	_gmMsgRegister   *network.MessageRegister
)

func registerGateH() {
	_gmMsgRegister = network.CreatMessageRegister("GM_CN")
	_gmMsgRegister.RegisteCMDHandler("HeartBeatRequest", doCNHeartBeatRequest, false)
	_gmMsgRegister.RegisteCMDHandler("GM_UpdateVersionResponse", doGM_UpdateVersionResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_UserDataResponse", doGM_UserDataResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_KickRoleResponse", doGM_KickRoleResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_NoTalkResponse", doGM_NoTalkResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_PassRookieResponse", doGM_PassRookieResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_UpdateUserDataResponse", doGM_UpdateUserDataResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_QuestDataResponse", doGM_QuestDataResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_ItemDataResponse", doGM_ItemDataResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_UpdateItemDataResponse", doGM_UpdateItemDataResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_StrengthenResponse", doGM_StrengthenResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_UpdateStrengthenResponse", doGM_UpdateStrengthenResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_HeroProgressResponse", doGM_HeroProgressResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_GetTerminateResponse", doGM_GetTerminateResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_SendRoleMailResponse", doGM_SendRoleMailResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_GetAllMailResponse", doGM_GetAllMailResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_DelMailResponse", doGM_DelMailResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_CreateNumberResponse", doGM_CreateNumberResponse, false)
	_gmMsgRegister.RegisteCMDHandler("GM_AddNoticeResponse", doGM_AddNoticeResponse, false)

	_gateMsgRegister = network.CreatMessageRegister("GAME-GT")
	//_gateMsgRegister.RegisteCMDHandler("HeartBeatRequest", doGTHeartBeatRequest, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_NewMailResponse", doPS_NewMailResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_GS_GmCmdResponse", doPS_GS_GmCmdResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_GS_GmCmdUpgradedResponse", doPS_GS_GmCmdUpgradedResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_KickUserResponse", doPS_KickUserResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_ModifyBlackListResponse", doPS_ModifyBlackListResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_SearchNameResponse", doPS_SearchNameResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_ModifyWhiteListResponse", doPS_ModifyWhiteListResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_ModifyWhiteListIpResponse", doPS_ModifyWhiteListIpResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_ModifyQaModeResponse", doPS_ModifyQaModeResponse, false)
	//_gateMsgRegister.RegisteCMDHandler("PS_GetAccountIdListResponse", doPS_GetAccountIdListResponse, false)
}
