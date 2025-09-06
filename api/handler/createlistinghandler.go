package handler

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"MarketPlace/kafka"
	"MarketPlace/logging"
	"MarketPlace/pkg/metrics"
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// پاسخ استاندارد با شماره کاربر
type ListingResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	ImagePath string `json:"image_path"`
	Phone     string `json:"phone"`
	Message   string `json:"message"`
}

func CreateListingHandler(c *gin.Context) {
	// 📌 گرفتن شماره کاربر از کانتکست (از JWT)
	phoneVal, exists := c.Get("userPhone")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "شماره کاربر یافت نشد"})
		return
	}
	phone := phoneVal.(string)

	// Parse form data
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse form"})
		return
	}

	title := c.PostForm("title")
	priceStr := c.PostForm("price")
	cityName := c.PostForm("city")
	categoryName := c.PostForm("category")
	description := c.PostForm("description")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid price"})
		return
	}

	database := db.GetDb()

	// 📊 متریک City
	var city model.City
	if err := database.Where("name = ?", cityName).First(&city).Error; err != nil {
		metrics.DbCall.WithLabelValues("find_city", "fail").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"message": "City not found"})
		return
	}
	metrics.DbCall.WithLabelValues("find_city", "success").Inc()

	// 📊 متریک Category
	var category model.Category
	if err := database.Where("name = ?", categoryName).First(&category).Error; err != nil {
		metrics.DbCall.WithLabelValues("find_category", "fail").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Category not found"})
		return
	}
	metrics.DbCall.WithLabelValues("find_category", "success").Inc()

	// آپلود عکس
	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image is required"})
		return
	}
	defer file.Close()

	os.MkdirAll("uploads", os.ModePerm)

	localPath := filepath.Join("uploads", handler.Filename)
	dst, err := os.Create(localPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot save image"})
		return
	}
	defer dst.Close()
	if _, err := dst.ReadFrom(file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot save image"})
		return
	}

	imageURL := "/uploads/" + handler.Filename

	listing := model.Listing{
		Title:       title,
		Price:       price,
		Description: description,
		ImagePath:   imageURL,
		CityID:      city.ID,
		CategoryID:  category.ID,
		Phone:       phone,
	}

	// 📊 متریک ایجاد آگهی
	if err := database.Create(&listing).Error; err != nil {
		metrics.DbCall.WithLabelValues("create_listing", "fail").Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot save listing"})
		return
	}
	metrics.DbCall.WithLabelValues("create_listing", "success").Inc()

	// ذخیره در MongoDB برای سرچ
	mongoClient := db.GetMongoClient()
	collection := mongoClient.Database("MarketplaceSearch").Collection("ads")

	mongoAd := model.MongoListing{
		ID:          listing.ID,
		Title:       listing.Title,
		Price:       listing.Price,
		Description: listing.Description,
		ImagePath:   listing.ImagePath,
		Phone:       listing.Phone,
		City:        city.Name,
		Category:    category.Name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, mongoAd); err != nil {
		logging.GetLogger().Errorw("failed to insert listing into MongoDB", "error", err)
	}

	newAd := kafka.Ad{
		ID:    strconv.Itoa(int(listing.ID)),
		Title: listing.Title,
	}
	kafka.ProduceAd(newAd)

	// پاسخ موفق
	c.JSON(http.StatusCreated, ListingResponse{
		ID:        listing.ID,
		Title:     listing.Title,
		ImagePath: listing.ImagePath,
		Phone:     phone,
		Message:   "آگهی با موفقیت ثبت شد!",
	})
}
