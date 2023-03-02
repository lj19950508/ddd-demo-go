// // Package v1 implements routing paths. Each services in own file.
package api

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	v1 "github.com/lj19950508/ddd-demo-go/internal/adapter/in/web/api/v1"
// 	"github.com/lj19950508/ddd-demo-go/internal/application/service"
// 	"github.com/lj19950508/ddd-demo-go/pkg/ioc"
// )

// // NewRouter -.
// // Swagger spec:
// // @title       Go Clean Template API
// // @description Using a translation service as an example
// // @version     1.0
// // @host        localhost:8080
// // @BasePath    /v1
// func NewRouter(handler *gin.Engine) {

// 	handler.Use(gin.Recovery())

// 	//访问日志 缺少入参处参
// 	handler.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
// 		if param.Latency > time.Minute {
// 			param.Latency = param.Latency.Truncate(time.Second)
// 		}
// 		return fmt.Sprintf(" %v | %3d | %13v | %15s | %-7s  %#v\n%s",
// 			param.TimeStamp.Format("2006/01/02T15:04:05"),
// 			param.StatusCode,
// 			param.Latency,
// 			param.ClientIP,
// 			param.Method,
// 			param.Path,
// 			param.ErrorMessage,
// 		)
// 	}))
// 	// K8s probe
// 	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

// 	// Routers
// 	h := handler.Group("/v1")
// 	// 为api配置路由，创建api
// 	v1.NewUserApi(h, ioc.Get[service.UserServiceImpl]())
// }
