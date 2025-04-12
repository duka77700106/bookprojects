package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique, notnull" json:"username"`
	Password string `json:"not null"`
	Role     string `json:"role"`
}
