package sellerhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"

	"github.com/gin-gonic/gin"
)

func Get_Product(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, &users)
}
