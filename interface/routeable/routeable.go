package routeable

import "github.com/gin-gonic/gin"

type Routeable interface {
	//groupinfo

	GetGroupPath() string
	GetHandleFunc() gin.RoutesInfo
}
