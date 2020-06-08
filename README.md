**使用golang和gin框架开发个人博客后台**

项目展示：[我的博客](https://blog.mittacy.com)

### 1. 项目结构

其中`js`、`css`、`index.html`为Vue打包文件，前端源码：[https://github.com/mittacy/blog-vue](https://github.com/mittacy/blog-vue)

~~~
.
├── main.go 主函数

├── router 路由目录
│   ├── middleware.go 中间件函数
│   └── router.go

├── controller 控制器目录
│   ├── admin_controller.go
│   ├── article_controller.go
│   ├── category_controller.go
│   └── log_controller.go

├── service 服务层目录，从缓存器和数据库获取数据
│   ├── admin_service.go
│   ├── article_service.go
│   └── category_service.go

├── repository 数据库操作目录
│   ├── admin_repository.go
│   ├── article_repository.go
│   └── category_repository.go

├── cache 缓存器
│   ├── admin_cache.go
│   ├── article_cache.go
│   ├── cache.go
│   ├── category_cache.go
│   ├── ip_cache.go
│   └── token.go

├── model 模型目录
│   ├── admin.go
│   ├── article.go
│   ├── category.go
│   └── id.go

├── database 数据库初始化目录
│   ├── mysql.go
│   └── redis.go

├── common 全局函数变量等
│   ├── json.go
│   ├── msg.go
│   ├── response.go
│   └── timer.go

├── config 设置目录
│   └── config.go

├── log 日志记录目录
│   ├── blogErr.log
│   └── log.go

└── utiles 工具目录
    └── tool.go
    
├── css 前端css
│   ├── ……
├── js 前端js
│   ├── ……
├── index.html 前端页面
├── go.mod
├── go.sum
├── README.md
~~~

一个api请求的访问顺序：`router` -> `controller` -> `service` -> `cache/repository` -> `model` -> `database`

### 2. 运行项目: 

#### 2.1 前提

+ golang环境

+ Redis环境

+ Mysql环境

  ~~~sql
  /* 创建数据库 */
  > create database blog;
  /* 使用数据库 */
  > use blog;
  /* 创建admin */
  > CREATE TABLE `admin` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `password` varchar(255) NOT NULL,
    `views` bigint unsigned NOT NULL DEFAULT '4224',
    `cname` varchar(50) DEFAULT NULL,
    `introduce` varchar(255) DEFAULT NULL,
    `github` varchar(100) DEFAULT NULL,
    `mail` varchar(100) DEFAULT NULL,
    `bilibili` varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
  /* 创建category */
  > CREATE TABLE `category` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(50) NOT NULL,
    `article_count` int unsigned DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `title` (`title`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
  /* 创建article */
  > CREATE TABLE `article` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `created_at` varchar(11) NOT NULL DEFAULT '2020-02-02',
    `updated_at` varchar(11) NOT NULL DEFAULT '',
    `category_id` int unsigned NOT NULL,
    `views` int unsigned DEFAULT '0',
    `title` varchar(50) NOT NULL,
    `content` longtext,
    PRIMARY KEY (`id`),
    KEY `category_id` (`category_id`),
    CONSTRAINT `article_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4; 
  /* 创建触发器1 */
  > create trigger tr_article_after_insert after insert
  on article for each row
  update category set article_count=article_count+1 where id = new.category_id;
  /* 创建触发器2 */
  > create trigger tr_article_after_delete after delete
  on article for each row
  update category set article_count=article_count-1 where id = old.category_id;
  /* 创建触发器3 */
  > DELIMITER ||
  
  > create trigger tr_article_after_update after update on article
  for each row
  BEGIN
  	update category set article_count=article_count-1 where id = old.category_id;
  	update category set article_count=article_count+1 where id = new.category_id;
  END||
  
  > DELIMITER ;
  ~~~

#### 2.2 获取

`go get "https://github.com/mittacy/blog-gin.git"`

#### 2.3 设置

1. 打开`config/config.go`

~~~go
const PageCategoryNums = 10	// 每页展示的分类数量
const PageArticleNums = 10 // 每页展示的文章数量
const ProhibitIPTime = 3600 // ip错误次数过多禁止访问时长, 单位为 秒
const IPPostTimes = 5	// ip ProhibitIPTime时间内连续请求次数

var InitAdmin = model.Admin {
	Name:      "admin",
	Password:  utiles.Encryption("admin"),
	Views:     7156,
	Cname:     "陈铭涛",
	Introduce: "就读佛山大学 - 大三 - 计算机系",
	Github:    "https://github.com/mittacy",
	Mail:      "mail@mittacy.com",
	Bilibili:  "https://space.bilibili.com/384942135",
}
~~~

2. 打开 `database/mysql.go` 设置mysql连接

   ~~~go
   const (
   	localhost			string = "127.0.0.1"
   	mysqlUser     string = "root"
   	mysqlPassword string = "**********"
   	mysqlHost     string = localhost + ":3306"
   	mysqlDatabase string = "blog"
   )
   ~~~

#### 2.4 运行

~~~shell
$ cd .../项目根目录
# 下载go依赖
$ go mod download
# 运行
$ go run main.go
~~~

如果监听端口为`3824`，浏览器打开http://localhost:3824

