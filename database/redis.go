package database

import "github.com/go-redis/redis"

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
	return nil
}

// CloseRedis 关闭redis连接
func CloseRedis() {
	RedisDB.Close()
}
