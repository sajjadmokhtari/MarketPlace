package services

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
)

func GetUserByPhone(phone string) (*model.User, error) {
	var user model.User

	database := db.GetDb()

	if err := database.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
