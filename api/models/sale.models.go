package models

import (
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ID        uint  `gorm:"type:integer; primaryKey"`
	Total     int64 `gorm:"type:integer; not null"`
	Cash      int64 `gorm:"type:integer; not null"`
	CreatedAt uint  `gorm:"type:integer; not null"`
	UpdatedAt uint  `gorm:"type:integer; not null"`
	DeletedAt uint  `gorm:"type:integer; default:null"`
	CashierID uint
	Cashiers  Cashier
}
