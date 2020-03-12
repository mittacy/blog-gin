package models

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

type Article struct {
	ID         uint32    `json:"id" db:"id"`
	CreatedAt  string    `json:"created_at" db:"created_at" binding:"required"`
	UpdatedAt  string	 	 `json:"updated_at" db:"updated_at"`
	CategoryID uint32    `json:"category_id" db:"category_id" binding:"required"`
	Views      uint32    `json:"views" db:"views"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Content    string    `json:"content" db:"content"`
}

// CreateArticle 创建文章model
func CreateArticle(article *Article) (string, error) {
	stmt, err := mysqlDB.Prepare("INSERT INTO article(created_at, category_id, title, content) values (?, ?,?,?)")
	if err != nil {
		return BACKERROR, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(article.CreatedAt, article.CategoryID, article.Title, article.Content)
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
	stmt, err := mysqlDB.Prepare("UPDATE article SET updated_at = ?, category_id = ?, title = ?, content = ? WHERE id = ?")
	if err != nil {
		return BACKERROR, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(article.UpdatedAt, article.CategoryID, article.Title, article.Content, article.ID)
	if err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// GetArticle 根据id获取文章
func GetArticle(articleID int) (*Article, string, error) {
	article := &Article{}
	err := mysqlDB.Get(article, "SELECT * FROM article WHERE id = ?", articleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NO_EXIST, err
		}
		return nil, BACKERROR, err
	}
	return article, CONTROLLER_SUCCESS, nil
}

// DeleteArticle 根据id删除文章
func DeleteArticle(articleID uint32) (string, error) {
	if _, err := mysqlDB.Exec("DELETE FROM article WHERE id = ?", articleID); err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// GetPageArticles 分页获取文章
func GetPageArticles(page, onePageArticlesCount int) ([]Article, string, error) {
	startIndex := strconv.Itoa(page * onePageArticlesCount)
	sql := "SELECT id, created_at, updated_at, title, views FROM article ORDER BY id DESC limit " + startIndex + ", " + strconv.Itoa(onePageArticlesCount)
	var articles []Article
	err := mysqlDB.Select(&articles, sql)
	if err != nil {
		return nil, BACKERROR, err
	}
	return articles, CONTROLLER_SUCCESS, nil
}

// GetArticlesCount 获取文章总数
func GetArticlesCount() (int, string, error) {
	var count int
	err := mysqlDB.QueryRow("SELECT count(*) FROM article").Scan(&count)
	return count, BACKERROR, err
}

// AddArticleViews 添加文章浏览数
func AddArticleViews(id int) (string, error) {
	if _, err := mysqlDB.Exec("UPDATE article SET views = views+1 WHERE id = ?", id); err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// GetRecentArticles 最近更新的五篇文章
func GetRecentArticles() ([]Article, string, error) {
	articlesJson, err := redisDB.Get(recentArticles).Bytes()
	if err != nil {
		return nil, BACKERROR, err
	}
	var articles []Article
	err = json.Unmarshal(articlesJson, &articles)
	if err != nil {
		return nil, BACKERROR, err
	}
	return articles, CONTROLLER_SUCCESS, nil
}