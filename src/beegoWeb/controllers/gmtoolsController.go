package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gm_server/src/beegoWeb/models"
	"gm_server/src/beegoWeb/netmsg"
	network_message "gm_server/src/network/message"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type GMToolsController struct {
	GMBaseController
}

type DicResult struct {
	C  string
	ID string
	DK string
}

/**
 * 错误记录
 * 无需 Post
 */
func (c *GMBaseController) TerminateGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.Data["pagetitle"] = "异常报错查询"

	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmTerminate)
	//models.AllPlantServerList()
	//c.Data["centerurl"] = models.PlantList

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/terminate.html"
}

func (c *GMBaseController) TerminatePost() {
	//fmt.Println("=======================TerminatePost========================")
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Num := c.GetString("num")
	numInt, _ := strconv.Atoi(Num)

	userData := c.GetSession("UserData").(AdminData)
	sessionId := netmsg.NewSession()
	//fmt.Printf("====sessionId:%v,Num:%v\n", sessionId, Num)
	ok := netmsg.SendMsgToGMServer(userData.ThisPlatformId, &network_message.GM_GetTerminateRequest{
		SessionId: sessionId,
		GetNum:    int32(numInt),
	})
	if ok {
		ret := netmsg.RecMsg(sessionId).(network_message.GM_GetTerminateResponse)
		c.Ctx.WriteString(ret.TerminateData)
		return
	}
	c.Ctx.WriteString("error")
}

/**
 * 字典查询 Get
 */
func (c *GMToolsController) DictionaryGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmDictionary)
	c.Data["pagetitle"] = "字典数据查询"

	onePar := DicResult{"", "", ""}
	c.Data["par"] = []DicResult{onePar}

	//c.Data["centerurl"] = models.PlantList

	//{ok, [{pagetitle, PageTitle},{alldict, 1},{par,[{"","",""}]},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/dictionary.html"
}

/**
 * 字典查询 Post
 */
func (c *GMToolsController) DictionaryPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	Roleid := c.GetString("roleid")
	Dictkey := c.GetString("dictkey")

	var UrlParam string = ""

	Dic1 := "228,190,139,58,32,100,105,99,116,95,99,114,101,97,116,117,114,101,95,105,100,32,228,184,186,231,169,186,230,151,182,230,159,165,232,175,162,229,133,168,233,131,168"

	if Dictkey == Dic1 {
		UrlParam = "zxlf_get?cmd=gm_get_role_dictionary&roleid=" + Roleid
	} else if Dictkey == "" {
		UrlParam = "zxlf_get?cmd=gm_get_role_dictionary&roleid=" + Roleid
	} else {
		UrlParam = "zxlf_get?cmd=gm_get_role_dictionary&roleid=" + Roleid + "&dictkey=[" + Dictkey + "]"
	}

	AllUrl := CmdUrl + UrlParam

	var reStr string

	result, error := models.HttpGet(AllUrl)

	if error != nil {
		reStr = error.Error()
	} else {
		reStr = string(result)
	}

	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmDictionary)
	c.Data["pagetitle"] = "字典数据查询"

	c.Data["alldict"] = reStr

	onePar := DicResult{CmdUrl, Roleid, Dictkey}
	c.Data["par"] = []DicResult{onePar}

	//c.Data["centerurl"] = models.PlantList

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/dictionary.html"

	//Par = [{CmdUrl,Roleid,Dictkey}],
	//{ok, [{alldict, Res},{par,Par},{centerurl,NewAllCenter}]}.
}

type RedisResult struct {
	CmdUrl string
	Roleid string
	Result string
}

/**
 * Redis 查询Get
 */
func (c *GMToolsController) RedisGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.Data["pagetitle"] = "Redis数据查询"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmRedis)
	//c.Data["allredis"] = 1;

	var mem RedisResult
	mem.Roleid = ""
	mem.CmdUrl = ""
	mem.Result = ""

	//models.GetPlantServerList(plantIndex,models.All_center[plantIndex]);
	reList := []RedisResult{mem}

	c.Data["par"] = reList
	//fmt.Printf("XXXXXXXXX:%v\n",models.PlantList)
	//c.Data["centerurl"] = models.PlantList

	//{ok, [{pagetitle, PageTitle},{allredis, 1},{par,[{"","",""}]},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/redis.html"
}

/**
 * Redis 查询Post
 */
func (c *GMToolsController) RedisPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	Roleid := c.GetString("roleid")
	Redis := c.GetString("redis")
	UrlParam := "zxlf_get?cmd=gm_get_role_redis&roleid=" + Roleid + "&redis=" + Redis
	AllUrl := CmdUrl + UrlParam

	var reStr string

	result, error := models.HttpGet(AllUrl)

	if error != nil {
		reStr = error.Error()
	} else {
		reStr = string(result)
	}
	c.Data["pagetitle"] = "Redis数据查询"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmRedis)
	c.Data["allredis"] = reStr
	//c.Data["centerurl"] = models.PlantList

	var mem RedisResult
	mem.Roleid = CmdUrl
	mem.CmdUrl = Roleid
	mem.Result = Redis

	reList := []RedisResult{mem}

	c.Data["par"] = reList

	//{ok, [{allredis, Res},{par,Par},{centerurl,NewAllCenter}]}.

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/redis.html"
}

/**
 * 在线查询Get
 * 无需 Post
 */
func (c *GMToolsController) ServerNumGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.Data["pagetitle"] = "在线角色查询"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmServerNum)

	//c.Data["centerurl"] = models.PlantList

	//{ok, [{pagetitle, PageTitle},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/serverNum.html"
}

type KickResult struct {
	CmdUrl string
	Roleid string
}

/**
 * 踢人 Get
 */
func (c *GMToolsController) KickGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.Data["pagetitle"] = "踢出在线玩家"
	//c.Data["online"] = -1;

	var mem KickResult
	mem.CmdUrl = ""
	mem.Roleid = ""
	var parList = []KickResult{mem}

	c.Data["par"] = parList

	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmKick)

	//c.Data["centerurl"] = models.PlantList

	//{ok, [{pagetitle, PageTitle},{online, -1},{par,[{"",""}]},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/kick.html"
}

type HttpKickResult struct {
	Error interface{}
	Data  interface{}
}

/**
 * 踢人 Post
 */
func (c *GMToolsController) KickPost() {

	fmt.Println("KickPost")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//cmdurl := c.GetString("cmdurl")
	ServerId, _ := c.GetInt32("serverid")
	Roleid, _ := c.GetUint64("roleid")

	userData := c.GetSession("UserData").(AdminData)
	sessionId := netmsg.NewSession()
	netmsg.SendMsgToGMServer(userData.ThisPlatformId, &network_message.GM_KickRoleRequest{
		SessionId: sessionId,
		ServerId:  ServerId,
		UserId:    Roleid,
	})
	ret := netmsg.RecMsg(sessionId).(network_message.GM_KickRoleResponse)
	reStr := "" + string(ret.ResState)

	c.Data["pagetitle"] = "踢出在线玩家"
	c.Data["userid"] = v
	//userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmKick)

	var mem KickResult
	//mem.CmdUrl = cmdurl
	mem.Roleid = strconv.FormatUint(Roleid, 10)
	var parList = []KickResult{mem}

	c.Data["par"] = parList

	c.Data["online"] = reStr
	//c.Data["centerurl"] = models.PlantList

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/kick.html"

	//返回的结果
	//Par = [{CmdUrl,Roleid}],
	//NewAllCenter = get_list(),
	//{ok, [{online, Res},{par,Par},{centerurl,NewAllCenter}]}
}

/**
 * 开放时间
 */
func (c *GMToolsController) ServerTimeGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.Data["pagetitle"] = "修改开服时间"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmServerTime)
	//c.Data["centerurl"] = models.PlantList

	//{ok, [{pagetitle, PageTitle},{result,-1},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/serverTime.html"
}

/**
 * 开放时间
 */
func (c *GMToolsController) ServerTimePost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	Settimer := c.GetString("settimer")
	UrlParam := "zxlf_get?cmd=gm_update_open_server_time&times=" + Settimer
	AllUrl := CmdUrl + UrlParam

	var reStr string

	result, error := models.HttpGet(AllUrl)

	if error != nil {
		reStr = error.Error()
	} else {
		reStr = string(result)
		//reStr = "修改成功，请查询确认！"
	}

	c.Data["pagetitle"] = "修改开服时间"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmServerTime)
	//c.Data["centerurl"] = models.PlantList

	c.Data["result"] = reStr

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/serverTime.html"

	//{ok, [{result,Res},{centerurl,NewAllCenter}]}
}

/**
 * 充值
 */
func (c *GMToolsController) ChargeGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	//models.AllPlantServerList()
	c.Data["pagetitle"] = " 充  值  "
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmCharge)

	//c.Data["centerurl"] = models.PlantList

	c.Data["style"] = "" //如果没有权限  "display:none"

	//{ok, [{pagetitle, PageTitle},{centerurl,NewAllCenter},{style,Style}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/charge.html"
}

/**
 * 充值
 */
func (c *GMToolsController) ChargePost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	Roleid := c.GetString("roleid")
	Chargeid := c.GetString("chargeid")
	SellRmb := c.GetString("sell_rmb")
	ChargeGold := c.GetString("chargegold")
	//SendTime := erlang:integer_to_list(util_time:now_to_s(util_time:time_now()))
	SendTime := models.GetNowNigxStr()
	GmUser := "gmtool"
	StrSid := "0"
	StrPlayerId := "GMYLY" + SendTime
	StrAuth := GmUser + StrSid + StrPlayerId + SendTime
	NewMd5 := models.Md5V(StrAuth)
	NewGmAuth := models.GetAuth("gmtool", SendTime)
	UrlParam := "zxlf_get?cmd=user_charge&roleid=IGMYLY" + Roleid + "&goods_id=" +
		Chargeid + "&rmb=" + SellRmb + "&gold=" + ChargeGold + "&gm_user=" + GmUser + "&gm_serverid=" + StrSid +
		"&gm_operateid=" + StrPlayerId + "&gm_time=" + SendTime + "&gm_session=" + NewMd5 +
		"&gm_plat_goodsid=0&gm_platcid=GMTools&gm=" + GmUser + "&time=" + SendTime +
		"&auth=" + NewGmAuth
	AllUrl := CmdUrl + UrlParam

	var reStr string

	result, error := models.HttpGet(AllUrl)

	if error != nil {
		reStr = error.Error()
	} else {
		reStr = string(result) //?ZH("充值成功，请查询确认！");  ZH("充值失败，请核实信息！")
	}

	c.Data["pagetitle"] = " 充  值  "
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmCharge)

	c.Data["result"] = reStr
	//c.Data["centerurl"] = models.PlantList

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/charge.html"
}

/**
 * 复制角色
 */
func (c *GMToolsController) CopyRoleGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	//models.AllPlantServerList()
	c.Data["pagetitle"] = "角色数据复制"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmCopyRole)
	//c.Data["centerurl"] = models.PlantList

	//{ok, [{result,-1},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/copyRole.html"
}

type CopyRolePostStruct struct {
	Cmd    string
	Roleid string
	Redis  string
}

/**
 * 复制角色
 */
func (c *GMToolsController) CopyRolePost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	RoleId := c.GetString("roleid")
	RoleData := c.GetString("sendroledata")
	SendStruct := CopyRolePostStruct{"set_role_redis_data", RoleId, RoleData}
	AllUrl := CmdUrl + "zxlf_gmpost?"

	//发送 Post
	jsonStr, _ := json.Marshal(SendStruct) //很方便的json化
	req, err := http.NewRequest("POST", AllUrl, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)

	//"发送失败，角色信息错误！"
	//"发送成功，请登陆确认！"
	//"发送失败，角色信息错误！"

	c.Data["pagetitle"] = "角色数据复制"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmCopyRole)
	//c.Data["centerurl"] = models.PlantList

	c.Data["send"] = content

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/copyRole.html"

	//{ok, [{pagetitle, PageTitle},{send,Res},{centerurl,NewAllCenter}]}
}

/**
 * GM账号
 */
func (c *GMToolsController) GMAccountGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	//models.AllPlantServerList()
	c.Data["pagetitle"] = "GM账号设置"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmAccount)

	//c.Data["centerurl"] = models.PlantList
	//{ok, [{pagetitle, PageTitle},{centerurl,NewAllCenter}]}

	//c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/gmAccount.html"
}

/**
 * GM账号
 */
func (c *GMToolsController) GMAccountPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	Account := c.GetString("account")
	Accountlist := strings.Split(Account, "@")

	vlen := len(Accountlist)

	var UserId string = v.(string)

	if vlen > 0 {

		Time := models.GetNowNigxStr()
		Auth := models.GetAuth(v.(string), Time)
		UrlParam := "zxlf_get?cmd=gm_white_gm_account&type=set&account=" + Account + "&auth=" + Auth + "&gm=" + UserId + "&time=" + Time
		AllUrl := CmdUrl + UrlParam

		var reStr string

		result, error := models.HttpGet(AllUrl)

		if error != nil {
			reStr = error.Error()
		} else {
			reStr = string(result) //?ZH("充值成功，请查询确认！");  ZH("充值失败，请核实信息！")
		}

		//("发送失败！")
		//("发送成功，请查询确认！")
		//("发送失败！")
		//("发送可能失败，请查询确认！")

		c.Data["send"] = reStr
	} else {

		c.Data["send"] = "账号输入错误请重新输入！"
	}

	//{ok, [{send,Res},{centerurl,NewAllCenter}]}

	//c.Data["centerurl"] = models.PlantList

	c.Data["pagetitle"] = "GM账号设置"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmAccount)

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/gmAccount.html"
}

/**
 * 封号 禁言
 */
func (c *GMToolsController) BlackAccountGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	//models.AllPlantServerList()
	c.Data["pagetitle"] = "封号禁言"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmBlackAccount)
	//c.Data["centerurl"] = models.PlantList

	//{ok, [{candoact, true},{pagetitle, PageTitle},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/blackAccount.html"
}

/**
 * 封号 禁言
 */
func (c *GMToolsController) BlackAccountPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	Account := c.GetString("roleid")
	Type := c.GetString("setstate1")    //0封号 1解封
	BanType := c.GetString("setstate2") //"1"login;  "2"talk

	var BanTime string
	if Type == "0" {
		BanTime = "315360000"
	} else {
		BanTime = "0"
	}

	var UserId string = v.(string)
	Time := models.GetNowNigxStr()
	Auth := models.GetAuth(v.(string), Time)
	UrlParam := "zxlf_get?cmd=gm_ban_role&type=" + Type + "&value=" + Account + "&bantime=" + BanTime + "&bantype=" + BanType + "&auth=" + Auth + "&gm=" + UserId + "&time=" + Time
	AllUrl := CmdUrl + UrlParam

	var reStr string

	result, error := models.HttpGet(AllUrl)

	if error != nil {
		reStr = error.Error()
	} else {
		reStr = string(result)
	}

	var Refmap string
	Refmap = "[]"

	//case lists:keyfind(<<"map">>, 1, UpVersionResult) of
	//false->
	//	Refmap="[]";
	//{_,BRefmap}->
	//	Refmap = erlang:binary_to_list(BRefmap)
	//Res = ?ZH("设置成功！");

	//Res = "设置失败！"
	//Res = "设置成功！";
	//Res = "设置失败！";
	//Res = "设置可能失败，请查询确认！"

	c.Data["pagetitle"] = "封号禁言"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmBlackAccount)
	//c.Data["centerurl"] = models.PlantList

	c.Data["send"] = reStr //Res
	c.Data["refmap"] = Refmap

	//{ok, [{pagetitle, PageTitle},{send,Res},{centerurl,NewAllCenter},{refmap, Refmap}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/blackAccount.html"
}

/**
 * 创建角色
 */
func (c *GMToolsController) SetCreateGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	//models.AllPlantServerList()
	c.Data["pagetitle"] = "创建角色"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmCanCreate)
	//c.Data["centerurl"] = models.PlantList

	//{ok, [{result,-1},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/setCreate.html"
}

/**
 * 创建角色
 */
func (c *GMToolsController) SetCreatePost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	Setstate := c.GetString("setstate")
	UrlParam := "zxlf_get?cmd=gm_set_can_create&data=" + Setstate
	AllUrl := CmdUrl + UrlParam

	//("修改成功，请查询确认！");

	var reStr string

	result, error := models.HttpGet(AllUrl)

	if error != nil {
		reStr = error.Error()
	} else {
		reStr = string(result)
	}

	//{ok, [{result,Res},{centerurl,NewAllCenter}]}

	//c.Data["centerurl"] = models.PlantList

	c.Data["pagetitle"] = "创建角色"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmCanCreate)
	c.Data["result"] = reStr

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/setCreate.html"
}

/**
 * 删除道具
 */
func (c *GMToolsController) DelItemGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	//models.AllPlantServerList()
	c.Data["pagetitle"] = "删除道具"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmDeleteItem)
	//c.Data["centerurl"] = models.PlantList

	//{ok, [{result,undefined},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/delItem.html"
}

/**
 * 删除道具
 */
func (c *GMToolsController) DelItemPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	var UserId string = v.(string)

	CmdUrl := c.GetString("cmdurl")
	RoleId := c.GetString("roleid")
	ProtoId := c.GetString("protoid")
	Level := c.GetString("level")
	ItemNum := c.GetString("itemnum")
	Time := models.GetNowNigxStr()
	Auth := models.GetAuth(UserId, Time)
	UrlParam := "zxlf_get?cmd=gm_del_item&roleid=" + RoleId + "&protoid=" + ProtoId + "&level=" + Level + "&itemnum=" + ItemNum + "&auth=" + Auth + "&gm=" + UserId + "&time=" + Time
	AllUrl := CmdUrl + UrlParam

	var reStr string

	result, error := models.HttpGet(AllUrl)

	if error != nil {
		reStr = error.Error()
	} else {
		reStr = string(result)
	}

	//c.Data["centerurl"] = models.PlantList

	//("发送失败！")
	//("删除成功,请确认！")
	//("删除失败,程序错误！")
	//("删除可能失败，请确认！")
	//{ok, [{result,Res},{centerurl,NewAllCenter}]}

	c.Data["pagetitle"] = "删除道具"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmDeleteItem)
	c.Data["result"] = reStr
	//c.Data["centerurl"] = models.PlantList

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "gmtools/delItem.html"
}

func (c *GMToolsController) UserDataPost() {
	v := c.GetSession("loginuser")
	if v == nil {
		//销毁全部的session
		c.DestroySession()
		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)
		c.PageLoginWitchError("身份过期")
		return
	}
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmDictionary)
	c.Data["pagetitle"] = "字典数据查询"

	// todo 這裏需要頁面配合！前個提交搜索頁面
	ServerId, _ := strconv.Atoi(c.GetString("serverid"))
	UserId, _ := strconv.ParseUint(c.GetString("roleid"), 10, 64)

	sessionId := netmsg.NewSession()
	netmsg.SendMsgToGMServer(userData.ThisPlatformId, &network_message.GM_UserDataRequest{
		SessionId: sessionId,
		ServerId:  int32(ServerId),
		UserId:    UserId,
	})
	ret := netmsg.RecMsg(sessionId).(network_message.GM_UserDataResponse)
	if ret.OneData == "" {
		c.PageError("找不到該用戶信息，請檢查")
		return
	}
	// todo 需頁面配合！后續的展示頁面
}
func (c *GMToolsController) UserDataGet() {
	v := c.GetSession("loginuser")
	if v == nil {
		//销毁全部的session
		c.DestroySession()
		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)
		c.PageLoginWitchError("身份过期")
		return
	}
	// todo 這裏需要頁面配合！前個模糊搜索頁面
	ServerId, _ := strconv.Atoi(c.GetString("serverid"))
	UserName := c.GetString("roleid")
	userData := c.GetSession("UserData").(AdminData)
	sessionId := netmsg.NewSession()
	netmsg.SendMsgToGMServer(userData.ThisPlatformId, &network_message.GM_UserDataRequest{
		SessionId: sessionId,
		ServerId:  int32(ServerId),
		UserName:  UserName,
	})
	ret := netmsg.RecMsg(sessionId).(network_message.GM_UserDataResponse)
	if ret.UserData == nil {
		c.PageError("找不到該用戶信息，請檢查")
		return
	}
	// todo 需頁面配合！后續的展示頁面
}
