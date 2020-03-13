package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	adminPassword string = "adminPassword"
	adminName	string = "mittacy"
	TokenName string = "adminToken"
	BlogViews string = "BlogViews"
	adminInfo string = "adminInfo"
	recentArticles string = "recentArticles"
)
// SavePassword 缓存密码到redis
func SavePassword(pwd string) (string, error) {
	return BACKERROR, redisDB.Set(adminPassword, pwd, 0).Err()
}
// SaveToken 缓存token到redis
func SaveToken(token string) (string, error) {
	return BACKERROR, redisDB.Set(TokenName, token, 6*time.Hour).Err()
}
// IncrBlogViews 缓存博客浏览增加量到redis
func IncrBlogViews() (string, error) {
	return BACKERROR, redisDB.Incr(BlogViews).Err()
}
// SaveBlogViews 将redis缓存的博客浏览量存到mysql
func SaveBlogViews() error {
	views, err := redisDB.GetSet(BlogViews, 0).Result()
	if err != nil {
		return err
	}
	addNum, _ := strconv.Atoi(views)
	if addNum == 0 {
		return nil
	}
	// 加到mysql里
	stmt, err := mysqlDB.Prepare("UPDATE admin SET views = views + ? limit 1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(addNum); err != nil {
		return err
	}
	return nil
}
// SaveAdminInfo 缓存admin信息到redis
func SaveAdminInfo(adminJson ...[]byte) (string, error) {
	if adminJson != nil {
		return BACKERROR, redisDB.Set(adminInfo, adminJson, 0).Err()
	}
	sql := "SELECT name, views, cname, introduce, github, mail, bilibili FROM admin limit 1"
	var admin Admin
	err := mysqlDB.Get(&admin, sql)
	if err != nil {
		return BACKERROR, err
	}
	adminJson_new, _ := json.Marshal(admin)
	err = redisDB.Set(adminInfo, adminJson_new, 0).Err()
	if err != nil {
		fmt.Println("err: ", err)
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}
// SaveRecentArticles 缓存首页最近更改的五篇文章
func SaveRecentArticles() (string, error) {
	sql := "SELECT id, created_at, updated_at, title, views FROM article ORDER BY updated_at DESC limit 5"
	var articles []Article
	err := mysqlDB.Select(&articles, sql)
	if err != nil {
		return BACKERROR, err
	}
	articlesJson, _ := json.Marshal(articles)
	err = redisDB.Set(recentArticles, articlesJson, 0).Err()
	if err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}