/*
 * @Author: kingford
 * @Date: 2022-09-20 20:30:53
 * @LastEditTime: 2022-09-20 20:32:39
 */
package router

import v1 "github.com/skingford/gin-server/internal/app/router/v1"

type RouterGroup struct {
	AuthRouter v1.AuthRouter
}

var RouterGroupApp = new(RouterGroup)
