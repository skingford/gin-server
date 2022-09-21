/*
 * @Author: kingford
 * @Date: 2022-09-21 10:24:05
 * @LastEditTime: 2022-09-21 10:44:11
 */
package app

import (
	"fmt"
	"runtime"

	"github.com/skingford/gin-server/internal/app/router"
	"github.com/skingford/gin-server/internal/global"
)

type server interface {
	ListenAndServe() error
}

var s server

func RunServer() {
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	r := router.SetupRouter()

	sysType := runtime.GOOS

	if sysType == "windows" {
		// windows系统
		s = initWinServer(address, r)
	} else {
		s = initOtherServer(address, r)
	}

	fmt.Printf(`
		默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
		默认前端文件运行地址:http://127.0.0.1:8080
	`, address)

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
