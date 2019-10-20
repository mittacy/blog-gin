**使用golang和gin框架开发个人博客后台，同时使用gorm关系数据库**

**API接口文档**

~~~go
// 获取管理员信息
GET			/api/admin
//登录，参数：admin, password
POST 		/api/admin
//修改管理员信息，参数：(admin/password/cname/introduce/github/mail)
PUT			/api/admin
// 博客浏览量加1
GET			/api/addviews


// 获取全部分类
GET  		/api/categories
// 显示某个分类所有的文章
GET			/api/category/:id
// 添加分类，参数：title
POST 		/api/category
// 修改分类名，参数：id, title
PUT			/api/category
// 删除分类，参数：id
Delete 	/api/category


// 获取全部文章
GET 		/api/articles
// 显示某篇文章
GET     /api/article/:id
// 添加文章，参数：title, category_id, content, attachment
POST		/api/article
// 修改文章，参数：id, title, category_id, content, attachment
PUT			/api/article
// 删除文章，参数：id
Delete	/api/article
~~~

