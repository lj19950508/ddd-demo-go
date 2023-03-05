package ginextends

import "github.com/gin-gonic/gin"

type RouterInfos []RouterInfo

type RequireLevel string

type RouterInfo struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
	NoAuth bool //不需要认证 默认为false  就是默认需要认证
}

type Routerable interface {
	Router() RouterInfos
}
