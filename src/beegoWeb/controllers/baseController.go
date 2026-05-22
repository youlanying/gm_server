// Package baseController implements boilerplate code for all baseControllers.
package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"gm_server/src/beegoWeb/beegodb"
	"io"
	"strconv"
	"strings"
)

const (
	AdminActionUserManager     = iota + 1 // 用户管理
	AdminActionPlatFormManager            // 平台管理
	AdminActionGmTools                    // GM工具
	AdminActionVersionManager             // 版本管理

	AdminActionGmDictionary   // 字典查询
	AdminActionGmRedis        // redis查询，這個已經沒有了吧？？
	AdminActionGmServerNum    // 在线查询
	AdminActionGmKick         // 踢人
	AdminActionGmLoginCount   // 登录统计
	AdminActionGmRegister     // 注册统计
	AdminActionGmServerTime   // 修改服务器时间
	AdminActionGmSelectCharge // 充值查询
	AdminActionGmCanCreate    // 查询/设置服务器是否可以创建新角色
	AdminActionGmDeleteItem   // 删除道具

	AdminActionOperatorSendMail      // 发送邮件
	AdminActionOperatorAuditMail     // 审核邮件
	AdminActionOperatorSendMarquee   // 发送跑马灯
	AdminActionOperatorAuditMarquee  // 审核跑马灯
	AdminActionOperatorAddProto      // 添加数据表
	AdminActionOperatorAddProtoData  // 表数据添加
	AdminActionOperatorProtoData     // 表数据信息
	AdminActionOperatorSendProtoData // 表数据发送

	AdminActionGmCopyRole     // 角色复制
	AdminActionGmGetRank      // 获取榜单，這個也沒了吧？
	AdminActionGmTerminate    // 报错检查
	AdminActionGmAccount      // 设置GM账号
	AdminActionGmBlackAccount // 封号禁言
	AdminActionGmCharge       // 在线充值
)

const (
	AdminActionBlackListAndWhiteList = 97 // 黑白名单
	AdminActionNotice                = 98 // 公告管理
	AdminActionCDKey                 = 99 // 各种码管理

	AdminActionGmPurViewServer1 = 101 // 平台1 公告及黑白名单管理权限
	AdminActionGmPurViewServer2 = 102 // 平台2 公告及黑白名单管理权限
	AdminActionGmPurViewServer3 = 103 // 平台3 公告及黑白名单管理权限
)

//-----------------------------------------------------------
//  移植GM后台
//-----------------------------------------------------------

type GMBaseController struct {
	beego.Controller
	controllerName string //当前控制名称
	actionName     string //当前action名称
	//curUser        models.BackendUser //当前用户信息
}

type AdminData struct {
	ID       int32
	UserId   string
	Password string
	//所属组ID
	GroupId int32
	//组权限ID
	ActionList []int32
	//侧边栏
	Sidebar string
	//平台map
	PlatformMap map[int32]*beegodb.TB_CENTER_LIST
	//已连接Center
	ThisPlatformId int32
}

// 重定向
func (c *GMBaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 重定向 去错误页
func (c *GMBaseController) PageError(msg string) {
	errorurl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}

// 重定向 去登录页 20200110 linzi发现 c.Data中的内容是不生效的
func (c *GMBaseController) PageLogin() {
	url := c.URLFor("LoginController.Get")
	c.Redirect(url, 302)
	c.StopRun()
}

//返回登录页 且 显示错误信息
func (c *GMBaseController) PageLoginWitchError(errMsg string) {
	c.Data["errMsg"] = errMsg
	c.TplName = "index.html"
	//c.StopRun()

	//url := c.URLFor("LoginController.Get");

}

func MD5(tmpPwd string) string {
	m := md5.New()
	io.WriteString(m, tmpPwd)
	md5Str := fmt.Sprintf("%x", m.Sum(nil))
	return md5Str
}

func DB_Str2ListInt(dbString string) []int32 {
	retList := make([]int32, 0, 0)
	tmpStr := strings.Trim(strings.Trim(dbString, "["), "]")
	strList := strings.Split(tmpStr, ",")
	for i := 0; i < len(strList); i++ {
		tmpValue, err := strconv.ParseInt(strList[i], 10, 32)
		if err != nil {
			fmt.Println("str change to int err, the value is:", strList[i])
			continue
		}
		retList = append(retList, int32(tmpValue))
	}
	//fmt.Printf("==DB_Str2ListInt==retList:%v\n",retList)
	return retList
}

func CheckRight(userData AdminData, actionId int32) bool {
	for _, tmpAction := range userData.ActionList {
		if tmpAction == actionId {
			return true
		}
	}
	return false
}

//-----------------------------------------------------------
/*
//** TYPES

type GmResp struct {
	Value interface{} `json:"values"`
}
type CommonResp struct {
	Message int16 `json:"code"`
}

type Response map[string]interface{}

func (resp Response) Push(key string, val interface{}) {
	resp[key] = val
}

type (
	// BaseController composes all required types and behavior.
	BaseController struct {
		beego.Controller
		//services.Service
	}
)

func (baseController *BaseController) GmResponse(response interface{}) {
	baseController.Data["json"] = response
	baseController.ServeJSON()
}

//** INTERCEPT FUNCTIONS

// Prepare is called prior to the baseController method.

func (baseController *BaseController) StartSession() (sess session.Store) {
	sess, _ = beego.GlobalSessions.SessionStart(baseController.Ctx.ResponseWriter, baseController.Ctx.Request)
	return

}

func (baseController *BaseController) NewResponse() Response {
	return make(Response)
}
func (baseController *BaseController) Prepare() {
	baseController.UserID = baseController.GetString("userID")
	if baseController.UserID == "" {
		baseController.UserID = baseController.GetString(":userID")
	}
	if baseController.UserID == "" {
		baseController.UserID = "Unknown"
	}

	baseController.Data["region"] = baseController.Ctx.Input.Header("Region-Id")

	//url := baseController.Ctx.Input.URL()
	//ip := baseController.Ctx.Input.IP()
	//	if authService.IsForbidCmd(url) && misc.IsTrustedIP(ip) == false {
	//		baseController.Ctx.Output.SetStatus(protocol.CMD_FORBIDED) // 527.禁用命令
	//		baseController.ServeJson()
	//		return
	//	}

	//	if err := baseController.Service.Prepare(); err != nil {
	//		baseController.ServeError(err)
	//		return
	//	}

}

// Finish is called once the baseController method completes.
func (baseController *BaseController) Finish() {
	defer func() {
		if baseController.MongoSession != nil {
			mongo.CloseSession(baseController.UserID, baseController.MongoSession)
			baseController.MongoSession = nil
		}
	}()

}

//** VALIDATION

// ParseAndValidate will run the params through the validation framework and then
// response with the specified localized or provided message.
func (baseController *BaseController) ParseAndValidate(params interface{}) bool {
	// This is not working anymore :(
	if err := baseController.ParseForm(params); err != nil {
		baseController.ServeError(err)
		return false
	}

	var valid validation.Validation
	ok, err := valid.Valid(params)
	if err != nil {
		baseController.ServeError(err)
		return false
	}

	if ok == false {
		// Build a map of the Error messages for each field
		messages2 := make(map[string]string)

		val := reflect.ValueOf(params).Elem()
		for i := 0; i < val.NumField(); i++ {
			// Look for an Error tag in the field
			typeField := val.Type().Field(i)
			tag := typeField.Tag
			tagValue := tag.Get("Error")

			// Was there an Error tag
			if tagValue != "" {
				messages2[typeField.Name] = tagValue
			}
		}

		// Build the Error response
		var errors []string
		for _, err := range valid.Errors {
			// Match an Error from the validation framework Errors
			// to a field name we have a mapping for

			// No match, so use the message as is
			errors = append(errors, err.Message)
		}

		baseController.ServeValidationErrors(errors)
		return false
	}

	return true
}

//** EXCEPTIONS

// ServeError prepares and serves an Error exception.
func (baseController *BaseController) ServeError(err error) {
	baseController.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	baseController.Ctx.Output.SetStatus(500)
	baseController.ServeJSON()
}

// ServeError prepares and serves an Error exception.
func (baseController *BaseController) ServeRsp(err error) {
	baseController.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	baseController.Ctx.Output.SetStatus(200)
	baseController.ServeJSON()
}

// ServeValidationErrors prepares and serves a validation exception.
func (baseController *BaseController) ServeValidationErrors(Errors []string) {
	baseController.Data["json"] = struct {
		Errors []string `json:"Errors"`
	}{Errors}
	baseController.Ctx.Output.SetStatus(409)
	baseController.ServeJSON()
}

//** CATCHING PANICS

// CatchPanic is used to catch any Panic and log exceptions. Returns a 500 as the response.
func (baseController *BaseController) CatchPanic(functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		baseController.ServeError(fmt.Errorf("%v", r))
	}
}

//** AJAX SUPPORT

// AjaxResponse returns a standard ajax response.
func (baseController *BaseController) AjaxResponse(resultCode int, resultString string, data interface{}) {
	response := struct {
		Result       int
		ResultString string
		ResultObject interface{}
	}{
		Result:       resultCode,
		ResultString: resultString,
		ResultObject: data,
	}

	baseController.Data["json"] = response
	baseController.ServeJSON()
}

func (this *BaseController) Display(tpl string) {
	//this.Layout = "layout.html"
	this.TplName = tpl + ".html"
}

func (this *BaseController) FormatJson(data interface{}) string {
	jsonExpr := string("{}")
	if b, err := json.Marshal(data); err != nil {
		logger.LogErrf("json.Marshal falied: %s", err)
	} else {
		jsonExpr = string(b)
	}
	return jsonExpr
}
*/
