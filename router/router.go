package router

import (
	"github.com/crazychat/blog-gin/controller"
	"github.com/crazychat/blog-gin/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	adminController := controller.GetAdminController()
	// 不需要登录验证的api
	api := router.Group("/api")
	{
		api.POST("/admin", controllers.PostAdmin)
		api.GET("/verify", controllers.Verify)
		api.GET("/admin", adminController.GetAdmin)
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
	//apiVerfiry := router.Group("/api")
	//apiVerfiry.Use(controllers.CheckAdmin())
	//{
	//	// 日志文件
	//	apiVerfiry.GET("/errlog", controllers.GetErrorLog)
	//	// 管理员
	//	apiVerfiry.PUT("/admin", controllers.PutAdmin)
	//	apiVerfiry.PUT("/admin/setpwd", controllers.PutAdminPwd)
	//	// 分类
	//	apiVerfiry.POST("/category", controllers.CreateCategory)
	//	apiVerfiry.PUT("/category", controllers.UpdataCategory)
	//	apiVerfiry.DELETE("/category", controllers.DeleteCategory)
	//	// 文章
	//	apiVerfiry.POST("/article", controllers.CreateArticle)
	//	apiVerfiry.PUT("/article", controllers.UpdateArticle)
	//	apiVerfiry.DELETE("/article", controllers.DeleteArticle)
	//}
}