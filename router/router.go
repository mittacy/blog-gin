package router

import (
	"github.com/crazychat/blog-gin/controller"
	"github.com/crazychat/blog-gin/log"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	adminController := controller.NewAdminController()
	if err := adminController.InitAdmin(); err != nil {
		log.ErrLogger.Fatalln(err)
	}
	categoryController := controller.NewCategoryController()
	articleController := controller.NewArticleController()
	// 不需要登录验证的api
	api := router.Group("/api")
	{
		api.POST("/admin", adminController.Post)
		api.GET("/verify", adminController.Verify)
		api.GET("/admin", adminController.Get)
		// 分类
		api.GET("/category_name/:id", categoryController.GetByID)
		api.GET("/categories", categoryController.GetAll)
		api.GET("/category_page/:num", categoryController.GetByPage)
		//api.GET("/category/:id/:num", controllers.GetCategoy)
		// 文章
		api.GET("/articles_recent", articleController.GetRecent)
		api.GET("/article/:id", articleController.GetByID)
		api.GET("/article_page/:num", articleController.GetByPage)
	}
	// 需要登录验证的api
	apiVerify := router.Group("/api")

	apiVerify.Use(VerifyMiddleware())
	{
		// 日志文件
		apiVerify.GET("/errlog", controller.GetErrorLog)
		// 管理员
		apiVerify.PUT("/admin", adminController.Put)
		apiVerify.PUT("/admin/setpwd", adminController.PutPassword)
		// 分类
		apiVerify.POST("/category", categoryController.Post)
		apiVerify.PUT("/category", categoryController.Put)
		apiVerify.DELETE("/category", categoryController.Delete)
		// 文章
		apiVerify.POST("/article", articleController.Post)
		apiVerify.PUT("/article", articleController.Put)
		apiVerify.DELETE("/article", articleController.Delete)
	}
}