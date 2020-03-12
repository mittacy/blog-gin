package models

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	RUN_MODE     string = "debug"
	SQL_TYPE     string = "mysql"
	SQL_USER     string = "root"
	SQL_PASSWORD string = "aini1314584"
	SQL_HOST     string = "127.0.0.1:3306"
	SQL_DATABASE string = "blog"
)

var redisDB	*redis.Client
var mysqlDB *sqlx.DB

// ConnectMysql 连接mysql数据库
func ConnectMysql() error {
	// 获取配置文件数据
	par := SQL_USER + ":" + SQL_PASSWORD + "@tcp(" + SQL_HOST + ")/" + SQL_DATABASE
	mysqlDB, _ = sqlx.Open(SQL_TYPE, par)
	if err := mysqlDB.Ping(); err != nil {
		return err
	}
	// 创建管理员信息
	if _, err := CreateAdmin(); err != nil {
		return err
	}
	return nil
}
// CloseMysql 关闭mysl
func CloseMysql() {
	redisDB.Close()
}
// ConnectRedis 获取Redis连接
func ConnectRedis() error {
	redisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
// CloseRedis 关闭redis连接
func CloseRedis() {
	redisDB.Close()
}
// startTimer 定时清理缓存任务任务
func StartTimer() {
	go func() {
		for {
			SaveBlogViews()
			SaveRecentArticles()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 4, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
