package cache

import (
	"fmt"
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
	fmt.Println("缓存admin成功，缓存器如下:")
	fmt.Println("adminCache", adminCache)
	fmt.Println("adminViewCache", adminViewCache)
	// 2. category 缓存
	control := repository.NewCategoryRepository("category")
	categories, err := control.Select()
	if err != nil {
		return err
	}
	InitCategoryCache(categories)
	fmt.Println("缓存categories成功，缓存器如下:")
	fmt.Println("categoryCache", categoryCache)
	fmt.Println("categoryCacheIndex", categoryCacheIndex)
	// 3. article 缓存
	return nil
}


