package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/netmsg"
	"gm_server/src/beegoWeb/table"
	network_message "gm_server/src/network/message"
	"strconv"
	"strings"
)

//-- 玩家信息查询 --

/**
 * 玩家信息查询
 */
type InfoController struct {
	GMBaseController
}

//-- 玩家 --

func (c *InfoController) Index() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderIndex(v, "")
}

type SimpleUsr struct {
	network_message.GM_RetUsers
	SexStr string
}

func (c *InfoController) Ask() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	asktype := c.GetString("asktype")
	askname := c.GetString("askname")
	askid, _ := c.GetInt("askid")

	msg := &network_message.GM_UserDataRequest{}
	if asktype == "I" {
		msg.UserId = uint64(askid)
	} else {
		msg.UserName = askname
	}

	sessionId := netmsg.NewSession()
	userData := c.GetSession("UserData").(AdminData)

	msg.SessionId = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_UserDataResponse)

	//成功了
	if ret.Errorcode == 1 {

		var finalArr []*SimpleUsr

		retListlen := len(ret.UserData)

		for i := 0; i < retListlen; i++ {

			newData := &SimpleUsr{}
			newData.GM_RetUsers = *ret.UserData[i]

			if newData.Sex == 0 {
				newData.SexStr = "男"
			} else {
				newData.SexStr = "女"
			}

			finalArr = append(finalArr, newData)
		}

		c.Data["admin"] = v
		////userData := c.GetSession("UserData").(AdminData)
		c.Data["candoact"] = true
		c.Data["audit_res"] = "查询成功"
		c.Data["pagetitle"] = "玩家信息查询"
		c.Data["usrlist"] = finalArr
		c.Layout = "basetemplate/basetemplate.html"
		c.TplName = "info/index.html"
	} else {

		//for test
		var finalArr []*SimpleUsr
		newData := &SimpleUsr{}
		newData.ServerId = 1
		newData.SexStr = "男"
		newData.RoleId = 232
		newData.UserName = "EVE"
		newData.Gold = 188
		newData.HeroNum = 133
		finalArr = append(finalArr, newData)

		c.Data["admin"] = v
		////userData := c.GetSession("UserData").(AdminData)
		c.Data["candoact"] = true
		c.Data["audit_res"] = "查询成功"
		c.Data["pagetitle"] = "玩家信息查询"
		c.Data["usrlist"] = finalArr
		c.Layout = "basetemplate/basetemplate.html"
		c.TplName = "info/index.html"

		//c.RenderIndex(v, "查询失败 Code=" + strconv.Itoa(int(ret.Errorcode)))
	}
}

/**
 * 踢掉某个玩家 这个用 ajax 解决
 */
func (c *InfoController) Kick() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	usrid, _ := c.GetInt("usrid")
	serverid, _ := c.GetInt("serverid")

	msg := &network_message.GM_KickRoleRequest{}
	msg.UserId = uint64(usrid)
	msg.ServerId = int32(serverid)

	sessionId := netmsg.NewSession()
	userData := c.GetSession("UserData").(AdminData)

	msg.SessionId = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_KickRoleResponse)

	result := int(ret.ResState)

	c.Ctx.WriteString(strconv.Itoa(result))
}

func (c *InfoController) RenderIndex(adminName interface{}, Res string) {
	//models.AllPlantServerList()
	//AuditMarquee = util_tools:check_user_group(SessionId, ?ADMIN_OPERATOR_AUDIT_MARQUEE)
	c.Data["admin"] = adminName
	//userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = true
	//c.Data["centerurl"] = models.PlantList
	c.Data["audit_res"] = Res
	c.Data["pagetitle"] = "玩家信息查询"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/index.html"
}

func (c *InfoController) Search(adminName interface{}, Res string) {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderIndex(v, "")
}

func (c *InfoController) Change(adminName interface{}, Res string) {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderIndex(v, "")
}

/**
 * 玩家界面-禁言/解除禁言 1个用户
 */
func (c *InfoController) MuteOneUsr(adminName interface{}, Res string) {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderIndex(v, "")
}

/**
 * 玩家界面-封号/解除封号 1个用户
 */
func (c *InfoController) SealOneUsr(adminName interface{}, Res string) {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderIndex(v, "")
}

//-- 禁言 --

func (c *InfoController) Mute() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderMute(v, "")
}

func (c *InfoController) RenderMute(adminName interface{}, Res string) {
	//models.AllPlantServerList()
	//AuditMarquee = util_tools:check_user_group(SessionId, ?ADMIN_OPERATOR_AUDIT_MARQUEE)
	c.Data["admin"] = adminName
	////userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = true
	//c.Data["centerurl"] = models.PlantList
	//c.Data["marquee_result"] = c.GetMarqueeRecords(true)
	c.Data["audit_res"] = Res

	//{ok, [{candoact, true},{centerurl,NewAllCenter},{marquee_result, MarqueeRecords},{audit_res, Res}]}

	c.Data["pagetitle"] = "信息查询-禁言"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/mute.html"
}

/**
 * 禁言界面-禁言/解除禁言 1个用户
 */
func (c *InfoController) MuteOneUsr2(adminName interface{}, Res string) {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderMute(v, "")
}

//-- 封号 --

func (c *InfoController) Seal() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderSeal(v, "")
}

func (c *InfoController) RenderSeal(adminName interface{}, Res string) {
	//models.AllPlantServerList()
	//AuditMarquee = util_tools:check_user_group(SessionId, ?ADMIN_OPERATOR_AUDIT_MARQUEE)
	c.Data["admin"] = adminName
	////userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = true
	//c.Data["centerurl"] = models.PlantList
	//c.Data["marquee_result"] = c.GetMarqueeRecords(true)
	c.Data["audit_res"] = Res

	//{ok, [{candoact, true},{centerurl,NewAllCenter},{marquee_result, MarqueeRecords},{audit_res, Res}]}

	c.Data["pagetitle"] = "信息查询-封号"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/seal.html"
}

/**
 * 封号界面-封号/解除封号 1个用户
 */
func (c *InfoController) SealOneUsr2(adminName interface{}, Res string) {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderSeal(v, "")
}

//-- 锁定玩家后的二级跳转 --

func (c *InfoController) Index2() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.Data["admin"] = v
	////userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = true
	//c.Data["centerurl"] = models.PlantList
	c.Data["pagetitle"] = "信息查询-菜单"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/index2.html"
}

//-- 背包 --

type BagItem struct {
	ID        uint64
	Num       int32
	ProtoID   int32
	TimeLimit uint64
	MainType  int32
	Qua       int32
	ItemName  string
}

func (c *InfoController) Bag() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")

	c.RenderBag(v.(string), serverid, usrid)
}

func (c *InfoController) RenderBag(admin string, serverid, usrid int) {

	userData := c.GetSession("UserData").(AdminData)

	msg := &network_message.GM_ItemDataRequest{}
	msg.ServerId = int32(serverid)
	msg.UserId = uint64(usrid)

	sessionId := netmsg.NewSession()
	msg.SessionId = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_ItemDataResponse)

	fmt.Println("查询道具成功")

	var reItemList []*BagItem
	itemslen := len(ret.ItemList)

	for i := 0; i < itemslen; i++ {
		taritem := ret.ItemList[i]
		newMem := &BagItem{}
		newMem.ID = taritem.ItemId
		sData, ok := table.ItemInfo_Get(taritem.ProtoId)
		if ok {
			newMem.ItemName = sData.Name
		}
		newMem.ProtoID = taritem.ProtoId
		newMem.MainType = taritem.Type
		newMem.Num = taritem.Num
		newMem.Qua = taritem.Rare
		newMem.TimeLimit = taritem.TimeLimit

		reItemList = append(reItemList, newMem)
	}

	c.Data["admin"] = admin
	c.Data["candoact"] = true
	c.Data["resultstr"] = "查询成功"
	c.Data["itemlist"] = reItemList
	c.Data["pagetitle"] = "信息查询-背包"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/bag.html"
}

func (c *InfoController) BagDelete() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	userData := c.GetSession("UserData").(AdminData)

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")
	itemid, _ := c.GetInt("itemid")

	msg := &network_message.GM_UpdateItemDataRequest{}
	msg.ServerId = int32(serverid)
	msg.UserId = uint64(usrid)
	msg.ItemId = uint64(itemid)
	msg.Num = 0

	sessionId := netmsg.NewSession()
	msg.SessionId = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_UpdateItemDataResponse)

	if ret.ResState != 1 {
		c.Ctx.WriteString("0")
		return
	}

	fmt.Println("删除道具成功")

	c.RenderBag(v.(string), serverid, usrid)
}

//-- 好友 --

/**
 * 20210804 好友暂时 还没开放
 */
func (c *InfoController) Friend() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")

	c.RenderFriend(v.(string), serverid, usrid)
}

func (c *InfoController) RenderFriend(admin string, serverid, usrid int) {

	c.Data["admin"] = admin
	c.Data["candoact"] = true
	c.Data["pagetitle"] = "信息查询-好友"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/friend.html"
}

//-- 任务 --

func (c *InfoController) Task() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")

	c.RenderTask(v.(string), serverid, usrid)
}

type QuestItem struct {
	ID        int32
	MainType  int32
	Name      string
	Condition string
	State     int32
	StateStr  string
}

func (c *InfoController) RenderTask(admin string, serverid, usrid int) {
	userData := c.GetSession("UserData").(AdminData)

	msg := &network_message.GM_QuestDataRequest{}
	msg.ServerId = int32(serverid)
	msg.UserId = uint64(usrid)

	sessionId := netmsg.NewSession()
	msg.SessionId = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_QuestDataResponse)

	fmt.Println("查询道具成功")

	var reList []*QuestItem
	slen := len(ret.QuestList)

	for i := 0; i < slen; i++ {
		taritem := ret.QuestList[i]
		newMem := &QuestItem{}
		newMem.ID = taritem.Id
		newMem.Name = taritem.Name
		newMem.MainType = taritem.Type
		newMem.Condition = taritem.Condition
		newMem.State = taritem.State
		if newMem.State == 0 {
			newMem.StateStr = "未完成"
		} else if newMem.State == 1 {
			newMem.StateStr = "进行中"
		} else {
			newMem.StateStr = "已完成"
		}

		reList = append(reList, newMem)
	}
	c.Data["questlist"] = reList
	c.Data["admin"] = admin
	c.Data["candoact"] = true
	c.Data["resultstr"] = "查询成功"
	c.Data["pagetitle"] = "信息查询-任务"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/task.html"
}

func (c *InfoController) TaskDelete() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	userData := c.GetSession("UserData").(AdminData)

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")
	itemid, _ := c.GetInt("itemid")

	//TODO 目前没这个接口
	msg := &network_message.GM_UpdateItemDataRequest{}
	msg.ServerId = int32(serverid)
	msg.UserId = uint64(usrid)
	msg.ItemId = uint64(itemid)
	msg.Num = 0

	sessionId := netmsg.NewSession()
	msg.SessionId = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_UpdateItemDataResponse)

	if ret.ResState != 1 {
		c.Ctx.WriteString("0")
		return
	}

	fmt.Println("删除道具成功")

	c.RenderTask(v.(string), serverid, usrid)
}

//-- 邮件 --

func (c *InfoController) Mail() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")

	c.RenderMail(v.(string), serverid, usrid)
}

func (c *InfoController) MailDelete() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	mailid, _ := c.GetInt("mailid")
	usrid, _ := c.GetInt("usrid")
	serverid, _ := c.GetInt("serverid")

	msg := &network_message.GM_DelMailRequest{}
	msg.RoleID = uint64(usrid)
	msg.MailID = uint64(mailid)

	sessionId := netmsg.NewSession()
	msg.SessionID = sessionId

	userData := c.GetSession("UserData").(AdminData)
	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_DelMailResponse)

	fmt.Println("发送删除邮件成功")

	if ret.ErrorCode != network_message.Result_SUCCESS {
		c.Ctx.WriteString("删除邮件失败")
		return
	}

	c.RenderMail(v.(string), serverid, usrid)
}

type MailItem struct {
	*network_message.MailInfo
}

func (c *InfoController) RenderMail(admin string, serverid, usrid int) {
	userData := c.GetSession("UserData").(AdminData)

	msg := &network_message.GM_GetAllMailRequest{}
	msg.RoleID = uint64(usrid)

	sessionId := netmsg.NewSession()
	msg.SessionID = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_GetAllMailResponse)

	fmt.Println("查询邮件列表成功")

	if ret.ErrorCode != network_message.Result_SUCCESS {
		c.Data["admin"] = admin
		c.Data["candoact"] = true
		c.Data["resultstr"] = "查询失败"

		c.Data["pagetitle"] = "信息查询-邮件"
		c.Layout = "basetemplate/basetemplate.html"
		c.TplName = "info/mail.html"
		return
	}

	var reList []*MailItem
	slen := len(ret.MailInfo)

	for i := 0; i < slen; i++ {
		taritem := ret.MailInfo[i]
		newMem := &MailItem{}
		newMem.MailInfo = taritem

		reList = append(reList, newMem)
	}

	c.Data["questlist"] = reList
	c.Data["admin"] = admin
	c.Data["candoact"] = true
	c.Data["resultstr"] = "查询成功"

	c.Data["pagetitle"] = "信息查询-邮件"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/mail.html"
}

//-- 强化 --

func (c *InfoController) Strength() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")

	c.RenderStrength(v.(string), serverid, usrid)
}

func (c *InfoController) StrengthUpdate() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")

	c.RenderStrength(v.(string), serverid, usrid)
}

type HeroDetail struct {
	HeroId   uint64
	ProtoId  int32
	Level    int32
	Rank     int32
	Star     int32
	HeroName string

	SkillP string
	Skill1 string
	Skill2 string
	SkillS string

	EquipLv  string
	EquipQua string
}

type HeroUnGet struct {
	ProtoId  int32
	HeroName string
	NumStr   string
	Num      int
	Max      int
}

func (c *InfoController) RenderStrength(admin string, serverid, usrid int) {
	userData := c.GetSession("UserData").(AdminData)

	msg := &network_message.GM_StrengthenRequest{}
	msg.ServerId = int32(serverid)
	msg.UserId = uint64(usrid)

	sessionId := netmsg.NewSession()
	msg.SessionId = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_StrengthenResponse)

	fmt.Println("查询已有英雄成功")

	var reList []*HeroDetail
	slen := len(ret.HeroData)

	for i := 0; i < slen; i++ {
		taritem := ret.HeroData[i]
		newMem := &HeroDetail{}

		newMem.HeroId = taritem.HeroId
		newMem.ProtoId = taritem.ProtoId
		newMem.Level = taritem.Level
		newMem.Rank = taritem.Rank
		newMem.Star = taritem.Star

		heroSDdata, ok := table.Hero_information_Get(newMem.ProtoId)
		if ok {
			newMem.HeroName = heroSDdata.Name
		}

		newMem.SkillP = "1/9"
		skill1arr := []string{"1/9", "1/9", "1/9"}
		skill2arr := []string{"1/9", "1/9", "1/9"}
		skillsarr := []string{"1/9", "1/9", "1/9"}

		for _, mem := range taritem.Skills {

			if mem.SlotId > 100 {
				fmt.Println("heroskill soltid=", mem.SlotId, "pass!")
				continue
			}

			if mem.SlotId == 1 {
				newMem.SkillP = strconv.Itoa(int(mem.Level)) + "/9"
				continue
			}

			if mem.SlotId <= 4 {
				skill1arr[mem.SlotId-2] = strconv.Itoa(int(mem.Level)) + "/9"
				continue
			}

			if mem.SlotId <= 7 {
				skill2arr[mem.SlotId-5] = strconv.Itoa(int(mem.Level)) + "/9"
				continue
			}

			skillsarr[mem.SlotId-8] = strconv.Itoa(int(mem.Level)) + "/9"
		}

		newMem.Skill1 = strings.Join(skill1arr, ",")
		newMem.Skill2 = strings.Join(skill2arr, ",")
		newMem.SkillS = strings.Join(skillsarr, ",")

		equip1arr := []string{"0", "0", "0"} //等级
		equip2arr := []string{"0", "0", "0"} //品质

		for _, equip := range taritem.Equips {
			if equip.SlotId == 1 {
				equip1arr[0] = strconv.Itoa(int(equip.Level))
				equip2arr[0] = strconv.Itoa(int(equip.Rare))
				continue
			}

			if equip.SlotId == 2 {
				equip1arr[1] = strconv.Itoa(int(equip.Level))
				equip2arr[1] = strconv.Itoa(int(equip.Rare))
				continue
			}

			if equip.SlotId == 3 {
				equip1arr[2] = strconv.Itoa(int(equip.Level))
				equip2arr[2] = strconv.Itoa(int(equip.Rare))
				continue
			}
		}

		newMem.EquipLv = strings.Join(equip1arr, ",")
		newMem.EquipQua = strings.Join(equip2arr, ",")

		reList = append(reList, newMem)
	}

	c.Data["herolist"] = reList
	c.Data["admin"] = admin
	c.Data["candoact"] = true
	c.Data["resultstr"] = "查询成功"

	c.Data["pagetitle"] = "信息查询-强化"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/strength.html"
}

//-- 英雄 --

func (c *InfoController) Hero() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	serverid, _ := c.GetInt("serverid")
	usrid, _ := c.GetInt("usrid")

	c.RenderHero(v.(string), serverid, usrid)
}

func (c *InfoController) RenderHero(admin string, serverid, usrid int) {

	userData := c.GetSession("UserData").(AdminData)

	msg := &network_message.GM_StrengthenRequest{}
	msg.ServerId = int32(serverid)
	msg.UserId = uint64(usrid)

	sessionId := netmsg.NewSession()
	msg.SessionId = sessionId

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_StrengthenResponse)

	fmt.Println("查询已有英雄成功")

	var reHeroList []*HeroDetail
	slen := len(ret.HeroData)

	for i := 0; i < slen; i++ {
		taritem := ret.HeroData[i]
		newMem := &HeroDetail{}

		newMem.HeroId = taritem.HeroId
		newMem.ProtoId = taritem.ProtoId
		newMem.Level = taritem.Level
		newMem.Rank = taritem.Rank
		newMem.Star = taritem.Star

		heroSDdata, ok := table.Hero_information_Get(newMem.ProtoId)
		if ok {
			newMem.HeroName = heroSDdata.Name
		}

		newMem.SkillP = "1/9"
		skill1arr := []string{"1/9", "1/9", "1/9"}
		skill2arr := []string{"1/9", "1/9", "1/9"}
		skillsarr := []string{"1/9", "1/9", "1/9"}

		for _, mem := range taritem.Skills {

			if mem.SlotId > 100 {
				fmt.Println("heroskill soltid=", mem.SlotId, "pass!")
				continue
			}

			if mem.SlotId == 1 {
				newMem.SkillP = strconv.Itoa(int(mem.Level)) + "/9"
				continue
			}

			if mem.SlotId <= 4 {
				skill1arr[mem.SlotId-2] = strconv.Itoa(int(mem.Level)) + "/9"
				continue
			}

			if mem.SlotId <= 7 {
				skill2arr[mem.SlotId-5] = strconv.Itoa(int(mem.Level)) + "/9"
				continue
			}

			skillsarr[mem.SlotId-8] = strconv.Itoa(int(mem.Level)) + "/9"
		}

		newMem.Skill1 = strings.Join(skill1arr, ",")
		newMem.Skill2 = strings.Join(skill2arr, ",")
		newMem.SkillS = strings.Join(skillsarr, ",")

		equip1arr := []string{"0", "0", "0"} //等级
		equip2arr := []string{"0", "0", "0"} //品质

		for _, equip := range taritem.Equips {
			if equip.SlotId == 1 {
				equip1arr[0] = strconv.Itoa(int(equip.Level))
				equip2arr[0] = strconv.Itoa(int(equip.Rare))
				continue
			}

			if equip.SlotId == 2 {
				equip1arr[1] = strconv.Itoa(int(equip.Level))
				equip2arr[1] = strconv.Itoa(int(equip.Rare))
				continue
			}

			if equip.SlotId == 3 {
				equip1arr[2] = strconv.Itoa(int(equip.Level))
				equip2arr[2] = strconv.Itoa(int(equip.Rare))
				continue
			}
		}

		newMem.EquipLv = strings.Join(equip1arr, ",")
		newMem.EquipQua = strings.Join(equip2arr, ",")

		reHeroList = append(reHeroList, newMem)
	}

	msg2 := &network_message.GM_HeroProgressRequest{}
	msg2.ServerId = int32(serverid)
	msg2.UserId = uint64(usrid)

	sessionId2 := netmsg.NewSession()
	msg2.SessionId = sessionId2

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg2)
	ret2 := netmsg.RecMsg(sessionId2).(network_message.GM_HeroProgressResponse)

	fmt.Println("查询未获得英雄列表成功")

	var unList []*HeroUnGet
	slen2 := len(ret2.List)

	for i := 0; i < slen2; i++ {
		taritem := ret2.List[i]
		newMem := &HeroUnGet{}
		newMem.ProtoId = taritem.ProtoId
		heroSDdata, ok := table.Hero_information_Get(newMem.ProtoId)
		if ok {
			newMem.HeroName = heroSDdata.Name
		}
		newMem.Num = int(taritem.Num)
		newMem.Max = int(taritem.Max)
		newMem.NumStr = strconv.Itoa(newMem.Num) + "/" + strconv.Itoa(newMem.Max)

		unList = append(unList, newMem)
	}

	c.Data["herolist"] = reHeroList
	c.Data["unherolist"] = unList
	c.Data["admin"] = admin
	c.Data["candoact"] = true
	c.Data["resultstr"] = "查询成功"

	c.Data["pagetitle"] = "信息查询-英雄"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "info/hero.html"
}

//-- ends --
