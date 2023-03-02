package main

import (
	"context"
	"net/http"
	"time"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/gin-gonic/gin"
	v1 "github.com/lj19950508/ddd-demo-go/adapter/in/web/api/v1"
	"github.com/lj19950508/ddd-demo-go/adapter/out/persistent/grails"
	"github.com/lj19950508/ddd-demo-go/application/service"
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
	options = append(options, base())
	options = append(options, apis())
	options = append(options, services())
	options = append(options, repositorys())
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
		fx.Annotate(httpServerProvider, fx.ParamTags(``,``,``,``,`name:"systemPool"`)),
		dbProvider,
		fx.Annotate(systemPoolProvider, fx.ResultTags(`name:"systemPool"`)),
	)

}

func apis() fx.Option {
	return fx.Provide(
		asRoute(v1.NewUserApi),
	)

}

func services() fx.Option {
	return fx.Provide(service.NewUserServiceImpl)

}

func repositorys() fx.Option {
	return fx.Provide(grails.NewUserRepositoryImpl)

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
	handler.Use(gin.Logger())
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	for _, routerGroup := range routers {
		for _, routerItem := range routerGroup.Router() {
			handler.Handle(routerItem.Method, routerItem.Path, routerItem.Handle)
		}
	}
	return handler
}

var httpServerProvider = func(lc fx.Lifecycle, cfg *config.Config, handler http.Handler, logger logger.Interface, pool gopool.Pool) *httpserver.Server {
	s := httpserver.New(handler, httpserver.Port(cfg.Port))

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
		//被需要的时候只会执行一次
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

//httpserverProvider

var loggerProvider = func(cfg *config.Config) logger.Interface {
	return logger.New(cfg.Log.Level)
}

//就在main层 声明Provid
