package xredis

import (
	"github.com/go-redis/redis"
	"github.com/myskull/common/xLog"
	"github.com/myskull/common/xconfig"
)

type Redis struct {
	// redis的操作
	client *redis.Client
}

var redisConn = Redis{}

// 持久化处理
func New() error {
	if redisConn.client != nil {
		return nil
	}
	address := xconfig.GetStr("redis", "address", "127.0.0.1:6379")
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})
	redisConn.client = client
	err := ping()
	if err != nil {
		xLog.Error("redis链接出错:%v", err.Error())
	}
	return err
}

func ping() error {
	cmd := redisConn.client.Ping()
	return cmd.Err()
}
