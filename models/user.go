package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}
