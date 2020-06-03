package model

type Category struct {
	ID           uint32 `json:"id" db:"id"`
	Title        string `json:"title" db:"title" binding:"required"`
	ArticleCount uint32 `json:"article_count" db:"article_count"`
}
