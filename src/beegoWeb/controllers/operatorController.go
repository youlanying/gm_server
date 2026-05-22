package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/models"
	"strconv"
)

/**
 * 运营相关 邮件与跑马灯
 */
type OperatorController struct {
	GMBaseController
}

//---- 跑马灯 ----

/**
 * Get 获取所有跑马灯
 */
func (c *OperatorController) Send_marquee() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderMarquee(v, "")
}

/**
 * 渲染走马灯列表
 */
func (c *OperatorController) RenderMarquee(adminName interface{}, Res string) {
	//models.AllPlantServerList()
	//AuditMarquee = util_tools:check_user_group(SessionId, ?ADMIN_OPERATOR_AUDIT_MARQUEE)
	c.Data["admin"] = adminName
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionOperatorSendMarquee)
	//c.Data["centerurl"] = models.PlantList
	c.Data["marquee_result"] = c.GetMarqueeRecords(true)
	c.Data["audit_res"] = Res

	//{ok, [{candoact, true},{centerurl,NewAllCenter},{marquee_result, MarqueeRecords},{audit_res, Res}]}

	c.Data["pagetitle"] = "跑马灯"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/send_marquee.html"
}

func (c *OperatorController) GetMarqueeRecords(Limit bool) []models.MarInfo {

	//page := models.GetMarqueeWithPage(1, 10)

	var reList []models.MarInfo

	vvList, _ := beegodb.TB_SEND_MARQUEEReadBySQL("")

	for _, vmem := range vvList {

		var newMem models.MarInfo
		newMem.TB_SEND_MARQUEE = *vmem

		//var Stb string
		var StrTime string
		var StrAuditTime string
		var Disabled0 string

		if vmem.MTYPE == 0 {
			newMem.TypeStr = "普通跑马灯"
		} else {
			newMem.TypeStr = "重要中央提示"
		}

		if vmem.SENDOK == "no" {
			//Stb = "C/"
			newMem.Mclass = "btn btn-danger" //Btn_danger
			Disabled0 = ""
			newMem.Style = ""
			newMem.SENDOK = "待审核"
		} else {
			//Stb = "F/"
			newMem.Style = "display:none"
			Disabled0 = "disabled='disabled'"
			newMem.Mclass = "btn btn-success" //Btn_danger
			newMem.SENDOK = "已发送"
		}

		if Limit == true {
			newMem.Disabled = Disabled0
		} else {
			newMem.Disabled = "disabled='disabled'"
		}

		if newMem.CREATETIME == "" {
			StrTime = "NULL"
		} else {
			//Now_time = calendar:now_to_local_time(CREATETIME),
			//StrTime = util_string:datetime_to_string(Now_time)
			StrTime = models.GetNowStr()
		}

		if newMem.AUDITTIME == "" {
			StrAuditTime = "NULL"
		} else {
			//Now_time1 = calendar:now_to_local_time(AUDITTIME)
			//StrAuditTime = util_string:datetime_to_string(Now_time1)
			StrAuditTime = models.GetNowStr()
		}

		//if(lists:keyfind(SERVERURL, 2, AllServerName)  == false)
		//SERVERURL1 = SERVERURL;
		//{Sname,_}->
		//SERVERURL1 = Sname

		newMem.AUDITTIME = StrAuditTime
		newMem.CREATETIME = StrTime

		reList = append(reList, newMem)
	}

	return reList
}

/**
 *  Post 编辑一条 跑马灯 并提交
 */
func (c *OperatorController) Send_marqueePost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Content := c.GetString("content")             //内容
	Marqueetype, _ := c.GetInt("marqueetype")     //类型  普通/重要中央提示
	MarNum, _ := c.GetInt("select_terminate_num") // 播放次数
	CmdUrl := c.GetString("cmdurl")
	timeNow := models.GetNowStr()

	orm := orm.NewOrm()

	//insert(0,list_to_integer(Marqueetype),list_to_integer(MarNum),Content, util_time:time_now(),"no", 0, CmdUrl, UserId, "Null");
	//r, err := db.Exec("insert into tb_send_marquee(type,num,content,createtime,sendok,audittime,serverurl,sendid,auditid)values(?,?,?,?,?,?,?,?,?)",
	//	0,Marqueetype,MarNum,Content,timeNow,"no",0,CmdUrl,UserId,"Null");

	var mem beegodb.TB_SEND_MARQUEE
	mem.ID = 1
	mem.MTYPE = int32(Marqueetype)
	mem.NUM = int32(MarNum)
	mem.CONTENT = Content

	mem.CREATETIME = timeNow
	mem.SENDOK = "no"
	mem.AUDITTIME = "0"
	mem.SERVERURL = CmdUrl
	mem.SENDID = v.(string)
	mem.AUDITID = "Null"

	id, err := orm.Insert(&mem)
	if err == nil {
		fmt.Println("insert succ:", id)
	} else {
		fmt.Println("insert fail:", err)
	}

	c.RenderMarquee(v, "新增走马灯成功")
}

/**
 *  审核一条走马灯
 */
func (c *OperatorController) Audit_marquee() {

	fmt.Println("Audit_marquee")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Marid, _ := c.GetInt("marid")
	var ReStr string

	sqlStr := "where id =" + strconv.Itoa(Marid)
	reList, ok := beegodb.TB_SEND_MARQUEEReadBySQL(sqlStr)

	if !ok || len(reList) == 0 {
		ReStr = "审核失败！已消失,请重新添加！"
	} else { //else 不能出现在下一行 草草草

		marqueeStruct := reList[0]

		NewContent := marqueeStruct.CONTENT

		if marqueeStruct.SENDOK == "no" {
			Time := models.GetNowStr()
			UserId := v.(string)
			Auth := models.GetAuth(UserId, Time)
			Type := strconv.Itoa(int(marqueeStruct.MTYPE))
			Num := strconv.Itoa(int(marqueeStruct.NUM))
			UrlParam := "zxlf_get?cmd=gm_send_broadcast&content=" + NewContent + "&type=" + Type + "&num=" + Num + "&auth=" + Auth + "&gm=" + UserId + "&time=" + Time
			AllUrl := marqueeStruct.SERVERURL + UrlParam

			fmt.Println("即将发送http请求")
			byteArr, err1 := models.HttpGet(AllUrl)

			if err1 != nil {
				ReStr = "连接错误,请检查服务器URL或内容去除一些符号再试!!"
			} else {
				fmt.Println("拿到请求结果继续处理")

				result := models.HttpResult{}
				err2 := json.Unmarshal(byteArr, &result)
				if err2 == nil {
					fmt.Println("jsonStruct.result=" + result.Result)
					if result.Result == "ok" {

						marqueeStruct.AUDITID = v.(string)
						marqueeStruct.AUDITTIME = models.GetNowStr()
						marqueeStruct.SENDOK = "yes"
						//tb_send_marquee:update_by_marid(MARID, [{sendok, "yes"},{auditid, UserId},{audittime,util_time:time_now()}]),

						_, OK := beegodb.TB_SEND_MARQUEEUpdateBy(marqueeStruct)

						if OK {
							ReStr = "审核成功 结果存入数据库失败！"
						} else {
							ReStr = "恭喜，审核发送成功！"
						}

					} else if result.Result == "false" {
						ReStr = "审核失败,请检查跑马灯信息！"
					} else {
						fmt.Println("失败原因:" + result.Result)
						ReStr = "审核可能失败,请进游戏确认！"
					}
				} else {
					ReStr = "审核失败,返回错误,请进游戏确认！"
				}
			}
		} else {
			ReStr = "审核失败！"
		}
	}

	c.RenderMarquee(v, ReStr)
}

func (c *OperatorController) DeleteMarquee() {

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

	sqlstr := "where id = " + strconv.Itoa(Did)
	_, ok := beegodb.TB_SEND_MARQUEEdeleteBy(sqlstr)

	if !ok {
		return
	}

	//db := models.GetDB();
	//models.DeleteFun(db,"tb_send_marquee","mailid",Did);
	c.RenderMarquee(v, "删除待审核邮件成功")
}

//---- get rank ----

func (c *OperatorController) GetRank() {

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
	//NewAllCenter := models.PlantList
	//{ok, [{centerurl,NewAllCenter}]}
	//c.Data["centerurl"] = NewAllCenter
	c.Data["pagetitle"] = "获取榜单"

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/get_rank.html"
}

func (c *OperatorController) GetRankPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	UserId := v.(string)
	CmdUrl := c.GetString("cmdurl")
	RType := c.GetString("rtype")
	RankType := c.GetString("ranktype")
	Time := models.GetNowStr()
	Auth := models.GetAuth(UserId, Time)
	UrlParam := "zxlf_get?cmd=gm_get_rank&type=" + RType + "&ranktype=" + RankType + "&gm=" + UserId + "&time=" + Time + "&auth=" + Auth

	AllUrl := CmdUrl + UrlParam

	var ResultStr string
	Result, error := models.HttpGet(AllUrl)

	if error != nil {
		ResultStr = "获取失败！"
	} else {
		ResultStr = string(Result)
	}

	//NewAllCenter := models.PlantList
	//c.Data["centerurl"] = NewAllCenter
	c.Data["allrank"] = ResultStr

	//{ok, [{centerurl,NewAllCenter},{allrank,Result}]}

	c.Data["pagetitle"] = "获取榜单"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/get_rank.html"
}
