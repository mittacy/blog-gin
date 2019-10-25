package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

var cfg *ini.File // 配置文件
var db *sqlx.DB   // gorm连接

var (
	adminTableSQL string = `CREATE TABLE admin (
		id tinyint NOT NULL AUTO_INCREMENT,
		name varchar(50) NOT NULL,
		password varchar(255) NOT NULL,
		views int unsigned DEFAULT "0",
		cname varchar(50) DEFAULT NULL,
		introduce varchar(255) DEFAULT NULL,
		github varchar(100) DEFAULT NULL,
		mail varchar(100) DEFAULT NULL,
		PRIMARY KEY (id),
		UNIQUE KEY name (name)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	categoryTableSQL string = `CREATE TABLE category (
		id tinyint unsigned NOT NULL AUTO_INCREMENT,
		title varchar(50) NOT NULL,
		article_count smallint unsigned DEFAULT 0,
		PRIMARY KEY (id),
		UNIQUE KEY title (title)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	contentTableSQL string = `CREATE TABLE content (
		id smallint unsigned NOT NULL AUTO_INCREMENT,
		con text,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	articleTableSQL string = `CREATE TABLE article (
		id int(10) NOT NULL AUTO_INCREMENT,
		created_at datetime NOT NULL,
		category_id tinyint unsigned NOT NULL,
		title varchar(100) NOT NULL,
		content_id smallint unsigned NOT NULL,
		views mediumint unsigned DEFAULT 0,
		assists mediumint unsigned DEFAULT 0,
		PRIMARY KEY (id),
		foreign key(category_id) references category(id),
		foreign key(content_id) references content(id)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
)

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

// OpenConn 连接mysql数据库
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
	db, err = sqlx.Open(sqlType, par)
	if err != nil {
		fmt.Println("连接数据库失败", err.Error())
		panic("Failed to connect database")
	}
	fmt.Println("连接数据库成功")
}

// CreateTables 创建表格
func CreateTables() {
	// 创建表格
	fmt.Println("创建表格...")
	tables := []string{"admin", "category", "content", "article"}
	tableSQLs := []string{adminTableSQL, categoryTableSQL, contentTableSQL, articleTableSQL}
	for i := 0; i <= 3; i++ {
		if err := CreateTable(tables[i], tableSQLs[i]); err != nil {
			panic(err)
		}
	}
	fmt.Println("创建表格成功")
	// 创建管理员信息
	// if msg, err := CreateAdmin(); err != nil {
	// 	panic(msg)
	// }
}

// IsTableExist 判断表格是否存在
func IsTableExist(tableName string) bool {
	row := db.QueryRow("select table_name from information_schema.TABLES WHERE table_name = ?;", tableName)
	if err := row.Scan(); err == sql.ErrNoRows {
		return false
	}
	return true
}

// CreateTable 创建单个表格
func CreateTable(tableName, sql string) error {
	if IsTableExist(tableName) {
		fmt.Println(tableName, "表格已存在")
		return nil
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println(tableName, "创建表格成功")
	return nil
}

// GetDB 返回db
func GetDB() *sqlx.DB {
	if db != nil {
		return db
	}
	return nil
}

// GetCfg 获取配置文件连接
func GetCfg() *ini.File {
	return cfg
}
