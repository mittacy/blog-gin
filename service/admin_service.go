package service

import (
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/repository"
)

type IAdminService interface {
	CreateAdmin(*model.Admin) error
	UpdateAdminInfo(*model.Admin) error
	UpdatePassword(*model.Admin) error
	GetAdminInfo() (*model.Admin, error)
	GetAdminPassword() (string, error)
}

func NewAdminService(repository repository.IAdminRepository) IAdminService {
	return &AdminService{repository}
}

type AdminService struct {
	AdminRepository repository.IAdminRepository
}
// CreateAdmin 启动服务，创建管理员，存在则跳过
func (as *AdminService) CreateAdmin(admin *model.Admin) error {
	if as.AdminRepository.SelectExist() { // 管理员存在则跳过
		return nil
	}
	return as.AdminRepository.Add(admin)
}
// UpdateAdminInfo 更新管理员信息
func (as *AdminService) UpdateAdminInfo(admin *model.Admin) error {
	// 1. 更新到缓冲器
	cache.UpdateAdminInfo(admin)
	// 2. 更新到数据库
	return as.AdminRepository.Update(admin)
}
// UpdatePassword 更新管理员密码
func (as *AdminService) UpdatePassword(admin *model.Admin) error {
	// 1. 更新到缓冲器
	cache.UpdateAdminPwd(admin.Password)
	// 2. 更新到数据库
	return as.AdminRepository.UpdatePassword(admin)
}
// GetAdminInfo 获取管理员信息(不包含密码)
func (as *AdminService) GetAdminInfo() (*model.Admin, error) {
	// 1. 从缓存器取数据
	admin, isExist := cache.GetAdminInfo()
	if isExist {
		return admin, nil
	}
	// 2. 从数据库取数据，并缓存到缓存器
	admin, err := as.AdminRepository.Select()
	if err != nil {
		return admin, err
	}
	admin.Password = "**********" // 隐藏密码
	cache.UpdateAdminInfo(admin)	// 隐藏密码后缓存，保证缓存器中的admin不包含密码
	return admin, nil
}
// GetAdminPassword 获取管理员加密的密码
func (as *AdminService) GetAdminPassword() (string, error) {
	// 1. 从缓存器取数据
	pwd, isExist := cache.GetAdminPwd()
	if isExist {
		return pwd, nil
	}
	// 2. 从数据库取数据，并缓存到缓存器
	admin, err := as.AdminRepository.Select()
	if err != nil {
		return "", err
	}
	cache.UpdateAdminPwd(admin.Password)
	return admin.Password, nil
}

