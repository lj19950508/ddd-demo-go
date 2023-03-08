package ginextends

import (
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

var TokenWithPermissionHandler = func(ctx *gin.Context) {
	//TODO ，根据前缀判断用户权限，根据prefix判断是否需要权限，然后根据ctx里面的用户信息判断是否拥有权限

}

var TokenHandler = func(ctx *gin.Context) {
}
