package di

import (
	"ddd-demo1/application/service"
	"ddd-demo1/domain/biz1/repository"
	"ddd-demo1/infrastructrue/middleware"
	"ddd-demo1/infrastructrue/persistent/gorm"
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
	this.produceBean()
	this.produceRouter()
	return func(c *gin.Context) {
		c.Next()
	}
}

func (this *BeanFactory) produceRouter() {
	this.routes = []routeable.Routeable{
		rest.NewUserController(this.beans["userService"].(service.UserService)),
	}
}
func (this *BeanFactory) produceBean() {
	this.beans = make(map[string]interface{})

	this.beans["gormDatasource"] = middleware.NewGormResource()
	this.beans["userRepository"] = gorm.NewUserRepositoryImpl(this.beans["gormDatasource"].(*middleware.GormResource))

	this.beans["userService"] = service.NewUserServiceImpl(this.beans["userRepository"].(repository.UserRepository))
}

func (this *BeanFactory) GetRoutes() []routeable.Routeable {
	return this.routes
}
