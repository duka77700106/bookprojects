package models

type Book struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
