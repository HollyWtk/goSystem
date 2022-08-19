package web

import (
	"github.com/gin-gonic/gin"
	"merchantService/db"
	"merchantService/server/web/apply/controller"
	"merchantService/server/web/apply/middleware"
)

func Init(router *gin.Engine) {
	//测试数据库，并且初始化数据库
	db.TestDB()
	//还有别的初始化方法
	initRouter(router)
}

func initRouter(router *gin.Engine) {
	router.Use(middleware.Cors())
	router.POST("/merchant/apply", controller.DefaultApplyController.AddApply)
	router.GET("/merchant/apply", controller.DefaultApplyController.QueryApply)
}
