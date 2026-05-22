package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/models"
	"gm_server/src/beegoWeb/table"
	"gm_server/src/tools"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type AdminController struct {
	GMBaseController
}

func (c *AdminController) Index() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	_basePath := path.Join(path.Dir(os.Args[0]), "../beegoWeb/static/images/index")
	files, _ := ioutil.ReadDir(_basePath)
	nameList := make([]string, 0)
	for i := 0; i < len(files); i++ {
		nameList = append(nameList, files[i].Name())
	}
	jsonImgList := models.ToJson(nameList)
	c.Data["pagetitle"] = "首页"
	c.Data["jsonImgList"] = jsonImgList
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "admin/index.html"
}

//---- User部分 ----

func (c *AdminController) ShowUser() {
	fmt.Println("Admin show user")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.updateUser()
}

func (c *AdminController) updateUser() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	UserMapData := beegodb.TbAdminUserAllData
	userListData := make([]beegodb.TB_ADMIN_USER, 0)
	for _, user := range UserMapData {
		userListData = append(userListData, *user)
	}

	groupListData := make([]beegodb.TB_ADMIN_GROUP, 0)
	for _, group := range beegodb.TbAdminGroupAllData {
		groupListData = append(groupListData, *group)
	}

	c.Data["pagetitle"] = "用户管理"

	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionUserManager)
	c.Data["userlist"] = userListData
	c.Data["grouplist"] = groupListData //组列表
	c.Data["defgroup"] = 3              //默认选中哪个索引 3是新用户

	//是否是 没有列表的情况
	c.Data["type"] = 1

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "admin/showuser.html"
}

//创建新用户
func (c *AdminController) NewUser() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	fmt.Println("Admin new user")
	vnewName := c.GetString("id")
	tmpPwd := c.GetString("pwd")
	vpassword := MD5(tmpPwd)
	vgroupid, _ := c.GetInt("group")
	fmt.Println("input group = ", vgroupid)

	// 判断是否重名
	_, ok := beegodb.TbAdminUserAllData[vnewName]
	if ok {
		fmt.Println("name is already had, please change it")
		return
	}
	_, ok1 := beegodb.TB_ADMIN_USERInsertAuto(vnewName, vpassword, int32(vgroupid), "")
	if !ok1 {
		fmt.Println("exec failed, ")
		return
	}
	beegodb.InitTbAdminUser()

	c.updateUser()
}

//删除用户
func (c *AdminController) DeleteUser() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	fmt.Println("Admin delete user")

	//需不需要 先查一下？

	//strconv.Itoa 数字转化为字符串

	vid, _ := c.GetInt("id")

	_, ok := beegodb.TB_ADMIN_USERdeleteBy("id=" + strconv.Itoa(vid))
	if !ok {
		fmt.Println("delete failed, ")
		return
	}
	beegodb.InitTbAdminUser()

	fmt.Println("delete succ:")

	c.updateUser()
}

// EditUser 用户 编辑
func (c *AdminController) EditUser() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	fmt.Println("Admin edituser Get")

	userid, err := c.GetInt("id")
	if err != nil {
		return
	}
	fmt.Println("input", userid)

	userData0, ok := beegodb.TbAdminUserIntMap[int32(userid)]
	if !ok {
		return
	}

	c.Data["pagetitle"] = "用户管理"

	//c.Data["res"] = 0; 这个模板中用 if判断后 可以放心食用
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionUserManager)
	// todo 下面兩個也不對！！！！
	c.Data["ismaster"] = true
	c.Data["isself"] = userData0.USERID == "admin"
	c.Data["edituserid"] = userData0.USERID
	c.Data["cid"] = userData0.ID      //当前的 ID
	c.Data["gid"] = userData0.GROUPID //当前的组ID
	c.Data["grouplist"] = beegodb.TbAdminGroupAllData

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "admin/edituser.html"
}

func (c *AdminController) DoEditUser() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	fmt.Println("Admin DoEditUser Post")

	userid := c.GetString("id")
	oldpwd := c.GetString("oldpwd")
	newpwd := c.GetString("newpwd")
	newgroup := c.GetString("group")

	fmt.Printf("==userid:%v, oldpwd:%v, newpwd:%v, newgroup:%v==\n", userid, oldpwd, newpwd, newgroup)

	//用权限判断更好些
	if len(newpwd) > 0 && oldpwd != newpwd {
		//新旧密码的判断 等还需要补充

		//修改一行
		row, ok := beegodb.TB_ADMIN_USERUpdateByKey("id="+userid, "password="+newpwd+",groupid="+newgroup)
		fmt.Printf("insert succ 2 param==row:%v, ok:%v\n", row, ok)
		beegodb.InitTbAdminUser()
	} else {
		//修改一行
		row, ok := beegodb.TB_ADMIN_USERUpdateByKey("id="+userid, "groupid="+newgroup)
		fmt.Printf("insert succ 1 param==row:%v, ok:%v\n", row, ok)
		beegodb.InitTbAdminUser()
	}

	//修改一行
	//c.updateUser()
	c.Redirect("/admin/showuser", 302)
}

//---- Group部分 ----

func (c *AdminController) Showgroup() {

	fmt.Println("Admin showgroup")

	v := c.GetSession("loginuser")

	if v == nil {
		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	actionAllList := getActionAll()
	c.Data["pagetitle"] = "组管理"
	// todo 上面的不對吧...
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionUserManager)
	vvList := getGroupMem()
	c.Data["grouplist"] = models.ToJson(vvList)
	c.Data["actlist"] = actionAllList

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "admin/showgroup.html"
}

func getGroupMem() []models.GroupMem {
	var vvList []models.GroupMem
	for _, mem := range beegodb.TbAdminGroupAllData {
		var gmem models.GroupMem
		gmem.Gid = mem.ID
		gmem.Gname = mem.NAME
		gmem.Gactlist = actionListToStr(mem.ACTIONLIST)
		vvList = append(vvList, gmem)
	}
	return vvList
}

//权限组ID转名称
func actionListToStr(list []int32) string {
	var strAction string
	for i := 0; i < len(list); i++ {
		k := list[i]
		actionData, ok := table.Admin_action_Get(k)
		if ok {
			strAction += "<" + actionData.Name + ">"
		}
		centerData, ok1 := beegodb.TbCenterListAllData[k]
		if ok1 {
			strAction += "<" + "平台：" + centerData.NAME + ">"
		}
	}
	return strAction
}

func getActionAll() []*table.Admin_action {
	actionAllList := table.Admin_action_All
	for _, centerData := range beegodb.TbCenterListAllData {
		actionAllList = append(actionAllList, &table.Admin_action{Id: centerData.ID, Name: "平台:" + centerData.NAME})
	}
	return actionAllList
}

func (c *AdminController) Newgroup() {
	fmt.Println("Admin newgroup")
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	newName := c.GetString("gname")
	showList := c.GetStrings("gactlist")
	fmt.Printf("insert succ=newName:%v,showList:%v\n", newName, showList)
	groupIdList := tools.ListStringToListInt32(showList)
	_, ok := beegodb.TB_ADMIN_GROUPInsertAuto(newName, groupIdList)
	if ok {
		beegodb.InitTbAdminGroup()
		vvList := getGroupMem()
		jsongroup := models.ToJson(vvList)
		c.Ctx.WriteString(jsongroup)
	} else {
		c.Ctx.WriteString("")
	}

}

func (c *AdminController) Deletegroup() {
	fmt.Println("Admin deletegroup")
	gid, _ := c.GetInt("gid")
	strGid := c.GetString("gid")

	notIn := check_guoup_user(int32(gid))
	fmt.Println("delete succ:", gid, notIn)
	if notIn {
		_, ok := beegodb.TB_ADMIN_GROUPdeleteBy("id=" + strGid)
		if ok {
			beegodb.InitTbAdminGroup()
			vvList := getGroupMem()
			jsongroup := models.ToJson(vvList)
			c.Ctx.WriteString(jsongroup)
		} else {
			c.Ctx.WriteString("")
		}
		return
	}
	c.Ctx.WriteString(models.ToJson("error"))
}

func check_guoup_user(gid int32) bool {
	ok := true
	for _, user := range beegodb.TbAdminUserAllData {
		if user.GROUPID == gid {
			ok = false
		}
	}
	return ok
}

/**
 * Editgroup Get
 */
func (c *AdminController) Editgroup() {
	fmt.Println("Admin editgroup")

	gid, _ := c.GetInt("gid")
	GroupData := beegodb.TbAdminGroupAllData[int32(gid)]
	c.Data["pagetitle"] = "组管理"
	// todo 上面的腫麽可能~~
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionUserManager)
	c.Data["gid"] = gid
	c.Data["gname"] = GroupData.NAME
	c.Data["showlist"] = GroupData.ACTIONLIST
	actionAllList := getActionAll()
	c.Data["actlist"] = actionAllList

	c.TplName = "admin/editgroup.html"
}

/**
 * Editgroup Post 修改群组权限
 */
func (c *AdminController) DoEditgroup() {

	fmt.Println("Admin DoEditgroup")

	newName := c.GetString("gname")
	vID, _ := c.GetInt("gid")

	//将数组[1,2,3] 变成 这样的字符串
	var values []string = c.Input()["gactlist"]
	fmt.Printf("=====================%v\n", values)
	groupIdList := tools.ListStringToListInt32(values)
	group := &beegodb.TB_ADMIN_GROUP{
		ID:         int32(vID),
		NAME:       newName,
		ACTIONLIST: groupIdList,
	}
	beegodb.UpdateTbAdminGroup(group)

	c.Redirect("/admin/showgroup", 302)
}
func (c *AdminController) updateGroup() {

	//rows, err := db.Query("SELECT * FROM `tb_admin_group`") //获取所有数据

	//if err != nil {
	//	fmt.Println("-----------------")
	//	fmt.Println(err)
	//	return
	//} else { //else 不能出现在下一行 草草草
	//	fmt.Println("select success")
	//}

	var reList []models.GroupMem

	//var vvid int
	//var vvname string
	//var quanxianList string
	//
	//for rows.Next() {
	//	rows.Scan(&vvid, &vvname, &quanxianList)
	//	vmem := models.GroupMem{vvid, vvname, models.ReplaceAdminAction(quanxianList)}
	//	reList = append(reList, vmem)
	//}

	c.Data["pagetitle"] = "组管理"
	c.Data["userid"] = "admin"
	// todo 所以上面的爲啥這樣寫...
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionUserManager)
	c.Data["grouplist"] = reList
	//c.Data["actlist"] = models.AdminActions

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "admin/showgroup.html"
}

func (c *AdminController) Logout() {
	fmt.Println("Admin logout")
}
