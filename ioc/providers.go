package ioc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"github.com/lj19950508/ddd-demo-go/pkg/grpcextends"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/route"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var ginHandlerProvider = func(routers []route.Routeable, lc fx.Lifecycle, cfg *config.Config, logger logger.Interface) *gin.Engine {
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
		for _, routerItem := range *routerGroup.Route() {
			handlerfunc, ok := routerItem.Handler.(func(*gin.Context))
			if ok {
				pattern := strings.Fields(routerItem.Pattern)
				handler.Handle(pattern[0], pattern[1], handlerfunc)
			}
		}
	}
	httpserver := httpserver.New(handler, httpserver.Port(cfg.HttpServer.Port))

	httpServerOnStart := func(ctx context.Context) error {
		httpserver.Start()
		logger.Info("httpserver start finished on port:%s", cfg.HttpServer.Port)
		//这里要开异步去监听 http是否报错,报错了调用appStop 关闭全部
		go (func() {
			err := <-httpserver.Notify()
			//被信号关闭了
			logger.Info("%s", err)
			if err = app.Stop(ctx); err != nil {
				logger.Error("%s", err)
			}
		})()

		return nil
	}
	httpServerOnStop := func(ctx context.Context) error {
		if err := httpserver.Shutdown(ctx); err != nil {
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
	return handler
}

var grpcProvider = func(grpcHandlers []grpcextends.GrpcHandler, lc fx.Lifecycle, cfg *config.Config, logger logger.Interface) *grpc.Server {

	grpcServer := grpc.NewServer()

	// server := httpserver.New(grpcServer, httpserver.Port(cfg.GrpcServer.Port))
	grpcServerOnStart := func(ctx context.Context) error {
		for _, handler := range grpcHandlers {
			handler.Register(grpcServer)
		}
		// server.Start()
		go func() {
			ls, err := net.Listen("tcp", ":"+cfg.GrpcServer.Port)
			if err != nil {
				logger.Fatal("grpcServer start err %s", err)
			}
			logger.Info("grpcserver start finished on port:%s", cfg.GrpcServer.Port)
			grpcServer.Serve(ls)
			if err != nil {
				//server会自动重连然 后报错
				logger.Fatal("grpcServer connect err %s", err)
			}

		}()

		return nil
	}
	grpcServerOnStop := func(ctx context.Context) error {
		grpcServer.Stop()
		logger.Info("grpcserver stop finished")

		return nil
	}

	lc.Append(fx.Hook{
		OnStart: grpcServerOnStart,
		OnStop:  grpcServerOnStop,
	})
	return grpcServer
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
			} else {
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
