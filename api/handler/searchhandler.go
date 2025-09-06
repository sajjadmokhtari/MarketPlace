package handler

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// SearchListingsHandler جستجو در آگهی‌ها بر اساس شهر و دسته‌بندی
func SearchListingsHandler(c *gin.Context) {
	cityQuery := c.Query("city")
	categoryQuery := c.Query("category")

	mongoClient := db.GetMongoClient()
	collection := mongoClient.Database("MarketplaceSearch").Collection("ads")

	// ساخت فیلتر با regex (case-insensitive)
	filter := bson.M{}
	if cityQuery != "" {
		filter["city"] = bson.M{"$regex": cityQuery, "$options": "i"}
	}
	if categoryQuery != "" {
		filter["category"] = bson.M{"$regex": categoryQuery, "$options": "i"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusOK, []gin.H{}) // اگر جستجو خطا داشت، آرایه خالی برگردون
		return
	}

	var mongoResults []model.MongoListing
	if err := cursor.All(ctx, &mongoResults); err != nil {
		c.JSON(http.StatusOK, []gin.H{}) // اگر خواندن نتایج خطا داشت، آرایه خالی برگردون
		return
	}

	// تبدیل داده‌ها برای فرانت
	results := make([]gin.H, 0, len(mongoResults))
	for _, l := range mongoResults {
		results = append(results, gin.H{
			"Title":       l.Title,
			"Description": l.Description,
			"Price":       l.Price,
			"ImagePath":   l.ImagePath,
			"Phone":       l.Phone,
			"City":        gin.H{"Name": l.City},       // همیشه آبجکت بساز
			"Category":    gin.H{"Name": l.Category},   // همیشه آبجکت بساز
		})
	}

	c.JSON(http.StatusOK, results)
}
