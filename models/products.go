package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	Price    int    `json:"price"`
	SellerId string `json:"sellerid"`
	Category string `json:"category"`
}
