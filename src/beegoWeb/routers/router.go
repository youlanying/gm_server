package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gm_server/src/beegoWeb/controllers"
)

func InitRouters() {

	//挪到 beegoWeb下面了 https://192.168.19.2:8443/svn/ZXLF_NWGD/goServer/beegoWeb/static
	//beego.SetStaticPath("/static", "gm_server/beegoWeb/static") //第一个是访问的路径，第二个是根下目录

	//-- 登录 --
	beego.Router("/", &controllers.GMBaseController{}, "get:Get")
	beego.Router("/logout", &controllers.GMBaseController{}, "get:Logout")
	beego.Router("/login", &controllers.GMBaseController{}, "post:Post")
	beego.Router("/SelectPlatform", &controllers.GMBaseController{}, "get:SelectPlatform")
	//-- 用户管理 --
	beego.Router("/admin/showgroup", &controllers.AdminController{}, "get:Showgroup")
	beego.Router("/admin/newgroup", &controllers.AdminController{}, "post:Newgroup")
	beego.Router("/admin/deletegroup", &controllers.AdminController{}, "post:Deletegroup")
	beego.Router("/admin/editgroup", &controllers.AdminController{}, "post:Editgroup")
	beego.Router("/admin/doeditgroup", &controllers.AdminController{}, "post:DoEditgroup")

	beego.Router("/admin/index", &controllers.AdminController{}, "get:Index")
	beego.Router("/admin/showuser", &controllers.AdminController{}, "get:ShowUser")
	beego.Router("/admin/newuser", &controllers.AdminController{}, "post:NewUser")
	beego.Router("/admin/deleteuser", &controllers.AdminController{}, "post:DeleteUser")
	beego.Router("/admin/edituser", &controllers.AdminController{}, "post:EditUser")
	beego.Router("/admin/doedituser", &controllers.AdminController{}, "post:DoEditUser")

	//-- 平台管理 --
	beego.Router("/pfselect/index", &controllers.PfSelectController{}, "get:Index")
	beego.Router("/pfselect/ajaxindex", &controllers.PfSelectController{}, "post:AjaxIndex")
	beego.Router("/pfselect/update", &controllers.PfSelectController{}, "post:Update")
	beego.Router("/pfselect/createnew", &controllers.PfSelectController{}, "post:CreateNew")
	beego.Router("/pfselect/delete", &controllers.PfSelectController{}, "post:Delete")

	//-- 公告管理 --
	beego.Router("/notice/index", &controllers.NoticeController{}, "get:Index")
	beego.Router("/notice/add", &controllers.NoticeController{}, "post:Add")
	beego.Router("/notice/del", &controllers.NoticeController{}, "post:Del")
	beego.Router("/notice/audit", &controllers.NoticeController{}, "post:Audit")

	//-- 黑白名单 --
	beego.Router("/whiteblack/top", &controllers.WhiteBlackController{}, "get:Top")
	beego.Router("/whiteblack/index", &controllers.WhiteBlackController{}, "get:Index")
	beego.Router("/whiteblack/show", &controllers.WhiteBlackController{}, "get:Show")
	beego.Router("/whiteblack/create", &controllers.WhiteBlackController{}, "post:Create")
	beego.Router("/whiteblack/delete", &controllers.WhiteBlackController{}, "post:Delete")

	//-- 各种码管理 --
	beego.Router("/number/top", &controllers.NumberController{}, "get:Number")
	beego.Router("/number/index", &controllers.NumberController{}, "get:Number")
	beego.Router("/number/number", &controllers.NumberController{}, "get:Number")

	beego.Router("/number/newnumber", &controllers.NumberController{}, "get:NewNumber;post:NewNumber")
	beego.Router("/number/newgift", &controllers.NumberController{}, "get:NewGift;post:NewGift")

	//-- GM TOOLS --
	//beego.Router("/gmtools/dictionary", &controllers.GMToolsController{}, "get:DictionaryGet;post:DictionaryPost") //注意这里是;
	//beego.Router("/gmtools/redis", &controllers.GMToolsController{}, "get:RedisGet;post:RedisPost")
	beego.Router("/gmtools/userData", &controllers.GMToolsController{}, "get:UserDataGet;post:UserDataPost")
	beego.Router("/gmtools/servernum", &controllers.GMToolsController{}, "get:ServerNumGet")
	beego.Router("/gmtools/kick", &controllers.GMToolsController{}, "get:KickGet;post:KickPost")

	beego.Router("/gmtools/servertime", &controllers.GMToolsController{}, "get:ServerTimeGet;post:ServerTimePost")
	beego.Router("/gmtools/charge", &controllers.GMToolsController{}, "get:ChargeGet;post:ChargePost")
	beego.Router("/gmtools/copyrole", &controllers.GMToolsController{}, "get:CopyRoleGet;post:CopyRolePost")
	beego.Router("/gmtools/terminate", &controllers.GMBaseController{}, "get:TerminateGet;post:TerminatePost")
	beego.Router("/gmtools/gm_account", &controllers.GMToolsController{}, "get:GMAccountGet;post:GMAccountPost")
	beego.Router("/gmtools/black_account", &controllers.GMToolsController{}, "get:BlackAccountGet;post:BlackAccountPost")
	beego.Router("/gmtools/setcreate", &controllers.GMToolsController{}, "get:SetCreateGet;post:SetCreatePost")
	beego.Router("/gmtools/delitem", &controllers.GMToolsController{}, "get:DelItemGet;post:DelItemPost")

	//-- new 打点与统计 --
	beego.Router("/gmtools/logincount", &controllers.StatisticsController{}, "get:LoginGet;post:LoginPost")
	beego.Router("/gmtools/register", &controllers.StatisticsController{}, "get:RegisterGet")

	//-- 运营相关 跑马灯 --
	beego.Router("/operator/send_marquee", &controllers.OperatorController{}, "get:Send_marquee")
	beego.Router("/operator/send_marquee", &controllers.OperatorController{}, "post:Send_marqueePost")
	beego.Router("/operator/audit_marquee", &controllers.OperatorController{}, "post:Audit_marquee")
	beego.Router("/operator/delete_marquee", &controllers.OperatorController{}, "get:DeleteMarquee")

	//-- 运营相关 邮件 --
	beego.Router("/operator/send_mail", &controllers.OperatorController{}, "get:ShowMail")
	beego.Router("/operator/send_mail", &controllers.OperatorController{}, "post:Send_mailPost")
	beego.Router("/operator/quick_createmails", &controllers.OperatorController{}, "post:QuickCreateMails")
	beego.Router("/operator/audit_mail", &controllers.OperatorController{}, "post:Audit_mail")
	beego.Router("/operator/delete_mail", &controllers.OperatorController{}, "get:DeleteMail")

	//-- 运营相关 日志查询 --

	beego.Router("/log/index", &controllers.LogController{}, "get:ShowLog")

	//-- 玩家信息查询 --
	beego.Router("/info/index", &controllers.InfoController{}, "get:Index")
	beego.Router("/info/mute", &controllers.InfoController{}, "get:Mute")
	beego.Router("/info/seal", &controllers.InfoController{}, "get:Seal")
	beego.Router("/info/bag", &controllers.InfoController{}, "get:Bag")
	beego.Router("/info/friend", &controllers.InfoController{}, "get:Friend")
	beego.Router("/info/task", &controllers.InfoController{}, "get:Task")
	beego.Router("/info/mail", &controllers.InfoController{}, "get:Mail")
	beego.Router("/info/strength", &controllers.InfoController{}, "get:Strength")
	beego.Router("/info/hero", &controllers.InfoController{}, "get:Hero")

	beego.Router("/info/index", &controllers.InfoController{}, "get:Index")
	beego.Router("/info/index2", &controllers.InfoController{}, "get:Index2")
	beego.Router("/info/search", &controllers.InfoController{}, "get:Search;post:Search")
	beego.Router("/info/change", &controllers.InfoController{}, "get:Change;post:Change")
	beego.Router("/info/muteoneusr", &controllers.InfoController{}, "get:MuteOneUsr;post:MuteOneUsr")
	beego.Router("/info/sealoneusr", &controllers.InfoController{}, "get:SealOneUsr;post:SealOneUsr")
	beego.Router("/info/muteoneusr2", &controllers.InfoController{}, "get:MuteOneUsr2;post:MuteOneUsr2")
	beego.Router("/info/sealoneusr2", &controllers.InfoController{}, "get:SealOneUsr2;post:SealOneUsr2")

	beego.Router("/info/ask", &controllers.InfoController{}, "get:Ask;post:Ask")
	beego.Router("/info/kick", &controllers.InfoController{}, "get:Kick;post:Kick")
	beego.Router("/info/bagdelete", &controllers.InfoController{}, "get:BagDelete;post:BagDelete")
	beego.Router("/info/taskdelete", &controllers.InfoController{}, "get:TaskDelete;post:TaskDelete")
	beego.Router("/info/maildelete", &controllers.InfoController{}, "get:MailDelete;post:MailDelete")
	beego.Router("/info/strengthupdate", &controllers.InfoController{}, "get:StrengthUpdate;post:StrengthUpdate")

	//-- 运营相关 添加表 --
	beego.Router("/operator/add_proto", &controllers.TableController{}, "get:AddProtoGet")
	beego.Router("/operator/add_proto", &controllers.TableController{}, "post:AddProtoPost")
	beego.Router("/operator/delete_proto", &controllers.TableController{}, "get:AddProtoDelete")
	beego.Router("/operator/up_add_proto", &controllers.TableController{}, "post:AddProtoUpdate")

	//-- 运营相关 表数据添加 --
	beego.Router("/operator/up_proto", &controllers.TableController{}, "get:UpProtoGet")
	beego.Router("/operator/select_table", &controllers.TableController{}, "post:UpProtoPost")
	beego.Router("/operator/add_proto_data", &controllers.TableController{}, "post:UpProtoCheck")

	//-- 运营相关 表数据信息 --
	beego.Router("/operator/send_proto", &controllers.TableController{}, "get:SendProtoGet")
	beego.Router("/operator/send_proto", &controllers.TableController{}, "post:SendProtoPost")
	beego.Router("/operator/deletetable", &controllers.TableController{}, "post:SendProtoDeleteTable")
	beego.Router("/operator/delete_proto_data", &controllers.TableController{}, "get:SendProtoDeleteTableData")

	//-- 获得排行 --
	beego.Router("/operator/get_rank", &controllers.OperatorController{}, "get:GetRank")
	beego.Router("/operator/get_rank", &controllers.OperatorController{}, "post:GetRankPost")

	//-- 版本管理 --
	beego.Router("/version/server_version", &controllers.VersionController{}, "get:Show")
	beego.Router("/version/new_server", &controllers.VersionController{}, "post:NewServerPOST")
	beego.Router("/version/update_server", &controllers.VersionController{}, "get:UpdateServerGET;post:UpdateServerPOST")
	beego.Router("/version/delete_server", &controllers.VersionController{}, "get:DeleteServerGET;post:DeleteServerPOST")

	beego.Router("/version/new_version", &controllers.VersionController{}, "post:NewVersionPOST")
	beego.Router("/version/update_version", &controllers.VersionController{}, "post:UpdateVersionPOST")

	beego.Router("/version/server_version_log", &controllers.VersionController{}, "post:ServerVersionLog")

	fmt.Println("====beego.Router===ok")
}
