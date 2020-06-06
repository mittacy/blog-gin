package main

import (
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/controllers"
	"github.com/crazychat/blog-gin/database"
	"github.com/crazychat/blog-gin/log"
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
	log.InitLog()
	defer log.CloseLogFile()
	// 3. 中间件
	r.Use(router.StaticMiddleware())
	r.Use(router.CorsMiddleware())
	//r.Use(router.CorsMiddleware())	// todo 上线前关闭跨域允许
	// 4. 数据库连接
	if err := database.ConnectRedis(); err != nil {
		log.ErrLogger.Fatalln(err)
	}
	defer database.CloseRedis()
	if err := database.ConnectMysql(); err != nil {
		log.ErrLogger.Fatalln(err)
	}
	defer database.CloseMysql()
	// **********
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
	// 8. 设置缓存
	if err := cache.InitCache(); err != nil {
		log.ErrLogger.Fatalln(err)
	}
	// 9. 启动服务
	s := &http.Server{
		Addr:           ":3824",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

