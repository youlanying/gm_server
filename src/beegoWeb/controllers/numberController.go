package controllers

import (
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gm_server/src/beegoWeb/models"
	"gm_server/src/beegoWeb/netmsg"
	network_message "gm_server/src/network/message"
	"math/big"
	"strconv"
	"time"
)

type NumberController struct {
	GMBaseController
}

type NumStruct struct {
	Link     string
	Type     string
	Platform int
}

//---- 码生成 ----

func (c *NumberController) Number() {

	fmt.Println("Num Number")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderNumber(v.(string), "")
}

type NumberResult struct {
	Name string
	Desc string
}

func (c *NumberController) RenderNumber(userName string, msg string) {

	switch msg {
	case "1":
		c.Data["msg"] = "对不起，当前有其他管理员正在使用，请5分钟以后再试！..."
		break
	case "2":
		c.Data["msg"] = "您没有此权限"
		break
	default:
		break
	}

	c.Data["pagetitle"] = "激活码&礼包码"
	c.Data["userid"] = userName
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionCDKey)

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "number/number.html"
}

/**
 * 生成码 - 激活码
 */
func (c *NumberController) NewNumber() {

	fmt.Println("Num NewNumber")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	vnum, _ := c.GetInt("num")
	vparam := c.GetString("param")
	vtime, _ := c.GetInt("time")
	svrunique, _ := c.GetInt("svrunique")

	msg := &network_message.GM_CreateNumberRequest{}
	sessionId := netmsg.NewSession()
	userData := c.GetSession("UserData").(AdminData)

	msg.SessionId = sessionId
	msg.NumberType = 1            //码类型，目前暂定 1为激活码，2为礼包码
	msg.ScrapTime = int32(vtime)  // 失效时间，单位为秒，0为永远有效
	msg.Unique = int32(svrunique) // 是否唯一，唯一则只能被1个玩家使用
	//激活码, 原则上来说可以随机生成，目前仅保留2位渠道id，且总长8位
	msg.Param = vparam    // 包含礼包、互斥组等信息的字符串
	msg.Num = int32(vnum) // 生成数量

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_CreateNumberResponse)

	c.Data["pagetitle"] = "激活码&礼包码"
	c.Data["userid"] = v
	c.Data["candoact"] = CheckRight(userData, AdminActionCDKey)
	c.Data["giftlist"] = ret.NumberList
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "number/number.html"
}

/**
 * 生成码 - 新手礼包
 */
func (c *NumberController) NewGift() {

	fmt.Println("Num NewGift")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	vnum, _ := c.GetInt("num")
	vparam := c.GetString("param")
	vtime, _ := c.GetInt("time")
	svrunique, _ := c.GetInt("svrunique")

	msg := &network_message.GM_CreateNumberRequest{}
	sessionId := netmsg.NewSession()
	userData := c.GetSession("UserData").(AdminData)

	msg.SessionId = sessionId
	msg.NumberType = 2            //码类型，目前暂定 1为激活码，2为礼包码
	msg.ScrapTime = int32(vtime)  // 失效时间，单位为秒，0为永远有效
	msg.Unique = int32(svrunique) // 是否唯一，唯一则只能被1个玩家使用
	//对于礼包码，1-2位为渠道id，3-5位为礼包码id，6-7位为互斥组id，8-12位为随机码（延用战舞规则，若想缩短，则需要改库表，比如8-10位总长，在库表里存相关信息）
	msg.Param = vparam    // 包含礼包、互斥组等信息的字符串
	msg.Num = int32(vnum) // 生成数量

	netmsg.SendMsgToGMServer(userData.ThisPlatformId, msg)
	ret := netmsg.RecMsg(sessionId).(network_message.GM_CreateNumberResponse)

	c.Data["pagetitle"] = "激活码&礼包码"
	c.Data["userid"] = v
	c.Data["candoact"] = CheckRight(userData, AdminActionCDKey)
	c.Data["codelist"] = ret.NumberList
	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "number/number.html"
}

//---- other ----

func (c *NumberController) batch_insert(Records []interface{}) {

	fmt.Println("batch_insert start")

	//orm := orm.NewOrm()
	//
	//p, err := orm.Raw("replace into `tb_number`(number,state,type,phasenum,createtime,scraptime,serverunique) values(?,?,?,?,?,?,?)").Prepare()
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//for _, vmem := range Records {
	//
	//	mem := vmem.(models.Tb_number)
	//
	//	_, err := p.Exec(mem.Number, mem.State, mem.Type, mem.Phasenum, mem.Createtime, mem.Scraptime, mem.Serverunique) //这个里面可以传入多个字段
	//
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//}

	fmt.Println("batch_insert end")
}

func (c *NumberController) check_lock_number_create() bool {

	return models.EtcLookUp(models.LOCK_ETS, models.LOCK_NUMBER_CREATE)
}

func (c *NumberController) set_lock_number_create(v string) {
	models.EtcInsert(models.LOCK_ETS, models.LOCK_NUMBER_CREATE, v)
}

func (c *NumberController) clear_lock_number_create() {
	models.EtcInsert(models.LOCK_ETS, models.LOCK_NUMBER_CREATE, "")
}

func (c *NumberController) create_gift_number(Channel string, Gift string, CreateNum int, PhaseNum int, Now time.Time, ScrapEnd time.Time, ServerUnique int) (A []interface{}, B []interface{}) {
	return c.create_gift_number1(Channel, Gift, "", CreateNum, PhaseNum, Now, ScrapEnd, ServerUnique)
}

func (c *NumberController) create_gift_number1(Channel string, Gift string, Mutex string, CreateNum int, PhaseNum int, Now time.Time, ScrapEnd time.Time, ServerUnique int) (A []interface{}, B []interface{}) {
	StrHead := Channel + Gift + Mutex

	var Arr []interface{}
	var Brr []interface{}

	return c.create_number(StrHead, PhaseNum, Now, ScrapEnd, ServerUnique, CreateNum, Arr, Brr)
}

func (c *NumberController) create_activation_number(Channel string, Param string, CreateNum int, PhaseNum int, Now time.Time, ScrapEnd time.Time, ServerUnique int) (A []interface{}, B []interface{}) {
	StrHead := Channel + Param

	var Arr []interface{}
	var Brr []interface{}

	return c.create_number(StrHead, PhaseNum, Now, ScrapEnd, ServerUnique, CreateNum, Arr, Brr)
}

func (c *NumberController) create_number(StrHead string, PhaseNum int, Now time.Time, ScrapEnd time.Time, ServerUnique int, LastNum int, NumberList []interface{}, NumberRecordLists []interface{}) (A []interface{}, B []interface{}) {

	if LastNum == 0 {
		return NumberList, NumberRecordLists
	}

	TempNumber := StrHead + c.Random5()

	Platform := c.GetSession("platform").(int)

	EtsName := strconv.Itoa(Platform) + "_number_ets"

	if models.EtcLookUp(EtsName, TempNumber) {
		return c.create_number(StrHead, PhaseNum, Now, ScrapEnd, ServerUnique, LastNum, NumberList, NumberRecordLists)
	} else {
		if models.ListsMember(TempNumber, NumberList) {
			return c.create_number(StrHead, PhaseNum, Now, ScrapEnd, ServerUnique, LastNum, NumberList, NumberRecordLists)
		} else {

			//var NumberRecord models.Tb_number
			//NumberRecord.Number = TempNumber
			//NumberRecord.State = 0
			//NumberRecord.Type = StrHead
			//NumberRecord.Phasenum = PhaseNum
			//NumberRecord.Createtime = Now
			//NumberRecord.Scraptime = ScrapEnd
			//NumberRecord.Serverunique = ServerUnique

			newNumberList := models.UnShift(TempNumber, NumberList)
			//newNumberRecordLists := models.UnShift(NumberRecord, NumberRecordLists)
			newNumberRecordLists := newNumberList

			return c.create_number(StrHead, PhaseNum, Now, ScrapEnd, ServerUnique, LastNum-1, newNumberList, newNumberRecordLists) //newNumberRecordLists
		}
	}
}

func (c *NumberController) get_insert_list(NumberList []interface{}, Length int, AllNum int) []interface{} {

	var Page int

	if AllNum%Length == 0 {
		Page = AllNum / Length
	} else {
		Page = AllNum/Length + 1
	}

	var resultList []interface{}
	return c.get_insert_list1(NumberList, Length, Page-1, resultList)
}

func (c *NumberController) get_insert_list1(NumberList []interface{}, Length int, Page int, resultList []interface{}) []interface{} {

	if Page == 0 {
		return models.ListConant(NumberList, resultList)
	}

	SubList, _ := models.ListsSplit(Length, NumberList)

	var newArr = models.ListConant(SubList, resultList)

	return c.get_insert_list1(NumberList, Length, Page-1, newArr)
}

type GetPhasenum struct {
	Result int
}

func (c *NumberController) proc_get_phasenum() int {

	SqlStr := "select max(phasenum) from tb_number"

	orm := orm.NewOrm()

	var result []GetPhasenum

	_, err := orm.Raw(SqlStr).QueryRows(&result)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	//如果是 null 也返回 1
	return result[0].Result
}

func (c *NumberController) Random5() string {
	CodeList := []string{"2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	lenNum := len(CodeList) - 1
	max := big.NewInt(int64(lenNum))

	var reStr string

	for i := 0; i < 5; i++ {
		r, _ := rand.Int(rand.Reader, max)
		reStr += strconv.Itoa(int(r.Int64()))
	}

	return reStr
}
