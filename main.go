package main

import (
	"blog-gin/controllers"
	"blog-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库连接
	db := models.GetDB()
	defer db.Close()

	router := gin.Default()
	// 不需要登录验证的api
	api := router.Group("/api")
	{
		api.POST("/admin", controllers.PostAdmin)
		api.GET("/admin", controllers.GetAdmin)
		api.PUT("/admin", controllers.PutAdmin)
		api.PUT("/admin/setpwd", controllers.PutAdminPwd)
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

	}
	// 需要登录验证的api
	apiAdmin := router.Group("/api")
	apiAdmin.Use(controllers.CheckAdmin())
	{

	}

	router.Run(":5201")
}
