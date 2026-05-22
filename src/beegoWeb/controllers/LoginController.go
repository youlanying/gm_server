package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/table"
	"gm_server/src/tools"
	"os"
	"path/filepath"
)

func (c *GMBaseController) Get() {
	fmt.Println("------Get OR Reridect------")
	path, _ := os.Executable()
	//if err != nil {
	//	log.Printf(err)
	//}
	dir := filepath.Dir(path)
	fmt.Println(path) // for example /home/user/main
	fmt.Println(dir)  // for example /home/user

	c.TplName = "index.html"
}

// Post 输入账号密码 登录后 跳转到 平台选择界面
func (c *GMBaseController) Post() {
	vName := c.GetString("username")
	vPwd := MD5(c.GetString("password"))
	userData, ok := beegodb.TbAdminUserAllData[vName]

	if !ok {
		c.PageLoginWitchError("没有这个用户")
		return
	}

	if userData.PASSWORD != vPwd {
		c.PageLoginWitchError("密码错误")
		return
	}

	fmt.Println("登录完成，当前的session:")
	fmt.Println(c.CruSession)

	userGroup, ok := beegodb.TbAdminGroupAllData[userData.GROUPID]
	if !ok {
		c.PageLoginWitchError("该用户所属用户组不正确")
		return
	}
	groupData := userGroup.ACTIONLIST
	strSidebar := getSidebar(groupData)
	platformMap := getPlatforms(groupData)

	adminData := AdminData{
		ID:          userData.ID,
		UserId:      userData.USERID,
		Password:    userData.PASSWORD,
		GroupId:     userData.GROUPID,
		ActionList:  groupData,
		Sidebar:     strSidebar,
		PlatformMap: platformMap,
	}
	c.SetSession("loginuser", vName)
	c.SetSession("UserData", adminData)
	c.Ctx.SetCookie("userid", vName, 36000, "/")
	c.Ctx.SetCookie("sidebar", strSidebar, 36000, "/")
	c.Data["userid"] = vName
	isTrue := CheckRight(adminData, AdminActionPlatFormManager)
	if isTrue {
		c.Data["isboss"] = isTrue
	}
	c.Layout = "basetemplate/platformbase.html"
	if len(platformMap) == 0 {
		c.Data["pagetitle"] = "平台管理"
		c.TplName = "pfselect/index.html"
	} else {
		c.Data["platforms"] = platformMap
		c.Data["pagetitle"] = "选择进入平台"
		c.TplName = "admin/pfselect.html"
	}
}

//设置选择的平台
func (c *GMBaseController) SelectPlatform() {

	fmt.Println("Login SelectCenter")
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
	plantIndex, err := c.GetInt("index")
	fmt.Println(plantIndex)
	if err != nil {
		c.Data["platforms"] = userData.PlatformMap
		c.Data["pagetitle"] = "选择进入平台"
		c.Layout = "basetemplate/platformbase.html"
		c.TplName = "admin/pfselect.html"
		return
	}
	centerData, _ := userData.PlatformMap[int32(plantIndex)]

	userData.ThisPlatformId = int32(plantIndex)
	c.SetSession("UserData", userData)
	c.SetSession("platform", plantIndex)
	c.Ctx.SetCookie("pfselectName", centerData.NAME, 36000, "/")
	c.Layout = "basetemplate/basetemplate.html"
	c.Redirect("/admin/index.html", 302)
}

func (c *GMBaseController) Logout() {

	fmt.Println("------Logout------")

	v := c.GetSession("loginuser")

	if v != nil {
		//删除指定的session
		c.DelSession("loginuser")
		c.DelSession("UserData")
		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)
		//删除指定的session
		c.Ctx.SetCookie("pfselectName", "", 36000, "/")
	}

	c.PageLogin()
}

//获取权限内平台
func getPlatforms(groupData []int32) map[int32]*beegodb.TB_CENTER_LIST {
	centermap := make(map[int32]*beegodb.TB_CENTER_LIST)
	for _, id := range groupData {
		if id >= 1000 {
			center, ok := beegodb.TbCenterListAllData[id]
			if ok {
				centermap[id] = center
			}
		}
	}
	return centermap
}

//获取侧边栏
func getSidebar(groupData []int32) string {
	sidebar := "<li><a href=\"/admin/index.html\"><i class='icon-home'></i>首页</a></li>\n"
	for _, acUi := range table.Admin_action_ui_All {
		if tools.ListsMember(acUi.Id, groupData) {
			action, _ := table.Admin_action_Get(acUi.Id)
			if len(acUi.Idlist) == 0 {
				sidebar += "<li><a href=" + action.Href + "><i class='" + action.Icon + "'></i>" + action.Name + "</a></li>\n"
			} else { //<ul class=\"closed\" style=\"display:none;\">"
				if acUi.Isfold == "true" {
					sidebar += "<li>\n<a href=\"#\"><i class=\"" + action.Icon + "\"></i>" + action.Name + "</a>\n<ul class=\"closed\" style=\"display: none;\">"
				} else {
					sidebar += "<li>\n<a href=\"#\"><i class=\"" + action.Icon + "\"></i>" + action.Name + "</a>\n<ul class style=\"display: block;\">"
				}
				for _, subI := range acUi.Idlist {
					if tools.ListsMember(subI, groupData) {
						subAction, _ := table.Admin_action_Get(subI)
						//sidebar += "<li><a href='#' onclick='OpenHtml("+subAction.Href+")'><i class='" + subAction.Icon + "'></i>" + subAction.Name + "</a></li>\n"
						sidebar += "<li><a href=" + subAction.Href + "><i class='" + subAction.Icon + "'></i>" + subAction.Name + "</a></li>\n"
					}
				}
				sidebar += " </ul>\n</li>\n"
			}
		}
	}
	return sidebar
}
