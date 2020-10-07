package router

import (
	"REDIS_IN_ACTION/controller"
	"github.com/gin-gonic/gin"
)

// InitUserRouter 用户路由模块
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("")
	{
		UserRouter.GET("/lessonString", controller.LessonString) // redis string数据类型
		UserRouter.GET("/lessonList", controller.LessonList)     // redis list数据类型
		UserRouter.GET("/lessonHash", controller.LessonHash)     // redis hash数据类型
		UserRouter.GET("/lessonSet", controller.LessonSet)       // redis set数据类型
		UserRouter.GET("/lessonZset", controller.LessonZset)     // redis set数据类型
		UserRouter.GET("/publisher", controller.Publisher)       // redis set数据类型
	}
}
