/*
 * @Author: kingford
 * @Date: 2022-09-20 20:18:55
 * @LastEditTime: 2022-09-21 10:24:37
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skingford/gin-server/internal/app/service"
	"github.com/skingford/gin-server/internal/global"
	"github.com/skingford/gin-server/internal/middleware"

	docs "github.com/skingford/gin-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var JwtService service.JWTService = service.NewJWTService()

// 初始化总路由

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	docs.SwaggerInfo.BasePath = "/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	routerGroup := RouterGroupApp

	// 无需要鉴权
	PublicGroup := v1.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		routerGroup.AuthRouter.InitAuthRouter(PublicGroup)

	}

	// 需要鉴权
	PrivateGroup := v1.Group("")
	PrivateGroup.Use(middleware.AuthorizeJWT(JwtService))
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{

		routerGroup.BookRouter.InitBookRouter(PrivateGroup)
		routerGroup.UserRouter.InitUserRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("router register success")
	return r
}
