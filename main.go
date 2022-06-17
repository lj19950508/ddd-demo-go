package main

import (
	"ddd-demo-go/adapter/in/web/rest"
	"ddd-demo-go/config"
	"github.com/gin-gonic/gin"
)

func main() {
	//load config from yml and env  (not a pointer)
	var cfg, err = config.NewConfig()
	if err != nil {
		panic(err)
	}

	//set logger

	//because not ioc, so code in one func , ioc is finding now.

	//out adapter init

	//repo init

	//service init

	//in adapter init\
	var webEngine = gin.New()
	webEngine.Use(gin.Logger())
	//error deal ， 只是为了稳定，千万别随便panic异常。
	webEngine.Use(gin.Recovery())

	group := webEngine.Group("/")
	var ctl = adapter.NewUserController(nil)
	group.GET("/route", ctl.Info)

	//threadlocal

}
