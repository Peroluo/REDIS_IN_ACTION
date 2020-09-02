package controller

import (
	"github.com/gin-gonic/gin"
	"REDIS_IN_ACTION/redisclient"
	"net/http"
)


// Regist 注册
func Regist(c *gin.Context) {
	var rdb = redisclient.RedisClient
	err := rdb.Set("key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get( "key").Result()
 
	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"msg":     "2323",
		"data":    val,
	})
}
