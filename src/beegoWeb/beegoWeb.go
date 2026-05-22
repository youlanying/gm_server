package beegoWeb

import (
	"fmt"
	"gm_server/src/beegoWeb/beegodb"
	"gm_server/src/beegoWeb/netmsg"
	"gm_server/src/beegoWeb/routers"
	"gm_server/src/beegoWeb/table"
	"gm_server/src/cfg"
	"gm_server/src/logger"
	"gm_server/src/logger_custom"
	"os"
	"path/filepath"
	"strconv"

	"github.com/astaxie/beego"
)

func webInit() {

	config := cfg.GetSection("BEEGOWEB")
	if config["httpport"] == "" {
		beego.Error("Web httpport config nil.")
	}

	viewDir := config["view"]
	if viewDir == "" {
		beego.Error("Web view_dir config nil.")
	}

	httpPort, err := strconv.Atoi(config["httpport"])
	if err != nil {
		beego.Error("convert httpport err", httpPort)
	}

	//session生效
	beego.BConfig.WebConfig.Session.SessionOn = true

	//开启热升级
	beego.BConfig.Listen.Graceful = true

	beego.BConfig.RunMode = "dev"

	//这里配置了 view的路径
	basePath, _ := os.Getwd()
	beego.BConfig.WebConfig.ViewsPath = basePath + "/" + viewDir
	//beego.BConfig.Listen.HTTPAddr = "192.168.20.61"
	beego.BConfig.Listen.HTTPPort = httpPort
	str, _ := os.Getwd()
	fmt.Printf("===========str:%v\n", str)
	//这里配置了 static的路径 开启静态资源,提供下载服务
	staticDir := config["static"]
	finalPath := basePath + "/" + staticDir
	beego.SetStaticPath("/static", finalPath)
}

func Main() {
	//与游戏服务器的连接（是一组游戏服务器而不是game）
	config := cfg.GetSection("LOG")
	logger.Log("-------------beego load config over---------------")
	logSize, _ := strconv.Atoi(config["log_maxsize"])
	level, _ := strconv.Atoi(config["log_level"])

	config = cfg.GetSection("BEEGOWEB")
	if config["log"] != "" {
		logger_custom.OpenDB()
		exePath, _ := os.Executable()
		logBasePath := filepath.Dir(exePath)
		logger.InitLogger(logBasePath, config["log"], int64(logSize))
		logger.SetLogLevel(level)
		logger.Logf("\n\n")
	} else {
		logger.LogFatal("[BEEGOWEB] log node is not find.")
	}

	config = cfg.GetSection("GM")
	if config["ip"] == "" || config["port"] == "" {
		logger.LogFatal("GM bip not find.")
	}

	webInit()
	netmsg.InitSession()

	isLoad := table.LoadAllTable()
	if !isLoad {
		os.Exit(0)
	}

	go netmsg.SignalProcHandler()

	netmsg.RegisterInit()
	//初始化数据库信息
	beegodb.OpenDB()
	netmsg.InitServerManager()
	routers.InitRouters()
	beego.Run()

}
