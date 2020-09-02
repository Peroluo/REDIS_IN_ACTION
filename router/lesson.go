package router

import (
	"REDIS_IN_ACTION/controller"
	"github.com/gin-gonic/gin"
)

// InitUserRouter 用户路由模块
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("changePassword", controller.Regist)     // 修改密码
	}
}
