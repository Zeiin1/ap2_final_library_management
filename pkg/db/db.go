package db

import (
	"github.com/sat0urn/book-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.StockDecreaseLog{})

	return Handler{
		DB: db,
	}
}
