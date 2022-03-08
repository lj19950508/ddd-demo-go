package bean_factory

import (
	"ddd-demo1/interface/facade/rest"
	"ddd-demo1/interface/routeable"
	"github.com/gin-gonic/gin"
)

type BeanFactory struct {
	routes []routeable.Routeable
	beans  map[string]interface{}
}

func NewBeanFactory() *BeanFactory {
	return &BeanFactory{}
}

func (this *BeanFactory) Run() gin.HandlerFunc {
	this.produceRouter()
	return func(c *gin.Context) {
		c.Next()
	}
}

func (this *BeanFactory) produceRouter() {
	this.routes = []routeable.Routeable{
		rest.NewUserController(),
	}
}
func (this *BeanFactory) produceBean() {

}

func (this *BeanFactory) GetRoutes() []routeable.Routeable {
	return this.routes
}
