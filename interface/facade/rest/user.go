package rest

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
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
	ctx.JSON(200, 200)
}
