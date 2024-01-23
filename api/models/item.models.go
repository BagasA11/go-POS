package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:string; not null; size:50"`
	Price      int64  `gorm:"type:integer; not null"`
	Stock      int16  `gorm:"type:integer; not null"`
	Unit       string `gorm:"type:string; not null; size:7"`
	RetailID   uint   //Item is BelongsTo Retail
	Retail     Retail
	SaleDetail []SaleDetail //Item is HasMany of SaleDetail
	ItemDetail []ItemDetail //Item is HasMany of ItemDetail
	Returns    []Return     //Item is HasMany of Return
	CreatedAt  uint         `gorm:"type:integer; not null"`
	UpdatedAt  uint         `gorm:"type:integer; not null"`
	DeletedAt  uint         `gorm:"type:integer; default:null"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {

	// generate time
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
