package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/gin-gonic/gin"
	v1 "github.com/lj19950508/ddd-demo-go/adapter/in/api/v1"
	"github.com/lj19950508/ddd-demo-go/adapter/out/queryimpl"
	"github.com/lj19950508/ddd-demo-go/adapter/out/repositoryimpl"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"github.com/lj19950508/ddd-demo-go/pkg/ginextends"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"go.uber.org/fx"
)

var app *fx.App

func main() {

	app = fx.New(
		options()...,
	)
	app.Run()
}

//---------实例注册声明------------//

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
	// 初始化根层 如 httpservcer socketserver
	options = append(options, invoke())
	return options
}

func invoke() fx.Option {
	return fx.Invoke(func(*httpserver.Server) {})
}

func base() fx.Option {
	return fx.Provide(
		config.New,
		loggerProvider,
		fx.Annotate(httpHandlerProvider, fx.ParamTags(`group:"routes"`)),
		fx.Annotate(httpServerProvider, fx.ParamTags(``, ``, ``, ``, `name:"systemPool"`)),
		dbProvider,
		fx.Annotate(systemPoolProvider, fx.ResultTags(`name:"systemPool"`)),
	)

}

func apis() fx.Option {
	return fx.Provide(
		asRoute(v1.NewUserApi),
	)

}

func queryService() fx.Option {
	return fx.Provide(queryimpl.NewUserQueryServiceimpl)
}

func cmdService() fx.Option {
	return fx.Provide(queryimpl.NewUserQueryServiceimpl)
}

func repositorys() fx.Option {
	return fx.Provide(repositoryimpl.NewUserRepositoryImpl)

}

func asRoute(f any) any {
	return fx.Annotate(f, fx.As(new(ginextends.Routerable)), fx.ResultTags(`group:"routes"`))
}

//
// ------------------------------构造函数声明-----------------------
//
//

var systemPoolProvider = func() gopool.Pool {
	pool := gopool.NewPool("system", int32(100), &gopool.Config{ScaleThreshold: 80})
	return pool
}

var httpHandlerProvider = func(routers []ginextends.Routerable, cfg *config.Config) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	handler.Use(gin.Recovery())

	handler.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[REQ] %s | %-6s| %s | %s | %d %s  %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.ClientIP,
			// param.Request.UserAgent(),
			param.Path,
			// param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	authGroup := handler.Group("")
	authGroup.Use(ginextends.TokenHandler)

	for _, routerGroup := range routers {
		for _, routerItem := range routerGroup.Router() {
			//区分不需登录 和需要登录的接口即可
			if routerItem.NoAuth {
				handler.Handle(routerItem.Method, routerItem.Path, routerItem.Handle)
			}else{
				authGroup.Handle(routerItem.Method, routerItem.Path, routerItem.Handle)
			}

		}
	}
	return handler
}

var httpServerProvider = func(lc fx.Lifecycle, cfg *config.Config, handler http.Handler, logger logger.Interface, pool gopool.Pool) *httpserver.Server {
	s := httpserver.New(handler, httpserver.Port(cfg.HttpServer.Port))

	httpServerOnStart := func(ctx context.Context) error {
		s.Start()
		logger.Info("httpserver start finished")
		//这里要开异步去监听 http是否报错,报错了调用appStop 关闭全部
		pool.Go(func() {
			logger.Info("pool [%s] execute start", pool.Name())

			err := <-s.Notify()
			logger.Error("%s", err)
			if err := app.Stop(ctx); err != nil {
				logger.Error("%s", err)
			}
		})

		return nil
	}
	httpServerOnStop := func(ctx context.Context) error {
		if err := s.Shutdown(ctx); err != nil {
			return err
		}
		logger.Info("httpserver stop finished")
		return nil
	}

	lc.Append(fx.Hook{
		OnStart: httpServerOnStart,
		OnStop:  httpServerOnStop,
	})
	return s
}

var dbProvider = func(lc fx.Lifecycle, cfg *config.Config, logger logger.Interface) *db.DB {
	db := db.New(cfg.Mysql.Url, db.MaxIdleConns(10), db.MaxOpenConns(100), db.ConnMaxLifetime(time.Hour))
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.Open(); err != nil {
				return err
			}
			logger.Info("database connection opend")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := db.SqlDb.Close(); err != nil {
				return err
			}
			logger.Info("database connection closed")
			return nil
		},
	})
	return db
}

var loggerProvider = func(cfg *config.Config) logger.Interface {
	return logger.New(cfg.Log.Level)
}
