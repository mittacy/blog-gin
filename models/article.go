package models

import (
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type Article struct {
	ID         uint32    `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	CategoryID uint32    `json:"category_id" db:"category_id" binding:"required"`
	Views      uint32    `json:"views" db:"views"`
	Assists    uint32    `json:"assists" db:"assists"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Content    string    `json:"content" db:"content"`
}

// CreateArticle 创建文章model
func CreateArticle(article *Article) (string, error) {
	stmt, err := db.Prepare("INSERT INTO article(category_id, title, content) values (?,?,?)")
	if err != nil {
		return BACKERROR, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(article.CategoryID, article.Title, article.Content)
	if err != nil {
		return BACKERROR, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return BACKERROR, err
	}
	article.ID = uint32(id)
	return CONTROLLER_SUCCESS, nil
}

// UpdateArticle 修改文章model
func UpdateArticle(article *Article) (string, error) {
	stmt, err := db.Prepare("UPDATE article SET category_id = ?, title = ?, content = ? WHERE id = ?")
	if err != nil {
		return BACKERROR, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(article.CategoryID, article.Title, article.Content, article.ID)
	if err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// GetArticle 根据id获取文章
func GetArticle(articleID int) (*Article, string, error) {
	article := Article{}
	var createdAt string
	row := db.QueryRow("SELECT * FROM article WHERE id = ?", articleID)
	err := row.Scan(&article.ID, &createdAt, &article.CategoryID, &article.Title, &article.Content, &article.Views, &article.Assists)
	if err == sql.ErrNoRows {
		return nil, NO_EXIST, err
	}
	if err != nil {
		return nil, BACKERROR, err
	}
	if article.CreatedAt, err = time.ParseInLocation("2006-01-02", createdAt, time.Local); err != nil {
		return nil, BACKERROR, err
	}
	return &article, CONTROLLER_SUCCESS, nil
}

// DeleteArticle 根据id删除文章
func DeleteArticle(articleID uint32) (string, error) {
	if _, err := db.Exec("DELETE FROM article WHERE id = ?", articleID); err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// GetPageArticles 分页获取文章
func GetPageArticles(page, onePageArticlesCount int) ([]Article, string, error) {
	startIndex := strconv.Itoa(page * onePageArticlesCount)
	sql := "SELECT id, created_at, title, views, assists FROM article limit " + startIndex + ", " + strconv.Itoa(onePageArticlesCount)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, BACKERROR, err
	}
	defer rows.Close()

	articles := make([]Article, 0)
	IsEmptyRows := true
	for rows.Next() {
		article := Article{}
		var createdAt string
		if err = rows.Scan(&article.ID, &createdAt, &article.Title, &article.Views, &article.Assists); err != nil {
			return nil, BACKERROR, err
		}
		if article.CreatedAt, err = time.ParseInLocation("2006-01-02 15:04:05", createdAt, time.Local); err != nil {
			return nil, BACKERROR, err
		}
		articles = append(articles, article)
		IsEmptyRows = false
	}
	if IsEmptyRows {
		return nil, NO_EXIST, errors.New(NO_EXIST)
	}
	return articles, CONTROLLER_SUCCESS, nil
}

// GetArticlesCount 获取文章总数
func GetArticlesCount() (int, string, error) {
	var count int
	err := db.QueryRow("SELECT count(*) FROM article").Scan(&count)
	return count, BACKERROR, err
}

// AddArticleViews 添加文章浏览数
func AddArticleViews(id uint32) (string, error) {
	if _, err := db.Exec("UPDATE article SET views = views+1 WHERE id = ?", id); err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// AddArticleAssists 添加文章点赞数
func AddArticleAssists(id uint32) (string, error) {
	if _, err := db.Exec("UPDATE article SET assists = assists+1 WHERE id = ?", id); err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}
