package cache

import (
	"github.com/crazychat/blog-gin/model"
)

var adminInfo *model.Admin	// 保存admin的信息(包括加密的密码)
// CacheAdminInfo 设置admin缓存
func CacheAdminInfo(admin *model.Admin) {
	adminInfo = admin
}
// GetCacheAdminInfo 获取admin缓存
func GetCacheAdminInfo() (*model.Admin, bool) {
	if adminInfo == nil {
		return adminInfo, false
	}
	return adminInfo, true
}

