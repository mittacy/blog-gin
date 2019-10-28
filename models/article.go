package models

import (
	"time"
)

type Article struct {
	ID         uint
	CreatedAt  time.Time
	CategoryID int    `json:"category_id" binding:"required"`
	Title      string `binding:"required"`
	Content    string
	Views      int
	Assists    int
}

// CreateArticle 创建文章model
func CreateArticle(article *Article) (string, error) {
	stmt, err := db.Prepare("INSERT INTO article(category_id, title, content) values (?,?,?)")
	if err != nil {
		return SQL_ERROR, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(article.CategoryID, article.Title, article.Content)
	if err != nil {
		return CHECKCONTENT, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return SQL_ERROR, err
	}
	article.ID = uint(id)
	return "", nil
}

// // GetArticles 获取全部文章
// func GetArticles(startID, endID int) {
// }

// // GetArticle 根据id获取文章
// func GetArticle(articleID int) (*Article, string, error) {
// 	article := Article{}
// 	var createdAt string
// 	row := sqlDb.QueryRow("SELECT * FROM article WHERE id = ?", articleID)
// 	err := row.Scan(&article.ID, &createdAt, &article.CategoryID, &article.Title, &article.Content, &article.Views, &article.Assists)
// 	if err == sql.ErrNoRows {
// 		return nil, ARTICLE_NO_EXIST, err
// 	}
// 	if err != nil {
// 		return nil, SQL_ERROR, err
// 	}
// 	if article.CreatedAt, err = time.ParseInLocation("2006-01-02 15:04:05", createdAt, time.Local); err != nil {
// 		return nil, SQL_ERROR, err
// 	}
// 	return &article, "", nil
// }
