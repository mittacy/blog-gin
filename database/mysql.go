package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	localhost			string = "127.0.0.1"
	mysqlUser     string = "root"
	mysqlPassword string = "aini1314584"
	mysqlHost     string = localhost + ":3306"
	mysqlDatabase string = "blog"
)

var MysqlDB *sqlx.DB

// ConnectMysql 连接mysql数据库
func ConnectMysql() error {
	// 获取配置文件数据
	par := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ")/" + mysqlDatabase
	var err error
	MysqlDB, err = sqlx.Open("mysql", par)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}
	if err = MysqlDB.Ping(); err != nil {
		return err
	}
	return nil
}
// CloseMysql 关闭mysl
func CloseMysql() {
	MysqlDB.Close()
}