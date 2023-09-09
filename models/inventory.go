package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ProductID Product `json:"productid"`
}
