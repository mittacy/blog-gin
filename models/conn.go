package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var cfg *ini.File // 配置文件
var db *gorm.DB   // gorm连接
var sqlDb *sql.DB // sql连接

func init() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	// 连接数据库
	OpenConn()
	// 创建表格
	CreateTables()
}

// OpenConn gorm连接mysql数据库
func OpenConn() {
	// 获取配置文件数据
	sec, _ := cfg.GetSection("database")
	sqlType := sec.Key("TYPE").MustString("debug")
	sqlName := sec.Key("USER").MustString("debug")
	sqlPassword := sec.Key("PASSWORD").MustString("debug")
	sqlHOST := sec.Key("HOST").MustString("debug")
	sqlDatabase := sec.Key("DATABASE").MustString("debug")
	par := sqlName + ":" + sqlPassword + "@tcp(" + sqlHOST + ")/" + sqlDatabase + "?charset=utf8"

	var err error
	db, err = gorm.Open(sqlType, par)
	if err != nil {
		fmt.Println("gorm连接数据库失败", err.Error())
		panic("gorm failed to connect database")
	}
	sqlDb = db.DB()
	fmt.Println("gorm连接数据库成功")
}

// CreateTables 创建表格
func CreateTables() {
	// 创建表格
	fmt.Println("创建表格...")
	db.SingularTable(true)
	if err := db.CreateTable(&Admin{}, &Category{}, &Article{}).Error; err != nil {
		fmt.Println("表格以及存在")
	} else {
		fmt.Println("创建表格成功...")
	}
	// 创建管理员信息
	if msg, err := CreateAdmin(); err != nil {
		panic(msg)
	}
}

// GetDB 返回gorm.DB
func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	return nil
}

// GetSQLDB 返回sql.DB
func GetSQLDB() *sql.DB {
	if sqlDb == nil {
		sqlDb = db.DB()
	}
	return sqlDb
}

// GetCfg 获取配置文件连接
func GetCfg() *ini.File {
	return cfg
}
