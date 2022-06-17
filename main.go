package main

import (
	"ddd-demo-go/adapter/in/web"
	"ddd-demo-go/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	//load config from yml and env  (not a pointer)

	//var cfg2 infrastructure.Config
	//container.Resolve(&cfg2)
	//container.Call(func(cfg infrastructure.Config) {
	//
	//})
	//set logger
	infrastructure.Produce()

	//转换

	//web server do
	var webEngine = gin.New()
	webEngine.Use(gin.Logger())
	webEngine.Use(gin.Recovery())
	//error deal ， 只是为了稳定，千万别随便panic异常。 异常 return到controller 处理
	//webEngine.Use(ThreadLocal)
	web.InitRoute(webEngine)

}
