package models

type Category struct {
	ID           int
	title        string `gorm:"unique;not ull;size:50"`
	ArticleCount int    `gorm:"default:0"`
	Articles     []Article
}
