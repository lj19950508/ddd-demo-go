package rest

import (
	"ddd-demo1/application/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSerivce service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userSerivce: userService,
	}
}

func (this *UserController) GetGroupPath() string {
	return "/user"
}
func (this *UserController) GetHandleFunc() gin.RoutesInfo {
	routeInfo := gin.RoutesInfo{
		{
			Method: "GET", Path: "/info", HandlerFunc: this.info,
		},
	}
	return routeInfo
}

func (this *UserController) info(ctx *gin.Context) {
	this.userSerivce.Hello()
	ctx.JSON(200, 200)
}
