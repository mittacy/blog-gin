package models

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB // DB连接

// OpenConn 连接mysql数据库
func OpenConn() error {
	// 获取配置文件数据
	par := SQL_USER + ":" + SQL_PASSWORD + "@tcp(" + SQL_HOST + ")/" + SQL_DATABASE
	db, _ = sqlx.Open(SQL_TYPE, par)
	if err := db.Ping(); err != nil {
		return err
	}
	// 创建管理员信息
	if _, err := CreateAdmin(); err != nil {
		return err
	}
	return nil
}

// GetDB 返回db
func GetDB() *sqlx.DB {
	if db != nil {
		return db
	}
	return nil
}