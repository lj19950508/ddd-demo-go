package factory

import (
	"ddd-demo-go/adapter/in/web/api"
	"ddd-demo-go/adapter/out/persistent/grails"
	"ddd-demo-go/application/service"

	"gorm.io/gorm"
)

// import "ddd-demo-go/adapter/in/web/api"

// "ddd-demo-go/adapter/in/web/api"
// 	"ddd-demo-go/application/service"

func Init() {

	//factoryhe ioc必须抽象 ， 不然会导致循环饮用

	Register(grails.NewUserRepositoryImpl(Get[gorm.DB]()))
	Register(service.NewUserServiceImpl(Get[grails.UserRepositoryImpl]()))

	Register(api.NewUserApi(Get[service.UserServiceImpl]()))
}
