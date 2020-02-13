package main

import (
	"blog-gin/asset"
	"blog-gin/controllers"
	"blog-gin/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库连接
	db := models.GetDB()
	defer db.Close()
	// 创建日志文件
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
	// 释放静态文件
	isSuccess := true
	dirs := []string{"css", "js", "img", "fonts", "index.html"}
	for _, dir := range dirs {
		if err := asset.RestoreAssets("./", dir); err != nil {
			isSuccess = false
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("./", dir))
		}
	}
	router := gin.New()
	// 静态文件
	router.Static("/css", "./css")
	router.Static("/fonts", "./fonts")
	router.Static("/img", "./img")
	router.Static("/js", "./js")
	router.Static("/index.html", "./index.html")
	router.LoadHTMLFiles("index.html")
	// 过滤api请求
	router.Use(TransparentStatic())
	// 后端路由
	router.Use(CorsMiddleware())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \" - err: \"%s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
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
		api.GET("/article/:id", controllers.GetArticle)
		api.GET("/article_page/:num", controllers.GetPageArticle)
		api.POST("/article/addViews", controllers.AddArticleViews)
		api.GET("/admin/article_id", controllers.GetArticleID)
	}
	// 需要登录验证的api
	apiAdmin := router.Group("/api")
	apiAdmin.Use(controllers.CheckAdmin())
	{
		// 管理员
		apiAdmin.PUT("/admin", controllers.PutAdmin)
		apiAdmin.PUT("/admin/setpwd", controllers.PutAdminPwd)
		// 分类
		apiAdmin.POST("/category", controllers.CreateCategory)
		apiAdmin.PUT("/category", controllers.UpdataCategory)
		apiAdmin.DELETE("/category", controllers.DeleteCategory)
		// 文章
		apiAdmin.POST("/article", controllers.CreateArticle)
		apiAdmin.PUT("/article", controllers.UpdateArticle)
		apiAdmin.DELETE("/article", controllers.DeleteArticle)
		api.PUT("/admin/article_id", controllers.PutArticleID)
	}
	s := &http.Server{
		Addr:           ":5201",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func TransparentStatic() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		if strings.Index(url, "api") >= 0 {
			c.Next()
			return
		}
		// 增加博客访问量
		if !models.AddViews() {
			fmt.Println("增加访问量失败")
		}
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
