package cache

import (
	"github.com/crazychat/blog-gin/model"
)

var articles []*model.Article
// CacheArticles 缓存全部文章（除了内容）
func CacheArticles(as []*model.Article) {
	articles = as
}

func GetCacheArticlesByPage(page int) ([]*model.Article, bool) {
	if articles == nil {
		return nil, false
	}
	// todo 修改页面
	return articles[:], true
}
