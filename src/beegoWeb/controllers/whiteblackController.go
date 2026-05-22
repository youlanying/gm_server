package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/models"
)

type WhiteBlackController struct {
	GMBaseController
}

/**
 * er - top Ger方法
 */
func (c *WhiteBlackController) Top() {

	fmt.Println("whiteblack Top Get")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//c.Data["platforms"] = models.All_center
	c.Data["pagetitle"] = "黑白名单"
	c.Data["userid"] = v
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionBlackListAndWhiteList)

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "whiteblack/top.html"
}

/**
 * 是 get方法 2个index合并
 */
func (c *WhiteBlackController) Index() {

	fmt.Println("whiteblack Index Get")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//plantIndex := c.GetSession("platform")

	//if(plantIndex == nil){
	plantIndex := c.GetString("index")
	c.SetSession("platform", plantIndex)
	//}
	fmt.Println("===Index===", plantIndex)
	c.do_index(v.(string), plantIndex)
}

type LinkStruct struct {
	CName    string
	Show     string
	Platform string
	EName    string
}

func (c *WhiteBlackController) do_index(userName string, Platform string) {

	PageTitle := "名单类型选择"
	Link1 := "黑白名单"
	Link2 := "黑白IP"

	Type1 := "special_account"
	Type2 := "special_ip"

	mem1 := LinkStruct{Link1, "show", Platform, Type1}
	mem2 := LinkStruct{Link2, "show", Platform, Type2}

	Links := []LinkStruct{mem1, mem2}

	c.Data["links"] = Links
	c.Data["selectpf"] = Platform
	c.Data["pagetitle"] = PageTitle
	c.Data["userid"] = userName
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionBlackListAndWhiteList)

	//{ok, [{'links', Links}, {'selectpf', Platform}, {'pagetitle', PageTitle}, {'userid', UserId}]}.
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "whiteblack/index.html"
}

func (c *WhiteBlackController) Show() {

	fmt.Println("whiteblack Show")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Type := c.GetString("type")
	Platform := c.GetSession("platform")
	if Platform == nil {
		Platform = c.GetString("platform")
	}
	fmt.Println("===Show===", Type, Platform)
	c.RenderShow(v, Platform.(string), Type)
}

type LinkMem struct {
	Link     string
	Index    string
	Platform string
}

func (c *WhiteBlackController) RenderShow(userName interface{}, Platform string, vtype string) {

	var PageTitle string
	var TypeTitle string

	//var orm = orm2.NewOrm();

	var TypeList interface{}
	//TableName := c.GetTableNameByType(vtype);

	var isIp bool
	//db := models.GetCenterDB(Platform)
	switch vtype {

	case "special_account":
		PageTitle = "黑白名单管理"
		TypeTitle = "名单"
		//var list1 []models.Tb_special_account
		//var sqlStr = "SELECT * FROM `tb_special_account`"
		//
		//rows, err := db.Query(sqlStr)
		////_,err := orm.QueryTable("tb_special_account").All(&list1);
		//if err != nil {
		//	fmt.Println(vtype, "RenderShow error"+err.Error())
		//	return
		//}
		//var account string
		//var atype int
		//var comment string
		//for rows.Next() { //循环显示所有的数据
		//	rows.Scan(&account, &atype, &comment)
		//	fmt.Println(account, atype, comment, "--")
		//	vmem := models.Tb_special_account{account, atype, comment}
		//	list1 = append(list1, vmem)
		//}
		//TypeList = list1
		isIp = false
		break

	case "special_ip":
		PageTitle = "黑白IP管理"
		TypeTitle = "IP"
		//var list1 []models.Tb_special_ip
		//var sqlStr = "SELECT * FROM `tb_special_ip`"
		//rows, err := db.Query(sqlStr)
		////_,err := orm.QueryTable("tb_special_ip").All(&list1);
		//if err != nil {
		//	fmt.Println(vtype, "RenderShow error"+err.Error())
		//	return
		//}
		//var ip string
		//var atype int
		//var comment string
		//for rows.Next() { //循环显示所有的数据
		//	rows.Scan(&ip, &atype, &comment)
		//	fmt.Println(ip, atype, comment, "--")
		//	vmem := models.Tb_special_ip{ip, atype, comment}
		//	list1 = append(list1, vmem)
		//}
		//TypeList = list1
		isIp = true
		break

		//case "whiteaccount":
		//	PageTitle = "白名单管理"
		//	TypeTitle = "名单"
		//	var list1 []models.Tb_white_account;
		//	_,err := orm.QueryTable("tb_white_account").All(&list1);
		//	if(err != nil) {
		//		fmt.Println(vtype,"RenderShow error" + err.Error())
		//		return;
		//	}
		//	TypeList = list1;
		//
		//	isIp = false;
		//	break;
		//
		//case "whiteip":
		//	PageTitle = "白IP管理"
		//	TypeTitle = "IP"
		//	var list1 []models.Tb_white_ip;
		//	_,err := orm.QueryTable("tb_white_ip").All(&list1);
		//	if(err != nil) {
		//		fmt.Println(vtype,"RenderShow error" + err.Error())
		//		return;
		//	}
		//	TypeList = list1;
		//
		//	isIp = true;
		//	break;
	}

	var linkMem LinkMem
	linkMem.Link = "返回"
	linkMem.Index = "index"
	linkMem.Platform = Platform
	Links := []LinkMem{linkMem}

	c.Data["isIP"] = isIp

	c.Data["links"] = Links
	c.Data["selectpf"] = Platform
	c.Data["pagetitle"] = PageTitle
	c.Data["type"] = vtype
	c.Data["type_title"] = TypeTitle
	c.Data["type_list"] = TypeList
	c.Data["userid"] = userName
	c.Data["special_list"] = models.SpecialList
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionBlackListAndWhiteList)

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "whiteblack/show.html"
}

func (c *WhiteBlackController) Create() {

	fmt.Println("whiteblack Create")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Platform := c.GetSession("platform")

	if Platform == nil {
		Platform = c.GetString("platform")
	}

	Type := c.GetString("type")
	Name := c.GetString("name")
	Desc := c.GetString("desc")
	//Special := c.GetString("select_special_num")

	fmt.Println(Type, Name, Desc)

	//db := models.GetCenterDB(Platform.(string))
	//
	//var err error
	//
	//switch Type {
	//
	//case "special_account":
	//	var sqlStr = "insert into tb_special_account(account, type, comment)values(?,?,?)"
	//	_, err = db.Exec(sqlStr, Name, Special, Desc)
	//	break
	//
	//case "special_ip":
	//	var sqlStr = "insert into tb_special_ip(ip, type, comment)values(?,?,?)"
	//	_, err = db.Exec(sqlStr, Name, Special, Desc)
	//	break
	//default:
	//	fmt.Println("type error")
	//	return
	//}

	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}

	fmt.Println("insert succ:")

	//centerUrl := models.All_center[Platform.(string)].Center_node_api
	switch Type {

	case "special_account":
		//models.HttpGet(centerUrl + "zxlf_update_special_account")
		break
	case "special_ip":
		//models.HttpGet(centerUrl + "zxlf_update_special_ip")
		break
	default:
		fmt.Println("type error")
		return
	}

	c.Data["action"] = "show"
	c.Data["platform"] = Platform
	c.Data["type"] = Type
	//{redirect, [{action, "show"}, {platform, Platform}, {type, Type}]}.
	c.RenderShow(v, Platform.(string), Type)
}

func (c *WhiteBlackController) Delete() {

	fmt.Println("whiteblack Create")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Platform := c.GetSession("platform")

	if Platform == nil {
		Platform = c.GetString("platform")
	}

	Type := c.GetString("type")
	//Key := c.GetString("id")
	//centerUrl := models.All_center[Platform.(string)].Center_node_api
	//db := models.GetCenterDB(Platform.(string))
	//switch Type {
	//
	//case "special_account":
	//	var sqlStr = "delete from tb_special_account where account=?"
	//	db.Exec(sqlStr, Key)
	//	//models.HttpGet(centerUrl + "zxlf_update_special_account")
	//	break
	//case "special_ip":
	//	var sqlStr = "delete from tb_special_ip where ip=?"
	//	db.Exec(sqlStr, Key)
	//	//models.HttpGet(centerUrl + "zxlf_update_special_ip")
	//	break
	//default:
	//	fmt.Println("type error")
	//	return
	//}

	//{redirect, [{action, "show"}, {platform, Platform}, {type, Type}]}.
	c.RenderShow(v, Platform.(string), Type)
}
