package web

import (
	"context"
	"ddd-demo-go/config"
	"ddd-demo-go/factory"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"github.com/gin-gonic/gin"
)

//启动web
//监听web
//web路由
func initWeb(cfg *config.Config) *http.Server {

	router := gin.New()
	router.Use(gin.Recovery())
	//请求日志accessLog 

	// group:=router.Group()
	
	userApi:=factory.GetUserApi()
	router.GET("/", userApi.Info)

	//这里依赖到controller.

	webServer := &http.Server{
		Addr:    ":" + fmt.Sprint(cfg.Port),
		Handler: router,
	}
	return webServer

}

//启动容器并监听关闭事件，这个没依赖到gin
func StartWeb(cfg *config.Config) {
	//获取信号上下文，和重置回调 ,这个信号可以被k8s等掌控到
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//延后执行stop 重制信号
	defer stop()
	// 定义web套接字以及 路由
	webServer := initWeb(cfg)
	// 异步起动web容器 防止影响下列行为
	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//等待上下文关闭信号
	<-ctx.Done()
	//重制os.Interrupt的默认行为
	stop()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	//给予程序最多5秒的时间处理正在服务的请求 ,可以延长
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//释放资源
	defer cancel()
	//执行上面的退出规范

	if err := webServer.Shutdown(ctx); err != nil {
		//如果有异常则抛出
		log.Fatal("Server forced to shutdown: ", err)
	}
	//
	log.Println("Server exiting")
	//优雅关闭pod流程
	//1.切流量
	//2.关闭（如果有长时间工作的进程则可以选择 超时关闭 或者无限等待）
}
