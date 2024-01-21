package models

// "fmt"
// "time"
// // "gorm.io/gorm"
// // "gorm.io/driver/mysql"
// "golang.org/x/crypto/bcrypt"

type User struct {
	ID        uint   `gorm:"type:integer; primaryKey"`
	Name      string `gorm:"type:string; size:20; not null;"`
	Email     string `gorm:"type:string; size:30; not null; unique"`
	owner     string `gorm:"type:boolean; not null;default:false"`
	Password  string `gorm:"type:string; size:72; not null"`
	CreatedAt uint   `gorm:"type:integer; not null"`
	UpdatedAt uint   `gorm:"type:integer; not null"`
	DeletedAt uint   `gorm:"type:integer; default:null"`
	RetailID  uint
}
