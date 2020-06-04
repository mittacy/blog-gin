package cache

import (
	"github.com/crazychat/blog-gin/model"
)

var adminInfo *model.Admin	// 保存admin的信息
var adminPwd string	// admin加密的密码

// UpdateAdminInfo 设置admin缓存
func UpdateAdminInfo(admin *model.Admin) {
	admin.Password = "**********"
	adminInfo = admin
}
// GetAdminInfo 获取admin缓存
func GetAdminInfo() (*model.Admin, bool) {
	if adminInfo == nil {
		return adminInfo, false
	}
	return adminInfo, true
}
// UpdateAdminPwd 缓存admin密码
func UpdateAdminPwd(str string) {
	adminPwd = str
}
// GetAdminPwd 获取admin密码
func GetAdminPwd() (string, bool) {
	if adminPwd == "" {
		return adminPwd, false
	}
	return adminPwd, true
}

