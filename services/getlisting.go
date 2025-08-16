package services

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
)

func GetAllListings(categoryID string) ([]model.Listing, error) {
    var listings []model.Listing
    query := db.GetDb().Preload("City").Preload("Category")

    if categoryID != "" {
        query = query.Where("category_id = ?", categoryID)
    }

    if err := query.Find(&listings).Error; err != nil {
        return nil, err
    }
    return listings, nil
}
