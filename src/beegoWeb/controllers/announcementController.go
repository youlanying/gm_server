package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/models"
	"strconv"
)

//---- 公告 玩家登陆进入游戏后 在首页弹出的公告 ----

/**
 * 显示 公告界面
 */
func (c *OperatorController) ShowAnnouncement() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderAnnouncement(v, "")
}

/**
 * 渲染 公告界面
 */
func (c *OperatorController) RenderAnnouncement(adminName interface{}, Res string) {
	c.Data["pagetitle"] = "发送公告"
	c.Data["admin"] = adminName
	//userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = 1
	//c.Data["centerurl"] = models.PlantList
	c.Data["noticelist"] = c.GetNoticeDBList(true)
	c.Data["audit_res"] = Res

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/send_announcement.html"
}

/**
 * 核心逻辑 获取邮件列表
 * @param Limit 代表session是否有权限操作 邮件
 */
func (c *OperatorController) GetNoticeDBList(Limit bool) []models.PublicNoticeMem {

	var reList []models.PublicNoticeMem

	vvList, OK := beegodb.TB_PUBLICNOTICEReadBySQL("")

	if !OK {
		return reList
	}

	for _, vmem := range vvList {

		var newMem models.PublicNoticeMem

		//旧属性重新赋值
		newMem.TB_PUBLICNOTICE = *vmem

		//if vmem.STATUS == 0 {
		//	newMem.Mclass = "btn btn-danger" //Btn_danger
		//	newMem.Disabled = ""
		//	newMem.Style = ""
		//	newMem.Sendok = "待审核"
		//} else {
		newMem.Style = "display:none" //Btn_danger
		newMem.Disabled = "disabled='disabled'"
		newMem.Mclass = "btn btn-success"
		newMem.Sendok = "已发送"
		//}

		reList = append(reList, newMem)
	}

	return reList
}

/**
 * 发送公告 Post
 */
func (c *OperatorController) SendAnnouncement() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//CmdUrl := c.GetString("cmdurl")
	//level, ok := c.GetInt("level")
	//pagelable := c.GetString("pagelable")
	//flag := c.GetString("flag")

	//t1 := c.GetString("t1")
	//t2 := c.GetString("t2")

	Title1 := c.GetString("title")
	titleshort := c.GetString("titleshort")
	Content := c.GetString("content")

	//timeNow := models.GetNowStr()

	var mem beegodb.TB_PUBLICNOTICE
	mem.TITLE = Title1
	mem.TITLESHORT = titleshort
	mem.CONTENT = Content

	//mem.PAGETYPE = pagelable
	//mem.FLAG = flag
	//mem.LEVEL = int32(level)

	//mem.CREATETIME = timeNow
	//mem.STATUS = 0
	//mem.SERVERURL = CmdUrl
	mem.SENDID = v.(string)
	mem.AUDITID = ""
	//mem.STARTTIME = t1
	//mem.ENDTIME = t2

	id, OK := beegodb.TB_PUBLICNOTICE_Insert(&mem)
	if OK {
		fmt.Println("insert succ:", id)
	} else {
		fmt.Println("insert fail:", id)
	}

	c.RenderAnnouncement(v, "新增公告成功")
}

/**
 * 这个实际是 点击 待审核后 发送审核公告
 */
func (c *OperatorController) AuditAnnouncement() {

	fmt.Println("Audit_mail ING")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//审核公告ID
	vID, _ := c.GetInt("ID")

	var sqlstr = "where id=" + strconv.Itoa(vID)

	maillist, ok := beegodb.TB_PUBLICNOTICEReadBySQL(sqlstr)

	var ReStr string

	if !ok || len(maillist) == 0 {
		ReStr = "审核失败！公告已消失,请重新添加！"
	} else { //else 不能出现在下一行 草草草

		//if maillist[0].STATUS == 0 {
		//	fmt.Println("发送申请结果检测")
		//}

		//TODO
	}

	c.RenderAnnouncement(v, ReStr)
}

/**
 * 删除公告
 */
func (c *OperatorController) DeleteAnnouncement() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Did, _ := c.GetInt("ID")
	sqlstr := "id=" + strconv.Itoa(Did)
	beegodb.TB_PUBLICNOTICEdeleteBy(sqlstr)

	c.RenderAnnouncement(v, "删除待审核公告成功")
}
