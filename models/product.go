package models

type Product struct {
    ID       uint    `gorm:"primaryKey"`
    Name     string  `gorm:"not null" validate:"required"`
    Price    float64 `gorm:"not null;default:0" validate:"required,gt=0"`
    UserID   uint    `gorm:"not null"`
}