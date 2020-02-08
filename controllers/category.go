package controllers

import (
	"blog-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	onePageCategoryNum = 10
	categoryCount = 0
)

// GetCategories 获取所有分类id和title
func GetCategories(c *gin.Context) {
	// 查询分类及总数
	categories, msg, err := models.GetCategories()
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, categories)
}

// GetPageCategory 分页获取分类
func GetPageCategory(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Param("num"))
	if !CheckErr(err) {
		RejectResult(c, models.FAILEDERROR)
		return
	}
	categories, msg, err := models.GetPageCategories(pageNum, onePageCategoryNum)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	// 查询第0页时更新分类总数
	if pageNum == 0 {
		count, msg, err := models.GetCategoriesCount()
		if !CheckErr(err) {
			RejectResult(c, msg)
			return
		}
		categoryCount = count
	}
	result := make(map[string]interface{}, 0)
	result["categoryCount"] = categoryCount
	result["categories"] = categories
	ResolveResult(c, models.CONTROLLER_SUCCESS, result)
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
	ResolveResult(c, msg, msg)
}

// GetCategoy 获取某个分类及其所有文章
func GetCategoy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if !CheckErr(err) {
		RejectResult(c, models.FAILEDERROR)
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
	ResolveResult(c, msg, msg)
}
