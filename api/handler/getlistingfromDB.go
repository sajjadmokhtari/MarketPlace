package handler

import (
	"MarketPlace/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListingsHandler(c *gin.Context) {
    categoryID := c.Query("category") // گرفتن پارامتر از URL

    listings, err := services.GetAllListings(categoryID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, listings)
}
