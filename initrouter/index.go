package initrouter

import (
	"REDIS_IN_ACTION/router"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() (Router *gin.Engine) {
	 Router = gin.Default()
	group := Router.Group("") 
	router.InitUserRouter(group)              
	return 
}
