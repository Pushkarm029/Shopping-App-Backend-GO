package models

import "gorm.io/gorm"

type eachCartItem struct {
	gorm.Model
	ProductID Product `json:"productid"`
	Quantity  int     `json:"quantity"`
}
