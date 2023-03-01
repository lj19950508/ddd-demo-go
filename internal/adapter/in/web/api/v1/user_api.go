package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/internal/adapter/in/web/api/wrapper"
	"github.com/lj19950508/ddd-demo-go/internal/application/service"
)

type UserApi struct {
	userService service.UserService
}

func NewUserApi(handler *gin.RouterGroup, userService service.UserService) {
	userApi := &UserApi{
		userService: userService,
	}
	routerGroup := handler.Group("/user")
	{
		// routerGroup.GET("", userApi.Page)
		routerGroup.GET("/info", userApi.Info)
		// routerGroup.POST("", userApi.Create)
		// routerGroup.PUT("", userApi.Update)
		// routerGroup.DELETE("", userApi.Delete)
	}

}

func (t *UserApi) Info(ctx *gin.Context) {

	//指责
	//1.转换参数成dto
	//2.打印参数(不知道能否logger实现) 出参 入参
	//3.验证参数
	//4.dto->domain
	//5.domian->dto
	//6.处理异常病打印堆栈
	//7.(业务吗与异常系统)
	//TODO 8.ResultDTO

	//abouterror
	//野生的异常会打印堆栈，已知的异常靠自己处理。频繁panic性能不好.
	//自己处理异常 BizErrorHandler

	// dealerr(ctx)
	//
	// var a *int

	user, err := t.userService.Info(100)
	if err != nil {
		//HandlerError(err)
		// return
		//TODO if erroris .... 怎样怎样
		//对err判断并处理
		// ctx.Status()
		// ctx.JSON(http.StatusBadRequest, err.Error())
		// httpCode,result := WrapperError(err)
		// ctx.JSON(wrapper.WrapperError(err))

		//如果是errorNotFind异常则要转换成业务异常。
		ctx.JSON(wrapper.Error(err))
		return
	}

	ctx.JSON(wrapper.ResultData(user))

}
