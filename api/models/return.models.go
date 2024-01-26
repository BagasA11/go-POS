package models

import (
	"time"

	"gorm.io/gorm"
)

type Return struct {
	gorm.Model
	ID        uint `gorm:"type:integer; primaryKey"`
	SaleID    uint
	Sale      Sale
	ItemID    uint
	Items     Item
	Qty       uint `gorm:"type:integer; not null"`
	CreatedAt uint `gorm:"type:integer; not null"`
	UpdatedAt uint `gorm:"type:integer; not null"`
	DeletedAt uint `gorm:"type:integer; null"`
}

func (rtn *Return) BeforeCreate(tx *gorm.DB) error {
	// generate time
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
