package model

type JsonID struct {
	ID           uint32 `json:"id" db:"id" binding:"required"`
}
