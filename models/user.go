package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null;unique" validate:"required"`
	Email    string `gorm:"not null;unique" validate:"email"`
	Password string `gorm:"not null" validate:"required"`
}
