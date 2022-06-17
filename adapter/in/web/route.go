package web

import "github.com/gin-gonic/gin"

func InitRoute(engine *gin.Engine) {
	userRouter(engine)
	roleRouter(engine)
}

func userRouter(engine *gin.Engine) {
	group := engine.Group("/")
	group.GET("/route", ctl.Info)
}

func roleRouter(engine *gin.Engine) {
	group := engine.Group("/")
	group.GET("/route", ctl.Info)
}
