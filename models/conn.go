package models

import (
	"github.com/jmoiron/sqlx"
	
	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB // DB连接

func init() {
	// 连接数据库
	OpenConn()
	// 创建管理员信息
	if msg, err := CreateAdmin(); err != nil {
		panic(msg)
	}
}

// OpenConn 连接mysql数据库
func OpenConn() {
	// 获取配置文件数据
	par := SQL_USER + ":" + SQL_PASSWORD + "@tcp(" + SQL_HOST + ")/" + SQL_DATABASE + "?charset=utf8"

	var err error
	db, err = sqlx.Open(SQL_TYPE, par)
	if err != nil {
		panic("Failed to connect database")
	}
}

// GetDB 返回db
func GetDB() *sqlx.DB {
	if db != nil {
		return db
	}
	return nil
}