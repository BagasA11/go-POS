package models

import (
	"time"

	"gorm.io/gorm"
)

type Return struct {
	gorm.Model
	ID         uint `gorm:"type:integer; primaryKey"`
	SaleID     uint
	Sale       Sale
	ItemID     uint
	Items      Item
	Qty        uint `gorm:"type:integer; not null"`
	ReturnDate uint `gorm:"type:integer; not null"`
}

func (rtn *Return) BeforeCreate(tx *gorm.DB) error {
	// generate time
	tx.Statement.SetColumn("ReturnDate", time.Now().Unix())
	return nil
}
