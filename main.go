/*
 * @Author: kingford
 * @Date: 2022-08-29 16:26:58
 * @LastEditTime: 2022-09-20 20:18:02
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skingford/gin-server/config"
	"github.com/skingford/gin-server/internal/app/controller"
	"github.com/skingford/gin-server/internal/app/service"
	"github.com/skingford/gin-server/internal/core"
	"github.com/skingford/gin-server/internal/global"
	"github.com/skingford/gin-server/internal/middleware"
	"github.com/skingford/gin-server/internal/repository"

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
	// viper
	global.GVA_VP = core.NewViper()

	// zap
	global.GVA_LOG = core.Zap() // 初始化zap日志库

	// db
	global.GVA_DB = config.SetupDatabaseConnection()

	global.GVA_LOG.Info("main init log complete:" + global.GVA_CONFIG.System.DbType)

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
