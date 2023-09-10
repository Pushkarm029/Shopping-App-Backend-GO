package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	// User       User           `gorm:"foreignKey:UserID" json:"user"`
	// UserID     uint           `json:"userid"`
	// EachCartID uint           `json:"eachcartid"`
	// CartItems  []EachCartItem `gorm:"foreignKey:EachCartID" json:"cartitems"`
}
