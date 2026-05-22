package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/models"
	"gm_server/src/beegoWeb/netmsg"
	network_message "gm_server/src/network/message"
	"strconv"
)

type NoticeController struct {
	GMBaseController
}

func (c *NoticeController) Index() {

	fmt.Println("Notice Index")

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
	noticeList := getPlatformNotice(userData.ThisPlatformId)
	jsonNotice := models.ToJson(noticeList)

	c.Data["pagetitle"] = "公告管理"
	c.Data["jsonNotice"] = jsonNotice

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "notice/index.html"
}

func (c *NoticeController) Add() {

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
	noticetype, _ := c.GetInt("noticetype") //类型
	label, _ := c.GetInt("label")           //标签
	priority, _ := c.GetInt("priority")     //优先级
	titleshort := c.GetString("titleshort") //短标题
	title := c.GetString("title")           //标题
	content := c.GetString("content")       //内容
	t1 := c.GetString("t1")                 //开始时间
	t2 := c.GetString("t2")                 //结束时间
	//serverid:0,noticetype:1,label:1,priority:1,titleshort:fghsdghdfg,title:fgsdfgsdf,content:dfgdfgsdfg,t1:2021/08/11 19:36,t2:2021/08/11 19:36
	fmt.Printf("==Add====serverid:%v,noticetype:%v,label:%v,priority:%v,titleshort:%v,title:%v,content:%v,t1:%v,t2:%v\n", serverid, noticetype, label, priority, titleshort, title, content, t1, t2)
	starttime := models.StrToUnixTime(t1 + ":00")
	endtime := models.StrToUnixTime(t2 + ":00")
	createtime := models.GetNow()
	userData := c.GetSession("UserData").(AdminData)
	_, ok := beegodb.TB_PUBLICNOTICEInsertAuto(userData.ThisPlatformId, int32(serverid), int32(noticetype), int32(label), int32(priority), titleshort, title, content, starttime, endtime, createtime, 0, userData.UserId, "")
	if ok {
		fmt.Println("insert succ")
		noticeList := getPlatformNotice(userData.ThisPlatformId)
		jsonNotice := models.ToJson(noticeList)
		c.Ctx.WriteString(jsonNotice)
	} else {
		c.Ctx.WriteString("")
	}
	c.Ctx.WriteString("")
}

func (c *NoticeController) Del() {

	fmt.Println("Notice del")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	NoticeId := c.GetString("noticeid")
	fmt.Println("del  id", NoticeId)
	_, ok := beegodb.TB_PUBLICNOTICEdeleteBy("id=" + NoticeId)
	if ok {
		userData := c.GetSession("UserData").(AdminData)
		noticeList := getPlatformNotice(userData.ThisPlatformId)
		jsonNotice := models.ToJson(noticeList)
		c.Ctx.WriteString(jsonNotice)
		return
	}
	c.Ctx.WriteString("")
}

func (c *NoticeController) Audit() {
	fmt.Println("Notice Audit")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	strNoticeId := c.GetString("noticeid")
	NoticeId, _ := c.GetInt32("noticeid")
	fmt.Println("Notice Audit NoticeId:", strNoticeId)
	userData := c.GetSession("UserData").(AdminData)
	noticeMap, dbOk := beegodb.TB_PUBLICNOTICEReadByid("id=" + strNoticeId)
	if dbOk {
		sessionId := netmsg.NewSession()
		noticeData := noticeMap[NoticeId]
		notice := &network_message.NoticeInfo{
			Serverid:   noticeData.SERVERID,
			NoticeType: noticeData.NOTICETYPE,
			Label:      noticeData.LABEL,
			Priority:   noticeData.PRIORITY,
			MiniTitle:  noticeData.TITLESHORT,
			Title:      noticeData.TITLE,
			Notice:     noticeData.CONTENT,
			StartTime:  noticeData.STARTTIME,
			EndTime:    noticeData.ENDTIME,
		}
		msg := &network_message.GM_AddNoticeRequest{SessionId: sessionId, ConnectId: 0, Notice: notice}
		ok := netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
		if ok {
			ret := netmsg.RecMsg(sessionId).(network_message.GM_AddNoticeResponse)
			if ret.IsSend {
				audittime := models.GetNow()
				_, ok = beegodb.TB_PUBLICNOTICEUpdateByKey("id="+strNoticeId, "audittime="+strconv.FormatInt(audittime, 10)+",auditid='"+userData.UserId+"'")
				if ok {
					noticeList := getPlatformNotice(userData.ThisPlatformId)
					jsonNotice := models.ToJson(noticeList)
					c.Ctx.WriteString(jsonNotice)
					return
				}
			}
		}
	}
	c.Ctx.WriteString("")
}

//查询platformId最新的100条
func getPlatformNotice(platformId int32) []*beegodb.TB_PUBLICNOTICE {
	fmt.Println("=======getPlatformNotice====", platformId)
	mapNotice, _ := beegodb.TB_PUBLICNOTICEReadBySQL("WHERE platformid=" + strconv.FormatInt(int64(platformId), 10) + " ORDER BY `id` DESC LIMIT 0,100")
	fmt.Println("=======getPlatformNotice====len:", len(mapNotice))
	noticeList := make([]*beegodb.TB_PUBLICNOTICE, 0)
	for _, notice := range mapNotice {
		noticeList = append(noticeList, notice)
	}
	return noticeList
}
