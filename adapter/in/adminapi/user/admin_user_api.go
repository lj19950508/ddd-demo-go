package adminapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	command "github.com/lj19950508/ddd-demo-go/application/command/user"
	query "github.com/lj19950508/ddd-demo-go/application/query/user"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/resultpkg"
	"github.com/lj19950508/ddd-demo-go/pkg/route"
)

//后台用户对App用户的接口
type AdminUserApi struct {
	userCommandService command.UserCommandService
	userQueryService   query.UserQueryService
	logger             logger.Interface
}

func (t *AdminUserApi) Route() *route.HttpRoutes {
	return &route.HttpRoutes{
		//默认使用user吧
		{Pattern:"GET /admin/users",Handler: t.List},
		{Pattern:"POST /admin/users",Handler: t.Create},
		{Pattern:"PUT /admin/users",Handler: t.Update},
		{Pattern:"DELETE /admin/users/:id",Handler: t.Delete},
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
	err := c.Bind(&cond)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultpkg.Fail(err.Error()))
		return
	}
	users, err := t.userQueryService.FindList(&cond)
	if err != nil {
		c.JSON(resultpkg.Error(err))
		return
	}
	c.JSON(http.StatusOK, resultpkg.OkData(users))
}

func (t *AdminUserApi) Create(c *gin.Context) {
	var cmd command.CreateCommand
	err := c.Bind(&cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultpkg.Fail(err.Error()))
		return
	}

	err = t.userCommandService.Create(&cmd)
	if err != nil {
		c.JSON(resultpkg.Error(err))
		return
	}
	c.JSON(http.StatusOK, resultpkg.Ok())
}

func (t *AdminUserApi) Update(c *gin.Context) {
	var cmd command.UpdateCommand
	err := c.Bind(&cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultpkg.Fail(err.Error()))
		return
	}
	err = t.userCommandService.Update(&cmd)
	if err != nil {
		c.JSON(resultpkg.Error(err))
		return
	}
	c.JSON(http.StatusOK, resultpkg.Ok())
}

func (t *AdminUserApi) Delete(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultpkg.Fail(err.Error()))
		return
	}

	err = t.userCommandService.Delete(userId)
	if err != nil {
		c.JSON(resultpkg.Error(err))
		return
	}
	c.JSON(http.StatusOK, resultpkg.Ok())
}
