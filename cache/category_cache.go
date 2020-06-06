package cache

import (
	"github.com/crazychat/blog-gin/config"
	"github.com/crazychat/blog-gin/model"
)

var categoryCache []model.Category
var categoryCacheIndex = make(map[uint32]int)	// 记录category在categoryCache中的位置

// InitCategoryCache 初始化所有分类缓存
func InitCategoryCache(categories []model.Category) {
	categoryCache = categories
	categoryCacheIndex = make(map[uint32]int, len(categoryCache))
	for i, v := range categoryCache {
		categoryCacheIndex[v.ID] = i
	}
}
// AddCategoryCache 增加分类
func AddCategoryCache(category model.Category) {
	categoryCacheIndex[category.ID] = len(categoryCache)
	categoryCache = append(categoryCache, category)
}
// DeleteCategoryCache 删除分类
func DeleteCategoryCache(id uint32) {
	index := categoryCacheIndex[id]
	categoryCache = append(categoryCache[:index], categoryCache[index+1:]...)
	delete(categoryCacheIndex, id)
}
// UpdateCategoryCache 更新分类的title
func UpdateCategoryCache(category model.Category) {
	index := categoryCacheIndex[category.ID]
	categoryCache[index].Title = category.Title
}
// GetCategoryCacheByID 根据 id 获取分类信息
func GetCategoryCacheByID(id uint32) (*model.Category, bool) {
	index, exist := categoryCacheIndex[id]
	if !exist {
		return nil, false
	}
	return &categoryCache[index], true
}
// GetCategoryCacheByPage 通过Page获取分类
func GetCategoriesCacheByPage(page int) ([]model.Category, int, bool) {
	if categoryCache == nil {
		return nil, 0, false
	}
	start := page * config.PageCategoryNums
	end := page *config.PageCategoryNums + config.PageCategoryNums
	length := len(categoryCache)
	if start >  len(categoryCache) {
		return make([]model.Category, 0), length, true
	}
	if end > len(categoryCache) {
		end = len(categoryCache)
	}
	return categoryCache[start:end], length, true
}
// GetCategoriesCache 获取全部分类信息
func GetCategoriesCache() ([]model.Category, bool) {
	if categoryCache == nil {
		return nil, false
	}
	return categoryCache, true
}