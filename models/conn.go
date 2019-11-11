package models

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

var cfg *ini.File // 配置文件
var db *sqlx.DB   // DB连接

func init() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
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

/*	1. 创建数据库blog	*/
//	create database blog;

/*  2. 创建表格admin, category, article	*/
//	use blog;
/*	adminTableSQL
CREATE TABLE admin (
id integer unsigned NOT NULL AUTO_INCREMENT,
name varchar(50) NOT NULL,
password varchar(255) NOT NULL,
views integer unsigned DEFAULT "0",
cname varchar(50) DEFAULT NULL,
introduce varchar(255) DEFAULT NULL,
github varchar(100) DEFAULT NULL,
mail varchar(100) DEFAULT NULL,
PRIMARY KEY (id),
UNIQUE KEY name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/
/*	categoryTableSQL
CREATE TABLE category (
id integer unsigned NOT NULL AUTO_INCREMENT,
title varchar(50) NOT NULL,
article_count integer unsigned DEFAULT 0,
PRIMARY KEY (id),
UNIQUE KEY title (title)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/
/*	articleTableSQL
CREATE TABLE article (
id integer unsigned NOT NULL AUTO_INCREMENT,
created_at datetime DEFAULT NOW(),
category_id integer unsigned NOT NULL,
title varchar(100) NOT NULL,
content text,
views integer unsigned DEFAULT 0,
assists integer unsigned DEFAULT 0,
PRIMARY KEY (id),
foreign key(category_id) references category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/

/*  3. 创建触发器	*/
/*	tr_article_after_insert: 添加文章后，自动修改分类文章数
create trigger tr_article_after_insert after insert
on article for each row
update category set article_count=article_count+1 where id = new.category_id;
*/

/*	tr_article_after_delete: 删除文章后，自动修改分类文章数
create trigger tr_article_after_delete after delete
on article for each row
update category set article_count=article_count-1 where id = old.category_id;
*/

/*	tr_article_after_update: 修改文章分类后，自动修改分类文章数
DELIMITER ||

create trigger tr_article_after_update after update on article
for each row
BEGIN
	update category set article_count=article_count-1 where id = old.category_id;
	update category set article_count=article_count+1 where id = new.category_id;
END;
*/
