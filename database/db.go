package database

import (
	"log"

	"github.com/talles-morais/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("error connecting to database")
	}
	DB.AutoMigrate(&models.Student{})
}
