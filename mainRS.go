package main

import (
	"log"

	"your_project/repositories"
	"your_project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// کانکشن به MySQL
	dsn := "username:password@tcp(localhost:3306)/your_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ساخت Repository و Service‌ها
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)

	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)

	// ایجاد Router
	router := gin.Default()
	api := NewApi(userService, productService)
	api.RegisterRoutes(router)

	// اجرای سرور
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
