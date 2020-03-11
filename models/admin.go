package models

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"io"
)

// SQL查询语句
const (
	SQL_GETADMIN string = "SELECT name, views, cname, introduce, github, mail, bilibili FROM admin limit 1"
	SQL_PUTADMIN string = "UPDATE admin SET cname = ?, introduce = ?, github = ?, mail = ?, bilibili = ? limit 1"
	SQL_PUTPASSWORD string = "UPDATE admin SET password = ? limit 1"
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
}

// CreateAdmin 创建管理员信息
func CreateAdmin() (string, error) {
	// admin是否存在，不存在则创建
	if err := mysqlDB.QueryRow("SELECT id FROM admin limit 1").Scan(); err != sql.ErrNoRows {
		return CONTROLLER_SUCCESS, nil
	}
	admin := Admin{
		Name:      "Mittacy",
		Password:  Encryption("admin"),
		Cname:     "陈铭涛",
		Introduce: "就读佛山大学 - 大三 - 计算机系",
		Github:    "https://github.com/crazychat",
		Mail:      "mail@mittacy.com",
		Bilibili:  "https://space.bilibili.com/384942135",
	}
	stmt, err := mysqlDB.Prepare("INSERT INTO admin(name, password, cname, introduce, github, mail, bilibili) values (?,?,?,?,?,?,?)")
	if err != nil {
		return BACKERROR, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Name, admin.Password, admin.Cname, admin.Introduce, admin.Github, admin.Mail, admin.Bilibili)
	if err != nil {
		return BACKERROR, err
	}
	// 缓存密码到redis
	if _, err := SavePassword(admin.Password); err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// CheckPassword 检验密码是否正确
func CheckPassword(admin *Admin) (string, error) {
	// 验证用户名
	if admin.Name != adminName {
		return NAMEERROR, errors.New(NAMEERROR)
	}
	// 获取redis缓存的admin密码
	pwd, err := redisDB.Get(adminPassword).Result()
	if err != nil {
		return NO_EXIST, errors.New(NO_EXIST)
	}
	// 验证密码
	if Encryption(admin.Password) != pwd {
		return PASSWORDERROR, errors.New(PASSWORDERROR)
	}
	return CONTROLLER_SUCCESS, nil
}

// GetAdmin 获取管理员信息
func GetAdmin() (*Admin, string, error) {
	var admin Admin
	err := mysqlDB.Get(&admin, SQL_GETADMIN)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NO_EXIST, err
		}
		return nil, BACKERROR, err
	}
	admin.Password = "**********"
	return &admin, CONTROLLER_SUCCESS, nil
}

// SetAdmin 修改管理员信息
func SetAdmin(admin *Admin) (string, error) {
	stmt, err := mysqlDB.Prepare(SQL_PUTADMIN)
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
	return CONTROLLER_SUCCESS, nil
}

// SetPassword 修改管理员密码
func SetPassword(password string) (string, error) {
	stmt, err := mysqlDB.Prepare(SQL_PUTPASSWORD)
	if err != nil {
		if err == sql.ErrNoRows {
			return NO_EXIST, err
		}
		return FAILEDERROR, err
	}
	defer stmt.Close()
	pwd := Encryption(password)
	if _, err = stmt.Exec(pwd); err != nil {
		return FAILEDERROR, err
	}
	// 修改redis缓存密码
	if _, err := SavePassword(pwd); err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// Encryption 密码加密
func Encryption(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Verify 验证登录状态
func Verify(token string) bool {
	result, err := redisDB.Get(TokenName).Result()
	if err != nil || result != token{
		return false
	}
	return true
}
