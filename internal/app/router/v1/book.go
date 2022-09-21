/*
 * @Author: kingford
 * @Date: 2022-09-21 09:50:25
 * @LastEditTime: 2022-09-21 10:09:03
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
	bookRepository repository.BookRepository = repository.NewBookRepository(global.GVA_DB)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	bookController controller.BookController = controller.NewBookController(bookService, JwtService)
)

type BookRouter struct{}

func (s *BookRouter) InitBookRouter(Router *gin.RouterGroup) {
	bookRoute := Router.Group("books")
	{
		bookRoute.GET("/", bookController.All)
		bookRoute.POST("/", bookController.Insert)
		bookRoute.GET("/:id", bookController.FindByID)
		bookRoute.PUT("/:id", bookController.Update)
		bookRoute.DELETE("/:id", bookController.Delete)
	}
}
