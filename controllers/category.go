package controllers

import (
	"blog-gin/models"

	"github.com/gin-gonic/gin"
)

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	cate := models.Category{}
	if !AnalysisJSON(c, &cate) {
		return
	}
	if msg, err := models.CreateCate(&cate); err != nil {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, cate)
}

// UpdataCategory 更新分类
func UpdataCategory(c *gin.Context) {
	cate := models.Category{}
	if !AnalysisJSON(c, &cate) {
		return
	}
	// 分类是否存在
	isExist := models.IsCateExist(&cate)
	if !isExist {
		RejectResult(c, 404, NO_EXIST)
		return
	}
	// 分类存在, 修改分类title
	msg, err := models.UpdateCate(&cate)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, cate)
}
