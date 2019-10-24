package controllers

import (
	"blog-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
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
