package models

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"
)

type Admin struct {
	ID        int
	Name      string
	Password  string
	Views     int
	Cname     string
	Introduce string
	Github    string
	Mail      string
}

// CreateAdmin 创建管理员信息
func CreateAdmin() (string, error) {
	row := db.QueryRow("SELECT id FROM admin limit 1")
	if err := row.Scan(); !(err == sql.ErrNoRows) {
		fmt.Println("管理员已存在")
		return "管理员已存在", nil
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
	stmt, err := db.Prepare("INSERT INTO admin(name, password, views, cname, introduce, github, mail) values (?,?,?,?,?,?,?)")
	if err != nil {
		return "创建管理员失败", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Name, admin.Password, admin.Views, admin.Cname, admin.Introduce, admin.Github, admin.Mail)
	if err != nil {
		return "创建管理员失败", err
	}
	return "创建管理员成功", nil
}

// // IsRightAdmin 检验密码是否正确
// func IsRightAdmin(admin *Admin) (string, error) {
// 	pwd, err := GetAdminPassword()
// 	if err != nil {
// 		return ADMIN_NO_EXIST, err
// 	}
// 	if Encryption(admin.Password) == pwd {
// 		return "登录成功", nil
// 	}
// 	return "密码错误", errors.New("密码错误")
// }

// // GetAdmin 获取管理员信息
// func GetAdmin() (*Admin, string, error) {
// 	name, err := GetAdminName()
// 	if err != nil {
// 		return nil, SQL_ERROR, err
// 	}
// 	admin := &Admin{}
// 	err = db.Where("name = ?", name).First(admin).Error
// 	if err != nil {
// 		return nil, SQL_ERROR, err
// 	}
// 	admin.Password = "**********"
// 	return admin, "", nil
// }

// // GetAdminName 获取管理员名字
// func GetAdminName() (string, error) {
// 	sec, err := cfg.GetSection("app")
// 	if err != nil {
// 		return NOKNOW_ERROR, err
// 	}
// 	name := sec.Key("ADMIN").MustString("debug")
// 	return name, nil
// }

// // GetAdminPassword 获取管理员密码
// func GetAdminPassword() (string, error) {
// 	name, err := GetAdminName()
// 	if err != nil {
// 		return "", err
// 	}
// 	var pwd string
// 	row := sqlDb.QueryRow("SELECT password FROM admin where name = ? limit 1", name)
// 	err = row.Scan(&pwd)
// 	return pwd, err
// }

// Encryption 密码加密
func Encryption(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
