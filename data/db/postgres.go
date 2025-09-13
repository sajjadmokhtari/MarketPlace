package db

import (
	"MarketPlace/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() error {
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=MarketPlace_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.GetLogger().Errorw("❌ failed to connect to database", "error", err)
		return err
	}
	logging.GetLogger().Infow("✅ connected to database successfully", "dsn", dsn)

	DB = db
	return nil
}

func GetDb() *gorm.DB {
	return DB
}
