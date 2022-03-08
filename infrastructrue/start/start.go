package main

import (
	"ddd-demo1/infrastructrue/bean_factory"
	"ddd-demo1/infrastructrue/route_register"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	//todo 1.装配配置文件
	engine := gin.New()
	engine.SetTrustedProxies(nil)
	//使用recover功能
	engine.Use(gin.Recovery())
	//2.装配日志
	engine.Use(gin.Logger())

	facotry := bean_factory.NewBeanFactory()
	//3. todo 使用bean工厂生产bean与装配基本bean 单例
	engine.Use(facotry.Run())
	engine.Use(route_register.NewRouteRigster(facotry.GetRoutes()).Run(engine))
	//4. todo使用router工厂配合bean的controller 注册路由
	//装配过滤器？ 可能依赖于group 先不做
	//装配数据库
	//装配第三方组件
	//engine.

	engine.Run()
}
