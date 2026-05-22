package main

import (
	"gm_server/src/beegoWeb"
	"gm_server/src/logger"
)

func main() {
	logger.Log("-----------------beego start---------------")
	beegoWeb.Main()
}
