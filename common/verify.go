package common

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/database"
	"io"
	"time"
)

const (
	ipMaxTimes string = "5"
)

// Encryption 密码加密
func Encryption(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
// CreateToken 生成token
func CreateToken(pwd string) (string, error) {
	encrpty := []byte(time.Now().String() + pwd)
	h := sha256.New()
	_, err := h.Write(encrpty)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
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
	times, err := database.RedisGet(ip)
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
	return database.RedisIncr(ip)
}
// DelIP 删除ip记录
func DelIP(ip string) error {
	return database.RedisDel(ip)
}

