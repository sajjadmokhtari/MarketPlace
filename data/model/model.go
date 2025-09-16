package model

import (
	"time"

	"gorm.io/gorm"
)

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
	Title       string   `gorm:"type:varchar(255);not null" json:"Title"` // عنوان
	Description string   `gorm:"type:text" json:"Description"`            // توضیحات
	Price       float64  `gorm:"not null" json:"Price"`                   // قیمت
	ImagePath   string   `gorm:"type:varchar(255)" json:"ImagePath"`      // مسیر عکس
	CityID      uint     `gorm:"not null" json:"CityID"`                  // شناسه شهر
	City        City     `gorm:"foreignKey:CityID" json:"City"`           // ارتباط با جدول شهر
	CategoryID  uint     `gorm:"not null" json:"CategoryID"`              // شناسه دسته‌بندی
	Category    Category `gorm:"foreignKey:CategoryID" json:"Category"`   // ارتباط با جدول دسته‌بندی
	Phone       string   `gorm:"type:varchar(20);not null" json:"Phone"`  // شماره تماس

	UserID uint `gorm:"not null" json:"UserID"`        // 👈 شناسه کاربر
	User   User `gorm:"foreignKey:UserID" json:"User"` // 👈 ارتباط با جدول کاربر
}

type MongoListing struct {
	ID          uint    `bson:"id"`
	Title       string  `bson:"title"`
	Price       float64 `bson:"price"`
	Description string  `bson:"description"`
	ImagePath   string  `bson:"image_path"`
	Phone       string  `bson:"phone"`
	City        string  `bson:"city"`
	Category    string  `bson:"category"`
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Phone     string `gorm:"uniqueIndex;not null"`
	Role      string `gorm:"default:user"`
	IsBlocked bool   `gorm:"default:false"`
	CreatedAt time.Time
	LastLogin *time.Time

	Listings []Listing `gorm:"foreignKey:UserID"` // 👈 آگهی‌های متعلق به کاربر
}
