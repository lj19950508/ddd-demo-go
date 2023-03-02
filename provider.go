package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/lj19950508/ddd-demo-go/adapter/in/web/api/v1"
	"github.com/lj19950508/ddd-demo-go/adapter/out/persistent/grails"
	"github.com/lj19950508/ddd-demo-go/application/service"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"go.uber.org/fx"
)

var httpHandlerProvider = func(userApi *v1.UserApi, cfg *config.Config) http.Handler {
	//todo 优化
	gin.SetMode(gin.DebugMode)
	handler := gin.New()
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())

	//TODO这里优化一下
	// Register(userApi)
	// Register()
	// Register()
	for _, v := range userApi.Router() {
		handler.Handle(v.Method, v.Path, v.Handle)
	}
	return handler
}

var httpServerProvider = func(lc fx.Lifecycle,cfg *config.Config,handler http.Handler){
		s:=httpserver.New(handler,httpserver.Port(cfg.Port))
		lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			s.Start()
			//在哪里notiyfy err呢
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.Shutdown()
			return nil
		},
	})
}


//httpserverProvider



var loggerProvider = func (cfg *config.Config) logger.Interface  {
	return logger.New(cfg.Log.Level)
}


//就在main层 声明Provid
func base() []fx.Option {
	return []fx.Option{
		
		//TODO CONFIG 优化
		fx.Provide(config.New),
		fx.Provide(loggerProvider),
		fx.Provide(db.New),
		fx.Provide(httpServerProvider),
		//handler
		fx.Provide(httpHandlerProvider),

		//依赖的末端
		fx.Invoke(func(*httpserver.Server) {}),

	}
}
func option() []fx.Option {
	options := []fx.Option{}
	options = append(options, repositorys()...)
	options = append(options, services()...)
	options = append(options, apis()...)
	options = append(options, base()...)
	return options
}

func apis() []fx.Option {
	return []fx.Option{
		fx.Provide(v1.NewUserApi),
	}
}

func services() []fx.Option {
	return []fx.Option{
		fx.Provide(service.NewUserServiceImpl),
	}
}

func repositorys() []fx.Option {
	return []fx.Option{
		fx.Provide(grails.NewUserRepositoryImpl),
	}
}


