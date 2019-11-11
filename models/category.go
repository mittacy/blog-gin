package models

import (
	"database/sql"
	"time"
)

var (
	GETCATEGORIESSQL string = "SELECT id, title, article_count FROM category"
)

type Category struct {
	ID           uint32 `json:"id"`
	Title        string `json:"title" binding:"required"`
	ArticleCount uint32 `json:"article_count" db:"ArticleCount"`
}

// CreateCate 创建分类
func CreateCate(cate *Category) (string, error) {
	stmt, err := db.Prepare("INSERT INTO category(title) values (?)")
	if err != nil {
		return SQL_ERROR, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(cate.Title)
	if err != nil {
		return SQL_ERROR, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return SQL_ERROR, err
	}
	cate.ID = uint32(id)
	return "", nil
}

// GetCategories 获取全部分类
func GetCategories() (*[]Category, string, error) {
	var categories []Category
	err := db.Select(&categories, GETCATEGORIESSQL)
	if err != nil {
		return nil, SQL_ERROR, err
	}
	return &categories, "", nil
	// categories := make([]Category, 0)
	// count := 0
	// rows, err := db.Query("SELECT id, title, article_count FROM category")
	// if err != nil {
	// 	return nil, count, SQL_ERROR, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var cate Category
	// 	if err := rows.Scan(&cate.ID, &cate.Title, &cate.ArticleCount); err != nil {
	// 		return nil, count, SQL_ERROR, err
	// 	}
	// 	categories = append(categories, cate)
	// 	count++
	// }
	// return &categories, count, "", nil
}

// GetCategory 获取id分类及其所有文章
func GetCategory(cate *Category) (map[string]interface{}, string, error) {
	var result = make(map[string]interface{})
	// 获取分类的ArticleCount
	row := db.QueryRow("SELECT title, article_count FROM category WHERE id = ? limit 1", cate.ID)
	if err := row.Scan(&cate.Title, &cate.ArticleCount); err != nil {
		if err == sql.ErrNoRows {
			return result, CATE_NO_EXIST, err
		}
		return result, SQL_ERROR, err
	}
	result["CateTitle"] = cate.Title
	result["ArticleCount"] = cate.ArticleCount
	if cate.ArticleCount == 0 {
		return result, "", nil
	}
	// 查找id为category_id的所有文章
	var article Article
	var articleTime string
	articles := make([]Article, 0)
	rows, err := db.Query("SELECT id, created_at, title, views, assists FROM article WHERE category_id = ?", cate.ID)
	if err != nil {
		return result, SQL_ERROR, err
	}
	defer rows.Close()
	article.CategoryID = cate.ID
	for rows.Next() {
		if rows.Scan(&article.ID, &articleTime, &article.Title, &article.Views, &article.Assists); err != nil {
			return result, SQL_ERROR, err
		}
		article.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", articleTime)
		articles = append(articles, article)
	}
	result["articles"] = articles
	return result, "", nil
}

// UpdateCate 更新分类
func UpdateCate(cate *Category) (string, error) {
	stmt, err := db.Prepare("UPDATE category SET title = ? WHERE id = ?")
	if err != nil {
		return SQL_ERROR, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cate.Title, cate.ID)
	if err != nil {
		return CHECKCONTENT, err
	}
	return CONTROLLER_SUCCESS, nil
}

// DeleteCategory 删除分类同时删除分类里的所有文章
func DeleteCategory(cateID uint32) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return SQL_ERROR, err
	}
	if _, err = tx.Exec("DELETE FROM article WHERE category_id = ?", cateID); err != nil {
		tx.Rollback()
		return SQL_ERROR, err
	}
	if _, err = tx.Exec("DELETE FROM category WHERE id = ?", cateID); err != nil {
		tx.Rollback()
		return SQL_ERROR, err
	}
	if err = tx.Commit(); err != nil {
		return SQL_ERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}
