package controllers

import (
	"blog-gin/models"
	"strconv"

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

// GetCategories 获取所有分类title
func GetCategories(c *gin.Context) {
	data := make([]models.Category, 0)
	data, msg, err := models.GetCategories(data)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
	}
	ResolveResult(c, 200, data)
}

// GetCategoy 获取某个分类的所有文章
func GetCategoy(c *gin.Context) {
	cateID, err := strconv.Atoi(c.Param("id"))
	if !CheckErr(err) {
		RejectResult(c, 400, NOKNOW_ERROR)
		return
	}
	articles, msg, err := models.GetCategory(cateID)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, articles)
}

// DeleteCategory 删除分类同时删除分类里的所有文章
func DeleteCategory(c *gin.Context) {
	cateID, err := strconv.Atoi(c.Param("id"))
	msg, err := models.DeleteCategory(cateID)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, "")
}
