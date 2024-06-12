package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "host=localhost user=saturna port=5432 dbname=grpc password=saturna"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("cannot start db because %w", err)
	}

	return db
}