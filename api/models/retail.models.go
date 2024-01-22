package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Retail struct {
	gorm.Model
	ID       uint      `gorm:"type:integer; primaryKey"`
	Email    string    `gorm:"type:string; size:30; not null; unique"`
	contact  string    `gorm:"type:string; size:30; not null; unique"`
	Password string    `gorm:"type:string; not null"`
	Cashiers []Cashier `gorm:"foreignKey:RetailID"`
	Items    []Item    `gorm:"foreignKey:ItemId"`
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

	// // generate time
	// tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	// tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
