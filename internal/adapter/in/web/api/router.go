// Package v1 implements routing paths. Each services in own file.
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/lj19950508/ddd-demo-go/internal/adapter/in/web/api/v1"
	"github.com/lj19950508/ddd-demo-go/internal/application/service"
	"github.com/lj19950508/ddd-demo-go/pkg/ioc"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine) {

	handler.Use(gin.Recovery())

	//访问日志 缺少入参处参
	handler.Use(gin.Logger())
	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Routers
	h := handler.Group("/v1")
	// 为api配置路由，创建api
	v1.NewUserApi(h, ioc.Get[service.UserServiceImpl](),ioc.Get[logger.Logger]())
}
