/*
 * @Author: kingford
 * @Date: 2022-09-20 19:38:56
 * @LastEditTime: 2022-09-21 10:20:04
 */
package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/skingford/gin-server/internal/app/controller"
	"github.com/skingford/gin-server/internal/app/service"
)

var (
	authService    service.AuthService       = service.NewAuthService(UserRepository)
	authController controller.AuthController = controller.NewAuthController(authService, JwtService)
)

type AuthRouter struct{}

func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
	}
}
