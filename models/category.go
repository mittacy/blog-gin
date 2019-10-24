package models

import (
	"database/sql"
)

type Category struct {
	ID           int
	Title        string    `gorm:"unique;not null;size:50" binding:"required"`
	ArticleCount int       `gorm:"default:0"`
	Articles     []Article `gorm:"foreignkey:CategoryID"`
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

// GetCategories 获取全部分类
func GetCategories(cates []Category) ([]Category, string, error) {
	rows, err := sqlDb.Query("SELECT id, title, article_count FROM category")
	if err != nil {
		return nil, SQL_ERROR, err
	}
	defer rows.Close()

	for rows.Next() {
		var cate Category
		if err := rows.Scan(&cate.ID, &cate.Title, &cate.ArticleCount); err != nil {
			return nil, SQL_ERROR, err
		}

		cates = append(cates, cate)
	}
	return cates, "", nil
}
