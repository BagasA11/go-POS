package models

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ID         uint  `gorm:"type:integer; primaryKey"`
	Total      int64 `gorm:"type:integer; not null"`
	Cash       int64 `gorm:"type:integer; not null"`
	CreatedAt  uint  `gorm:"type:integer; not null"`
	UpdatedAt  uint  `gorm:"type:integer; not null"`
	DeletedAt  uint  `gorm:"type:integer; default:null"`
	CashierID  uint
	Cashiers   Cashier
	Return     []Return
	SaleDetail []SaleDetail
}

func (Sale *Sale) BeforeCreate(tx *gorm.DB) error {
	// generate time
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
