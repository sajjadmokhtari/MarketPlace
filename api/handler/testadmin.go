package handler

import "github.com/gin-gonic/gin"

func Test(c *gin.Context){

	c.JSON(200,Response{Valid: true,Message: "admin is ok "})

}