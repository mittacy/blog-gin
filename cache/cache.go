package cache

import (
	"github.com/crazychat/blog-gin/repository"
)

// InitCache 启动服务时的缓存操作
func InitCache() error {
	// 1. admin 缓存
	admin, err := repository.NewAdminRepository("admin").Select()
	if err != nil {
		return err
	}
	UpdateAdminCache(admin)
	// 2. category 缓存
	// 3. article 缓存
	return nil
}


