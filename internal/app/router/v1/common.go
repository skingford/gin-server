/*
 * @Author: kingford
 * @Date: 2022-09-21 09:55:22
 * @LastEditTime: 2022-09-21 10:19:47
 */
package v1

import (
	"github.com/skingford/gin-server/internal/app/service"
	"github.com/skingford/gin-server/internal/global"
	"github.com/skingford/gin-server/internal/repository"
)

var (
	JwtService     service.JWTService        = service.NewJWTService()
	UserRepository repository.UserRepository = repository.NewUserRepository(global.GVA_DB)
)
