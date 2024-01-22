package models

import (
	"gorm.io/gorm"
)

type SaleDetail struct {
	gorm.Model
	ID     uint  `gorm:"type:integer; primaryKey"`
	Qty    int16 `gorm:"type:integer; not null"`
	SaleID uint
	Sale   Sale
	ItemId string
	Item   Item
}
