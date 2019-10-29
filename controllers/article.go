package controllers

import (
	"blog-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	onePageArticlesNum = 10
)

// CreateArticle 创建文章controller
func CreateArticle(c *gin.Context) {
	article := models.Article{}
	if !AnalysisJSON(c, &article) {
		return
	}
	msg, err := models.CreateArticle(&article)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, article)
}

// UpdateArticle 修改文章controller
func UpdateArticle(c *gin.Context) {
	article := models.Article{}
	if !AnalysisJSON(c, &article) {
		return
	}
	msg, err := models.UpdateArticle(&article)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, msg)
}

// GetArticle 根据id获取文章
func GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RejectResult(c, 400, NOKNOW_ERROR)
	}
	article, msg, err := models.GetArticle(articleID)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, article)
}

// DeleteArticle 根据id删除文章
func DeleteArticle(c *gin.Context) {
	getID := GetID{}
	if !AnalysisJSON(c, &getID) {
		return
	}
	msg, err := models.DeleteArticle(getID.ID)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, msg)
}

// GetPageArticle 分页获取文章
func GetPageArticle(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		RejectResult(c, 400, NOKNOW_ERROR)
	}
	articls, msg, err := models.GetPageArticles(pageNum, onePageArticlesNum)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, articls)
}
