package database

import (
	"github.com/crazychat/blog-gin/model"
	"github.com/jmoiron/sqlx"
)

const (
	localhost			string = "127.0.0.1"
	mysqlUser     string = "root"
	mysqlPassword string = "aini1314584"
	mysqlHost     string = localhost + ":3306"
	mysqlDatabase string = "blog"
)

var InitAdmin = model.Admin {
	Name:      "mittacy",
	Password:  "123456",
	Views:		 7156,
	Cname:     "陈铭涛",
	Introduce: "就读佛山大学 - 大三 - 计算机系",
	Github:    "https://github.com/crazychat",
	Mail:      "mail@mittacy.com",
	Bilibili:  "https://space.bilibili.com/384942135",
}

var MysqlDB *sqlx.DB

// ConnectMysql 连接mysql数据库
func ConnectMysql() error {
	// 获取配置文件数据
	par := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ")/" + mysqlDatabase
	MysqlDB, _ = sqlx.Open("mysql", par)
	if err := MysqlDB.Ping(); err != nil {
		return err
	}
	return nil
}
// CloseMysql 关闭mysl
func CloseMysql() {
	MysqlDB.Close()
}