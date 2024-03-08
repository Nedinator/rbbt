package data

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenDB(dsn string) {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&User{}, Url{})
}

func DB() *gorm.DB {
	return db
}
