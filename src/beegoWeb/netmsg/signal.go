package netmsg

import (
	"gm_server/src/cfg"
	"gm_server/src/ctimer"
	"gm_server/src/logger"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

var (
	SIGTERM  int32
	SIGINT   int32
	STOP_SIG = make(chan int, 1)
)

func setQuitTimer() {
	ctimer.CreateTimer(0, 0, ctimer.CTData{Action: ctimer.ACTION_SHUTDOWN})
}

// 进程信号处理，主要捕获异常，退出之前做一些持久化处理。
func SignalProcHandler() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	for {
		msg := <-ch
		switch msg {
		case syscall.SIGHUP:
			//终端控制进程结束(终端连接断开)
			logger.Log("\033[043;1m[SIGHUP]\033[0m")
			cfg.Reload()

		case syscall.SIGTERM:
			//结束程序(可以被捕获、阻塞或忽略)
			atomic.StoreInt32(&SIGTERM, 1)
			sigQuit(msg)

		case syscall.SIGINT:
			//用户发送INTR字符(Ctrl+C)触发
			atomic.StoreInt32(&SIGINT, 1)
			sigQuit(msg)
		default:
			logger.Log("signal=", msg)
		}
	}
}

func sigQuit(msg os.Signal) {
	logger.Logf("\033[043;1m[%v, quit]\033[0m", msg)

	setQuitTimer()
	<-STOP_SIG

	os.Exit(-1)
}
