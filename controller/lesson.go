package controller

import (
	"REDIS_IN_ACTION/redisclient"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"reflect"
)

type doctor struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// LessonString Redis string数据类型
func LessonString(c *gin.Context) {
	jsonDoc, _ := json.Marshal(doctor{"钟南山", 18})

	redisclient.RedisClient.Set("string:doctor", jsonDoc, 0)

	val, _ := redisclient.RedisClient.Get("string:doctor").Result()

	var d2 doctor

	json.Unmarshal([]byte(val), &d2)

	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"msg":     "redis数据类型string",
		"data":    d2,
	})
}

// LessonList Redis list数据类型 存储字符串数组
func LessonList(c *gin.Context) {
	listKey := "list:doctors"

	jsonDoc, _ := json.Marshal(doctor{"钟南山", 18})

	redisclient.RedisClient.RPush(listKey, jsonDoc) // 添加数据

	data, _ := redisclient.RedisClient.LRange(listKey, 0, -1).Result() // 获取列表

	doctors := []doctor{}

	for _, v := range data {
		d := doctor{}
		json.Unmarshal([]byte(v), &d)
		doctors = append(doctors, d)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"msg":     "redis数据类型list",
		"data":    doctors,
	})
}

// LessonHash Redis hash数据类型
func LessonHash(c *gin.Context) {
	hashKey := "hash:doctor"

	mset := make(map[string]interface{})

	doctor := doctor{
		"罗国雄222",
		28,
	}

	//反射
	t := reflect.TypeOf(doctor)
	v := reflect.ValueOf(doctor)
	fmt.Println(t)
	fmt.Println(v)
	for k := 0; k < t.NumField(); k++ {
		// fmt.Println(t.Field(k).Name)            // 键名
		// fmt.Println(v.Field(k).Interface())     // 属性值
		// fmt.Println(t.Field(k).Tag.Get("json")) // tag
		mset[t.Field(k).Name] = v.Field(k).Interface()
	}

	redisclient.RedisClient.HMSet(hashKey, mset)

	data, _ := redisclient.RedisClient.HGetAll(hashKey).Result()

	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"msg":     "redis数据类型hash",
		"data":    data,
	})
}

// LessonSet Redis set数据类型
func LessonSet(c *gin.Context) {
	setKey := "set:doctor"
	newDoctor := doctor{
		"罗国雄2",
		28,
	}
	doctors := []doctor{}

	jsonDoctor, _ := json.Marshal(newDoctor)

	// 新增
	isHas, err := redisclient.RedisClient.SAdd(setKey, jsonDoctor).Result() //新增
	fmt.Println(isHas, err)

	//删除
	// isHas, err := redisclient.RedisClient.SRem(setKey, jsonDoctor).Result()
	// fmt.Println(isHas, err) 1 删除

	// 获取全部
	setList, _ := redisclient.RedisClient.SMembers(setKey).Result()

	for _, v := range setList {
		d := doctor{}
		json.Unmarshal([]byte(v), &d)
		doctors = append(doctors, d)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"msg":     "redis数据类型set",
		"data":    doctors,
	})
}

// LessonZset Redis zset数据类型
func LessonZset(c *gin.Context) {
	zsetKey := "zset:doctor"

	ranking := []redis.Z{
		redis.Z{
			Score: 20, Member: "钟南山2",
		},
	}

	redisclient.RedisClient.ZAdd(zsetKey, ranking...)

	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"msg":     "redis数据类型zset",
		"data":    nil,
	})
}
