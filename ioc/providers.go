package ioc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/gin-gonic/gin"
	"github.com/grafana/grafana/pkg/bus"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"github.com/lj19950508/ddd-demo-go/pkg/ginextends"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"go.uber.org/fx"
)


var systemPoolProvider = func() gopool.Pool {
	//有没有释放
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
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	handler.Use(func(ctx *gin.Context) {
		//这里实现可以轻松提到api网关的操作
		//todo if prefix with admin
		//else if prefix with permit(放行)
		//else 需要登录才可以访问
	})

	for _, routerGroup := range routers {
		for _, routerItem := range routerGroup.Router() {
			handler.Handle(routerItem.Method, routerItem.Path, routerItem.Handle)
		}
	}
	return handler
}

var httpServerProvider = func(lc fx.Lifecycle, cfg *config.Config, handler http.Handler, logger logger.Interface, pool gopool.Pool) *httpserver.Server {
	s := httpserver.New(handler, httpserver.Port(cfg.HttpServer.Port))

	httpServerOnStart := func(ctx context.Context) error {
		s.Start()
		logger.Info("httpserver start finished on port:%s",cfg.HttpServer.Port)
		//这里要开异步去监听 http是否报错,报错了调用appStop 关闭全部
		pool.Go(func() {
			logger.Info("pool [%s] execute start", pool.Name())

			err := <-s.Notify()
			//被信号关闭了
			logger.Info("%s", err)
			if err = app.Stop(ctx); err != nil {
				logger.Error("%s", err)
			}
		})

		return nil
	}
	httpServerOnStop := func(ctx context.Context) error {
		if err := s.Shutdown(ctx); err != nil {
			logger.Error("%s", err)
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
				logger.Error("%s", err)
				return err
			} else {
				logger.Info("database connection opend")
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if db.SqlDb != nil {
				if err := db.SqlDb.Close(); err != nil {
					logger.Error("%s", err)
					return err
				}
				logger.Info("database connection closed")
			}else{
				logger.Info("database connection closed when no opend")
			}
			return nil

		},
	})
	return db
}

var loggerProvider = func(cfg *config.Config) logger.Interface {
	return logger.New(cfg.Log.Level)
}

// 线程内eventbug
var inProcEventBusProvider  = func (cfg *config.Config) bus.Bus{
	return bus.New()
}

// var eventBusHandle  = func (lc fx.Lifecycle, cfg *config.Config) eventEventHandler{
// 	return bus.New()
// }