package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/blog-gin/common"
	"github.com/mittacy/blog-gin/config"
	"github.com/mittacy/blog-gin/log"
	"github.com/mittacy/blog-gin/model"
	"github.com/mittacy/blog-gin/repository"
	"github.com/mittacy/blog-gin/service"
	"strconv"
)

type ICategoryController interface {
 Post(c *gin.Context)
 Delete(c *gin.Context)
 Put(c *gin.Context)
 GetAll(c *gin.Context)
 GetByID(c *gin.Context)
 GetByPage(c *gin.Context)
}

type CategoryController struct {
	CategoryService service.ICategoryService
}

func NewCategoryController() ICategoryController {
	repo := repository.NewCategoryRepository("category")
	categoryService := service.NewCategoryService(repo)
	return &CategoryController{categoryService}
}
// Post 添加分类
func (cc *CategoryController) Post(c *gin.Context) {
	// 1. 解析请求
	cate := model.Category{}
	if err := c.ShouldBindJSON(&cate); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Category{})
		return
	}
	// 2. 操作数据库
	if err := cc.CategoryService.CreateCategory(cate); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, cate)
}
// Delete 删除分类
func (cc *CategoryController) Delete(c *gin.Context) {
	// 1. 解析json数据到结构体
	cate := &model.JsonID{}
	if err := c.ShouldBindJSON(cate); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Category{})
		return
	}
	// 2. 操作数据库
	if err := cc.CategoryService.DeleteCategory(int(cate.ID)); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, nil)
}
// Put 修改文章
func (cc *CategoryController) Put(c *gin.Context) {
	// 1. 解析json数据到结构体
	cate := model.Category{}
	if err := c.ShouldBindJSON(&cate); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Category{})
		return
	}
	// 2. 操作数据库
	if err := cc.CategoryService.UpdateCategory(cate); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, cate)
}
// GetAll 获取全部分类信息
func (cc *CategoryController) GetAll(c *gin.Context) {
	// 1. 操作数据库
	categories, err := cc.CategoryService.GetCategories()
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, categories)
}
// GetByID 根据id获取分类
func (cc *CategoryController) GetByID(c *gin.Context) {
	// 1. 解析参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.FAILEDERROR, &model.Category{})
		return
	}
	// 2. 操作数据库
	cate, err := cc.CategoryService.GetCategoryByID(id)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Category{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, cate)
}
// GetByPage 根据page获取分类
func (cc *CategoryController) GetByPage(c *gin.Context) {
	// 1. 解析参数
	page, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &[]model.Category{})
		return
	}
	// 2. 操作数据库
	categories, categoryCount, err := cc.CategoryService.GetCategoriesByPage(page, config.PageCategoryNums)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &[]model.Category{})
		return
	}
	// 3. 返回结果
	result := make(map[string]interface{}, 0)
	result["categoryCount"] = categoryCount
	result["categories"] = categories
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, result)
}