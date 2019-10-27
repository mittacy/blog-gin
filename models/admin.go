package models

import (
	"crypto/md5"
	"database/sql"
	"errors"
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

// IsRightAdmin 检验密码是否正确
func IsRightAdmin(admin *Admin) (string, error) {
	var adminName, adminPwd string
	row := db.QueryRow("SELECT name, password FROM admin limit 1")
	err := row.Scan(&adminName, &adminPwd)
	if err != nil {
		if err == sql.ErrNoRows {
			return ADMIN_NO_EXIST, err
		}
		return SQL_ERROR, err
	}
	if adminName != admin.Name || Encryption(admin.Password) != adminPwd {
		return "密码错误", errors.New("密码错误")
	}
	return "登录成功", nil
}

// GetAdmin 获取管理员信息
func GetAdmin() (*Admin, string, error) {
	admin := Admin{}
	row := db.QueryRow("SELECT name, views, cname, introduce, github, mail FROM admin limit 1")
	err := row.Scan(&admin.Name, &admin.Views, &admin.Cname, &admin.Introduce, &admin.Github, &admin.Mail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ADMIN_NO_EXIST, err
		}
		return nil, SQL_ERROR, err
	}
	admin.Password = "**********"
	return &admin, "", nil
}

// SetAdmin 修改管理员信息
func SetAdmin(admin *Admin) (string, error) {
	stmt, err := db.Prepare("UPDATE admin SET cname = ?, introduce = ?, github = ?, mail = ? limit 1")
	if err != nil {
		if err == sql.ErrNoRows {
			return ADMIN_NO_EXIST, err
		}
		return SQL_ERROR, err
	}
	_, err = stmt.Exec(admin.Cname, admin.Introduce, admin.Github, admin.Mail)
	if err != nil {
		return SQL_ERROR, err
	}
	admin.Password = "**********"
	return "", nil
}

// Encryption 密码加密
func Encryption(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
