package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//todo 1.装配配置文件
	//todo 2.装配日志

	//启动web服务
	//加载第三方中间件（自动重连）

	//要有一个管理Controller Service 的工具
	//gin.SetMode(gin.DebugMode)
	engine := gin.New()
	//engine.Use(gin.Recovery())
	//engine.SetTrustedProxies(nil)

	//注册理由
	//engine.Routes()
	_ = http.Server{
		//Handler: engine,
		//ReadTimeout:,
		//WriteTimeout: ,
		//Addr:

	}

	//使用recover功能

	//2.装配日志
	//engine.Use(gin.Logger())

}
