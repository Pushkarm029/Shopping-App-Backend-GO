package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	User       User           `json:"user" gorm:"foreignKey:UserID"`
	UserID     uint           `json:"userid"`
	EachCartId uint           `json:"eachcartid"`
	CartItems  []EachCartItem `json:"cartitems" gorm:"foreignKey:EachCartId"`
}
