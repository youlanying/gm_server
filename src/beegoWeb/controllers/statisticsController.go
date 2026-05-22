package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/models"
	"net/http"
	"strings"
	"time"
)

/**
 * 统计后台
 */
type StatisticsController struct {
	GMBaseController
}

/**
 * 登录统计
 */
func (c *StatisticsController) LoginGet() {

	fmt.Println("Statistics LoginGet")

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
	c.Data["pagetitle"] = "登陆统计"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmLoginCount)
	//c.Data["centerurl"] = models.PlantList

	//{ok, [{pagetitle, PageTitle},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "statistics/login.html"
}

/**
 * 登录统计  是否存在？
 */
func (c *StatisticsController) LoginPost() {

	fmt.Println("Statistics LoginPost")

	UserId := "admin"
	CmdUrl := c.GetString("cmdurl")
	Account := c.GetString("account")
	Accountlist := strings.Split(Account, "@")

	var vLen = len(Accountlist)

	if vLen > 0 {

		Time := models.GetNowStr()
		Auth := models.GetAuth(UserId, Time)
		UrlParam := "zxlf_get?cmd = gm_white_gm_account&type=set&account=" + Account + "&auth=" + Auth + "&gm=" + UserId + "&time=" + Time
		AllUrl := CmdUrl + UrlParam
		client := http.Client{Timeout: 10 * time.Second}
		resp, error := client.Get(AllUrl)

		defer resp.Body.Close()

		if error != nil {
			panic(error)
		}

		c.Data["send"] = resp

		//?ZH("发送失败！");
		//?ZH("发送失败！");
		//?ZH("发送成功，请查询确认！");
		//?ZH("发送失败！");
		//?ZH("发送可能失败，请查询确认！")

	} else {
		c.Data["send"] = "账号输入错误请重新输入！"
	}

	//{ok, [{send,Res},{centerurl,NewAllCenter}]}
	c.Data["pagetitle"] = "统计平台"
	c.Data["userid"] = "admin"
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmLoginCount)

	//NewAllCenter := models.PlantList
	//c.Data["centerurl"] = NewAllCenter

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "statistics/login.html"

}

/**
 * 注册统计
 * 没有 Post
 */
func (c *StatisticsController) RegisterGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	fmt.Println("Statistics RegisterGet")
	//models.AllPlantServerList()
	c.Data["pagetitle"] = "注册统计"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionGmRegister)
	//c.Data["centerurl"] = models.PlantList

	//{ok, [{pagetitle, PageTitle},{centerurl,NewAllCenter}]}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "statistics/register.html"
}
