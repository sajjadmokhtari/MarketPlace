package handler

import (
	"MarketPlace/validations"
	"github.com/gin-gonic/gin"
	"net/http"
)


func CheckPhoneHandler(c *gin.Context) {
	var req PhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Valid:   false,
			Message: "درخواست نامعتبر",
		})
		return
	}

	if !validations.CheckIranianMobile(req.Phone) {
		c.JSON(http.StatusBadRequest, Response{
			Valid:   false,
			Message: "شماره موبایل معتبر نیست",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Valid: true,
	})
}
