package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/config"
	v1 "github.com/lj19950508/ddd-demo-go/adapter/in/web/api/v1"
	"github.com/lj19950508/ddd-demo-go/adapter/out/persistent/grails"
	"github.com/lj19950508/ddd-demo-go/application/service"
	"github.com/lj19950508/ddd-demo-go/pkg/httpserver"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/mysql"
	"go.uber.org/fx"
)


func main() {
	fx := fx.New(
		option()...,
	)
	fx.Run()
}

func option() []fx.Option {
	options := []fx.Option{}
	options = append(options, repositorys()...)
	options = append(options, services()...)
	options = append(options, apis()...)
	options = append(options, base()...)
	// options = append(options, routers()...)
	return options
}

func base() []fx.Option {
	return []fx.Option{
		
		//TODO CONFIG 优化
		fx.Provide(config.New),
		fx.Provide(logger.New),
		fx.Provide(mysql.New),
		fx.Provide(httpserver.New),
		//handler
		fx.Provide(func (userApi *v1.UserApi,cfg *config.Config) http.Handler {
			//TODO youhua
			gin.SetMode(gin.DebugMode)
			handler := gin.New()
			handler.Use(gin.Recovery())
			handler.Use(gin.Logger())

			//TODO这里优化一下
			// Register(userApi)
			// Register()
			// Register()
			for _,v := range userApi.Router() {
				handler.Handle(v.Method,v.Path,v.Handle)
			}
			return handler
		}),

		//依赖的末端
		fx.Invoke(func(*httpserver.Server) {}),

	}
}

// func routers() []fx.Option {
// 	return []fx.Option{
// 		fx.Provide(v1.NewUserRouter),
// 	}
// }
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


