package ginextends

import "github.com/gin-gonic/gin"

type RouterInfos []RouterInfo



type RouterInfo struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
	// 接口标志
}

type Routerable interface {
	Router() RouterInfos
}

