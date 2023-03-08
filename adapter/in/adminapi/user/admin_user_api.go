package adminapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	command "github.com/lj19950508/ddd-demo-go/application/command/user"
	query "github.com/lj19950508/ddd-demo-go/application/query/user"
	"github.com/lj19950508/ddd-demo-go/pkg/ginextends"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/resultpkg"
)

//后台用户对App用户的接口
type AdminUserApi struct {
	userCommandService command.UserCommandService
	userQueryService   query.UserQueryService
	logger             logger.Interface
}

func (t *AdminUserApi) Router() ginextends.RouterInfos {
	return ginextends.RouterInfos{
		//默认使用user吧
		{Method: "GET", Path: "/admin/users", Handle: t.List},
		{Method: "POST", Path: "/admin/users", Handle: nil},
		{Method: "PUT", Path: "/admin/users", Handle: nil},
		{Method: "DELETE", Path: "/admin/users/:id", Handle: nil},
	}
}

func NewAdminUserApi(userCommandService command.UserCommandService, userQueryService query.UserQueryService, logger logger.Interface) *AdminUserApi {
	return &AdminUserApi{
		userCommandService: userCommandService,
		userQueryService:   userQueryService,
		logger:             logger,
	}
}


func (t *AdminUserApi) List(c *gin.Context) {
	var cond query.UserPageQuery
	err:=c.Bind(&cond)
	// validator.New().Struct(&cond)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultpkg.Fail(err.Error()))
		return
	}
	users, err := t.userQueryService.FindList(&cond)
	if err != nil {
		c.JSON(resultpkg.Error(err))
		return
	}
	c.JSON(http.StatusOK, resultpkg.Ok(users))

}


