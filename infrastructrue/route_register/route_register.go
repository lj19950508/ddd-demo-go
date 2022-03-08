package route_register

import (
	"ddd-demo1/interface/routeable"
	"github.com/gin-gonic/gin"
)

type RouteRegister struct {
	controllers []routeable.Routeable
}

func NewRouteRigster(routes []routeable.Routeable) *RouteRegister {
	return &RouteRegister{
		controllers: routes,
	}
}

func (this *RouteRegister) Run(engine *gin.Engine) gin.HandlerFunc {
	this.register(engine)
	return func(ctx *gin.Context) {
		ctx.Next()
	}

}

func (this *RouteRegister) register(engine *gin.Engine) {

	for i := range this.controllers {
		group := engine.Group(this.controllers[i].GetGroupPath())
		for j := range this.controllers[i].GetHandleFunc() {
			funcItem := this.controllers[i].GetHandleFunc()[j]
			group.Handle(funcItem.Method, funcItem.Path, funcItem.HandlerFunc)
		}
	}
}
