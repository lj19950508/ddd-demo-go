package main

import (
	"ddd-demo-go/adapter/in/web"
	"ddd-demo-go/adapter/out/persistent/grails"
	"ddd-demo-go/config"
	"ddd-demo-go/factory"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	//加载通用中间件 如日志。。。
	//errors.Wrap()
	//set logger
	//加载ioc容器
	//ioc
	// repo := gorm.NewUserRepositoryImpl()
	// svc := service.NewUserService(repo)
	// ppi := api.NewUserApi(svc)
	//初始化具体组件 如mysql web socket mq
	//web server do

	// in -> iioc
	// ioc -> in/out/
	// main -> in/out /ioc
	web.StartWeb(cfg)

	//解决ioc问题
	factory.Register(grails.StartOrm())

}

//有一个创建返回controler的
//启动web容器
