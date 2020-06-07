package repository

import (
	"database/sql"
	"github.com/crazychat/blog-gin/database"
	"github.com/crazychat/blog-gin/model"
	"github.com/jmoiron/sqlx"
)

type IAdminRepository interface {
	Conn() error
	Add(admin *model.Admin) error
	Update(admin *model.Admin) error
	UpdatePassword(admin *model.Admin) error
	Select() (*model.Admin, error)
	SelectExist() bool
	UpdateViews(views int) error
}

func NewAdminRepository(table string) IAdminRepository {
	return &AdminManagerRepository{table, database.MysqlDB}
}

type AdminManagerRepository struct {
	table string
	mysqlConn *sqlx.DB
}
// Conn 确保数据库连接正常
func (amr *AdminManagerRepository) Conn() error {
	if amr.mysqlConn == nil {
		if err := database.ConnectMysql(); err != nil {
			return err
		}
		amr.mysqlConn = database.MysqlDB
	}
	if amr.table == "" {
		amr.table = "admin"
	}
	return nil
}
// Add 添加管理员
func (amr *AdminManagerRepository) Add(admin *model.Admin) error {
	if err := amr.Conn(); err != nil {
		return err
	}

	sqlStr := "insert into " + amr.table + "(name, password, views, cname, introduce, github, mail, bilibili) values (?,?,?,?,?,?,?,?)"
	stmt, err := amr.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Name, admin.Password, admin.Views, admin.Cname, admin.Introduce, admin.Github, admin.Mail, admin.Bilibili)
	if err != nil {
		return err
	}
	return nil
}
// Update 更新管理员除密码外的其他信息
func (amr *AdminManagerRepository) Update(admin *model.Admin) error {
	if err := amr.Conn(); err != nil {
		return err
	}
	sqlStr := "UPDATE " + amr.table + " SET name = ?, cname = ?, introduce = ?, github = ?, mail = ?, bilibili = ? limit 1"
	stmt, err := amr.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Name, admin.Cname, admin.Introduce, admin.Github, admin.Mail, admin.Bilibili)
	if err != nil {
		return err
	}
	return nil
}
// UpdatePassword 更新管理员密码
func (amr *AdminManagerRepository) UpdatePassword(admin *model.Admin) error {
	if err := amr.Conn(); err != nil {
		return err
	}
	sqlStr := "UPDATE " + amr.table + " SET password = ? limit 1"
	stmt, err := amr.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Password)
	if err != nil {
		return err
	}
	return nil
}
// Select 获取管理员所有信息(包括加密的密码)
func (amr *AdminManagerRepository) Select() (*model.Admin, error) {
	if err := amr.Conn(); err != nil {
		return &model.Admin{}, err
	}
	sqlStr := "select * from " + amr.table + " limit 1"
	admin := &model.Admin{}
	err := amr.mysqlConn.Get(admin, sqlStr)
	return admin, err
}
// SelectExist 管理员是否存在
func (amr *AdminManagerRepository) SelectExist() bool {
	if err := amr.Conn(); err != nil {
		return false
	}
	if err := amr.mysqlConn.QueryRow("select id from " + amr.table + " limit 1").Scan(); err != sql.ErrNoRows {
		return true
	}
	return false
}
// UpdateViews 更新博客访问量
func (amr *AdminManagerRepository) UpdateViews(views int) error {
	if err := amr.Conn(); err != nil {
		return err
	}
	sqlStr := "update " + amr.table + " set views = views + ? limit 1"
	stmt, err := amr.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(views)
	if err != nil {
		return err
	}
	return nil
}