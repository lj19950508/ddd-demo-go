package ioc

import (
	adminapi "github.com/lj19950508/ddd-demo-go/adapter/in/adminapi/user"
	api "github.com/lj19950508/ddd-demo-go/adapter/in/api/user"
	evthandler "github.com/lj19950508/ddd-demo-go/adapter/in/eventhandler/user"
	queryimpl "github.com/lj19950508/ddd-demo-go/adapter/out/queryimpl/user"
	repositoryimpl "github.com/lj19950508/ddd-demo-go/adapter/out/repositoryimpl/user"
	command "github.com/lj19950508/ddd-demo-go/application/command/user"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/pkg/eventbus"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/route"
	"go.uber.org/fx"
)

var app *fx.App

func Run() {
	

	//根据环境
	//config
	//fx.add in oc config
	app = fx.New(
		options()...,	
	)
	app.Run()
}

func apis() fx.Option {
	return fx.Provide(
		asRoute(api.NewUserApi),
		asRoute(adminapi.NewAdminUserApi),
	)
}

func queryService() fx.Option {
	return fx.Provide(queryimpl.NewUserQueryServiceImpl)
}

func cmdService() fx.Option {
	return fx.Provide(command.NewUserCommandImpl)
}

func repositorys() fx.Option {
	return fx.Provide(repositoryimpl.NewUserRepositoryImpl)
}

func eventhandler() fx.Option{
	return fx.Provide(
		asEventHandler(evthandler.NewUserEventHandler),
	)
}

func options() []fx.Option {
	options := []fx.Option{}

	cfg :=config.New()
	if(cfg.Log.Level!="debug"){
		options = append(options, fx.NopLogger)
	}
	options = append(options, fx.Supply(cfg))
	

	// 基础实现层 如http mysql ，redis ，web
	options = append(options, base())
	// api接口
	options = append(options, apis())
	// service cqrs的体现
	// queryservice 注入 queryserviceimpl 注入  读库的 db,es,redis
	options = append(options, queryService())
	options = append(options, cmdService())
	// 仓储注入 writedb
	options = append(options, repositorys())
	options = append(options, eventhandler())
	// 初始化根层 如 httpservcer socketserver
	options = append(options, invoke())
	//IF DEV  option 要在ioc之前
	// fx.Populate(cfg),

	// options = append(options, fx.NopLogger)


	return options
}

func invoke() fx.Option {
	return fx.Invoke(
		func(*httpserver.Server) {},
		//为了启动eventbus持续消费
		func(eventbus.EventBus) {},
	)

}

func base() fx.Option {
	return fx.Provide(		
		loggerProvider,
		fx.Annotate(httpHandlerProvider, fx.ParamTags(`group:"routes"`)),
		fx.Annotate(httpServerProvider, fx.ParamTags(``, ``, ``, ``, `name:"systemPool"`)),
		dbProvider,
		ginHandlerProvider,
		fx.Annotate(mqRpcEventBusProvider, fx.ParamTags(``,`group:"eventhandler"`)),
		fx.Annotate(systemPoolProvider, fx.ResultTags(`name:"systemPool"`)),
	)

}

func asRoute(f any) any {
	return fx.Annotate(f, fx.As(new(route.Routeable)), fx.ResultTags(`group:"routes"`))
}

func asEventHandler(f any) any {
	return fx.Annotate(f, fx.As(new(eventbus.Dispatcher)), fx.ResultTags(`group:"eventhandler"`))
}
