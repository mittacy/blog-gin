package service

import (
	"github.com/crazychat/blog-gin/database"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/repository"
)

type IAdminService interface {
	CreateAdmin() error
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
func (as *AdminService) CreateAdmin() error {
	if as.AdminRepository.SelectExist() { // 管理员存在则跳过
		return nil
	}
	return as.AdminRepository.Add(&database.InitAdmin)
}
// UpdateAdminInfo 更新管理员信息
func (as *AdminService) UpdateAdminInfo(admin *model.Admin) error {
	return as.AdminRepository.Update(admin)
}
// UpdatePassword 更新管理员密码
func (as *AdminService) UpdatePassword(admin *model.Admin) error {
	return as.AdminRepository.UpdatePassword(admin)
}
// GetAdminInfo 获取管理员信息
func (as *AdminService) GetAdminInfo() (*model.Admin, error) {
	admin, err := as.AdminRepository.Select()
	if err == nil {
		admin.Password = "**********"
	}
	return admin, err
}
// GetAdminPassword 获取管理员加密的密码
func (as *AdminService) GetAdminPassword() (string, error) {
	admin, err := as.AdminRepository.Select()
	if err == nil {
		return admin.Password, nil
	}
	return "", err
}
