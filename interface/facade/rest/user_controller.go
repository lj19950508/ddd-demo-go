package rest

import (
	"ddd-demo1/application/service"
	"github.com/gin-gonic/gin"
	"strconv"
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
		{Method: "GET", Path: "/info/:id", HandlerFunc: this.info},
		{Method: "GET", Path: "/test", HandlerFunc: this.test},
	}
	return routeInfo
}

func (this *UserController) info(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, strconv.IntSize)
	user, err := this.userSerivce.Info(uint(id))
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(200, user)
	}
}

func (this *UserController) test(ctx *gin.Context) {
	ctx.JSON(200, 200)
}
