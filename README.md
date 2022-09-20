<!--
 * @Author: kingford
 * @Date: 2022-08-29 16:31:14
 * @LastEditTime: 2022-09-20 10:46:57
-->
# gin-server

## 启动项目

1. 修改.env 配置文件

```base
DB_USER = postgres
DB_PASS = fD52ki3qPbH21ORY6QyZYt6CCdUC
DB_HOST = 127.0.0.1
DB_PORT = 5432
DB_NAME = gin-server
```

2. 安装包

```bash
# 配置代理
go env -w GOPROXY=https://goproxy.cn 

go mod tidy
```

3. 运行项目

```bash
go run main.go

```

## git

 git清除本地缓存

```bash
git rm -r --cached .
```

## 推荐阅读

- [gorm.io](https://gorm.io/zh_CN/docs/create.html)
- [gin](https://gin-gonic.com/zh-cn/docs/)
- 日志：使用[zap](https://github.com/uber-go/zap) 实现日志记录。
- 配置文件：[viper](https://github.com/spf13/viper) 和 [fsnotify](https://github.com/fsnotify/fsnotify) 实现yaml格式的配置文件
- [jwt]()和[casbin]()
- API文档：使用Swagger构建自动化文档
