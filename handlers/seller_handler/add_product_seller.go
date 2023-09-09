package sellerhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"

	"github.com/gin-gonic/gin"
)

func Add_Product(c *gin.Context) {
	var users []models.Product
	config.DB.Find(&users)
	c.JSON(200, &users)
}
