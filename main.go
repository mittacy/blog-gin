package main

import (
	"fmt"
	"github.com/mittacy/blog-gin/cache"
	"github.com/mittacy/blog-gin/common"
	"github.com/mittacy/blog-gin/database"
	"github.com/mittacy/blog-gin/log"
	"github.com/mittacy/blog-gin/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	// 1. 创建 Gin 框架 todo 上线改成gin.New()
	r := gin.Default()
	// 2. 设置日志
	if err := log.InitLog(); err != nil {
		fmt.Println(err)
	}
	defer log.CloseLogFile()
	// 3. 数据库连接
	if err := database.ConnectRedis(); err != nil {
		log.ErrLogger.Fatalln(err)
	}
	defer database.CloseRedis()
	if err := database.ConnectMysql(); err != nil {
		log.ErrLogger.Fatalln(err)
	}
	defer database.CloseMysql()
	// 4. 加载静态文件
	r.Static("/css", "./css")
	r.Static("/js", "./js")
	r.Static("/index.html", "./index.html")
	r.LoadHTMLFiles("index.html")
	// 5. 中间件
	r.Use(gin.Recovery())
	r.Use(router.StaticMiddleware())
	//r.Use(router.CorsMiddleware()) // todo 上线前关闭跨域允许
	// 6. 设置路由
	router.Router(r)
	// 7. 设置深夜定时更新缓存到数据库
	go common.StartTimer()
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

