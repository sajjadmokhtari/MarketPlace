package services

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"MarketPlace/pkg/metrics"
)

func GetAllListings(categoryID string) ([]model.Listing, error) {
	var listings []model.Listing
	query := db.GetDb().Preload("City").Preload("Category")

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// دیتابیس کال
	if err := query.Find(&listings).Error; err != nil {
		metrics.DbCall.WithLabelValues("find_listings", "fail").Inc()
		return nil, err
	}

	// موفقیت
	metrics.DbCall.WithLabelValues("find_listings", "success").Inc()
	return listings, nil
}
