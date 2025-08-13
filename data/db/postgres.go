package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() error {
	dsn := "host=localhost user=postgres password=admin dbname=MarketPlace_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

// func CloseDb() {

// 	if sqlDB, err := dbClient.DB(); err == nil {
// 		sqlDB.Close()
// 	} else {
// 		log.Println("Error closing the database:", err)
// 	}
// }

func GetDb() *gorm.DB {
	return DB
}
