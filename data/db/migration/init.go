package migration

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"MarketPlace/logging"

	"gorm.io/gorm"
)

func Up_1() {
	database := db.GetDb()

	err := database.AutoMigrate(
		&model.City{},
		&model.Category{},
		&model.Listing{},
	)
	if err != nil {

		logging.GetLogger().Errorw("failed to migrate database: %v", err)
	}

	// بعد داده‌های پیش‌فرض رو اضافه کن
	CreateCity(database)
	CreateCategory(database)
}

func CreateCity(database *gorm.DB) {
	// چک کن که جدول City خالیه یا نه
	var count int64
	if err := database.Model(&model.City{}).Count(&count).Error; err != nil {

		logging.GetLogger().Errorw("error counting cities: %v", err)
		return
	}

	if count == 0 {
		cities := []model.City{
			{Name: "آذربایجان شرقی"},
			{Name: "آذربایجان غربی"},
			{Name: "اردبیل"},
			{Name: "اصفهان"},
			{Name: "البرز"},
			{Name: "ایلام"},
			{Name: "بوشهر"},
			{Name: "تهران"},
			{Name: "چهارمحال و بختیاری"},
			{Name: "خراسان جنوبی"},
			{Name: "خراسان رضوی"},
			{Name: "خراسان شمالی"},
			{Name: "خوزستان"},
			{Name: "زنجان"},
			{Name: "سمنان"},
			{Name: "سیستان و بلوچستان"},
			{Name: "فارس"},
			{Name: "قزوین"},
			{Name: "قم"},
			{Name: "کردستان"},
			{Name: "کرمان"},
			{Name: "کرمانشاه"},
			{Name: "کهگیلویه و بویراحمد"},
			{Name: "گلستان"},
			{Name: "گیلان"},
			{Name: "لرستان"},
			{Name: "مازندران"},
			{Name: "مرکزی"},
			{Name: "هرمزگان"},
			{Name: "همدان"},
			{Name: "یزد"},
		}

		if err := database.Create(&cities).Error; err != nil {
			logging.GetLogger().Errorw("error Creating cities: %v", err)

		} else {

			logging.GetLogger().Infow("cities inserted successfully")
		}
	}
}

func CreateCategory(database *gorm.DB) {
	var count int64
	if err := database.Model(&model.Category{}).Count(&count).Error; err != nil {

		logging.GetLogger().Errorw("error counting categories: %v", err)
		return
	}

	if count == 0 {
		categories := []model.Category{
			{Name: "الکترونیک و دیجیتال"},
			{Name: "موبایل و تبلت"},
			{Name: "لپ‌تاپ و کامپیوتر"},
			{Name: "پوشاک و مد"},
			{Name: "لوازم منزل و دکوراسیون"},
			{Name: "مبلمان"},
			{Name: "کودک و نوزاد"},
			{Name: "ورزشی و سفر"},
			{Name: "کتاب، فیلم و موسیقی"},
			{Name: "وسایل نقلیه (ماشین، موتور، قطعات)"},
			{Name: "خودرو و موتور"},
			{Name: "آلات و تجهیزات صنعتی"},
			{Name: "حیوانات خانگی و لوازم آن‌ها"},
			{Name: "سرگرمی و بازی"},
			{Name: "خرید و فروش آثار قدیمی و آنتیک"},
			{Name: "ابزار و لوازم تعمیرات"},
		}

		if err := database.Create(&categories).Error; err != nil {

			logging.GetLogger().Errorw("error creating categories: %v", err)
		} else {

			logging.GetLogger().Infow("categories inserted successfully")
		}
	}
}
