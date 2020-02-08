**使用golang和gin框架开发个人博客后台**

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


// 分页获取文章, 参数: pageIndex
GET			/api/article_page/:num
// 显示某篇文章
GET     /api/article/:id
// 添加文章，参数: title, category_id, content
POST		/api/article
// 修改文章，参数: id, title, category_id, content, attachment
PUT			/api/article
// 删除文章，参数: id
Delete	/api/article
// 文章添加访问量, 参数: id
POST		/api/article/addViews
// 文章添加点赞量, 参数: id
PSOT		/api/article/addAssists
~~~

