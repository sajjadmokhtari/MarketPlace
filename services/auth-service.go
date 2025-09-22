package services

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"MarketPlace/logging"
	"time"

	"gorm.io/gorm"
)

func UserRole(phone string) string {
	if phone == "09911732328" {
		return "admin"
	} else {
		return "user"
	}
}

func UpdateOrCreateUser(phone, role string) (*model.User, error) {
    database := db.GetDb()
    var user model.User

    err := database.Where("phone = ?", phone).First(&user).Error
    if err == gorm.ErrRecordNotFound {
        user = model.User{
            Phone:     phone,
            Role:      role,
            CreatedAt: time.Now(),
        }
        if err := database.Create(&user).Error; err != nil {
            logging.GetLogger().Errorw("Error creating new user", "phone", phone, "error", err)
            return nil, err
        }
        logging.GetLogger().Infow("New user created", "phone", phone, "role", role)
        return &user, nil
    }

    if err != nil {
        logging.GetLogger().Errorw("Error fetching user", "phone", phone, "error", err)
        return nil, err
    }

    database.Model(&user).Update("last_login", time.Now())
    logging.GetLogger().Infow("User login time updated", "phone", phone)
    return &user, nil
}

