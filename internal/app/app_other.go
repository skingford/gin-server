/*
 * @Author: kingford
 * @Date: 2022-09-21 10:33:05
 * @LastEditTime: 2022-09-21 10:35:53
 */
package app

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func initOtherServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
