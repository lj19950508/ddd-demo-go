package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/lj19950508/ddd-demo-go/application/command"
	"github.com/lj19950508/ddd-demo-go/application/query"
	"github.com/lj19950508/ddd-demo-go/pkg/ginextends"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/resultpkg"
)

type UserApi struct {
	userCommandService command.UserCommandService
	userQueryService query.UserQueryService
	logger      logger.Interface
}

func (t *UserApi) Router() ginextends.RouterInfos {
	return ginextends.RouterInfos{
		//默认使用user吧
		{Method: "GET", Path: "/v1/users/:id", Handle: t.Info,NoAuth: true},
		{Method: "GET", Path: "/v1/user/:id", Handle: t.Info},
	}
}

func NewUserApi(userCommandService command.UserCommandService,userQueryService query.UserQueryService, logger logger.Interface) *UserApi {
	return &UserApi{
		userCommandService: userCommandService,
		userQueryService: userQueryService,
		logger:      logger,
	}
}

func (t *UserApi) Delete(c *gin.Context){}

func (t *UserApi) Info(c *gin.Context) {
	var userQuery query.UserQuery
	err :=c.ShouldBindQuery(&userQuery) //根据情况should活着 must
	if err != nil {
		c.JSON(http.StatusBadRequest,resultpkg.Fail(err.Error()))
		return
	}
	// currentUserId:=c.GetInt64("userId") //注入当前用户 数据权限的方法
	// userQuery.ID=sql.NullInt64{currentUserId,true}

	t.logger.Info("[访问用户信息-入参];%v", userQuery)
	user, err := t.userQueryService.FindOne(&userQuery)
	if err != nil {
		t.logger.Info("[访问用户信息-错误] err:%s", err)
		c.JSON(resultpkg.Error(err))
		return
	}
	t.logger.Info("[访问用户信息-返回]:%+v", user)
	c.JSON(http.StatusOK, resultpkg.Ok(user))

}
