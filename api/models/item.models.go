package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id         string `gorm:"primaryKey"`
	Name       string `gorm:"type:string; not null; size:50"`
	Price      int64  `gorm:"type:integer; not null"`
	Stock      int16  `gorm:"type:integer; not null"`
	Unit       string `gorm:"type:string; not null; size:7"`
	RetailID   uint
	Retail     Retail
	SaleDetail []SaleDetail
	CreatedAt  uint `gorm:"type:integer; not null"`
	UpdatedAt  uint `gorm:"type:integer; not null"`
	DeletedAt  uint `gorm:"type:integer; default:null"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	item.Id = uuid.NewString()
	// generate time
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
