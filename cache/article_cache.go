package cache

import (
	"errors"
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/model"
)

var articles []*model.Article
// CacheArticles 缓存全部文章（除了内容）
func CacheArticles(as []*model.Article) {
	articles = as
}

func GetCacheArticlesByPage(page int) ([]*model.Article, error) {
	if articles == nil {
		return nil, errors.New(common.NO_CACHE)
	}
	// todo 修改页面
	return articles[:], nil
}
