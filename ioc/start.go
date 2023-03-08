package ioc

import (
	adminapi "github.com/lj19950508/ddd-demo-go/adapter/in/adminapi/user"
	api "github.com/lj19950508/ddd-demo-go/adapter/in/api/user"
	evthandler "github.com/lj19950508/ddd-demo-go/adapter/in/eventhandler/user"
	queryimpl "github.com/lj19950508/ddd-demo-go/adapter/out/queryimpl/user"
	repositoryimpl "github.com/lj19950508/ddd-demo-go/adapter/out/repositoryimpl/user"
	command "github.com/lj19950508/ddd-demo-go/application/command/user"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/pkg/ginextends"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"go.uber.org/fx"
)

var app *fx.App

func Run() {
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
	return fx.Provide(evthandler.NewUserEventHandler)
}

func options() []fx.Option {
	options := []fx.Option{}
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
	return options
}

func invoke() fx.Option {
	return fx.Invoke(
		func(*httpserver.Server) {},
		func(*evthandler.UserEventHandler) {},
	)

}

func base() fx.Option {
	return fx.Provide(
		config.New,
		loggerProvider,
		fx.Annotate(httpHandlerProvider, fx.ParamTags(`group:"routes"`)),
		fx.Annotate(httpServerProvider, fx.ParamTags(``, ``, ``, ``, `name:"systemPool"`)),
		dbProvider,
		inProcEventBusProvider,
		fx.Annotate(systemPoolProvider, fx.ResultTags(`name:"systemPool"`)),
	)

}

func asRoute(f any) any {
	return fx.Annotate(f, fx.As(new(ginextends.Routerable)), fx.ResultTags(`group:"routes"`))
}
