package ioc

import (
	"github.com/gin-gonic/gin"
	adminapi "github.com/lj19950508/ddd-demo-go/adapter/in/adminapi/user"
	api "github.com/lj19950508/ddd-demo-go/adapter/in/api/user"
	"github.com/lj19950508/ddd-demo-go/adapter/in/grpc/user"
	queryimpl "github.com/lj19950508/ddd-demo-go/adapter/out/queryimpl/user"
	"github.com/lj19950508/ddd-demo-go/adapter/out/grpcclient/user"
	repositoryimpl "github.com/lj19950508/ddd-demo-go/adapter/out/repositoryimpl/user"
	"google.golang.org/grpc"

	command "github.com/lj19950508/ddd-demo-go/application/command/user"
	"github.com/lj19950508/ddd-demo-go/config"

	// "github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/grpcextends"
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

func grpcapi() fx.Option{
	return fx.Provide(
		asGrpcHandler(grpchandler.NewUserApi),
	)
}

func grpcsender() fx.Option{
	return fx.Provide(grpcclient.NewUserRpcSender)
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
	options = append(options, grpcapi())
	options = append(options, grpcsender())
	// 初始化根层 如 httpservcer socketserver
	options = append(options, invoke())
	//IF DEV  option 要在ioc之前
	// fx.Populate(cfg),

	// options = append(options, fx.NopLogger)


	return options
}

func invoke() fx.Option {
	return fx.Invoke(
		func(*gin.Engine) {},
		func(*grpc.Server) {},
		// func(*user.UserApi) {},
	)

}

func base() fx.Option {
	return fx.Provide(
		//定义服务名
		loggerProvider,
		dbProvider,
		fx.Annotate(ginHandlerProvider, fx.ParamTags(`group:"routes"`)),
		fx.Annotate(grpcProvider, fx.ParamTags(`group:"grpchandlers"`)),

	)

}

func asRoute(f any) any {
	return fx.Annotate(f, fx.As(new(route.Routeable)), fx.ResultTags(`group:"routes"`))
}

func asGrpcHandler(f any) any {
	return fx.Annotate(f, fx.As(new(grpcextends.GrpcHandler)), fx.ResultTags(`group:"grpchandlers"`))
}
