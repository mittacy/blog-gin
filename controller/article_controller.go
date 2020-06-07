package controller

import (
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/config"
	"github.com/crazychat/blog-gin/log"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/repository"
	"github.com/crazychat/blog-gin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IArticleController interface {
	Post(*gin.Context)
	Delete(*gin.Context)
	Put(*gin.Context)
	GetByID(*gin.Context)
	GetByPage(*gin.Context)
	GetRecent(*gin.Context)
	GetByCategoryID(*gin.Context)
}

func NewArticleController() IArticleController {
	repo := repository.NewArticleRepository("article")
	articleService := service.NewArticleService(repo)
	return &ArticleController{articleService}
}

type ArticleController struct {
	ArticleService service.IArticleService
}

func (ac *ArticleController) Post(c *gin.Context) {
	// 1. 解析请求
	article := model.Article{}
	if err := c.ShouldBindJSON(&article); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Category{})
		return
	}
	// 2. 操作数据库
	if err := ac.ArticleService.CreateArticle(article); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, article)
}

func (ac *ArticleController) Delete(c *gin.Context) {
	// 1. 解析请求
	article := model.JsonID{}
	if err := c.ShouldBindJSON(&article); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Category{})
		return
	}
	// 2. 操作数据库
	if err := ac.ArticleService.DeleteArticle(int(article.ID)); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, nil)
}

func (ac *ArticleController) Put(c *gin.Context) {
	// 1. 解析请求
	article := model.Article{}
	if err := c.ShouldBindJSON(&article); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Category{})
		return
	}
	// 2. 操作数据库
	if err := ac.ArticleService.UpdateArticle(article); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, article)
}

func (ac *ArticleController) GetByID(c *gin.Context) {
	// 1. 解析请求
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.FAILEDERROR, &model.Category{})
		return
	}
	// 2. 操作数据库
	article, err := ac.ArticleService.GetArticleByID(id)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, article)
}

func (ac *ArticleController) GetByPage(c *gin.Context) {
	// 1. 解析请求
	page, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &[]model.Category{})
		return
	}
	// 2. 操作数据库
	articles, articleCount, err := ac.ArticleService.GetArticlesByPage(page, config.PageArticleNums)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &[]model.Category{})
		return
	}
	// 3. 返回结果
	result := make(map[string]interface{}, 0)
	result["articleCount"] = articleCount
	result["articles"] = articles
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, result)
}

func (ac *ArticleController) GetRecent(c *gin.Context) {
	// 1. 解析请求
	// 2. 操作数据库
	articles, err := ac.ArticleService.GetRecent()
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &[]model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, articles)
}

func (ac *ArticleController) GetByCategoryID(c *gin.Context) {
	// 1. 解析请求
	cateID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &[]model.Article{})
		return
	}
	page, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &[]model.Category{})
		return
	}
	// 2. 操作数据库
	articles, articleCount, err := ac.ArticleService.GetArticlesByCateID(cateID, config.PageArticleNums ,page)
	// 3. 返回结果
	result := make(map[string]interface{}, 0)
	result["articleCount"] = articleCount
	result["articles"] = articles
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, result)
}