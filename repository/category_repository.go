package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mittacy/blog-gin/database"
	"github.com/mittacy/blog-gin/model"
	"strconv"
)

type ICategoryRepository interface {
	Conn() error
	Add(cate model.Category) error
	Delete(int) error
	Update(cate model.Category) error
	Select() ([]model.Category, error)
	SelectByID(id int) (*model.Category, error)
	SelectByPage(page, onePageCategoryNum int) ([]model.Category, int, error)
}

func NewCategoryRepository(table string) ICategoryRepository {
	return &CategoryManaerRepository{table, database.MysqlDB}
}

type CategoryManaerRepository struct {
	table string
	mysqlConn *sqlx.DB
}
// Conn 确保数据库连接正常
func (cmr *CategoryManaerRepository) Conn() error {
	if cmr.mysqlConn == nil {
		if err := database.ConnectMysql(); err != nil {
			return err
		}
		cmr.mysqlConn = database.MysqlDB
	}
	if cmr.table == "" {
		cmr.table = "category"
	}
	return nil
}
// Add 添加分类
func (cmr *CategoryManaerRepository) Add(cate model.Category) error {
	if err := cmr.Conn(); err != nil {
		return err
	}
	sqlStr := "insert into " + cmr.table + "(title) values (?)"
	stmt, err := cmr.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cate.Title)
	if err != nil {
		return err
	}
	return nil
}
// Delete 根据id删除分类
func (cmr *CategoryManaerRepository) Delete(id int) error {
	if err := cmr.Conn(); err != nil {
		return err
	}
	sqlStr := "delete from " + cmr.table + " where id = ?"
	stmt, err := cmr.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
// Update 修改分类title
func (cmr *CategoryManaerRepository) Update(cate model.Category) error {
	if err := cmr.Conn(); err != nil {
		return err
	}
	sqlStr := "update " + cmr.table + " set title = ? where id = ?"
	stmt, err := cmr.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cate.Title, cate.ID)
	if err != nil {
		return err
	}
	return nil
}
// Select 获取全部分类
func (cmr *CategoryManaerRepository) Select() (categories []model.Category, err error) {
	if err = cmr.Conn(); err != nil {
		return
	}
	sqlStr := "select * from " + cmr.table
	categories = []model.Category{}
	err = cmr.mysqlConn.Select(&categories, sqlStr)
	return
}
// SelectByID 通过id获取分类信息
func (cmr *CategoryManaerRepository) SelectByID(id int) (category *model.Category, err error) {
	if err = cmr.Conn(); err != nil {
		return
	}
	sqlStr := "select * from " + cmr.table + " where id = ?"
	category = &model.Category{}
	err = cmr.mysqlConn.Get(category, sqlStr, id)
	return
}
// SelectByPage 分页获取分类
func (cmr *CategoryManaerRepository) SelectByPage(page, onePageCategoryNum int) (categories []model.Category, categoryCount int, err error) {
	if err = cmr.Conn(); err != nil {
		return
	}
	// 1. 查询 categories
	startIndex := strconv.Itoa(page * onePageCategoryNum)
	sqlStr := "select * from " + cmr.table + " limit " + startIndex + ", " + strconv.Itoa(onePageCategoryNum)
	categories = []model.Category{}
	if err = cmr.mysqlConn.Select(&categories, sqlStr); err != nil {
		return
	}
	// 2. 查询 categoryCount
	sqlStr = "select count(*) from " + cmr.table
	err = cmr.mysqlConn.QueryRow(sqlStr).Scan(&categoryCount)
	return
}