package controllers

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/models"
	"gm_server/src/beegoWeb/netmsg"
	network_message "gm_server/src/network/message"
	"log"
	"strconv"
	"strings"
	"time"
)

//---- 邮件 ----

/**
 * 常规显示邮件列表
 */
func (c *OperatorController) ShowMail() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderMail(v, "")
}

/**
 * 渲染邮件列表
 */
func (c *OperatorController) RenderMail(adminName interface{}, Res string) {
	//CanDoAction = util_tools:check_user_group(SessionId, ?ADMIN_OPERATOR_SEND_MAIL),
	//models.AllPlantServerList()
	c.Data["pagetitle"] = "发送邮件"
	c.Data["admin"] = adminName
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionOperatorSendMail)
	//c.Data["centerurl"] = models.PlantList
	c.Data["mailresult"] = c.GetMailRecords(true)
	c.Data["audit_res"] = Res

	//{ok, [{candoact, true},{centerurl,NewAllCenter},{mailresult, MailRecords},{audit_res, Res}]}
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/send_mail.html"
}

/**
 * 核心逻辑 获取邮件列表
 * @param Limit 代表session是否有权限操作 邮件
 */
func (c *OperatorController) GetMailRecords(Limit bool) []models.MailMem {

	//page := models.GetMailsWithPage(1, 10)
	var reList []models.MailMem

	vvList, ok := beegodb.TB_SEND_MAILReadBySQL("")
	if !ok {
		return reList
	}

	for _, vmem := range vvList {

		newMem := models.MailMem{}
		newMem.TB_SEND_MAIL = *vmem

		var Stb string
		var StrTime string
		var Disabled0 string

		if vmem.ISALL == 1 {
			newMem.Recvname = "全服"
		}

		if vmem.STATUS == 0 {
			Stb = "C/"
			newMem.Mclass = "btn btn-danger" //Btn_danger
			Disabled0 = ""
			newMem.Style = ""
			newMem.Sendok = "待审核"
		} else {
			Stb = "F/"
			newMem.Style = "display:none" //Btn_danger
			Disabled0 = "disabled='disabled'"
			newMem.Mclass = "btn btn-success"
			newMem.Sendok = "已发送"
		}

		if Limit == true {
			newMem.Disabled = Disabled0
		} else {
			newMem.Disabled = "disabled='disabled'"
		}

		if len(vmem.CREATETIME) == 0 {
			StrTime = "NULL"
		} else {
			StrTime = models.GetNowStr()
		}

		vmem.CREATETIME = Stb + StrTime
		reList = append(reList, newMem)
	}

	return reList
}

/**
 * 发送邮件 Post
 */
func (c *OperatorController) Send_mailPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	title1 := c.GetString("title1")
	content := c.GetString("content")
	items := c.GetString("items")
	mailtype := c.GetString("mailtype")
	Roleids := c.GetString("roleids")
	cmdurl := c.GetString("cmdurl")
	sendname := c.GetString("sendname")
	lvstart, _ := c.GetInt("lvstart")
	lvend, _ := c.GetInt("lvend")
	sextype := c.GetString("sextype")
	timetype := c.GetString("timetype")
	t1 := c.GetString("t1")
	t2 := c.GetString("t2")
	t3 := c.GetString("t3")
	t4 := c.GetString("t4")

	mem := &beegodb.TB_SEND_MAIL{}
	mem.TITLE = title1
	mem.SENDID = sendname
	mem.CONTENT = content

	//全体
	if mailtype == "Q" {
		mem.RECVNAME = "ALL"

		mem.LVSTART = int32(lvstart)
		mem.LVEND = int32(lvend)

		if sextype == "M" {
			mem.SEX = 1
		} else if sextype == "W" {
			mem.SEX = 2
		} else {
			mem.SEX = 0
		}
	} else {
		mem.RECVNAME = Roleids
	}

	mem.ITEMLIST = items
	mem.CREATETIME = models.GetNowStr()
	mem.STATUS = 0

	mem.SERVERURL = cmdurl
	mem.SENDID = v.(string)
	mem.AUDITID = "Null"

	if timetype == "V2" {
		mem.TIMETYPE = 1
		mem.STARTTIME = t1
		mem.ENDTIME = t2
	} else if timetype == "V3" {
		mem.TIMETYPE = 2
		mem.STARTTIME = t3
		mem.ENDTIME = t4
	} else {
		mem.TIMETYPE = 0
	}

	result, ok := beegodb.TB_SEND_MAIL_Insert(mem)

	if ok {
		fmt.Println("insert succ:", result)
		c.RenderMail(v, "新增邮件成功")
	} else {
		fmt.Println("insert fail:", result)
		c.RenderMail(v, "新增邮件失败")
	}
}

/**
 * 这个实际是 点击 待审核后 发送审核邮件
 */
func (c *OperatorController) Audit_mail() {

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

	//审核邮件ID
	Mailid, _ := c.GetInt("mailid")

	var ReStr string

	sqlStr := "where mailid = " + strconv.Itoa(Mailid)

	vList, ok := beegodb.TB_SEND_MAILReadBySQL(sqlStr)

	if !ok || len(vList) == 0 {
		ReStr = "审核失败！邮件已消失,请重新添加！"
		c.RenderMail(v, ReStr)
		return
	}

	targetMail := vList[0]

	if targetMail == nil {
		ReStr = "审核失败！邮件已消失,请重新添加！"
		c.RenderMail(v, ReStr)
		return
	}

	if targetMail.STATUS == 1 {
		ReStr = "审核失败！邮件已经审核过了！"
		c.RenderMail(v, ReStr)
		return
	}

	fmt.Println("发送申请结果检测")

	sessionId := netmsg.NewSession()
	userData := c.GetSession("UserData").(AdminData)

	msg := &network_message.GM_SendRoleMailRequest{}
	msg.Title = targetMail.TITLE
	msg.GmSender = targetMail.SENDID
	msg.GmEditor = v.(string)
	msg.SendName = targetMail.SENDID
	msg.IsSendAll = targetMail.ISALL == 1

	if !msg.IsSendAll {

		strArr := strings.Split(targetMail.RECVNAME, ",")
		strLen := len(strArr)
		for i := 0; i < strLen; i++ {
			reID, _ := strconv.Atoi(strArr[i])
			msg.ToPlayers = append(msg.ToPlayers, uint64(reID))
		}
	}

	msg.SexType = targetMail.SEX
	msg.Title = targetMail.TITLE
	msg.ContentStr = targetMail.CONTENT

	strArr := strings.Split(targetMail.ITEMLIST, ";")
	strarrlen := len(strArr)
	for i := 0; i < strarrlen; i++ {
		memStr := strArr[i]
		memStr1 := memStr[1 : len(memStr)-1]
		memArr := strings.Split(memStr1, ",")

		num1, _ := strconv.Atoi(memArr[0])
		num2, _ := strconv.Atoi(memArr[1])
		num3, _ := strconv.Atoi(memArr[2])

		newMem := &network_message.Items{}
		newMem.ItemProtoid = int32(num1)
		newMem.ItemNum = int32(num2)
		newMem.ItemNum = int32(num3)

		msg.ItemList = append(msg.ItemList, newMem)
	}

	msg.TimeType = int32(targetMail.TIMETYPE)
	t1 := models.StringToTime(targetMail.STARTTIME)
	t2 := t1.UnixNano() / 1e6
	msg.StartTime = int64(t2)
	t3 := models.StringToTime(targetMail.ENDTIME)
	t4 := t3.UnixNano() / 1e6
	msg.EndTime = int64(t4)
	t5 := time.Now().UnixNano() / 1e6
	msg.SendTime = int64(t5)

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_SendRoleMailResponse)

	//成功了
	if ret.ErrorCode == 0 {

		faillen := len(ret.FailPlayers)

		if faillen == 0 {
			targetMail.STATUS = 1
			targetMail.AUDITID = v.(string)
			_, ok := beegodb.TB_SEND_MAILUpdateBy(targetMail)
			if ok {
				fmt.Println("审核更新成功")
			} else {
				fmt.Println("审核更新报错")
			}

			ReStr = "恭喜，审核发送成功！"
		} else {

			var failStr string = ""

			for i := 0; i < faillen; i++ {
				failStr += strconv.Itoa(int(ret.FailPlayers[i]))
			}

			targetMail.STATUS = 1
			targetMail.AUDITID = v.(string)
			targetMail.CREATETIME = models.GetNowStr()
			_, ok := beegodb.TB_SEND_MAILUpdateBy(targetMail)
			if ok {
				fmt.Println("审核更新成功")
				ReStr = "恭喜，审核发送成功！"
			} else {
				fmt.Println("审核更新报错")
				ReStr = "审核成功,但部分玩家未收到,请检查邮件信息后再发送！未发送成功ID:" + failStr
			}
		}
	} else {
		ReStr = "审核失败,请检查邮件信息！"
	}

	c.RenderMail(v, ReStr)
}

type MailShenHe struct {
	Roleid string
	Bool   string
}

/**
 * 删除邮件
 */
func (c *OperatorController) DeleteMail() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Did, _ := c.GetInt("id")
	sqlStr := "mailid=" + strconv.Itoa(Did)
	beegodb.TB_SEND_MAILdeleteBy(sqlStr)
	c.RenderMail(v, "删除待审核邮件成功")
}

/**
 * 快速批量创建 邮件发送请求
 */
func (c *OperatorController) QuickCreateMails() {

	fmt.Println("QuickCreateMails")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	// 读取文件信息
	f, h, err := c.GetFile("upfile")

	defer f.Close()

	if err != nil {
		log.Fatal("读取文件错误", err)
	}

	//timeUnix := time.Now().Format("2006-01-02 15:04:05")
	//fileNameArr := strings.Split(h.Filename,".")
	//fileNameArr[0] = fileNameArr[0] + timeUnix;
	//newFileName := strings.Join(fileNameArr,".")

	newFileName := h.Filename

	result := c.SaveToFile("upfile", "./"+newFileName)

	if result != nil {
		fmt.Println("SaveError", result.Error())
		return
	}
	xlFile, err := excelize.OpenFile("./" + newFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}

	var finalArr []*QuickMailStruct

	//通常第一个表是 有意义的表
	for _, mem := range xlFile.GetSheetList() {
		if mem == "" {
			continue
		}
		vArr := ReadOneTable(xlFile, mem)
		finalArr = append(finalArr, vArr...)
	}

	var sqlArr []beegodb.TB_SEND_MAIL

	arrlen := len(finalArr)

	for i := 0; i < arrlen; i++ {
		memData := finalArr[i]

		var mem beegodb.TB_SEND_MAIL
		mem.TITLE = memData.Title
		mem.SENDID = memData.SenderName

		var UserIDStr string = ""
		idlens := len(memData.UserID)
		for n := 0; n < idlens; n++ {

			if n != idlens-1 {
				UserIDStr += strconv.Itoa(memData.UserID[n])
			} else {
				UserIDStr += strconv.Itoa(memData.UserID[n]) + ","
			}
		}

		mem.RECVNAME = "[" + UserIDStr + "]"
		mem.CONTENT = memData.Content

		mem.ITEMLIST = memData.ItemListStr

		timeNow := models.GetNowStr()
		mem.CREATETIME = timeNow

		mem.STATUS = 0
		mem.SERVERURL = memData.ServerName
		mem.SENDID = v.(string)
		mem.AUDITID = "Null"

		sqlArr = append(sqlArr, mem)
	}

	////插入多行
	//num, err := orm.InsertMulti(len(sqlArr), sqlArr)
	//if err == nil {
	//	fmt.Println("insert more succ:", num)
	//} else {
	//	fmt.Println("insert more fail:", err)
	//}

	c.RenderMail(v, "快速批量发送邮件成功")
}

type QuickMailReward struct {
	ItemID  int
	ItemNum int
	ItemLv  int
}

type QuickMailStruct struct {
	ServerID    int
	ChannelName string
	ServerName  string
	UserID      []int
	UserName    []string
	Reward      []*QuickMailReward
	SenderName  string
	Title       string
	Content     string
	Reason      string
	TimeType    int
	TimeStart   int
	TimeEnd     int

	CompareRow  []string
	ItemListStr string
}

func ReadOneTable(xlFile *excelize.File, sheetName string) []*QuickMailStruct {

	fmt.Println("Read TableSheet ", sheetName)

	ReArr := make([]*QuickMailStruct, 0)

	var j int
	// 获取 Sheet 上所有单元格
	rows, err := xlFile.GetRows(sheetName)
	if err != nil {
		fmt.Printf("=====ReadOneTable===err:%v\n", err)
		return ReArr
	}
	if len(rows) > 2 {
		fmt.Printf("=====ReadOneTable===rows:0\n")
		return ReArr
	}
	//第二行起是 数据
	for i := 2; i < len(rows); i++ {

		//先读行
		row := rows[i]

		thenReArrLen := len(ReArr)
		for j = 0; j < thenReArrLen; j++ {
			//是相同的 做插入
			if CompareRow(ReArr[j].CompareRow, row) {
				text4 := row[3]
				text5 := row[4]
				vUserID, _ := strconv.Atoi(text4)
				InsertQuickMailStruct(ReArr[j], vUserID, text5)
				break
			}
		}

		//没有新增 那么新增
		if j == thenReArrLen {
			newMem := BuildNewQuickMailStruct(row)
			ReArr = append(ReArr, newMem)
		}
	}

	return ReArr
}

func CompareRow(a, b []string) bool {
	//前三个是 服务器
	if a[0] != b[0] {
		return false
	}

	if a[1] != b[1] {
		return false
	}

	if a[2] != b[2] {
		return false
	}

	//3 4是角色ID和名字 可以不一样

	//补偿列表 邮件标题 发件人名字
	if a[5] != b[5] {
		return false
	}
	if a[6] != b[0] {
		return false
	}
	if a[7] != b[7] {
		return false
	}

	//补偿原因 和 邮件内容 太大了 不比较

	//补偿类型
	if a[10] != b[10] {
		return false
	}
	if a[11] != b[11] {
		return false
	}
	if a[12] != b[12] {
		return false
	}

	return true
}

func BuildNewQuickMailStruct(row []string) *QuickMailStruct {
	mailAsk := &QuickMailStruct{}

	text1 := row[0]
	text2 := row[1]
	text3 := row[2]
	text4 := row[3]
	text5 := row[4]
	text6 := row[5]
	text7 := row[6]
	text8 := row[7]
	text9 := row[8]
	text10 := row[9]
	text11 := row[10]
	text12 := row[11]
	text13 := row[12]

	mailAsk.ServerID, _ = strconv.Atoi(text1)
	mailAsk.ChannelName = text2
	mailAsk.ServerName = text3
	vUserID, _ := strconv.Atoi(text4)
	mailAsk.UserID = append(mailAsk.UserID, vUserID)
	mailAsk.UserName = append(mailAsk.UserName, text5)

	mailAsk.ItemListStr = text6
	strlen := len(mailAsk.ItemListStr)
	//去掉[]
	reStr := mailAsk.ItemListStr[1 : strlen-1]
	//一维数组 去掉;
	strArr := strings.Split(reStr, ";")
	vLen := len(strArr)
	var strMem string
	var strMem1 string
	var memLen int

	for j := 0; j < vLen; j++ {

		strMem = strArr[j]

		//去掉{}
		strMem1 = strMem[1 : len(strMem)-1]

		//去掉,
		strArr2 := strings.Split(strMem1, ",")

		memReward := &QuickMailReward{}

		memLen = len(strArr2)

		val1, _ := strconv.Atoi(strArr2[0])
		val2, _ := strconv.Atoi(strArr2[1])
		memReward.ItemID = val1
		memReward.ItemNum = val2

		if memLen == 3 {
			val3, _ := strconv.Atoi(strArr2[2])
			memReward.ItemLv = val3
		} else {
			memReward.ItemLv = 0
		}

		mailAsk.Reward = append(mailAsk.Reward, memReward)
	}

	mailAsk.SenderName = text7
	mailAsk.Title = text8
	mailAsk.Content = text9
	mailAsk.Reason = text10
	mailAsk.TimeType, _ = strconv.Atoi(text11)
	mailAsk.TimeStart, _ = strconv.Atoi(text12)
	mailAsk.TimeEnd, _ = strconv.Atoi(text13)

	mailAsk.CompareRow = row

	return mailAsk
}

func InsertQuickMailStruct(vMem *QuickMailStruct, userid int, username string) {
	vMem.UserID = append(vMem.UserID, userid)
	vMem.UserName = append(vMem.UserName, username)
}
