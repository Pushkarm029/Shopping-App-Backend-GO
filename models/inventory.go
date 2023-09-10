package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	ProductID uint    `json:"productid"`
}
