package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
)

type Admin struct {
	ID        int
	Name      string `gorm:"unique;not null;size:10"`
	Password  string `gorm:"not null"`
	Views     int    `gorm:"default:0"`
	Cname     string
	Introduce string
	Github    string
	Mail      string
}

// CreateAdmin 创建管理员信息
func CreateAdmin() error {
	name, _ := GetAdminName()
	if row := sqlDb.QueryRow("SELECT id FROM admin WHERE name = ?", name); row != nil {
		fmt.Println("管理员已存在。")
		return nil
	}
	password := Encryption("admin")
	admin := Admin{
		Name:      "mittacy",
		Password:  password,
		Views:     2352,
		Cname:     "陈铭涛",
		Introduce: "就读佛山大学 - 大三 - 计算机系",
		Github:    "https://github.com/crazychat",
		Mail:      "mail@mittacy.com",
	}
	if err := gormDb.Create(&admin).Error; err != nil {
		fmt.Println("创建管理员失败")
		return err
	}
	fmt.Println("创建管理员成功")
	return nil
}

// IsRightAdmin 检验密码是否正确
func IsRightAdmin(admin *Admin) (string, error) {
	pwd, err := GetAdminPassword()
	if err != nil {
		return ADMIN_NO_EXIST, err
	}
	if Encryption(admin.Password) == pwd {
		return "登录成功", nil
	}
	return "密码错误", errors.New("密码错误")
}

// GetAdmin 获取管理员信息
func GetAdmin() (*Admin, string, error) {
	name, err := GetAdminName()
	if err != nil {
		return nil, SQL_ERROR, err
	}
	admin := &Admin{}
	err = gormDb.Where("name = ?", name).First(admin).Error
	if err != nil {
		return nil, SQL_ERROR, err
	}
	admin.Password = "**********"
	return admin, "", nil
}

// GetAdminName 获取管理员名字
func GetAdminName() (string, error) {
	sec, err := cfg.GetSection("app")
	if err != nil {
		return "", err
	}
	name := sec.Key("ADMIN").MustString("debug")
	return name, nil
}

// GetAdminPassword 获取管理员密码
func GetAdminPassword() (string, error) {
	name, err := GetAdminName()
	if err != nil {
		return "", err
	}
	var pwd string
	row := sqlDb.QueryRow("SELECT password FROM admin where name = ?", name)
	err = row.Scan(&pwd)
	return pwd, err
}

// Encryption 密码加密
func Encryption(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
