package model

import "gorm.io/gorm"

// جدول شهرها
type City struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null;unique"`
	Listings []Listing // ارتباط یک به چند با آگهی‌ها
}

// جدول دسته‌بندی‌ها
type Category struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null;unique"`
	Listings []Listing // ارتباط یک به چند با آگهی‌ها
}

// جدول آگهی‌ها / لیستینگ‌ها
type Listing struct {
	gorm.Model
	Title       string   `gorm:"type:varchar(255);not null"` // عنوان
	Description string   `gorm:"type:text"`                  // توضیحات
	Price       float64  `gorm:"not null"`                   // قیمت
	ImagePath   string   `gorm:"type:varchar(255)"`          // مسیر عکس
	CityID      uint     `gorm:"not null"`                   // شناسه شهر
	City        City     `gorm:"foreignKey:CityID"`          // ارتباط با جدول شهر
	CategoryID  uint     `gorm:"not null"`                   // شناسه دسته‌بندی
	Category    Category `gorm:"foreignKey:CategoryID"`      // ارتباط با جدول دسته‌بندی
}
