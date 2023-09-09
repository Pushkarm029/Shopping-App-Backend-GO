package models

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	ProductID Product `json:"productid"`
	OrderDate string  `json:"orderdate"`
	Status    string  `json:"status"`
	UserID    User    `json:"userid"`
	SellerID  Seller  `json:"sellerid"`
}
