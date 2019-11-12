package models

import (
	"fmt"

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
views integer unsigned DEFAULT 0,
assists integer unsigned DEFAULT 0,
title varchar(100) NOT NULL,
content text,
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
END||

DELIMITER ;
*/
