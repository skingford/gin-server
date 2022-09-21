/*
 * @Author: kingford
 * @Date: 2022-09-20 20:18:55
 * @LastEditTime: 2022-09-21 09:40:03
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skingford/gin-server/internal/global"

	docs "github.com/skingford/gin-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由

func Routers() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	docs.SwaggerInfo.BasePath = "/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	routerGroup := RouterGroupApp

	PublicGroup := v1.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{

		routerGroup.AuthRouter.InitAuthRouter(PublicGroup)
		// systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		// systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	// PrivateGroup := Router.Group("")
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	// {
	// 	systemRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由
	// 	systemRouter.InitJwtRouter(PrivateGroup)                 // jwt相关路由
	// 	systemRouter.InitUserRouter(PrivateGroup)                // 注册用户路由
	// 	systemRouter.InitMenuRouter(PrivateGroup)                // 注册menu路由
	// 	systemRouter.InitSystemRouter(PrivateGroup)              // system相关路由
	// 	systemRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
	// 	systemRouter.InitAutoCodeRouter(PrivateGroup)            // 创建自动化代码
	// 	systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
	// 	systemRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
	// 	systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)     // 自动化代码历史
	// 	systemRouter.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
	// 	systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
	// 	systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)  // 字典详情管理

	// 	exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
	// 	exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
	// 	exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	// }

	global.GVA_LOG.Info("router register success")
	return Router
}
