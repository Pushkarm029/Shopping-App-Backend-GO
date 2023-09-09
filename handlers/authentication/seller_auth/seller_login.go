package sellerauth

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type onlyEmailandPass struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var seller onlyEmailandPass
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	var dataFromDB models.Seller
	if err := config.DB.Where("email = ?", seller.Email).First(&dataFromDB).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	if !CheckPasswordHash(seller.Password, dataFromDB.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
