package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	// ProductID Product `json:"productid" gorm:"foreignKey:Email"`
	OrderDate *time.Time `json:"orderdate"`
	Status    string     `json:"status"`
	UserID    uint       `json:"userid"`
	User      User       `json:"user" gorm:"foreignKey:UserID"`
	// SellerID  Seller  `json:"sellerid"`
	SellerID  uint    `json:"sellerid"`
	Seller    Seller  `gorm:"foreignKey:SellerID" json:"seller"`
	ProductID uint    `json:"productid"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
