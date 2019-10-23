package models

import (
	"database/sql"
)

type Category struct {
	ID           int
	Title        string `gorm:"unique;not ull;size:50"`
	ArticleCount int    `gorm:"default:0"`
	Articles     []Article
}

// CreateCate 创建分类
func CreateCate(cate *Category) (string, error) {
	if err := db.Create(&cate).Error; err != nil {
		return CATE_EXIST, err
	}
	return "", nil
}

// IsCateExist 分类是否存在
func IsCateExist(cate *Category) bool {
	row := sqlDb.QueryRow("SELECT id FROM category WHERE id = ? limit 1", cate.ID)
	if err := row.Scan(); err == sql.ErrNoRows {
		return false
	}
	return true
}

// UpdateCate 更新分类
func UpdateCate(cate *Category) (string, error) {
	if err := db.Model(&cate).Update("title", cate.Title).Error; err != nil {
		return CATE_EXIST, err
	}
	return "", nil
}
