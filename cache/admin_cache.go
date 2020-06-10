package cache

import (
	"github.com/mittacy/blog-gin/model"
	"github.com/mittacy/blog-gin/repository"
)

var adminCache *model.Admin	// 保存admin的信息
var adminViewCache int	// admin新增访问量

func InitAdminCache() error {
	admin, err := repository.NewAdminRepository("admin").Select()
	if err != nil {
		return err
	}
	UpdateAdminCache(admin)
	adminViewCache = 0
	return nil
}

// UpdateAdminCache 设置admin缓存
func UpdateAdminCache(admin *model.Admin) {
	adminCache = admin
}
// GetAdminCache 获取admin缓存
func GetAdminCache() (model.Admin, bool) {
	if adminCache == nil {
		return *adminCache, false
	}
	return *adminCache, true
}
// AdminViewCacheIncr 新增博客访问量
func AdminViewCacheIncr() {
	adminViewCache++
}
// UpdateViewsToDatabase 更新博客新增访问量到数据库
func UpdateViewsToDatabase() error {
	repo := repository.NewAdminRepository("admin")
	return repo.UpdateViews(adminViewCache)
}

