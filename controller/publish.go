package controller

import (
	"REDIS_IN_ACTION/redisclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Publisher 发布消息
func Publisher(c *gin.Context) {
	err := redisclient.RedisClient.Publish("message", "有人发送消息了").Err()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": 200,
			"msg":     "redis publish 发布失败",
			"data":    nil,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"msg":     "redis publish 发布消息",
		"data":    nil,
	})
}
