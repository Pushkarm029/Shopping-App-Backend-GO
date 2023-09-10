package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Price int    `json:"price"`
	// SellerId Seller `json:"sellerid" gorm:"foreignKey:Email"`
	SellerID uint   // Define the foreign key here
	Seller   Seller `gorm:"foreignKey:SellerID" json:"seller"`
	Category string `json:"category"`
}
