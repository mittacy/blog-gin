package cache

import "github.com/mittacy/blog-gin/log"

// InitCache 启动服务时的缓存操作
func InitCache() (err error) {
	// 1. admin 缓存
	if err = InitAdminCache(); err != nil {
		log.RecordErr(err)
		return
	}
	// 2. category 缓存
	if err = InitCategoryCache(); err != nil {
		log.RecordErr(err)
		return
	}
	// 3. article 缓存
	if err = InitArticleCache(); err != nil {
		log.RecordErr(err)
	}
	return
}


