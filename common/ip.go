package common

import "github.com/crazychat/blog-gin/database"

func CheckIPRequestTimes(ip string) bool {
	// 判断ip是否存在
	exists, err := database.RedisDB.Exists(ip).Result()
	if err != nil {
		return false
	}
	if exists == 0 {
		// ip不存在, 允许请求
		return true
	}
	// ip存在，判断请求次数是否超过五次
	times, err := database.RedisDB.Get(ip).Result()
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
	return database.RedisDB.Incr(ip).Err()
}

func DelIP(ip string) error {
	return database.RedisDB.Del(ip).Err()
}
