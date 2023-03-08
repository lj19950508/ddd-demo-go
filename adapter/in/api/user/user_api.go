package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	command "github.com/lj19950508/ddd-demo-go/application/command/user"
	query "github.com/lj19950508/ddd-demo-go/application/query/user"
	"github.com/lj19950508/ddd-demo-go/pkg/ginextends"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/resultpkg"
)

type UserApi struct {
	userCommandService command.UserCommandService
	userQueryService   query.UserQueryService
	logger             logger.Interface
}

func (t *UserApi) Router() ginextends.RouterInfos {
	return ginextends.RouterInfos{
		//默认使用user吧
		{Method: "GET", Path: "/v1/users/:id", Handle: t.Info},
	}
}

func NewUserApi(userCommandService command.UserCommandService, userQueryService query.UserQueryService, logger logger.Interface) *UserApi {
	return &UserApi{
		userCommandService: userCommandService,
		userQueryService:   userQueryService,
		logger:             logger,
	}
}



func (t *UserApi) Info(c *gin.Context) {
	id :=c.Param("id")
	userId ,err:= strconv.ParseInt(id,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultpkg.Fail(err.Error()))
		return
	}
	
	user, err := t.userQueryService.FindOne(&query.UserQuery{
		IdEq: &userId,
	})
	if err != nil {
		c.JSON(resultpkg.Error(err))
		return
	}
	c.JSON(http.StatusOK, resultpkg.OkData(user))

}
