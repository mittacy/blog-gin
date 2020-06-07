package repository

import (
	"github.com/crazychat/blog-gin/database"
	"github.com/crazychat/blog-gin/model"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type IArticleRepository interface {
	Conn() error
	Add(article *model.Article) error
	Delete(int) error
	Update(article *model.Article) error
	UpdateViews(id uint32, views int) error
	Select() ([]model.Article, error)
	SelectByID(id int) (*model.Article, error)
	SelectByPage(page, onePageArticleCount int) ([]model.Article, int, error)
	SelectRecent() ([]model.Article, error)
	SelectByCategoryID(cateID, onePageArticleCount, page int) ([]model.Article, int, error)
}

func NewArticleRepository(table string) IArticleRepository {
	return &ArticleRepository{table, database.MysqlDB}
}

type ArticleRepository struct {
	table string
	mysqlConn *sqlx.DB
}

func (ar *ArticleRepository) Conn() error {
	if ar.mysqlConn == nil {
		if err := database.ConnectMysql(); err != nil {
			return err
		}
		ar.mysqlConn = database.MysqlDB
	}
	if ar.table == "" {
		ar.table = "article"
	}
	return nil
}

func (ar *ArticleRepository) Add(article *model.Article) error {
	if err := ar.Conn(); err != nil {
		return err
	}
	sql := "insert into "+ ar.table +"(created_at, updated_at, category_id, title, content) values(?,?,?,?,?)"
	stmt, err := ar.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(article.CreatedAt, article.CreatedAt, article.CategoryID, article.Title, article.Content)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	article.ID = uint32(id)
	return nil
}

func (ar *ArticleRepository) Delete(id int) error {
	if err := ar.Conn(); err != nil {
		return err
	}
	sqlStr := "delete from " + ar.table + " where id = ?"
	stmt, err := ar.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (ar *ArticleRepository) Update(article *model.Article) error {
	if err := ar.Conn(); err != nil {
		return err
	}
	sqlStr := "update " + ar.table + " set updated_at = ?, category_id = ?, title = ?, content = ? WHERE id = ?"
	stmt, err := ar.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(article.UpdatedAt, article.CategoryID, article.Title, article.Content, article.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ar *ArticleRepository) UpdateViews(id uint32, views int) error {
	if err := ar.Conn(); err != nil {
		return err
	}
	sqlStr := "update " + ar.table + " set views = views + ? WHERE id = ?"
	stmt, err := ar.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(views, id)
	if err != nil {
		return err
	}
	return nil
}

func (ar *ArticleRepository) Select() (articles []model.Article, err error) {
	if err = ar.Conn(); err != nil {
		return
	}
	sqlStr := "select id, created_at, updated_at, category_id, title, views FROM " + ar.table + " order by id DESC"
	articles = []model.Article{}
	err = ar.mysqlConn.Select(&articles, sqlStr)
	return
}

func (ar *ArticleRepository) SelectByID(id int) (article *model.Article, err error) {
	if err = ar.Conn(); err != nil {
		return
	}
	sqlStr := "select * from " + ar.table + " where id = ?"
	article = &model.Article{}
	err = ar.mysqlConn.Get(article, sqlStr, id)
	return
}

func (ar *ArticleRepository) SelectByPage(page, onePageArticleCount int) (articles []model.Article, articleCount int, err error) {
	if err = ar.Conn(); err != nil {
		return
	}
	// 1. 查询 articles
	startIndex := strconv.Itoa(page * onePageArticleCount)
	sqlStr := "select id, created_at, updated_at, title, views from " + ar.table + " order by id DESC limit " + startIndex + ", " + strconv.Itoa(onePageArticleCount)
	articles = []model.Article{}
	if err = ar.mysqlConn.Select(&articles, sqlStr); err != nil {
		return
	}
	// 2. 查询 articleCount
	sqlStr = "select count(*) from " + ar.table
	err = ar.mysqlConn.QueryRow(sqlStr).Scan(&articleCount)
	return
}

func (ar *ArticleRepository) SelectRecent() (articles []model.Article, err error) {
	if err = ar.Conn(); err != nil {
		return
	}
	sqlStr := "select id, created_at, updated_at, title, views from " + ar.table + " order by updated_at desc limit 0, 5"
	articles  = []model.Article{}
	err = ar.mysqlConn.Select(&articles, sqlStr)
	return
}

func (ar *ArticleRepository) SelectByCategoryID(cateID, onePageArticleCount, page int) (articles []model.Article, articleCount int, err error) {
	if err = ar.Conn(); err != nil {
		return
	}
	// 1. 查询 articles
	startIndex := strconv.Itoa(page * onePageArticleCount)
	sqlStr := "select id, created_at, updated_at, title, views from " + ar.table + " where category_id = ? order by id desc limit  " + startIndex + ", " + strconv.Itoa(onePageArticleCount)
	articles = []model.Article{}
	if err = ar.mysqlConn.Select(&articles, sqlStr, cateID); err != nil {
		return
	}
	// 2. 查询 articleCount
	sqlStr = "select count(*) from " + ar.table + " where category_id = ?"
	err = ar.mysqlConn.QueryRow(sqlStr, cateID).Scan(&articleCount)
	return
}