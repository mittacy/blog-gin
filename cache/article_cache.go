package cache

import (
	"fmt"
	"github.com/crazychat/blog-gin/log"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/repository"
)

var (
	articleCache []model.Article // 全部文章简介缓存
	articleCacheIndex = make(map[uint32]int)  // // 记录article在articleCache中的位置
	recentArticleCache []model.Article // 最新五篇文章缓存
)

func InitArticleCache() error {
	control := repository.NewArticleRepository("article")
	articles, err := control.Select()
	if err != nil {
		log.RecordErr(err)
		return err
	}
	SetArticleCache(articles)
	InitRecentArticleCache()
	return nil
}
// InitArticleCache 初始化全部文章缓存
func SetArticleCache(articles []model.Article) {
	articleCache = articles
	updateArticleCacheIndex()
}
// InitRecentArticleCache 初始化最新五篇文章缓存
func InitRecentArticleCache() {
	articles, err := repository.NewArticleRepository("article").SelectRecent()
	if err != nil {
		log.RecordErr(err)
		return
	}
	// 更新最近五篇文章
	recentArticleCache = articles
}
// AddArticleCache 添加文章
func AddArticleCache(article model.Article) {
	article.Content = ""
	articleSlice :=  []model.Article{article}
	articleCache = append(articleSlice, articleCache...)
	recentArticleCache = recentArticleCache[:4]
	recentArticleCache = append(articleSlice, recentArticleCache...)
	updateArticleCacheIndex()
}
// DeleteArticleCache 删除文章
func DeleteArticleCache(id uint32) {
	if index, isExist := articleCacheIndex[id]; isExist {
		fmt.Println("删除文章, id = ", id, ", index: ", index)
		articleCache = append(articleCache[:index], articleCache[index+1:]...)
		// 更新最近五篇文章
		InitRecentArticleCache()
		updateArticleCacheIndex()
	}

}
// UpdateArticleCache 更新文章内容
func UpdateArticleCache(article model.Article) {
	if index, isExist := articleCacheIndex[article.ID]; isExist {
		article.Content = ""
		articleCache[index] = article
		// 更新最近五篇文章
		InitRecentArticleCache()
	}
}
// GetArticleCacheByID 获取文章简介
func GetArticleCacheByID(id uint32) (model.Article, bool) {
	if index, isExist := articleCacheIndex[id]; isExist {
		return articleCache[index], true
	}
	fmt.Println("文章不存在")
	return model.Article{}, false
}
// GetArticleCacheByPage 通过Page获取文章
func GetArticleCacheByPage(page, pageArticleNums int) ([]model.Article, int, bool) {
	if articleCache == nil {
		return nil, 0, false
	}
	start := page * pageArticleNums
	end := page * pageArticleNums + pageArticleNums
	length := len(articleCache)
	if start > len(articleCache) {
		return make([]model.Article, 0), length, true
	}
	if end > len(articleCache) {
		end = len(articleCache)
	}
	return articleCache[start:end], length, true
}
// GetRecentArticleCache 获取最新五篇文章
func GetRecentArticleCache() ([]model.Article, bool) {
	if recentArticleCache == nil {
		return recentArticleCache, false
	}
	return recentArticleCache, true
}
// updateArticleCacheIndex 更新index
func updateArticleCacheIndex() {
	articleCacheIndex = make(map[uint32]int, len(articleCache))
	for i, v := range articleCache {
		articleCacheIndex[v.ID] = i
	}
}
