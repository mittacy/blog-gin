package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var RedisDB *redis.Client

// ConnectRedis 获取Redis连接
func ConnectRedis() error {
	RedisDB = redis.NewClient(&redis.Options{
		Addr: localhost + ":6379",
		Password: "",
		DB: 0,
	})
	_, err := RedisDB.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("redis连接成功")
	return nil
}

// CloseRedis 关闭redis连接
func CloseRedis() {
	RedisDB.Close()
}
// RedisIncr redis对象加1
func RedisIncr(key string) error {
	return RedisDB.Incr(key).Err()
}
// RedisIncr redis对象减1
func RedisDel(key string) error {
	return RedisDB.Del(key).Err()
}
// RedisIncr redis获取对象值
func RedisGet(key string) (string, error) {
	return RedisDB.Get(key).Result()
}
// RedisSet redis保存对象
func RedisSet(key, val string, duration time.Duration) error {
	return RedisDB.Set(key, val, duration).Err()
}
