**使用golang和gin框架开发个人博客后台**

**Mysql初始化**

~~~mysql
/* 创建数据库 */
$ create database blog;
/* 使用数据库 */
$ use blog;
/* 创建admin */
$ CREATE TABLE admin (
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
/* 创建category */
$ CREATE TABLE category (
id integer unsigned NOT NULL AUTO_INCREMENT,
title varchar(50) NOT NULL,
article_count integer unsigned DEFAULT 0,
PRIMARY KEY (id),
UNIQUE KEY title (title)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/* 创建article */
$ CREATE TABLE article (
id integer unsigned NOT NULL AUTO_INCREMENT,
created_at varchar(11) NOT NULL DEFAULT '2020-02-02',
updated_at varchar(11) NOT NULL DEFAULT '',
category_id integer unsigned NOT NULL,
views integer unsigned DEFAULT 0,
title varchar(50) NOT NULL,
content longtext,
PRIMARY KEY (id),
foreign key(category_id) references category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4; 
/* 创建触发器1 */
$ create trigger tr_article_after_insert after insert
on article for each row
update category set article_count=article_count+1 where id = new.category_id;
/* 创建触发器2 */
$ create trigger tr_article_after_delete after delete
on article for each row
update category set article_count=article_count-1 where id = old.category_id;
/* 创建触发器3 */
$ DELIMITER ||

$ create trigger tr_article_after_update after update on article
for each row
BEGIN
	update category set article_count=article_count-1 where id = old.category_id;
	update category set article_count=article_count+1 where id = new.category_id;
END||

$ DELIMITER ;
~~~

**API接口文档**

~~~go
// 验证登录状态
GET         /api/verify
// 获取管理员信息
GET			/api/admin
// 登录，参数: name, password
POST 		/api/admin
// 修改管理员信息, 参数: cname, introduce, github, mail
PUT			/api/admin
// 修改管理员密码, 参数: password
PUT         /api/admin/setpwd
// 博客浏览量加1
GET			/api/admin/addviews


// 分页获取分类, 参数: pageIndex
GET  		/api/category_page/:num
// 显示某个分类所有的文章, 参数: id
GET			/api/category/:id
// 添加分类, 参数: title
POST 		/api/category
// 修改分类名, 参数: id, title
PUT			/api/category
// 删除分类, 参数: id
Delete 	/api/category
// 根据id获取分类title
GET     /api/category_name/:id


// 分页获取文章, 参数: pageIndex
GET			/api/article_page/:num
// 显示某篇文章
GET     /api/article/:id
// 添加文章，参数: created_at, category_id, title , content
POST		/api/article
// 修改文章，参数: id, updated_at, category_id, title , content
PUT			/api/article
// 删除文章，参数: id
Delete	/api/article
// 文章添加访问量, 参数: id
POST		/api/article/addViews
~~~

