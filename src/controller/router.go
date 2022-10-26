package controller

import (
	"github.com/gin-gonic/gin"
	"reaven-server/src/controller/api"
	"reaven-server/src/controller/middleware"
)

func Router(r *gin.Engine) {
	// 全局错误捕获中间件
	r.Use(middleware.Recover)
	user := r.Group("api/v1/user")
	{
		// 不使用Token校验中间件的接口
		user.POST("/test1", api.Test)
		user.POST("/login", api.Login)
		// Token校验中间件
		user.Use(middleware.VerifyToken)
		// 使用Token校验中间件的接口
		user.POST("/test2", api.Test)

	}
}
