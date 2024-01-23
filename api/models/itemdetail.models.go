package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemDetail struct {
	gorm.Model
	UUID   string `gorm:"type:string; primaryKey"`
	ItemID uint
	item   Item
}

func (itdetail *ItemDetail) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	itdetail.UUID = uuid.NewString()
	return nil
}
