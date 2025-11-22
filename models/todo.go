package models

type Todo struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" binding:"required"`
	Status bool   `json:"status"`
}
