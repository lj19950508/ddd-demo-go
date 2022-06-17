package adapter

import (
	"ddd-demo-go/application/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserApi struct {
	userSerivce service.UserService
}

func NewUserApi(userService service.UserService) *UserApi {
	return &UserApi{
		userSerivce: userService,
	}
}

func (t *UserApi) Info(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, strconv.IntSize)
	user, err := t.userSerivce.Info(uint(id))
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(200, user)
	}
}

func (t *UserApi) test(ctx *gin.Context) {
	ctx.JSON(200, 200)
}
