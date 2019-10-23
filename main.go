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
	sqlDb := models.GetSQLDB()
	defer sqlDb.Close()

	router := gin.Default()
	// 不需要登录验证的api
	api := router.Group("/api")
	{
		api.POST("/admin", controllers.PostAdmin)
		api.GET("/admin", controllers.GetAdmin)
	}
	// 需要登录验证的api
	apiAdmin := router.Group("/api")
	apiAdmin.Use(controllers.CheckAdmin())
	{
	}

	router.Run(":5201")
}
