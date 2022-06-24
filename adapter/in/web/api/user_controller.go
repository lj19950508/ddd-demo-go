package api

import (
	"ddd-demo-go/application/service"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	userService service.UserService
}

func NewUserApi(userService service.UserService) UserApi {

	return UserApi{
		userService: userService,
	}
}

func (t *UserApi) Info(ctx *gin.Context) {
	//1.转换参数成dto
	//2.验证dto参数
	//3.传入参数dto
	user, err := t.userService.Info(1)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(200, user)
	}
	//4.处理返回
}

func (t *UserApi) test(ctx *gin.Context) {

	ctx.JSON(200, 200)
}
