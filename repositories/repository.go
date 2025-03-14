package main

import (
    "context"
    "fmt"
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/go-playground/validator/v10"
)

// Item مدل مربوط به آیتم است
type Item struct {
    ID    uint    `gorm:"primaryKey"`
    Name  string  `gorm:"not null" validate:"required"`
    Price float64 `validate:"required,gt=0"`
}

// IRepository اینترفیس برای انجام عملیات CRUD است
type IRepository interface {
    Create(ctx context.Context, item Item) (uint, error)
    Read(ctx context.Context, id uint) (Item, error)
    Update(ctx context.Context, item Item) error
    Delete(ctx context.Context, id uint) error
    List(ctx context.Context) ([]Item, error)
}

// Repository پیاده‌سازی IRepository
type Repository struct {
    DB        *gorm.DB
    Validator *validator.Validate
}

// NewRepository ساختار repository را ایجاد می‌کند
func NewRepository() (*Repository, error) {
    db, err := gorm.Open(sqlite.Open("items.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&Item{})
    if err != nil {
        return nil, err
    }

    return &Repository{
        DB:        db,
        Validator: validator.New(),
    }, nil
}

// Create آیتم را به دیتابیس اضافه می‌کند
func (r *Repository) Create(ctx context.Context, item Item) (uint, error) {
    // اعتبارسنجی
    if err := r.Validator.Struct(item); err != nil {
        return 0, err
    }

    result := r.DB.WithContext(ctx).Create(&item)
    return item.ID, result.Error
}

// Read آیتم با شناسه مشخص را برمی‌گرداند
func (r *Repository) Read(ctx context.Context, id uint) (Item, error) {
    var item Item
    result := r.DB.WithContext(ctx).First(&item, id)
    return item, result.Error
}

// Update آیتم به‌روز شده را ذخیره می‌کند
func (r *Repository) Update(ctx context.Context, item Item) error {
    if err := r.Validator.Struct(item); err != nil {
        return err
    }
    result := r.DB.WithContext(ctx).Save(&item)
    return result.Error
}

// Delete آیتم را بر اساس شناسه حذف می‌کند
func (r *Repository) Delete(ctx context.Context, id uint) error {
    result := r.DB.WithContext(ctx).Delete(&Item{}, id)
    return result.Error
}

// List تمام آیتم‌ها را برمی‌گرداند
func (r *Repository) List(ctx context.Context) ([]Item, error) {
    var items []Item
    result := r.DB.WithContext(ctx).Find(&items)
    return items, result.Error
}