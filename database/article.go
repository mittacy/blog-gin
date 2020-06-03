package database

type Article struct {
	ID         uint32    `json:"id" db:"id"`
	CreatedAt  string    `json:"created_at" db:"created_at" binding:"required"`
	UpdatedAt  string	 	 `json:"updated_at" db:"updated_at"`
	CategoryID uint32    `json:"category_id" db:"category_id" binding:"required"`
	Views      uint32    `json:"views" db:"views"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Content    string    `json:"content" db:"content"`
}
