package main

import (
	"go-store-filter/internal/infrastructure/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:p0o9i8u7@tcp(127.0.0.1:3306)/store_v2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&model.FilterType{}, &model.Filter{})
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	log.Println("Migration completed successfully!")
}
