package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/internal/application/service"
)

type UserApi struct {
	userService service.UserService
}

func NewUserApi(handler *gin.RouterGroup,userService service.UserService) {
	userApi:=&UserApi{
		userService: userService,
	}
	routerGroup:=handler.Group("/user")
	{
		routerGroup.GET("/info",userApi.Info)
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
