package models

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"io"
)

// SQL查询语句
var (
	GETADMINSQL string = "SELECT name, views, cname, introduce, github, mail FROM admin limit 1"
)

type Admin struct {
	ID        uint32
	Name      string
	Password  string
	Views     uint32
	Cname     string
	Introduce string
	Github    string
	Mail      string
}

// CreateAdmin 创建管理员信息
func CreateAdmin() (string, error) {
	// admin是否存在，不存在则创建
	row := db.QueryRow("SELECT id FROM admin limit 1")
	if err := row.Scan(); err != sql.ErrNoRows {
		fmt.Println(ADMIN_EXIST)
		return CONTROLLER_SUCCESS, nil
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
	return CONTROLLER_SUCCESS, nil
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
	if adminName != admin.Name {
		return NAMEERROR, errors.New(NAMEERROR)
	}
	if Encryption(admin.Password) != adminPwd {
		return PASSWORDERROR, errors.New(PASSWORDERROR)
	}
	return CONTROLLER_SUCCESS, nil
}

// GetAdmin 获取管理员信息
func GetAdmin() (*Admin, string, error) {
	var admin Admin
	err := db.Get(&admin, GETADMINSQL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ADMIN_NO_EXIST, err
		}
		return nil, SQL_ERROR, err
	}
	admin.Password = "**********"
	return &admin, CONTROLLER_SUCCESS, nil
}

// SetAdmin 修改管理员信息
func SetAdmin(admin *Admin) (string, error) {
	stmt, err := db.Prepare("UPDATE admin SET cname = ?, introduce = ?, github = ?, mail = ? limit 1")
	if err != nil {
		return SQL_ERROR, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Cname, admin.Introduce, admin.Github, admin.Mail)
	if err != nil {
		if err == sql.ErrNoRows {
			return ADMIN_NO_EXIST, err
		}
		return SQL_ERROR, err
	}
	admin.Password = "**********"
	return CONTROLLER_SUCCESS, nil
}

// SetPassword 修改管理员密码
func SetPassword(password string) (string, error) {
	pwd := Encryption(password)
	stmt, err := db.Prepare("UPDATE admin SET password = ? limit 1")
	if err != nil {
		if err == sql.ErrNoRows {
			return ADMIN_NO_EXIST, err
		}
		return SQL_ERROR, err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(pwd); err != nil {
		return SQL_ERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// AddAdminView 添加访问量
func AddAdminView() (string, error) {
	if _, err := db.Exec("UPDATE admin SET views = views + 1 limit 1"); err != nil {
		return SQL_ERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// Encryption 密码加密
func Encryption(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
