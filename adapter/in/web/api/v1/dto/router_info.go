package dto

import "github.com/gin-gonic/gin"

//TODO youhua
type RouterInfo struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
}