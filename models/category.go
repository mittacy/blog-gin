package models

import (
	"database/sql"
	"time"
)

type Category struct {
	ID           int
	Title        string `binding:"required"`
	ArticleCount int
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
	cate.ID = int(id)
	return "", nil
}

// GetCategories 获取全部分类
func GetCategories(cates []Category) ([]Category, string, error) {
	rows, err := db.Query("SELECT id, title, article_count FROM category")
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
		return err.Error(), err
	}
	return CONTROLLER_SUCCESS, nil
}

// // DeleteCategory 删除分类同时删除分类里的所有文章
// func DeleteCategory(cateID int) (string, error) {
// 	cate := Category{ID: cateID}
// 	tx := db.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	if err := tx.Error; err != nil {
// 		return SQL_ERROR, err
// 	}
// 	// 分类是否存在
// 	if !IsCateExist(&cate) {
// 		return CATE_NO_EXIST, errors.New(CATE_NO_EXIST)
// 	}
// 	// 开始事务
// 	// 删除分类
// 	if err := tx.Delete(&cate).Error; err != nil {
// 		tx.Rollback()
// 		return SQL_ERROR, err
// 	}
// 	// 删除文章
// 	if err := db.Where("category_id LIKE ?", cateID).Delete(Article{}).Error; err != nil {
// 		tx.Rollback()
// 		return SQL_ERROR, err
// 	}
// 	// 提交事务
// 	if err := tx.Commit().Error; err != nil {
// 		return SQL_ERROR, err
// 	}
// 	return "", nil
// }
