package dto

import "github.com/gin-gonic/gin"


type RouterInfo struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
}