package models

import "gorm.io/gorm"

type EachCartItem struct {
	gorm.Model
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	ProductID uint    `json:"productid"`
	Quantity  int     `json:"quantity"`
}
