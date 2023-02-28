package app

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/internal/adapter/in/web/api"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
)

func Run(cfg *config.Config) {
	var err error
	logger.New(cfg.Log.Level)
	wire()

	//初始化mysql
	//mysql.new...  defer close.

	//web服务
	handle := gin.New()
	api.NewRouter(handle)
	httpServer := httpserver.New(handle, httpserver.Port(cfg.Port))

	//监听信号
	// interrup t := make(chan os.Signal, 1)
	// signal.Noti(interrupt, os.Interrupt, syscall.SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// 	//延后执行stop 重制信号
	defer stop()

	//循环while直至收到线程关闭信号 或者httpserver关闭信号（这个不清楚）,或者rmq关闭信号
	//某个case执行成功就会往下执行，除非有for
	select {
	case s := <-ctx.Done():
		fmt.Println(s)
	case err = <-httpServer.Notify():
		fmt.Print("err http" + err.Error())
		// case err = ...
	}

	//执行关闭代码（安全关闭）
	err = httpServer.Shutdown()
	if err != nil {
		fmt.Println("hh")
	}
	//.....

}

//耻辱代码
// func Run(cfg *config.Config) {
// 	//初始化日志对象
// 	//从底层开始创建对象

// 	//创建对象
// 	//defer。Close

// 	//优雅关闭context

// 	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
// 	//延后执行stop 重制信号
// 	defer stop()
// 	// ---web开始
// 	router := gin.New()
// 	router.Use(gin.Recovery())

// 	webServer := &http.Server{
// 		Addr:    ":" + fmt.Sprint(cfg.Port),
// 		Handler: router,
// 	}
// 	// 异步起动web容器 防止影响下列行为
// 	go func() {
// 		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("listen: %s\n", err)
// 		}
// 	}()
// 	//--web结束

// 	//等待上下文关闭信号
// 	<-ctx.Done()
// 	//重制os.Interrupt的默认行为
// 	stop()

// 	log.Println("shutting down gracefully, press Ctrl+C again to force")
// 	//给予程序最多5秒的时间处理正在服务的请求 ,可以延长
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	//释放资源
// 	defer cancel()
// 	//执行上面的退出规范

// 	if err := webServer.Shutdown(ctx); err != nil {
// 		//如果有异常则抛出
// 		log.Fatal("Server forced to shutdown: ", err)
// 	}
// 	//
// 	log.Println("Server exiting")
// }
