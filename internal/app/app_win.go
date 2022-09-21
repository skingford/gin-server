/*
 * @Author: kingford
 * @Date: 2022-09-21 10:32:42
 * @LastEditTime: 2022-09-21 10:35:44
 */

package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initWinServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
