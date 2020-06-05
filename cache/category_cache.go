package cache

import (
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/model"
)

var categoryCache []*model.Category

// UpdateCategoryCache 设置所有分类缓存
func UpdateCategoryCache(categories []*model.Category) {
	categoryCache = categories
}
// GetCategoryByPageCache 通过Page获取分类
func GetCategoryByPageCache(page int) []*model.Category {
	start := page * common.PageCategoryNums
	end := page * common.PageCategoryNums + common.PageCategoryNums
	return categoryCache[start:end]
}