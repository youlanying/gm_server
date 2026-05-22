package controllers

import (
	"fmt"
	"gm_server/src/beegoWeb/models"
)

/**
 * 运营相关 表处理部分
 * add send 和 up 三个页面
 */
type TableController struct {
	GMBaseController
}

//---- add_proto 添加数据表 ----

/**
 * 展示数据表
 */
func (c *TableController) AddProtoGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderAddProto()
}

func (c *TableController) RenderAddProto() {

	//TRecords := c.GetProtoTable()

	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionOperatorAddProto)
	//c.Data["tableinfo"] = TRecords
	//{ok, [{candoact, true},{tableinfo,TRecords}]}

	c.Data["pagetitle"] = "添加数据表"
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/add_proto.html"
}

/**
 * 新增数据表
 */
func (c *TableController) AddProtoPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//Tableid,_ := c.GetInt("tableid")
	//Tablename := c.GetString("tablename")
	//Describe := c.GetString("describe")
	//Structure,_ := c.GetInt("list")

	//tb_proto_table:insert(0,Tableid,Tablename,Describe,Structure)

	//c.Data["action"] = "add_proto";

	c.RenderAddProto()
}

/**
 * 删除数据表
 * get only
 */
func (c *TableController) AddProtoDelete() {

	//ID, _ := c.GetInt("ID")
	//tb_proto_table:delete_by_id(list_to_integer(ID))
	//db := models.GetDB()

	//strconv.Itoa 数字转化为字符串

	//r, err := db.Exec("DELETE FROM tb_proto_data WHERE id=" + strconv.Itoa(ID))
	//if err != nil {
	//	fmt.Println("delete failed, ", err)
	//	return
	//}

	//fmt.Println(r)
	fmt.Println("delete succ:")

	//{redirect, [{action, "add_proto"}]}
	c.RenderAddProto()
}

/**
 * 修改数据表
 * post only
 */
func (c *TableController) AddProtoUpdate() {
	//ID, _ := c.GetInt("id")
	//Tid, _ := c.GetInt("tid")
	//Tablename := c.GetString("tname")
	//Describe := c.GetString("describe")
	//Structure := c.GetString("uplist")

	//tb_proto_table:update_by_id(ID, [{tid,Tid},{tname,Tablename},{tdescribe,Describe},{structure,Structure}])
	//db := models.GetDB()

	//stmt, _ := db.Prepare(`UPDATE tb_proto_data SET tid = ?,tname = ?,tdescribe = ?,structure = ? WHERE id=?`)
	//res, _ := stmt.Exec(Tid, Tablename, Describe, Structure, ID)
	//num, _ := res.RowsAffected() //影响行数
	//fmt.Println("update succ:", num)

	//{redirect, [{action, "add_proto"}]}
	c.RenderAddProto()
}

//---- up_proto 表数据添加 ----

/**
 * 表数据添加
 * up_proto Get
 * 这个 Erlang 有2个函数实现
 */
func (c *TableController) UpProtoGet() {

	fmt.Println("UpProtoGet")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.UpProtoRender()
}

/**
 * 不带参数渲染
 */
func (c *TableController) UpProtoRender() {

	StrID := c.GetSession("proto_id")

	var ID int
	var DESCRIBE string
	var TinfoKey []string
	var Thead []string

	//如果session中 没有proto_id
	if StrID == nil {
		ID = 0
		DESCRIBE = ""
	} else { //如果有
		ID = StrID.(int)

		if ID == 0 {
			DESCRIBE = ""
		} else {
			//var vData models.Tb_proto_table
			//vData.Id = ID
			//orm := orm2.NewOrm()
			//orm.Read(&vData)

			//DESCRIBE = vData.Tdescribe
			//STRUCTURE := models.StrToTableTitle(vData.Structure)

			//for _, sMem := range STRUCTURE {
			//	strMem := sMem.CS + ":" + sMem.Key + ":" + sMem.Type
			//	Thead = append(Thead, strMem)
			//	TinfoKey = append(Thead, sMem.Key)
			//}
		}
	}

	c.Data["pagetitle"] = "表数据添加"
	c.Data["result"] = -1
	//c.Data["centerurl"] = models.PlantList
	c.Data["tableinfo"] = TinfoKey
	c.Data["thead"] = Thead
	c.Data["tbname"] = DESCRIBE

	//{ok, [{centerurl,NewAllCenter},{tableinfo,TRecords},{thead,Thead},{tinfo, TinfoKey},{tbname, DESCRIBE},{dataid, ID}]}
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/up_proto.html"
}

/**
 * 带参渲染
 */
func (c *TableController) UpProtoRender2(StrID int) {

	ID := StrID

	var Thead []string

	if ID == 0 {
		//Thead = []
	} else {
		//var vData models.Tb_proto_table
		//vData.Id = ID
		//orm := orm2.NewOrm()
		//orm.Read(&vData)

		//STRUCTURE := models.StrToTableTitle(vData.Structure)

		//for _, sMem := range STRUCTURE {
		//	strMem := sMem.CS + ":" + sMem.Key + ":" + sMem.Type
		//	Thead = append(Thead, strMem)
		//}
	}

	c.Data["pagetitle"] = "表数据添加"
	c.Data["result"] = -1
	//c.Data["centerurl"] = models.PlantList
	//c.Data["tableinfo"] = c.GetProtoTable()
	c.Data["thead"] = Thead
	//{ok, [{result,-1},{centerurl,NewAllCenter},{tableinfo,TRecords},{thead,Thead}]}
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/up_proto.html"
}

/**
 * 表数据添加 原名Select_table
 * up_proto Post 行为
 * 先选择 要查询 哪个表 然后给出 改表的信息
 */
func (c *TableController) UpProtoPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	StrID, _ := c.GetInt("proto_id")
	//ID := erlang:list_to_integer(StrID)
	//boss_session:set_session_data(SessionId, "proto_id", StrID)
	c.SetSession("proto_id", StrID) //设置session

	//{redirect, [{action, "up_proto"},{thead, StrID}]}
	c.UpProtoRender2(StrID)
}

//-- 源码解读: C是客户端的意思 S是服务器的意思 CS就是客户端服务器都用  Cli是 client的缩写 Ser是 server的缩写 --

/**
 * 表POST 数据添加
 * 先读出对应表的 Tb_proto_table表结构 然后按照发的数值 存储到tb_proto_data数据库
 * up_proto Post 原名add_proto_data
 */
func (c *TableController) UpProtoCheck() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//ID, _ := c.GetInt("proto_data_id")

	//var vData models.Tb_proto_table
	//vData.Id = ID
	//orm := orm2.NewOrm()
	//orm.Read(&vData)

	//Tid := vData.Tid;
	//NAME := vData.Tname
	//DESCRIBE := vData.Tdescribe
	//STRUCTURE := vData.Structure //表结构如下 信息为 cs id type
	//
	//Fun = fun({CS,Key,Type},{Acc,AccS,OAcc,OAccS}) ->
	//	KeyData = Req:post_param(atom_to_list(Key))
	//	switch CS{
	//		case s:
	//			NewKeyData = format:cell_to_data(KeyData, Type)
	//			NewAccS = [{Key,NewKeyData}|AccS]
	//			NewOAccS = "{"+atom_to_list(Key)+","+KeyData+"}"+OAccS
	//			{Acc,NewAccS,OAcc,NewOAccS};
	//
	//		case c:
	//			NewKeyDataCli = format:cell_to_data_cli(KeyData, {Type,atom_to_list(Key)})
	//			NewAcc = [{Key,NewKeyDataCli}|Acc]
	//			NewOAcc = "{"+atom_to_list(Key)+","+KeyData+"}"++OAcc
	//			{NewAcc,AccS,NewOAcc,OAccS};
	//
	//		case cs:
	//			NewKeyDataCli = format:cell_to_data_cli(KeyData, {Type,atom_to_list(Key)})
	//			NewKeyData = format:cell_to_data(KeyData, Type)
	//			NewAccS = [{Key,NewKeyData}|AccS]
	//			NewAcc = [{Key,NewKeyDataCli}|Acc]
	//			NewOAccS = "{" + atom_to_list(Key) + "," + KeyData + "}" + OAccS
	//			NewOAcc = "{" + atom_to_list(Key) + "," + KeyData + "}" + OAcc
	//			{NewAcc,NewAccS,NewOAcc,NewOAccS}
	//		}
	//
	//{Cli,Ser,OCli,OSer} = lists:foldr(Fun, {[],[],"",""}, STRUCTURE);
	//
	//if(Cli == []){
	//	UrlCli = "";
	//} else{
	//	[{_,Cid}|_]=Cli,
	//	SendStructCli = {struct,Cli}
	//	QS = util_json:json_encode(SendStructCli)
	//	{ok,SendBinCli0} = QS
	//	SendBinCli = erlang:binary_to_list(SendBinCli0)
	//	StrSendBinCli = util_string:escape_uri(SendBinCli)
	//	UrlCli = "&client=" + StrSendBinCli + "&tid=" + integer_to_list(Tid) + "&cid=" + util_string:escape_uri(util_string:term_to_string(Cid))
	//}

	//if(Ser == []){
	//	UrlServer="";
	//}else{
	//	SendStructSer = {struct,Ser},
	//	{ok,SendBinSer0} = util_json:json_encode(SendStructSer),
	//	SendBinSer = erlang:binary_to_list(SendBinSer0),
	//	StrSendBinSer = util_string:escape_uri(SendBinSer),
	//	UrlServer ="&server="++StrSendBinSer
	//}
	//
	////实际并没有去请求 仅仅是存在数据库中 应该是由h5去请求审核了
	//Urldata := NAME + UrlServer + UrlCli

	//tb_proto_data:insert(0,Tid,NAME,DESCRIBE,OCli,OSer,Urldata,util_time:time_now(),null,"null","null")

	//{redirect, [{action, "up_proto"}]}
	c.UpProtoRender()
}

func (c *TableController) UpProtoFun() {

}

//---- send_proto 表数据信息 ----

/**
 * 表数据信息 Get
 */
func (c *TableController) SendProtoGet() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	isSend, _ := c.GetBool("Send")

	c.SendProtoRender(isSend)
}

func (c *TableController) SendProtoRender(isSend bool) {

	//TRecords := c.GetTableData()
	//TRecordsP := c.GetProtoTable()

	c.Data["pagetitle"] = "表数据信息"
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionOperatorSendProtoData)
	//c.Data["tabledata"] = TRecords
	//c.Data["tableinfo"] = TRecordsP
	c.Data["send"] = true
	//c.Data["centerurl"] = models.PlantList

	if !isSend {
		c.Data["send"] = isSend
	}

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "operator/send_proto.html"
}

/**
 * 表数据信息 Post
 * send_proto
 */
func (c *TableController) SendProtoPost() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//UserId := "admin"
	//ID, _ := c.GetInt("id")
	//CmdUrl := c.GetString("cmdurl")

	//db := models.GetDB()

	//rows, err2 := db.Query("SELECT * FROM `tb_proto_data` where id =" + strconv.Itoa(ID)) //获取所有数据

	//if err2 != nil {
	//	return
	//}

	//var vid int
	//var TID string
	//var TNAME int
	//var TDESCRIBE int
	//
	//var CDATA int
	//var SDATA int
	//var URLDATA int
	//var CREATETIME int
	//
	//var SENDTIME int
	//var SENDID int
	//var AUDITID int

	//rows.Next()
	//rows.Scan(&vid, &TID, &TNAME, &TDESCRIBE, &CDATA, &SDATA, &URLDATA, &CREATETIME, &SENDTIME, &SENDID, &AUDITID)

	//createtime := models.GetNowStr()

	//stmt, _ := db.Prepare(`UPDATE tb_proto_data SET sendtime = ?,sendid = ? WHERE id=?`)
	//res, _ := stmt.Exec(createtime, UserId, vid)
	//num, _ := res.RowsAffected() //影响行数
	//fmt.Println("insert succ:", num)

	//tb_proto_data:update_by_id(ID, [{sendtime,util_time:time_now()},{sendid,UserId}])
	//Bool = case util_http:http_get(AllUrl)
	//{redirect, [{action, "send_proto"},{send,Bool}]}
	c.SendProtoRender(true)
}

/**
 * 表数据信息
 * send_proto 删除表
 */
func (c *TableController) SendProtoDeleteTable() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	CmdUrl := c.GetString("cmdurl")
	vType := c.GetString("mailtype")

	var UrlParam string
	var Tablename string

	switch vType {
	case "Q":
		UrlParam = "zxlf_get?cmd=gm_empty_ets&type=all"
	case "D":
		Tablename = c.GetString("tablename")
		UrlParam = "zxlf_get?cmd=gm_empty_ets&type=table&table=" + Tablename
	case "K":
		Tablename = c.GetString("tablename")
		Tablekey := c.GetString("tablekey")
		UrlParam = "zxlf_get?cmd=gm_empty_ets&type=id&table=" + Tablename + "&key=" + Tablekey
	default:
		break
	}

	AllUrl := CmdUrl + UrlParam

	fmt.Printf(AllUrl)
	//Result,error := util_http:http_get(AllUrl);

	//{redirect, [{action, "send_proto"},{send,Bool}]}
	c.SendProtoRender(true)
}

/**
*  表数据信息
* send_proto 删除表数据
* Get Only
 */
func (c *TableController) SendProtoDeleteTableData() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//ID, _ := c.GetInt("ID")

	//tb_proto_data:delete_by_id(list_to_integer(ID));

	//db := models.GetDB()

	//strconv.Itoa 数字转化为字符串

	//r, err := db.Exec("DELETE FROM tb_proto_data WHERE id=" + strconv.Itoa(ID))
	//if err != nil {
	//	fmt.Println("delete failed, ", err)
	//	return
	//}

	//fmt.Println(r)
	fmt.Println("delete succ:")

	//{redirect, [{action, "send_proto"}]}
	c.SendProtoRender(true)
}

//---- 核心页查询 ----

//func (c *TableController) GetProtoTable() []models.Tb_proto_table {

//page := models.GetProtoTableWithPage(1, 10)

//var vvList []models.Tb_proto_table = page.List.([]models.Tb_proto_table)

//foldl 倒序
//var reList []models.Tb_proto_table;
//for _,vmem := range vvList{
//STRUCTURE1 = util_string:term_to_string(STRUCTURE) 对象转化为字符串？
//}

//vLen := len(vvList)

//if vLen <= 1 {
//	return vvList
//}

//var temp models.Tb_proto_table

//for i := 0; i < vLen/2; i++ {
//	temp = vvList[i]
//	vvList[i] = vvList[vLen-1-i]
//	vvList[vLen-1-i] = temp
//}

//return vvList
//}

func (c *TableController) GetTableData() []models.ProtoData {

	//page := models.GetTableDataWithPage(1, 10)
	//var page models.Page

	//var vvList []models.Tb_proto_data = page.List.([]models.Tb_proto_data)

	//先调转数组
	//vLen := len(vvList)

	//foldl 倒序 调转数组
	//if vLen > 1 {
	//	var temp models.Tb_proto_data
	//
	//	for i := 0; i < vLen/2; i++ {
	//		temp = vvList[i]
	//		vvList[i] = vvList[vLen-1-i]
	//		vvList[vLen-1-i] = temp
	//	}
	//}

	var reList []models.ProtoData

	//再处理
	//for _, vmem := range vvList {
	//	//Now_time = calendar:now_to_local_time(CREATETIME)
	//	//StrTime = util_string:datetime_to_string(Now_time)
	//	StrTime := models.TimeToString(vmem.Createtime)
	//
	//	var StrSendTime string
	//
	//	if vmem.Sendtime.IsZero() {
	//		StrSendTime = ""
	//	} else {
	//		//NowSend_time := calendar:now_to_local_time(SENDTIME);
	//		//StrSendTime = util_string:datetime_to_string(NowSend_time);
	//		StrSendTime = models.TimeToString(vmem.Sendtime)
	//	}
	//
	//	var newMem models.ProtoData
	//	newMem.Id = vmem.Id
	//	newMem.Tdescribe = vmem.Tdescribe
	//	newMem.Cdata = vmem.Cdata
	//	newMem.Sdata = vmem.Sdata
	//	newMem.CreatetimeStr = StrTime
	//	newMem.SendtimeStr = StrSendTime
	//
	//	reList = append(reList, newMem)
	//}

	return reList
}
