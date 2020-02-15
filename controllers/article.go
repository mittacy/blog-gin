package controllers

import (
	"blog-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	onePageArticlesNum = 10
	ArticleCount       = 0
)

// CreateArticle 创建文章controller
func CreateArticle(c *gin.Context) {
	article := models.Article{}
	if !AnalysisJSON(c, &article) { return }
	msg, err := models.CreateArticle(&article)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, article)
}

// UpdateArticle 修改文章controller
func UpdateArticle(c *gin.Context) {
	article := models.Article{}
	if !AnalysisJSON(c, &article) {
		return
	}
	msg, err := models.UpdateArticle(&article)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, msg)
}

// GetArticle 根据id获取文章
func GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if !CheckErr(err) {
		RejectResult(c, models.BACKERROR)
	}
	article, msg, err := models.GetArticle(articleID)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, article)
}

// DeleteArticle 根据id删除文章
func DeleteArticle(c *gin.Context) {
	getID := GetID{}
	if !AnalysisJSON(c, &getID) {
		return
	}
	msg, err := models.DeleteArticle(getID.ID)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, msg)
}

// GetPageArticle 分页获取文章
func GetPageArticle(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Param("num"))
	if !CheckErr(err) {
		RejectResult(c, models.BACKERROR)
		return
	}
	articls, msg, err := models.GetPageArticles(pageNum, onePageArticlesNum)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	// 查询第0页时更新文章总数
	if pageNum == 0 {
		count, msg, err := models.GetArticlesCount()
		if !CheckErr(err) {
			RejectResult(c, msg)
			return
		}
		ArticleCount = count
	}
	result := make(map[string]interface{}, 0)
	result["articleCount"] = ArticleCount
	result["articles"] = articls
	ResolveResult(c, models.CONTROLLER_SUCCESS, result)
}

// AddArticleViews 文章添加浏览量
func AddArticleViews(c *gin.Context) {
	getID := GetID{}
	if !AnalysisJSON(c, &getID) {
		return
	}
	msg, err := models.AddArticleViews(getID.ID)
	if CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, msg)
}

// RecentArticles 最近更新的五篇文章
func RecentArticles(c *gin.Context) {
	articls, msg, err := models.GetRecentArticles()
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, models.CONTROLLER_SUCCESS, articls)
}
