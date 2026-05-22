package controllers

import (
	"bufio"
	"fmt"
	"gm_server/src/cfg"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

//---- 日志查询 ----

const (
	/**
	 * 登陆打点类型 和 存放的 文件夹名字
	 */
	SignPointType1 = "clientlogin"

	/**
	 * 崩溃报错 打点类型 和 存放的 文件夹名字
	 */
	SignPointType2 = "clientbug"

	/**
	 * 普通打点类型 和 存放的 文件夹名字
	 */
	SignPointType3 = "clientsignpoint"
)

/**
 * 运营相关 邮件与跑马灯
 */
type LogController struct {
	GMBaseController
}

/**
 * 常规显示列表
 */
func (c *LogController) ShowLog() {

	v := c.GetSession("loginuser")

	if v == nil {

		//销毁全部的session
		c.DestroySession()

		fmt.Println("销毁后的session:")
		fmt.Println(c.CruSession)

		c.PageLoginWitchError("身份过期")
		return
	}

	c.RenderLog(v, "")
}

var waitGroup *sync.WaitGroup
var loginArr []string
var bugArr []string
var pointArr []string

/**
 * 渲染日志列表
 */
func (c *LogController) RenderLog(adminName interface{}, Res string) {

	if waitGroup != nil {
		return
	}

	waitGroup = &sync.WaitGroup{}
	waitGroup.Add(3)

	go ReadMain(SignPointType1)
	go ReadMain(SignPointType2)
	go ReadMain(SignPointType3)

	waitGroup.Wait()

	waitGroup = nil

	c.Data["pagetitle"] = "日志查询"
	c.Data["admin"] = adminName
	userData := c.GetSession("UserData").(AdminData)
	c.Data["candoact"] = CheckRight(userData, AdminActionOperatorSendMail)
	c.Data["audit_res"] = "查询成功"
	c.Data["loginArr"] = loginArr
	c.Data["bugArr"] = bugArr
	c.Data["pointArr"] = pointArr

	c.Layout = "basetemplate/basetemplate.html"
	c.TplName = "log/index.html"
}

func ReadMain(vtype string) {

	loginFile := cfg.GetBasePath() + "/log/" + vtype

	files, err := ioutil.ReadDir(loginFile) //读取目录下文件
	if err != nil {
		fmt.Println("read Dir error", err.Error())
		waitGroup.Done()
		return
	}

	vLen := len(files)

	fmt.Println("文件数量", vLen)

	// 初始化一些channel，用于数据传递
	logChannel := make(chan string, 10) //读取日志文件量更大，设置为3倍

	go DoRead(vtype, logChannel)

	for _, file := range files {

		vLen--

		name := file.Name()

		ReadOneTxt(loginFile, name, vtype, logChannel, vLen <= 0)
	}
}

func ReadOneTxt(url string, name string, vtype string, logChannel chan string, isEnd bool) {

	urlStr := url + "/" + name

	// 打日志
	fd, err := os.Open(urlStr) //打开go生成的日志
	defer fd.Close()           //关闭文件

	if err != nil {
		fmt.Println("read file error", err)
		waitGroup.Done()
		return
	}

	bufferRead := bufio.NewReader(fd)

	for {
		line, err := bufferRead.ReadString('\n') //一行行读

		if err != nil { //error部位空有两种情况，一种是错误，一种是读到尾部了
			if err == io.EOF { //读到尾部了(读完了)，休息3秒钟
				fmt.Println("readline:Finsh") //提醒在等待，已经读到了第n行

				if len(line) > 0 {
					logChannel <- line //读出一行写入一次logChannel
				}

			} else {
				fmt.Println("readline error", err) //错误则打出错误
			}

			if isEnd {
				close(logChannel)
			}
			break
		}

		logChannel <- line //读出一行写入一次logChannel
	}
}

func DoRead(vtype string, logChannel chan string) {

	var realArr []string

	for line := range logChannel {
		realArr = append(realArr, line)
	}

	if vtype == SignPointType1 {
		loginArr = realArr
	} else if vtype == SignPointType2 {
		bugArr = realArr
	} else if vtype == SignPointType3 {
		pointArr = realArr
	} else {
		return
	}

	fmt.Println(vtype, "len=", len(realArr))

	waitGroup.Done()
}
