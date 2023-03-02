package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/internal/application/service"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/wrapper"
)

type UserApi struct {
	userService service.UserService
	logger logger.Interface
}

func NewUserApi(userService service.UserService,logger logger.Interface) *UserApi {
	return &UserApi{
		userService: userService,
		logger: logger,
	}
}

func (t *UserApi) Router() *gin.RouterGroup{
	group:= &gin.RouterGroup{}
	group.Group("/user")
	{
		group.GET("/info/:id", t.Info)
	}
	return group
}

func (t *UserApi) Info(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest,wrapper.ResultMsg(err.Error()))
		return
	}
	t.logger.Info("[访问用户信息-入参] id:%d", id)
	//.. ctx.ShouldBind dto
	//.. 验证参数或者从query dto 标签验证
	// dto->domain

	user, err := t.userService.Info(id)

	//domain->dto
	if err != nil {
		//异常要不要输出堆栈的问题
		t.logger.Info("[访问用户信息-错误] err:%s",err)
		c.JSON(wrapper.Error(err))
		return
	}
	t.logger.Info("[访问用户信息-返回]:%+v", user)
	c.JSON(http.StatusOK,wrapper.ResultData(user))

}
