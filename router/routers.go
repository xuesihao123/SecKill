package router

import (
	"SecKill/controller"
	routers "SecKill/router/middleware"
	"SecKill/service"
	"github.com/gin-gonic/gin"
)

func RoutersInit()*gin.Engine{
	r := gin.Default()
	//Restful风格在API的确定上不能使用动词
	//进入主页后使用主页的API
	r.Use(routers.Cors())
	r.POST("/main",controller.Login)
	r.POST("/register",controller.Register)

	userGroup := r.Group("/user")
	userGroup.Use(routers.JWTAuthMiddleware())
	{
		userGroup.POST("/fail" , controller.FailUser)
		userGroup.GET("/seckill/:name",controller.SecKillDiscount)
		userGroup.POST("/discount",controller.CreateDiscount)
	}
	service.RunUpdate()
	return r
}
