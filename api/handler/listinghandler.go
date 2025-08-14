package handler

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// پاسخ استاندارد
type ListingResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	ImagePath string `json:"image_path"`
	Message   string `json:"message"`
}

func CreateListingHandler(c *gin.Context) {
	// Parse form data (برای فایل‌ها حتما max memory بگذاریم)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB
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

	var city model.City
	if err := database.Where("name = ?", cityName).First(&city).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "City not found"})
		return
	}

	var category model.Category
	if err := database.Where("name = ?", categoryName).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Category not found"})
		return
	}

	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image is required"})
		return
	}
	defer file.Close()

	os.MkdirAll("uploads", os.ModePerm)

	// مسیر ذخیره واقعی فایل
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

	// مسیر URL که فرانت می‌فهمه
	imageURL := "/uploads/" + handler.Filename

	listing := model.Listing{
		Title:       title,
		Price:       price,
		Description: description,
		ImagePath:   imageURL, // ✅ مسیر URL
		CityID:      city.ID,
		CategoryID:  category.ID,
	}

	if err := database.Create(&listing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot save listing"})
		return
	}

	c.JSON(http.StatusCreated, ListingResponse{
		ID:        listing.ID,
		Title:     listing.Title,
		ImagePath: listing.ImagePath,
		Message:   "آگهی با موفقیت ثبت شد!",
	})
}
