package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/adapter/in/web/api/v1/dto"
	"github.com/lj19950508/ddd-demo-go/adapter/in/web/api/wrapper"
	"github.com/lj19950508/ddd-demo-go/application/service"
	"github.com/lj19950508/ddd-demo-go/pkg/ginextends"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
)

type UserApi struct {
	userService service.UserService
	logger      logger.Interface
}

func (t *UserApi) Router() ginextends.RouterInfos {
	return ginextends.RouterInfos{
		{Method: "GET", Path: "/v1/users/:id", Handle: t.Info},
	}
}

func NewUserApi(userService service.UserService, logger logger.Interface) *UserApi {
	return &UserApi{
		userService: userService,
		logger:      logger,
	}
}

func (t *UserApi) Info(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, wrapper.ResultMsg(err.Error()))
		return
	}
	t.logger.Info("[访问用户信息-入参] id:%d", id)
	user, err := t.userService.Info(id)
	if err != nil {
		t.logger.Info("[访问用户信息-错误] err:%s", err)
		c.JSON(wrapper.Error(err))
		return
	}
	dto := dto.NewUser(user.Id,user.Name)
	t.logger.Info("[访问用户信息-返回]:%+v", dto)
	c.JSON(http.StatusOK, wrapper.ResultData(dto))

}
