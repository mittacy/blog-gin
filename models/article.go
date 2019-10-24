package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	CategoryID int    `gorm:"not null" binding:"required"`
	Title      string `gorm:"not null;size:100" binding:"required"`
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
	article := &Article{}
	err := db.First(article).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ARTICLE_NO_EXIST, err
		}
		return nil, SQL_ERROR, err
	}
	return article, "", nil
}
