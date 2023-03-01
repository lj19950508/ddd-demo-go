package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/internal/adapter/in/web/api"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/ioc"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/mysql"
)

func Run(cfg *config.Config) {
	var err error

	ioc.NewIOC()
	logger :=logger.New(cfg.Log.Level)

	//从下面开始可以使用 logger.Instance

	//注册第三方插件的start
	//初始化mysql
	//mysql.new...  defer close.
	mysql, err := mysql.New(cfg.Mysql.Url)
	//if err
	if err != nil {
		logger.Fatal("%s", err)
	}
	defer mysql.Close()

	//new一个ioc比较好
	ioc.Register(logger,mysql)
	wire()

	//web服务  socket等 输入服务输入服务写在下面
	handle := gin.New()
	api.NewRouter(handle)
	httpServer := httpserver.New(handle, httpserver.Port(cfg.Port))

	//监听信号
	// interrup t := make(chan os.Signal, 1)
	// signal.Noti(interrupt, os.Interrupt, syscall.SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	//循环while直至收到线程关闭信号 或者httpserver关闭信号（这个不清楚）,或者rmq关闭信号
	//某个case执行成功就会往下执行，除非有for
	select {
	case <-ctx.Done():
		//ctrl c 会走到这
		logger.Info("context done")
	case err = <-httpServer.Notify():
		logger.Error("httpServer.shutdown:%s", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		//关闭如果出错 则会走这
		logger.Error("httpServer.shutdown:%w", err)
	}
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
