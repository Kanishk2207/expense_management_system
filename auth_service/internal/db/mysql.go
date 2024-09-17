package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dsn string) {
	var err error

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	log.Printf("working fine on new %v", dsn)
	if err := DB.Ping(); err != nil {
		log.Printf("Could not ping the database: %v", err)
	}
	AutoMigrate()
}
