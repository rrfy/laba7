package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"not null"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"not null"`
	Price      float64 `gorm:"not null"`
	CategoryID uint
}

func main() {
	// Подключение к базе данных
	dsn := "host=localhost user=username password=password dbname=mydb port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Миграция
	db.AutoMigrate(&Category{}, &Product{})

	// Создание категории и продуктов
	category := Category{Name: "Electronics", Products: []Product{
		{Name: "Smartphone", Price: 699.99},
		{Name: "Laptop", Price: 999.99},
	}}
	db.Create(&category)

	// Чтение продуктов по категории
	var categoryFromDb Category
	db.Preload("Products").First(&categoryFromDb, "name = ?", "Electronics")
	println("Category:", categoryFromDb.Name)
	for _, product := range categoryFromDb.Products {
		println("Product:", product.Name, "Price:", product.Price)
	}

	// Обновление категории у продукта
	var product Product
	db.First(&product, "name = ?", "Smartphone")
	var newCategory Category
	db.FirstOrCreate(&newCategory, Category{Name: "Gadgets"})
	db.Model(&product).Update("CategoryID", newCategory.ID)

	// Удаление категории и связанных продуктов
	db.Delete(&categoryFromDb)
}
