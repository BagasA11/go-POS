package models

type Retail struct {
	ID      uint   `gorm:"type:integer; primaryKey"`
	Email   string `gorm:"type:string; size:30; not null; unique"`
	contact string `gorm:"type:string; size:30; not null; unique"`
	subs    string `gorm:"type:boolean; null;"`
	Users   []User `gorm:"foreignKey:RetailID"`
}
