package models

import (
	"database/sql"
	"time"
)

type Article struct {
	ID         uint
	CreatedAt  time.Time `gorm:"not null"`
	CategoryID int       `gorm:"not null" binding:"required"`
	Title      string    `gorm:"not null;size:100" binding:"required"`
	Content    string
	Views      int `gorm:"default:0"`
	Assists    int `gorm:"default:0"`
}

// CreateArticle 创建文章model
func CreateArticle(article *Article) (string, error) {
	if err := db.Create(article).Error; err != nil {
		return SQL_ERROR, err
	}
	return "", nil
}

// GetArticle 根据id获取文章
func GetArticle(articleID int) (*Article, string, error) {
	article := Article{}
	var createdAt string
	row := sqlDb.QueryRow("SELECT * FROM article WHERE id = ?", articleID)
	err := row.Scan(&article.ID, &createdAt, &article.CategoryID, &article.Title, &article.Content, &article.Views, &article.Assists)
	if err == sql.ErrNoRows {
		return nil, ARTICLE_NO_EXIST, err
	}
	if err != nil {
		return nil, SQL_ERROR, err
	}
	if article.CreatedAt, err = time.ParseInLocation("2006-01-02 15:04:05", createdAt, time.Local); err != nil {
		return nil, SQL_ERROR, err
	}
	return &article, "", nil
}
