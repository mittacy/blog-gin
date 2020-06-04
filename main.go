package main

import (
	"github.com/crazychat/blog-gin/controllers"
	"github.com/crazychat/blog-gin/database"
	"github.com/crazychat/blog-gin/models"
	"github.com/crazychat/blog-gin/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	// 1. 创建 Gin 框架
	r := gin.Default()
	// 2. 设置日志
	// 日志处理	controllers/log.go init()已经打开文件, 此处关闭
	defer controllers.CloseLogFile()
	// 3. 中间件
	r.Use(router.StaticMiddleware())
	//r.Use(router.CorsMiddleware())	// todo 上线前关闭跨域允许
	// 4. 数据库连接
	if err := database.ConnectRedis(); err != nil {
		controllers.ErrLogger.Fatal(err)
	}
	defer database.CloseRedis()
	if err := database.ConnectMysql(); err != nil {
		controllers.ErrLogger.Fatal(err)
	}
	defer database.CloseMysql()

	if err := models.ConnectRedis(); err != nil {
		controllers.ErrLogger.Fatal(err)
	}
	defer models.CloseRedis()
	if err := models.ConnectMysql(); err != nil {
		controllers.ErrLogger.Fatal(err)
	}
	defer models.CloseMysql()
	// 5. 加载静态文件
	r.Static("/css", "./css")
	r.Static("/js", "./js")
	r.Static("/index.html", "./index.html")
	r.LoadHTMLFiles("index.html")
	// 6. 设置路由
	router.Router(r)
	// 7. todo 设置深夜定时更新缓存到数据库
	//go models.StartTimer()
	// 8. 启动服务
	s := &http.Server{
		Addr:           ":3824",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

//func Router(router *gin.Engine) {
//	// 不需要登录验证的api
//	api := router.Group("/api")
//	{
//		api.POST("/admin", controllers.PostAdmin)
//		api.GET("/verify", controllers.Verify)
//		api.GET("/admin", controllers.GetAdmin)
//		// 分类
//		api.GET("/category_name/:id", controllers.GetCategoryName)
//		api.GET("/categories", controllers.GetCategories)
//		api.GET("/category_page/:num", controllers.GetPageCategory)
//		api.GET("/category/:id/:num", controllers.GetCategoy)
//		// 文章
//		api.GET("/articles_recent", controllers.RecentArticles)
//		api.GET("/article/:id", controllers.GetArticle)
//		api.GET("/article_page/:num", controllers.GetPageArticle)
//	}
//	// 需要登录验证的api
//	apiVerfiry := router.Group("/api")
//	apiVerfiry.Use(controllers.CheckAdmin())
//	{
//		// 日志文件
//		apiVerfiry.GET("/errlog", controllers.GetErrorLog)
//		// 管理员
//		apiVerfiry.PUT("/admin", controllers.PutAdmin)
//		apiVerfiry.PUT("/admin/setpwd", controllers.PutAdminPwd)
//		// 分类
//		apiVerfiry.POST("/category", controllers.CreateCategory)
//		apiVerfiry.PUT("/category", controllers.UpdataCategory)
//		apiVerfiry.DELETE("/category", controllers.DeleteCategory)
//		// 文章
//		apiVerfiry.POST("/article", controllers.CreateArticle)
//		apiVerfiry.PUT("/article", controllers.UpdateArticle)
//		apiVerfiry.DELETE("/article", controllers.DeleteArticle)
//	}
//}

