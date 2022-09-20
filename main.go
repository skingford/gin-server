/*
 * @Author: kingford
 * @Date: 2022-08-29 16:26:58
 * @LastEditTime: 2022-09-20 10:11:01
 */
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/skingford/gin-server/config"
	"github.com/skingford/gin-server/controller"
	"github.com/skingford/gin-server/core"
	"github.com/skingford/gin-server/global"
	"github.com/skingford/gin-server/middleware"
	"github.com/skingford/gin-server/repository"
	"github.com/skingford/gin-server/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {
	global.GVA_VP = core.NewViper()

	fmt.Println("global config viper:", global.GVA_CONFIG.System.DbType)

	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	v1 := r.Group("/v1")

	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := v1.Group("/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	bookRoutes := v1.Group("/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	r.Run(":8080")
}
