package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    User           `json:"userid"`
	CartItems []eachCartItem `json:"cartitems"`
}
