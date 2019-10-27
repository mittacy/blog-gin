package models

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

// // GetCategory 获取id分类的所有文章
// func GetCategory(id int) ([]Article, string, error) {
// 	rows, err := sqlDb.Query("SELECT id, created_at, title FROM article WHERE category_id = ?", id)
// 	if err != nil {
// 		return nil, SQL_ERROR, err
// 	}
// 	defer rows.Close()

// 	articles := make([]Article, 0)
// 	for rows.Next() {
// 		article := Article{}
// 		var timeStr string
// 		if err := rows.Scan(&article.ID, &timeStr, &article.Title); err != nil {
// 			return nil, SQL_ERROR, err
// 		}
// 		if article.CreatedAt, err = time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local); err != nil {
// 			return nil, SQL_ERROR, err
// 		}
// 		articles = append(articles, article)
// 	}
// 	return articles, "", nil
// }

// // IsCateExist 分类是否存在
// func IsCateExist(cate *Category) bool {
// 	row := sqlDb.QueryRow("SELECT id FROM category WHERE id = ? limit 1", cate.ID)
// 	if err := row.Scan(); err == sql.ErrNoRows {
// 		return false
// 	}
// 	return true
// }

// // UpdateCate 更新分类
// func UpdateCate(cate *Category) (string, error) {
// 	if err := db.Model(&cate).Update("title", cate.Title).Error; err != nil {
// 		return CATE_EXIST, err
// 	}
// 	return "", nil
// }

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
