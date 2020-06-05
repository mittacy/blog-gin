package cache

import (
	"github.com/crazychat/blog-gin/model"
)

var adminCache *model.Admin	// 保存admin的信息
var adminViewCache int	// admin新增访问量

// UpdateAdminCache 设置admin缓存
func UpdateAdminCache(admin *model.Admin) {
	adminCache = admin
}
// GetAdminCache 获取admin缓存
func GetAdminCache() (*model.Admin, bool) {
	if adminCache == nil {
		return nil, false
	}
	return adminCache, true
}
// UpdateAdminViewCache 新增博客访问量
func UpdateAdminViewCache() {
	adminViewCache++
}
// GetAdminViewCache 获取博客缓存的单日新增访问量
func GetAdminViewCache() int {
	return adminViewCache
}

