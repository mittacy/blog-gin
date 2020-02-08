package models

import (
	"database/sql"
	"strconv"
	"time"
)

type Category struct {
	ID           uint32 `json:"id" db:"id"`
	Title        string `json:"title" binding:"required" db:"title"`
	ArticleCount uint32 `json:"article_count" db:"article_count"`
}

// CreateCate 创建分类
func CreateCate(cate *Category) (string, error) {
	stmt, err := db.Prepare("INSERT INTO category(title) values (?)")
	if err != nil {
		return BACKERROR, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(cate.Title)
	if err != nil {
		return EXISTED, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return BACKERROR, err
	}
	cate.ID = uint32(id)
	return CONTROLLER_SUCCESS, nil
}

// GetCategories 获取全部分类id和title
func GetCategories() (*[]Category, string, error) {
	var categories []Category
	err := db.Select(&categories, "SELECT id, title FROM category")
	if err != nil {
		return nil, BACKERROR, err
	}
	return &categories, CONTROLLER_SUCCESS, nil
}

// GetCategory 获取id分类及其所有文章
func GetCategory(cate *Category) (map[string]interface{}, string, error) {
	var result = make(map[string]interface{})
	// 获取分类的文章数量
	err := db.Get(cate, "SELECT title, article_count FROM category WHERE id = ?", cate.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, NO_EXIST, err
		}
		return result, BACKERROR, err
	}
	result["cateTitle"] = cate.Title
	result["articleCount"] = cate.ArticleCount
	if cate.ArticleCount == 0 {
		return result, EXISTED, nil
	}
	// 查找id为category_id的所有文章
	var article Article
	var articleTime string
	articles := make([]Article, 0)
	rows, err := db.Query("SELECT id, created_at, title, views, assists FROM article WHERE category_id = ?", cate.ID)
	if err != nil {
		return result, BACKERROR, err
	}
	defer rows.Close()
	article.CategoryID = cate.ID
	for rows.Next() {
		if rows.Scan(&article.ID, &articleTime, &article.Title, &article.Views, &article.Assists); err != nil {
			return result, BACKERROR, err
		}
		article.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", articleTime)
		articles = append(articles, article)
	}
	result["articles"] = articles
	return result, CONTROLLER_SUCCESS, nil
}

// UpdateCate 更新分类
func UpdateCate(cate *Category) (string, error) {
	stmt, err := db.Prepare("UPDATE category SET title = ? WHERE id = ?")
	if err != nil {
		return BACKERROR, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cate.Title, cate.ID)
	if err != nil {
		return EXISTED, err
	}
	return CONTROLLER_SUCCESS, nil
}

// DeleteCategory 删除分类同时删除分类里的所有文章
func DeleteCategory(cateID uint32) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return BACKERROR, err
	}
	if _, err = tx.Exec("DELETE FROM article WHERE category_id = ?", cateID); err != nil {
		tx.Rollback()
		return BACKERROR, err
	}
	if _, err = tx.Exec("DELETE FROM category WHERE id = ?", cateID); err != nil {
		tx.Rollback()
		return BACKERROR, err
	}
	if err = tx.Commit(); err != nil {
		return BACKERROR, err
	}
	return CONTROLLER_SUCCESS, nil
}

// GetPageCategories 分页获取分类
func GetPageCategories(page, onePageCategoryNum int) ([]Category, string, error) {
	startIndex := strconv.Itoa(page * onePageCategoryNum)
	sql := "SELECT * FROM category limit " + startIndex + ", " + strconv.Itoa(onePageCategoryNum)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, BACKERROR, err
	}
	defer rows.Close()

	categories := make([]Category, 0)
	for rows.Next() {
		category := Category{}
		if err = rows.Scan(&category.ID, &category.Title, &category.ArticleCount); err != nil {
			return nil, BACKERROR, err
		}
		categories = append(categories, category)
	}
	return categories, CONTROLLER_SUCCESS, nil
}

// GetCategoriesCount 获取分类总数
func GetCategoriesCount() (int, string, error) {
	var count int
	err := db.QueryRow("SELECT count(*) FROM category").Scan(&count)
	return count, BACKERROR, err
}