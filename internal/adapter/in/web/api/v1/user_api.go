package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/internal/application/service"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
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
	//如果是业务错误直接这么返回并打印堆栈没错（找不到数据库）， 如果是服务器错误也要打印堆栈吗
	//如何区分业务错误和服务器错误   业务错误(状态不对,数据库没数据) 服务器错误(nil指针，数据库连接出错，)
	//业务错误比如 状态不对 if status != 5 ，如何控制这个流程
	// 结论  有error声明的都是业务错误， 没有被error识别的 会被panic recover 为服务器错误
	user, err := t.userService.Info(1)
	if err != nil {
		logger.Instance.Error("%+v",err)
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, user)
	}
	
}
