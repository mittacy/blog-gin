package main

import (
	"blog-gin/controllers"
	"blog-gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库连接
	db := models.GetDB()
	defer db.Close()

	router := gin.Default()
	router.Static("/statics", "./statics")
	// 不需要登录验证的api
	api := router.Group("/api")
	{
		api.POST("/admin", controllers.PostAdmin)
		api.GET("/admin", controllers.GetAdmin)
		api.PUT("/admin", controllers.PutAdmin)
		api.PUT("/admin/setpwd", controllers.PutAdminPwd)
		api.GET("/admin/addviews", controllers.AddAdminView)
		// 分类
		api.GET("/categories", controllers.GetCategories)
		api.POST("/category", controllers.CreateCategory)
		api.PUT("/category", controllers.UpdataCategory)
		api.DELETE("/category", controllers.DeleteCategory)
		api.GET("/category/:id", controllers.GetCategoy)
		// 文章
		api.POST("/article", controllers.CreateArticle)
		api.GET("/article/:id", controllers.GetArticle)
		api.PUT("/article", controllers.UpdateArticle)
		api.DELETE("/article", controllers.DeleteArticle)
		api.GET("/article_page/:num", controllers.GetPageArticle)
		api.POST("/article/addViews", controllers.AddArticleViews)
		api.POST("/article/addAssists", controllers.AddArticleAssists)

	}
	// 需要登录验证的api
	apiAdmin := router.Group("/api")
	apiAdmin.Use(controllers.CheckAdmin())
	{

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
