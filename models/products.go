package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	Price    int    `json:"price"`
	SellerId string `json:"sellerid"`
}
