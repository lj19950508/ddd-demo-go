package app

import (
	"github.com/lj19950508/ddd-demo-go/internal/adapter/out/persistent/grails"
	"github.com/lj19950508/ddd-demo-go/internal/application/service"
	"github.com/lj19950508/ddd-demo-go/pkg/ioc"
	"github.com/lj19950508/ddd-demo-go/pkg/mysql"
)

// import "github.com/lj19950508/ddd-demo-go/pkg/logger"

//区分什么是不运行就关掉的   redis不必须， mysql 必须  ,mq（可选）必须。。等等  必须的 要监听错误的话得关闭当前整个服务以便于被观测到， 或者得发送通知，
//只会在两个地方register  wire和main
//注册的东西有service, repository , mysql
//在api wire service,或者 也是在adapter wire别的东西，

//注册第三方组件 如mysql  mq

//自动装配service（注册且注入）  repo/service
func wire() {
	//-------repository
	ioc.Register(grails.NewUserRepositoryImpl(ioc.Get[mysql.Mysql]()))

	//-------service
	ioc.Register(service.NewUserServiceImpl(ioc.Get[grails.UserRepositoryImpl]()))
	
}
