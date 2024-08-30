package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dns string) {
	var err error

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	log.Printf("Connected to DB")

}
