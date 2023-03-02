package ginextends

import "github.com/gin-gonic/gin"

type RouterInfos []RouterInfo


type RouterInfo struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
}

type Routerable interface{
	Router() RouterInfos 
}

