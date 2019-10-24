package models

import (
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
