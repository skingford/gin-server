/*
 * @Author: kingford
 * @Date: 2022-09-19 20:06:52
 * @LastEditTime: 2022-09-20 10:26:13
 */
package global

import (
	"sync"

	"github.com/skingford/gin-server/config"
	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB

	GVA_CONFIG config.Config

	GVA_VP *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger

	GVA_Concurrency_Control = &singleflight.Group{}

	lock sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbName string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbName]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbName string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbName]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
