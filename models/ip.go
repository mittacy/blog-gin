package models

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisDB *redis.Client

func GetRedisClient() (*redis.Client, error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err := redisDB.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return redisDB, err
}

func CheckIPRequestTimes(ip string) bool {
	// 判断ip是否存在
	exists, err := redisDB.Exists(ip).Result()
	if err != nil {
		return false
	}
	if exists == 0 {
		// ip不存在, 允许请求
		return true
	}
	// ip存在，判断请求次数是否超过五次
	times, err := redisDB.Get(ip).Result()
	if err != nil {
		return false
	}
	if times < "5" {
		// 少于5次, 允许请求
		return true
	}
	// 多于5次, 拒绝请求
	return false
}

func IncrIP(ip string) error {
	if err := redisDB.SetNX(ip, 0, 1*time.Minute).Err(); err != nil {
		return err
	}
	return redisDB.Incr(ip).Err()
}

func DelIP(ip string) error {
	return redisDB.Del(ip).Err()
}