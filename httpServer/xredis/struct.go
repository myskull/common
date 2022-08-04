package xredis

import (
	"fmt"
	"gitee.com/myskull/common/httpServer/xconfig"
	"github.com/go-redis/redis"
)

type Redis struct {
	// redis的操作
	client *redis.Client
}

var redisConn = Redis{}

// 持久化处理
func New() error {
	if redisConn.client != nil {
		err := ping()
		if err != nil {
			fmt.Println("redis链接已失效:", err.Error())
		} else {
			return nil
		}
	}
	address := xconfig.GetStr("redis", "address", "127.0.0.1:6379")
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})
	redisConn.client = client
	err := ping()
	if err != nil {
		fmt.Println("redis链接出错:", err.Error())
	}
	return err
}

func ping() error {
	cmd := redisConn.client.Ping()
	return cmd.Err()
}
