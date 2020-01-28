package controllers

import (
	"blog-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	categoryCount = 0
)

// GetCategories 获取所有分类title
func GetCategories(c *gin.Context) {
	// 查询分类及总数
	categories, msg, err := models.GetCategories()
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, categories)
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	cate := models.Category{}
	if !AnalysisJSON(c, &cate) {
		return
	}
	msg, err := models.CreateCate(&cate)
	if err != nil {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, cate)
}

// UpdataCategory 更新分类
func UpdataCategory(c *gin.Context) {
	cate := models.Category{}
	if !AnalysisJSON(c, &cate) {
		return
	}
	msg, err := models.UpdateCate(&cate)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, nil)
}

// GetCategoy 获取某个分类及其所有文章
func GetCategoy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if !CheckErr(err) {
		RejectResult(c, models.UNKONWNERROR)
		return
	}
	cate := models.Category{ID: uint32(id)}
	result, msg, err := models.GetCategory(&cate)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, result)
}

// DeleteCategory 删除分类同时删除分类里的所有文章
func DeleteCategory(c *gin.Context) {
	getID := GetID{}
	if !AnalysisJSON(c, &getID) {
		return
	}
	msg, err := models.DeleteCategory(getID.ID)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, nil)
}
