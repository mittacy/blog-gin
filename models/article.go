package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	CategoryID uint
	Title      string `gorm:"unique;not ull;"`
	Content    string
	Views      int `gorm:"default:0"`
	Assists    int `gorm:"default:0"`
}
