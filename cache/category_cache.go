package cache

import (
	"github.com/crazychat/blog-gin/config"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/repository"
)

var (
	categoryCache []model.Category
	categoryCacheIndex = make(map[uint32]int)	// 记录category在categoryCache中的位置
)

func InitCategoryCache() error {
	categories, err := repository.NewCategoryRepository("category").Select()
	if err != nil {
		return err
	}
	SetCategoryCache(categories)
	return nil
}

// SetCategoryCache 初始化所有分类缓存
func SetCategoryCache(categories []model.Category) {
	categoryCache = categories
	updateCategoryCacheIndex()
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
	updateCategoryCacheIndex()
}
// UpdateCategoryCache 更新分类的title
func UpdateCategoryCache(category model.Category) {
	index := categoryCacheIndex[category.ID]
	categoryCache[index].Title = category.Title
}
// UpdateCategoryCacheIncr 分类文章数+1
func UpdateCategoryCacheIncr(id uint32) bool {
	if index, isExist := categoryCacheIndex[id]; isExist {
		categoryCache[index].ArticleCount++
		return true
	}
	return false
}
// UpdateCategoryCacheDecr 分类文章数-1
func UpdateCategoryCacheDecr(id uint32) bool {
	if index, isExist := categoryCacheIndex[id]; isExist {
		categoryCache[index].ArticleCount--
		return true
	}
	return false
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
// updateCategoryCacheIndex 更新index
func updateCategoryCacheIndex() {
	categoryCacheIndex = make(map[uint32]int, len(categoryCache))
	for i, v := range categoryCache {
		categoryCacheIndex[v.ID] = i
	}
}