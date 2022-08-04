package xredis

import (
	"github.com/go-redis/redis"
	"time"
)

/**
expire:秒数
*/
func Set(key string, value interface{}, expire int64) error {
	err := New()
	if err != nil {
		return err
	}
	cmd := redisConn.client.Set(key, value, time.Second*time.Duration(expire))
	return cmd.Err()
}

func Get(key string) string {
	err := New()
	if err != nil {
		return ""
	}
	cmd := redisConn.client.Get(key)
	return cmd.Val()
}

func Del(key string) error {
	err := New()
	if err != nil {
		return err
	}
	cmd := redisConn.client.Del(key)
	return cmd.Err()
}

func HSet(key, field string, value interface{}) error {
	err := New()
	if err != nil {
		return err
	}
	cmd := redisConn.client.HSet(key, field, value)
	return cmd.Err()
}

func HGet(key, field string) string {
	err := New()
	if err != nil {
		return ""
	}
	cmd := redisConn.client.HGet(key, field)
	return cmd.Val()
}

func HDel(key string, field ...string) error {
	err := New()
	if err != nil {
		return err
	}
	cmd := redisConn.client.HDel(key, field...)
	return cmd.Err()
}

func SetNx(key string, value interface{}, expire int64) error {
	err := New()
	if err != nil {
		return err
	}
	cmd := redisConn.client.SetNX(key, value, time.Second*time.Duration(expire))
	return cmd.Err()
}

func HSetNx(key, field string, value interface{}) error {
	err := New()
	if err != nil {
		return err
	}
	cmd := redisConn.client.HSetNX(key, field, value)
	return cmd.Err()
}

func Client() *redis.Client {
	return redisConn.client
}
