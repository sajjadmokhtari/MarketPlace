package handler

import (
    "MarketPlace/data/db"
    "net/http"
    "context"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

func GetActiveListingsHandler(c *gin.Context) {
    mongoClient := db.GetMongoClient()
    collection := mongoClient.Database("MarketplaceSearch").Collection("ads")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // فرض می‌کنیم همه آگهی‌ها "فعال" هستند، اگر status داری می‌تونی اضافه کنی
    count, err := collection.CountDocuments(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"count": 0})
        return
    }

    c.JSON(http.StatusOK, gin.H{"count": count})
}
