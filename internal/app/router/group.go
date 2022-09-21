/*
 * @Author: kingford
 * @Date: 2022-09-20 20:30:53
 * @LastEditTime: 2022-09-21 10:21:32
 */
package router

import (
	v1 "github.com/skingford/gin-server/internal/app/router/v1"
)

type RouterGroup struct {
	AuthRouter v1.AuthRouter
	BookRouter v1.BookRouter
	UserRouter v1.UserRouter
}

var RouterGroupApp = new(RouterGroup)
