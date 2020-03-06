package models

import "time"

const (
	adminPassword string = "adminPassword"
	adminName	string = "mittacy"
	TokenName string = "adminToken"
	BlogViews string = "BlogViews"
)
// SavePassword 缓存密码到redis
func SavePassword(pwd string) error {
	if err := redisDB.Set(adminPassword, pwd, 0).Err(); err != nil {
		return err
	}
	return nil
}
// SaveToken 保存token到redis
func SaveToken(token string) error {
	// todo 修改时间为2 * hour
	if err := redisDB.Set(TokenName, token, 24*time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

// todo SaveBlogViews 定时将Redis的缓存添加到MongoDB