package controllers

import (
	"archive/zip"
	"bytes"
	"fmt"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/models"
	"gm_server/src/beegoWeb/netmsg"
	"gm_server/src/cfg"
	"gm_server/src/logger"
	network_message "gm_server/src/network/message"
	"gm_server/src/tools"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type VersionController struct {
	GMBaseController
}

//检查超时退出
func (c *VersionController) checkTimeOut() bool {
	v := c.GetSession("loginuser")
	if v == nil {
		//销毁全部的session
		c.DestroySession()
		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)
		c.PageLoginWitchError("身份过期")
		return false
	}
	return true
}

/**
 * server_version
 */
func (c *VersionController) Show() {
	//读取本地记录
	models.InitVerison()
	if !c.checkTimeOut() {
		return
	}
	c.RenderVersion("", nil, nil)
}
func (c *VersionController) RenderVersion(Res string, Updateres interface{}, Modelres interface{}) {

	fmt.Println("RenderVersion")

	reVs := renderDirServerVersion()

	userData := c.GetSession("UserData").(AdminData)

	var reList []models.ServerVersion

	center := userData.PlatformMap[userData.ThisPlatformId]
	newMem := models.ServerVersion{
		Id:        center.ID,
		CName:     center.NAME,
		Dir:       center.SERVERPATH,
		VersionId: models.Proc_get_version(center.ID),
	}

	reList = append(reList, newMem)

	c.Data["pagetitle"] = "版本管理"
	c.Data["serverinfo"] = reList
	c.Data["allversion"] = reVs //显示所有提交的版本

	c.Data["candoact"] = CheckRight(userData, AdminActionVersionManager)
	c.Data["updateres"] = Updateres //是否弹出更新列表
	c.Data["modelres"] = Modelres   //弹出的更新列表
	c.Data["Res"] = Res

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "version/server_version.html"
}

func renderDirServerVersion() []string {
	var baseDir string
	// 保存文件, 本地文件路径static/upload/上传文件名
	// 需要提前创建好static/upload目录
	if cfg.GetBasePath() == "." {
		baseDir = "./"
	} else {
		baseDir = cfg.GetBasePath() + "/"
	}
	SERVERVERSION_DIR := baseDir + "beegoWeb/server_version/"

	reVs := models.GetFileChildNames(SERVERVERSION_DIR)
	return reVs
}

func (c *VersionController) NewServerPOST() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//ServerId, _ := c.GetInt("serverid")
	//ServerName := c.GetString("servername")
	//ServerIp := c.GetString("serverip")
	//ServerDir := c.GetString("serverdir")
	//Severurl := c.GetString("severurl")

	//tb_server_version:insert(ServerId, ServerName, ServerIp, ServerDir, Severurl)

	//orm := orm.NewOrm()
	//
	//var mem models.Tb_server_version
	//mem.Serverid = ServerId
	//mem.ServerName = ServerName
	//mem.ServerIp = ServerIp
	//mem.ServerDir = ServerDir
	//mem.Severurl = Severurl

	//id, err := orm.Insert(&mem)
	//if err == nil {
	//	fmt.Println("insert succ:", id)
	//} else {
	//	fmt.Println("insert fail:", err)
	//}

	//gm_logger_op:version_new_server(UserId,ServerId, ServerName, ServerIp, ServerDir, Severurl)
	//{redirect, [{action, "server_version"}]}
	c.RenderVersion("", nil, nil)
}

func (c *VersionController) UpdateServerGET() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	ServerId, _ := c.GetInt("serid")
	ServerName := c.GetString("servername")
	ServerIp := c.GetString("serverip")
	ServerDir := c.GetString("serverdir")
	Severurl := c.GetString("severurl")

	fmt.Println(ServerId, ServerName, ServerIp, ServerDir, Severurl)

	//mailMem := models.Tb_server_version{Serverid: ServerId}

	//orm := orm.NewOrm()
	//err := orm.Read(&mailMem)

	//if err != nil {
	//	c.RenderVersion("找不到对应记录", nil, nil)
	//	return
	//} //是不是没有的话 要改成插入？
	//
	//mailMem.ServerName = ServerName
	//mailMem.ServerIp = ServerIp
	//mailMem.ServerDir = ServerDir
	//mailMem.Severurl = Severurl

	//_, err2 := orm.Update(&mailMem)
	//if err2 == nil {
	//	fmt.Println("update_server成功")
	//} else {
	//	fmt.Println("update_server报错")
	//}

	//io:format("update_server ############### ~p ~n", [{Id, Servername, Serverip, Serverdir, Severurl}]),
	//tb_server_version:write(Id, Servername, Serverip, Serverdir, Severurl)
	//{redirect, [{action, "server_version"}]}

	c.RenderVersion("", nil, nil)
}

/**
 * 修改一条 服务器信息
 */
func (c *VersionController) UpdateServerPOST() {

	fmt.Println("Update_serverPOST")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	centerId := c.GetString("id")
	ServerDir := c.GetString("dir")
	fmt.Println("UpdateServerPOST ", centerId, ServerDir)
	upNum, ok := beegodb.TB_CENTER_LISTUpdateByKey("id="+centerId, "serverpath='"+ServerDir+"'")
	fmt.Println("UpdateServerPOST ", upNum, ok)
	if !ok {
		c.Ctx.WriteString("false")
		return
	}
	beegodb.InitTbCenterList()
	c.Ctx.WriteString("ok")
}

func (c *VersionController) DeleteServerGET() {

	fmt.Println("Delete_serverGET")

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

	fmt.Println("deleteID = " + strconv.Itoa(Did))

	//models.DeleteFunNew("", "tb_server_version", "serverid", Did)

	//tb_server_version:delete_by_serverid(Id)
	//{redirect, [{action, "server_version"}]}

	c.RenderVersion("", nil, nil)
}

func (c *VersionController) DeleteServerPOST() {

	fmt.Println("Delete_serverPOST")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	Did, _ := c.GetInt("delid")

	fmt.Println("deleteID = " + strconv.Itoa(Did))

	//models.DeleteFunNew("", "tb_server_version", "serverid", Did)
	//tb_server_version:delete_by_serverid(Id)
	//gm_logger_op:delete_server(UserId, Id)
	//{redirect, [{action, "server_version"}]}
	c.RenderVersion("", nil, nil)
}

/**
 * 创建并上传 version 版本
 */
func (c *VersionController) NewVersionPOST() {
	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}
	// 读取文件信息
	f, h, err := c.GetFile("upfile")

	if err != nil {
		log.Fatal("读取文件错误", err)
	}
	result := isZip(f)
	if result != true {
		c.PageLoginWitchError("非ZIP文件")
		// 延迟关闭文件
		defer f.Close()
		return
	}
	// 延迟关闭文件
	defer f.Close()

	var baseDir string
	// 保存文件, 本地文件路径static/upload/上传文件名
	// 需要提前创建好static/upload目录
	if cfg.GetBasePath() == "." {
		baseDir = ""
	} else {
		baseDir = cfg.GetBasePath() + "/"
	}
	c.SaveToFile("upfile", baseDir+"beegoWeb/server_version/zipversion/"+h.Filename)
	filePath := baseDir + "beegoWeb/server_version/"
	fmt.Println("server_version/zipversion/" + h.Filename)

	UnZip(filePath, baseDir+"beegoWeb/server_version/zipversion/"+h.Filename)
	reVs := renderDirServerVersion()
	//gm_logger_op:new_version(UserId, FileName)
	htmlVersion := ""
	for i := 0; i < len(reVs); i++ {
		htmlVersion = "<option value='" + reVs[i] + "'>" + reVs[i] + "</option>" + htmlVersion
	}
	c.Ctx.WriteString(htmlVersion)
}

/**
 * 更新 version 版本
 */
func (c *VersionController) UpdateVersionPOST() {

	fmt.Println("UpdateServerVersion 1")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	SelVersion := c.GetString("selversion")
	strUpdateServerId := c.GetString("updateserverid")
	UpdateServerId, _ := c.GetInt("updateserverid")

	logger.Logf("UpdateServerVersion 2 SelVersion:%v, UpdateServerId:%v", SelVersion, UpdateServerId)

	var baseDir string
	if cfg.GetBasePath() == "." {
		baseDir = "./"
	} else {
		baseDir = cfg.GetBasePath() + "/"
	}

	userData := c.GetSession("UserData").(AdminData)
	thisCenterId := userData.ThisPlatformId
	centerData := userData.PlatformMap[thisCenterId]
	strIp := checkIp(centerData.IP)
	command := "./copy_up.escript"
	cmd := exec.Command("escript", command, "-id", strUpdateServerId, "-serverip", strIp, "-serverdir", centerData.SERVERPATH, "-serververion", SelVersion)
	logger.Log("UpdateServerVersion cmd ", cmd)
	logger.Logf("UpdateServerVersion 3 %v", baseDir)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = baseDir + "beegoWeb/"
	err := cmd.Run()
	if err != nil {
		logger.Log(fmt.Sprint(err) + ": " + stderr.String())
		logger.Log("-----------------error------------------")
		logger.Log(err.Error(), out.String())
		c.RenderVersion("更新Version失败 命令运行失败", nil, nil)
		return
	}
	logger.Log("Result: " + out.String())
	logger.Log("UpdateServerVersion 4")

	sessionId := netmsg.NewSession()
	netmsg.SendMsgToGMServer(userData.ThisPlatformId, &network_message.GM_UpdateVersionRequest{
		SessionId: sessionId,
		Id:        SelVersion,
	})
	ret := netmsg.RecMsg(sessionId).(network_message.GM_UpdateVersionResponse)
	if ret.State != 1 {
		c.Ctx.WriteString("更新Version失败 httpGet失败")
		return
	}
	logger.Log("UpdateServerVersion 8")

	//vlog := &models.Tb_server_version_log{}
	//vlog.Serverid, _ = strconv.Atoi(UpdateServerId)
	//vlog.Serverversion = SelVersion
	//vlog.Versiondate = time.Now()
	//
	//orm := orm.NewOrm()
	//id, err := orm.Insert(vlog)
	//if err == nil {
	//	logger.Log("insert succ:", id)
	//} else {
	//	logger.Log("insert fail:", err)
	//}

	logger.Log("UpdateServerVersion request over")
	c.Ctx.WriteString("更新成功~")

	//gm_logger_op:update_version(UserId, UpdateServerId, SelVersion, UpdateServerUrl)
}

type LogStruct struct {
	ServerId    int
	ServerName  string
	TmpVersion  string
	VersionDate string
}

func (c *VersionController) ServerVersionLog() {

	fmt.Println("ServerVersionLog")

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	//ServerId, _ := c.GetInt("versionserverid")
	//ServerName := c.GetString("versionservername")
	//SelectSql := "select * from tb_server_version_log WHERE serverid =" + strconv.Itoa(ServerId) + " ORDER BY id DESC LIMIT 0,30"
	//
	//
	//var AllTbVersionMap []LogStruct
	//for _, dbMem := range users {
	//
	//	var mem LogStruct
	//	mem.ServerId = dbMem.Serverid
	//	mem.ServerName = ServerName
	//	mem.TmpVersion = dbMem.Serverversion
	//	mem.VersionDate = models.TimeToString(dbMem.Versiondate)
	//
	//	AllTbVersionMap = append(AllTbVersionMap, mem)
	//}
	//
	//c.Data["versionlog"] = AllTbVersionMap

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "version/server_version_log.html"
}

func GET(url string) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	//Add 头协议
	//reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	//reqest.Header.Add("Connection", "keep-alive")
	//reqest.Header.Add("Cookie", "设置cookie")
	//reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	response, err := client.Do(reqest) //提交
	defer response.Body.Close()
	cookies := response.Cookies() //遍历cookies
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		// handle error
	}
	fmt.Println(string(body)) //网页源码

}

//解压缩文件
func UnZip(dst, src string) (err error) {
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return
	}

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}

		// 获取到 Reader
		fr, err := file.Open()
		if err != nil {
			return err
		}

		// 创建要写出的文件对应的 Write
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		n, err := io.Copy(fw, fr)
		if err != nil {
			return err
		}

		// 将解压的结果输出
		fmt.Printf("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)

		// 因为是在循环中，无法使用 defer ，直接放在最后
		// 不过这样也有问题，当出现 err 的时候就不会执行这个了，
		// 可以把它单独放在一个函数中，这里是个实验，就这样了
		fw.Close()
		fr.Close()
	}
	return nil
}

//检查zip文件
func isZip(f multipart.File) bool {
	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}
	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

//检查是否为本机IP
func checkIp(ip string) string {
	ipList := tools.GetLocalInternalIpList()
	for i := 0; i < len(ipList); i++ {
		if ipList[i] == ip {
			return "127.0.0.1"
		}
	}
	return ip
}
