package cache

import (
	"github.com/crazychat/blog-gin/config"
	"github.com/crazychat/blog-gin/model"
)

var articleCache []*model.Article // 全部文章简介缓存
var cateArticleCache = make(map[int][]*model.Article)  // 各个分类的文章缓存

// UpdateArticleCache 更新全部文章缓存
func UpdateArticleCache(articles []*model.Article) {
	articleCache = articles
}
// GetArticleByPageCache 通过page获取文章
func GetArticleByPageCache(page int) []*model.Article {
	start := page * config.PageArticleNums
	end := page *config.PageArticleNums + config.PageArticleNums
	return articleCache[start:end]
}

// DelCateArticleCache 清空各个分类文章缓存
func DelCateArticleCache() {
	cateArticleCache = make(map[int][]*model.Article)
}
// AddCateArticleCache 添加分类文章缓存
func AddCateArticleCache(cateId int, articles []*model.Article) {
	cateArticleCache[cateId] = articles
}
// GetCateArticleCache 获取分类所有文章
func GetCateArticleCache(cateId int) []*model.Article {
	return cateArticleCache[cateId]
}
