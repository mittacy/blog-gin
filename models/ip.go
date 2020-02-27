package models

import (
	"fmt"
	"github.com/go-redis/redis"
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
	fmt.Println("连接redis成功, redisDB: ", redisDB)
	return redisDB, err
}

func CheckIPRequestTimes(ip string) bool {
	// 判断ip是否存在
	exists, err := redisDB.Exists(ip).Result()
	if err != nil {
		fmt.Println("err: ", err)
		return false
	}
	if exists == 0 {
		fmt.Println("ip不存在, 允许请求...")
		return true
	}
	fmt.Println("ip存在")
	// ip存在，判断请求次数是否超过五次
	times, err := redisDB.Get(ip).Result()
	if err != nil {
		fmt.Println(ip, " err: ", err)
		return false
	}
	fmt.Println(ip, " request times: ", times)
	if times < "5" {
		fmt.Println("少于5次, 允许请求")
		return true
	}
	fmt.Println("多于5次, 拒绝请求")
	return false
}