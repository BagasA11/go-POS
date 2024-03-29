package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Retail struct {
	gorm.Model
	ID         uint      `gorm:"type:integer; primaryKey"`
	Email      string    `gorm:"type:string; size:30; not null; unique"`
	Contact    string    `gorm:"type:string; size:30; not null; unique"`
	Password   string    `gorm:"type:string; not null"`
	CashierAvb uint      `gorm:"type:integer; not null"`
	CreatedAt  uint      `gorm:"type:integer; not null"`
	UpdatedAt  uint      `gorm:"type:integer; not null"`
	DeletedAt  uint      `gorm:"type:integer; default:null"`
	Cashiers   []Cashier `gorm:"foreignKey:RetailID"`
	Items      []Item    `gorm:"foreignKey:ItemId"`
}

// Before
func (retail *Retail) BeforeCreate(tx *gorm.DB) error {
	// hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(retail.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Println(retail.Password)
	tx.Statement.SetColumn("Password", hash)
	tx.Statement.SetColumn("CashierAvb", 2)
	// // generate time
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
