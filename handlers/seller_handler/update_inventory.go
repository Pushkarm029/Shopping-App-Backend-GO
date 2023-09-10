package sellerhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update_Product(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusCreated, &users)
}
