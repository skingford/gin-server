/*
 * @Author: kingford
 * @Date: 2022-08-29 16:26:58
 * @LastEditTime: 2022-09-21 10:40:44
 */
package main

import (
	"github.com/skingford/gin-server/config"
	"github.com/skingford/gin-server/internal/app"
	"github.com/skingford/gin-server/internal/core"
	"github.com/skingford/gin-server/internal/global"
)

func main() {
	// viper
	global.GVA_VP = core.NewViper()

	// zap
	global.GVA_LOG = core.Zap() // 初始化zap日志库

	// db
	global.GVA_DB = config.SetupDatabaseConnection()

	defer config.CloseDatabaseConnection(global.GVA_DB)

	app.RunServer()
}
