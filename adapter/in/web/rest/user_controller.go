package adapter

import (
	"ddd-demo-go/application/service"
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

func (this *UserController) Info(ctx *gin.Context) {
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
