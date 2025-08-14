package handler

import (
	"MarketPlace/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListingsHandler(c *gin.Context) {
	listings, err := services.GetAllListings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, listings)
}
