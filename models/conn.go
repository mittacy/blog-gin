package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var cfg *ini.File   // 配置文件
var gormDb *gorm.DB // gorm连接
var sqlDb *sql.DB   // sql连接

func init() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	// 连接数据库
	Conn()
	// 创建表格
	CreateTables()
}

// Conn 打开sql、gorm数据库连接
func Conn() {
	// 获取配置文件数据
	sec, _ := cfg.GetSection("database")
	sqlType := sec.Key("TYPE").MustString("debug")
	sqlName := sec.Key("USER").MustString("debug")
	sqlPassword := sec.Key("PASSWORD").MustString("debug")
	sqlHOST := sec.Key("HOST").MustString("debug")
	sqlDatabase := sec.Key("DATABASE").MustString("debug")
	par := sqlName + ":" + sqlPassword + "@tcp(" + sqlHOST + ")/" + sqlDatabase + "?charset=utf8"
	GormConn(sqlType, par)
	SQLConn(sqlType, par)
}

// GormConn gorm连接mysql数据库
func GormConn(sqlType, par string) {
	var err error
	gormDb, err = gorm.Open(sqlType, par)
	if err != nil {
		fmt.Println("gorm连接数据库失败", err.Error())
		panic("gorm failed to connect database")
	}
	fmt.Println("gorm连接数据库成功")
}

// SQLConn sql连接mysql数据库
func SQLConn(sqlType, par string) {
	var err error
	sqlDb, err = sql.Open(sqlType, par)
	if err != nil {
		fmt.Println("sql连接数据库失败", err.Error())
		panic("sql failed to connect database")
	}
	fmt.Println("sql连接数据库成功")
}

// CreateTables 创建表格
func CreateTables() {
	// 创建表格
	fmt.Println("创建表格...")
	gormDb.SingularTable(true)
	if err := gormDb.CreateTable(&Admin{}).Error; err != nil {
		fmt.Println("表格以及存在")
	} else {
		fmt.Println("创建表格成功...")
	}
	// 创建管理员信息
	if err := CreateAdmin(); err != nil {
		panic("failed create admin")
	}
}

// GormDB 返回gorm.DB
func GormDB() *gorm.DB {
	if gormDb != nil {
		return gormDb
	}
	return nil
}

// SQLDB 返回sql.DB
func SQLDB() *sql.DB {
	if sqlDb != nil {
		return sqlDb
	}
	return nil
}

func Cfg() *ini.File {
	return cfg
}
