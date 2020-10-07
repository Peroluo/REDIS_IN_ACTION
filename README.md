# REDIS_IN_ACTION

### 一、Redis 基本数据类型

> 1. string： 存储的值是字符串
> 2. List： 链表存储的值是字符串
> 3. Set：独一无二的字符串链表
> 4. Hash: 无序散列列表
> 5. Zset: 由member和score组成的有序映射

### 二、Redis 发布订阅模式

1. 发布者

```go
client := redis.NewClient(&redis.Options{
  Addr:     "localhost:6379",
  Password: "",
  DB:       0,
})
err := client.Publish("message", "有人发送消息了").Err()
if err != nil {
  fmt.Println("redis publish 发布失败")
}
fmt.Println("redis publish 发布消息")
```

2. 订阅者

```go
client := redis.NewClient(&redis.Options{
  Addr:     "localhost:6379",
  Password: "",
  DB:       0,
})
pubsub := client.Subscribe("message")
_, err := pubsub.Receive()
if err != nil {
  fmt.Println(err)
  return
}
ch := pubsub.Channel()

for msg := range ch {
  fmt.Println(msg.Channel, msg.Payload)
}
```

> PUBLISH和SUBSCRIBE的缺陷在于客户端必须一直在线才能接收到消息，断线可能会导致客户端丢失消息，除此之外，旧版的redis可能会由于订阅者消费不够快而变的不稳定导致崩溃，甚至被管理员杀掉

### 三、Redis 事务



### 四、Redis 数据备份

1. save保存当前的数据

```shell
redis 127.0.0.1:6379> SAVE 
```

> save保存完数据后，会生产一个dump.rdb文件

> save数据恢复， 把上述的rdb文件移入redis安装目录，重启服务即可

