package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/models"
	"gm_server/src/beegoWeb/netmsg"
	"gm_server/src/logger"
)

/**
 * 平台选择 逻辑处理 这里面有很多 db的原始调用方法
 */
type PfSelectController struct {
	GMBaseController
}

/**
 * 所有平台展示页
 */
func (c *PfSelectController) Index() {

	fmt.Println("PfSelect Index")

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
	isTrue := CheckRight(userData, AdminActionPlatFormManager)
	if isTrue {

	}

	c.Data["pagetitle"] = "平台管理"

	centerList := make([]*beegodb.TB_CENTER_LIST, 0)
	for _, i2 := range userData.PlatformMap {
		centerList = append(centerList, i2)
	}
	c.Data["platforms"] = models.ToJson(centerList)
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "pfselect/index.html"
}

/**
 * 所有平台展示页
 */
func (c *PfSelectController) AjaxIndex() {

	fmt.Println("PfSelect ajaxIndex")

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
	isTrue := CheckRight(userData, AdminActionPlatFormManager)
	if isTrue {
		c.Data["pagetitle"] = "平台管理"

		centerList := make([]*beegodb.TB_CENTER_LIST, 0)
		for _, i2 := range userData.PlatformMap {
			centerList = append(centerList, i2)
		}
		c.Data["platforms"] = models.ToJson(centerList)
		c.TplName = "pfselect/index.html"
	} else {
		c.PageLoginWitchError("没有权限")
	}
}

/**
 * POST 原名 update
 */
func (c *PfSelectController) Update() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	Strcid := c.GetString("cid")
	cid, _ := c.GetInt("cid")
	centerName := c.GetString("center_name")
	ip := c.GetString("ip")
	port := c.GetString("port")
	//fmt.Println("Update=====", cid,centerName,ip,port)
	upNum, ok := beegodb.TB_CENTER_LISTUpdateByKey("id="+Strcid, "name='"+centerName+"',ip='"+ip+"',port='"+port+"'")
	logger.Logf("==Update===:%v", upNum)
	if ok {
		//关闭连接
		netmsg.CloseCenterLink(int32(cid))
		UpdateSessionUserData(c)
		//检查连接平台
		//netmsg.CheckCenterLink(int32(cid), ip+":"+port)
		c.Ctx.WriteString("true")
	} else {
		c.Ctx.WriteString("false")
	}
}

func UpdateSessionUserData(c *PfSelectController) {
	beegodb.InitTbCenterList()
	userData := c.GetSession("UserData").(AdminData)
	centerMap := getPlatforms(userData.ActionList)
	userData.PlatformMap = centerMap
	c.SetSession("UserData", userData)
}

/**
 * POST
 */
func (c *PfSelectController) Delete() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	id, _ := c.GetInt("cid")
	strId := c.GetString("cid")
	num, ok := beegodb.TB_CENTER_LISTdeleteBy("id=" + strId)
	logger.Logf("==Delete===:%v", num)
	if ok {
		//关闭连接
		netmsg.CloseCenterLink(int32(id))
		beegodb.DeleteGroupId(int32(id))
		userData := c.GetSession("UserData").(AdminData)
		delete(userData.PlatformMap, int32(id))
		c.SetSession("UserData", userData)
		centerList := make([]*beegodb.TB_CENTER_LIST, 0)
		for i, i2 := range userData.PlatformMap {
			if i != int32(id) {
				centerList = append(centerList, i2)
			}
		}
		c.Ctx.WriteString(models.ToJson(centerList))
	} else {
		c.Ctx.WriteString("")
	}
}

/**
 * POST
 */
func (c *PfSelectController) CreateNew() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	cname := c.GetString("cname")
	ip := c.GetString("ip")
	port := c.GetString("port")
	maxId := beegodb.GetMaxCenterId()
	fmt.Println("CreateNew=====", maxId, cname, ip, port)
	num, ok := beegodb.TB_CENTER_LISTInsert(maxId+1, cname, ip, port, "")
	if ok {
		fmt.Printf("CreateNew=====id:%v,cname:%v, ip:%v, port:%v, num:%v\n", maxId, cname, ip, port, num)
		beegodb.InitTbCenterList()
		userData := c.GetSession("UserData").(AdminData)
		beegodb.UpdateSuperAdminAuthority(userData.GroupId)
		userData.PlatformMap[maxId+1] = beegodb.TbCenterListAllData[maxId+1]
		c.SetSession("UserData", userData)
		centerList := make([]*beegodb.TB_CENTER_LIST, 0)
		for _, i2 := range userData.PlatformMap {
			centerList = append(centerList, i2)
		}
		c.Ctx.WriteString(models.ToJson(centerList))
		return
	}
	fmt.Println("Creat=====")
	c.Ctx.WriteString("")
}
