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
	GETADMINSQL string = "SELECT name, views, cname, introduce, github, mail, bilibili FROM admin limit 1"
)

type Admin struct {
	ID        uint32	`json:"id"`
	Name      string	`json:"name"`
	Password  string	`json:"password"`
	Views     uint32	`json:"views"`
	Cname     string	`json:"cname"`
	Introduce string	`json:"introduce"`
	Github    string	`json:"github"`
	Mail      string	`json:"mail"`
	Bilibili  string	`json:"bilibili"`
	Homearticle uint32	`json:"homearticle"`
}

// CreateAdmin 创建管理员信息
func CreateAdmin() (string, error) {
	// admin是否存在，不存在则创建
	row := db.QueryRow("SELECT id FROM admin limit 1")
	if err := row.Scan(); err != sql.ErrNoRows {
		return CONTROLLER_SUCCESS, nil
	}
	password := Encryption("admin")
	admin := Admin{
		Name:      "Mittacy",
		Password:  password,
		Cname:     "陈铭涛",
		Introduce: "就读佛山大学 - 大三 - 计算机系",
		Github:    "https://github.com/crazychat",
		Mail:      "mail@mittacy.com",
		Bilibili:  "https://space.bilibili.com/384942135",
	}
	stmt, err := db.Prepare("INSERT INTO admin(name, password, cname, introduce, github, mail, bilibili) values (?,?,?,?,?,?,?)")
	if err != nil {
		return FAILEDERROR, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Name, admin.Password, admin.Cname, admin.Introduce, admin.Github, admin.Mail, admin.Bilibili)
	if err != nil {
		return "创建管理员失败", err
	}
	return CONTROLLER_SUCCESS, nil
}

// GetArticleID 获取首页文章id
func GetArticleID() (*Admin, string, error) {
	var admin Admin
	err := db.Get(&admin, "SELECT homearticle FROM admin limit 1")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NO_EXIST, err
		}
		return nil, FAILEDERROR, err
	}
	return &admin, CONTROLLER_SUCCESS, nil
}

// IsRightAdmin 检验密码是否正确
func IsRightAdmin(admin *Admin) (string, error) {
	var adminName, adminPwd string
	row := db.QueryRow("SELECT name, password FROM admin limit 1")
	err := row.Scan(&adminName, &adminPwd)
	if err != nil {
		if err == sql.ErrNoRows {
			return NO_EXIST, err
		}
		return FAILEDERROR, err
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
			return nil, NO_EXIST, err
		}
		return nil, FAILEDERROR, err
	}
	admin.Password = "**********"
	return &admin, CONTROLLER_SUCCESS, nil
}

// SetAdmin 修改管理员信息
func SetAdmin(admin *Admin) (string, error) {
	stmt, err := db.Prepare("UPDATE admin SET cname = ?, introduce = ?, github = ?, mail = ?, bilibili = ? limit 1")
	if err != nil {
		return FAILEDERROR, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Cname, admin.Introduce, admin.Github, admin.Mail, admin.Bilibili)
	if err != nil {
		if err == sql.ErrNoRows {
			return NO_EXIST, err
		}
		return FAILEDERROR, err
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
			return NO_EXIST, err
		}
		return FAILEDERROR, err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(pwd); err != nil {
		return FAILEDERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// Encryption 密码加密
func Encryption(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SetArticleID 修改主页文章id
func SetArticleID(id uint32) (string, error) {
	stmt, err := db.Prepare("UPDATE admin SET homearticle = ? limit 1")
	if err != nil {
		return FAILEDERROR, err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return NO_EXIST, err
		}
		return FAILEDERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// AddViews 增加博客访问量
func AddViews() bool {
	if _, err := db.Exec("UPDATE admin SET views = views + 1 limit 1"); err != nil {
		return false
	}
	return true
}
