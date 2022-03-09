package rest

import (
	"ddd-demo1/application/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSerivce service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userSerivce: userService,
	}
}

func (this *UserController) GetGroupPath() string {
	return "/user"
}

func (this *UserController) GetHandleFunc() gin.RoutesInfo {
	routeInfo := gin.RoutesInfo{
		{Method: "GET", Path: "/info", HandlerFunc: this.info},
		{Method: "GET", Path: "/test", HandlerFunc: this.test},
	}
	return routeInfo
}

func (this *UserController) info(ctx *gin.Context) {
	user, err := this.userSerivce.Info(1)
	if err != nil {
		print(1)
		ctx.JSON(400, err.Error())
	} else {
		print(2)
		ctx.JSON(200, user)
	}
}

func (this *UserController) test(ctx *gin.Context) {
	ctx.JSON(200, 200)
}
