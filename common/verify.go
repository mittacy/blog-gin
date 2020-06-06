package common

import (
	"fmt"
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/database"
)

const (
	ipMaxTimes string = "5"
)
// SaveToken 保存token到redis
func SaveToken(token string) {
	cache.SetToken(token)
}
// CheckIPRequestPower 检查ip是否访问次数过多
func CheckIPRequestPower(ip string) bool {
	// todo 修复ip限制功能
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
	times, err := cache.RedisGet(ip)
	if err != nil {
		fmt.Println(times, err)
		return false	// redis崩溃, 停止所有ip登录
	}
	if times < "5" {
		// 少于5次, 允许请求
		return true
	}
	// 多于5次, 拒绝请求
	return false
}
// IncrIP ip记录加1
func IncrIP(ip string) error {
	return cache.RedisIncr(ip)
}
// DelIP 删除ip记录
func DelIP(ip string) error {
	return cache.RedisDel(ip)
}

