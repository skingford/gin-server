/*
 * @Author: kingford
 * @Date: 2022-09-21 09:50:25
 * @LastEditTime: 2022-09-21 10:20:29
 */
package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/skingford/gin-server/internal/app/controller"
	"github.com/skingford/gin-server/internal/app/service"
)

var (
	userService    service.UserService       = service.NewUserService(UserRepository)
	userController controller.UserController = controller.NewUserController(userService, JwtService)
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRoutes := Router.Group("users")
	{
		userRoutes.GET("/:id/profile", userController.Profile)
		userRoutes.PUT("/:id/profile", userController.Update)
	}
}
