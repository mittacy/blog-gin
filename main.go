package main

import (
	"fmt"
	"github.com/crazychat/blog-gin/controllers"
	"github.com/crazychat/blog-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func main() {
	// 日志处理	controllers/log.go init()已经打开文件, 此处关闭
	defer controllers.CloseLogFile()
	// 数据库连接
	if err := models.ConnectRedis(); err != nil {
		controllers.ErrLogger.Fatal(err)
	}
	defer models.CloseRedis()
	if err := models.ConnectMysql(); err != nil {
		controllers.ErrLogger.Fatal(err)
	}
	defer models.CloseMysql()
	// 加载静态文件
	router := gin.Default()
	router.Static("/css", "./css")
	router.Static("/js", "./js")
	router.Static("/index.html", "./index.html")
	router.LoadHTMLFiles("index.html")
	// 过滤前端请求
	router.Use(TransparentStatic())
	// api路由
	//router.Use(CorsMiddleware())	// todo 上线前关闭跨域允许
	router.Use(gin.Recovery())
	Router(router)
	s := &http.Server{
		Addr:           ":5201",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 深夜定点保存缓存到mysql
	models.StartTimer()
	s.ListenAndServe()
}

func Router(router *gin.Engine) {
	// 不需要登录验证的api
	api := router.Group("/api")
	{
		api.POST("/admin", controllers.PostAdmin)
		api.GET("/verify", controllers.Verify)
		api.GET("/admin", controllers.GetAdmin)
		// 分类
		api.GET("/category_name/:id", controllers.GetCategoryName)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/category_page/:num", controllers.GetPageCategory)
		api.GET("/category/:id/:num", controllers.GetCategoy)
		// 文章
		api.GET("/articles_recent", controllers.RecentArticles)
		api.GET("/article/:id", controllers.GetArticle)
		api.GET("/article_page/:num", controllers.GetPageArticle)
	}
	// 需要登录验证的api
	apiVerfiry := router.Group("/api")
	apiVerfiry.Use(controllers.CheckAdmin())
	{
		// 日志文件
		apiVerfiry.GET("/errlog", controllers.GetErrorLog)
		// 管理员
		apiVerfiry.PUT("/admin", controllers.PutAdmin)
		apiVerfiry.PUT("/admin/setpwd", controllers.PutAdminPwd)
		// 分类
		apiVerfiry.POST("/category", controllers.CreateCategory)
		apiVerfiry.PUT("/category", controllers.UpdataCategory)
		apiVerfiry.DELETE("/category", controllers.DeleteCategory)
		// 文章
		apiVerfiry.POST("/article", controllers.CreateArticle)
		apiVerfiry.PUT("/article", controllers.UpdateArticle)
		apiVerfiry.DELETE("/article", controllers.DeleteArticle)
	}
}

func TransparentStatic() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		if strings.Index(url, "api") >= 0 {
			c.Next()
			return
		}
		// 增加博客访问量
		models.IncrBlogViews()
		c.HTML(200, "index.html", gin.H{"msg": "Success"})
		return
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma, adminToken")
			// 允许跨域设置,可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
