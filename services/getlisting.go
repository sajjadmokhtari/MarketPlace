package services

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
)

func GetAllListings() ([]model.Listing, error) {
	var listings []model.Listing
	if err := db.GetDb().Preload("City").Preload("Category").Find(&listings).Error; err != nil {
		return nil, err
	}
	return listings, nil
}
