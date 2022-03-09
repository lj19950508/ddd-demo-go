package main

import (
	"ddd-demo1/infrastructrue/di"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	engine := gin.New()
	//engine.SetTrustedProxies(nil)

	//todo 1.装配配置文件
	//engine.Use(middleware.Run())
	//使用recover功能
	engine.Use(gin.Recovery())
	//2.装配日志
	engine.Use(gin.Logger())
	//3.生产装配bean 与路由
	facotry := di.NewBeanFactory()
	engine.Use(facotry.Run())
	engine.Use(di.NewRouteRigster(facotry.GetRoutes()).Run(engine))

	engine.Use()

	//4. todo使用router工厂配合bean的controller 注册路由
	//装配过滤器？ 可能依赖于group 先不做
	//装配数据库
	//装配第三方组件
	//engine.

	engine.Run()

}
