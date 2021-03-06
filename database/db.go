package database

import (
	"log"

	"github.com/david-luk4s/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() *gorm.DB {

	dsn := "host=localhost user=postgres password=postgres dbname=dbstudents port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic(err.Error())
	}
	DB.AutoMigrate(&models.Student{})
	return DB
}

func ConnectionDBTest() *gorm.DB {
	DB, err = gorm.Open(sqlite.Open("gorm_test.db"), &gorm.Config{})

	if err != nil {
		log.Panic(err.Error())
	}
	DB.AutoMigrate()
	return DB
}
