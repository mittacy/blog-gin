package service

import (
	"github.com/mittacy/blog-gin/cache"
	"github.com/mittacy/blog-gin/model"
	"github.com/mittacy/blog-gin/repository"
)

type IAdminService interface {
	CreateAdmin(admin *model.Admin) error
	UpdateAdminInfo(admin *model.Admin) error
	UpdateAdminPassword(admin *model.Admin) error
	GetAdminInfo() (*model.Admin, error)
	GetAdmin() (*model.Admin, error)
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
	// 1. 更新到缓冲器，不更新密码，所以需要保存一下密码
	temp, isExist := cache.GetAdminCache()
	if isExist {
		admin.Password = temp.Password
		cache.UpdateAdminCache(admin)
	}
	// 2. 更新到数据库
	return as.AdminRepository.Update(admin)
}
// UpdatePassword 更新管理员密码
func (as *AdminService) UpdateAdminPassword(admin *model.Admin) error {
	// 1. 更新到缓冲器，只更新密码
	temp, isExist := cache.GetAdminCache()
	if isExist {
		temp.Password = admin.Password
		cache.UpdateAdminCache(&temp)
	}
	// 2. 更新到数据库
	return as.AdminRepository.UpdatePassword(admin)
}
// GetAdminInfo 获取管理员信息(不包含密码)
func (as *AdminService) GetAdminInfo() (*model.Admin, error) {
	// 1. 从缓存器取数据
	admin, isExist := cache.GetAdminCache()
	if isExist {
		admin.Password = "**********"
		return &admin, nil
	}
	// 2. 从数据库取数据，并缓存到缓存器
	adminData, err := as.AdminRepository.Select()
	if err != nil {
		return adminData, err
	}
	cache.UpdateAdminCache(adminData)	// 缓存
	return adminData, nil
}
// GetAdmin 获取管理员信息(包含密码)
func (as *AdminService) GetAdmin() (*model.Admin, error) {
	// 1. 从缓存器取数据
	admin, isExist := cache.GetAdminCache()
	if isExist {
		return &admin, nil
	}
	// 2. 从数据库取数据，并缓存到缓存器
	adminData, err := as.AdminRepository.Select()
	if err != nil {
		return &model.Admin{}, err
	}
	cache.UpdateAdminCache(adminData)	// 缓存
	return adminData, nil
}

