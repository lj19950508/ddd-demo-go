package main

import (
	adapter "ddd-demo-go/adapter/in/web/api"
	"ddd-demo-go/adapter/out/persistent/gorm"
	"ddd-demo-go/application/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//load config from ym and env  (not a pointer)
	_, err := NewConfig()
	if err != nil {
		panic(err)
	}

	//errors.Wrap()

	//set logger

	repo := gorm.NewUserRepositoryImpl()
	service := service.NewUserService(repo)
	api := adapter.NewUserApi(service)

	//web server do
	var engine = gin.New()
	//engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	group := engine.Group("/")
	group.GET("/route", api.Info)

}
