package repository

// "BagasA11/go-POS/api/models"
// 	"BagasA11/go-POS/configs"
// 	"time"
// 	"gorm.io/gorm"
import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"

	"gorm.io/gorm"
)

type CashierRepository struct {
	Db *gorm.DB
}

/*create Cashier Repository instance*/
func NewCashierRepository() *CashierRepository {
	return &CashierRepository{
		Db: configs.GetDB(),
	}
}

/*create new record*/
func (CR *CashierRepository) Create(csr models.Cashier) error {
	tx := CR.Db.Begin()
	err := tx.Create(&csr).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

/*retrieve cashier object if id given value match to Cashier ID */
func (CR *CashierRepository) FindByID(id uint) (models.Cashier, error) {
	var csr models.Cashier
	err := CR.Db.Where("id = ?", id).First(&csr).Error
	return csr, err
}

/*find Cashier Object by username attribute*/
func (CR *CashierRepository) FindByUsername(username string) (models.Cashier, error) {
	var csr models.Cashier
	err := CR.Db.Where("username = ?", username).First(&csr).Error
	return csr, err
}

/*retrieve entire record of Cashier object*/
func (CR *CashierRepository) All() ([]models.Cashier, error) {
	var csr []models.Cashier
	err := CR.Db.Find(&csr).Error
	return csr, err
}
