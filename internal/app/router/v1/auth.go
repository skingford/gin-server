/*
 * @Author: kingford
 * @Date: 2022-09-20 19:38:56
 * @LastEditTime: 2022-09-20 20:31:54
 */
package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/skingford/gin-server/internal/app/controller"
	"github.com/skingford/gin-server/internal/app/service"
	"github.com/skingford/gin-server/internal/global"
	"github.com/skingford/gin-server/internal/repository"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository(global.GVA_DB)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

type AuthRouter struct{}

func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
	}
}
