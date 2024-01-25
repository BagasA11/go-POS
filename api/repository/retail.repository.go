package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"
	"time"

	"gorm.io/gorm"
)

/*Define Retail Repositry structure*/
type RetailRepository struct {
	db *gorm.DB
}

/*create Retail Repository instance*/
func NewRetailRepository() *RetailRepository {
	return &RetailRepository{
		db: configs.GetDB(),
	}
}

/*create new record*/
func (retailRepo *RetailRepository) Create(retail models.Retail) error {
	tx := retailRepo.db.Begin()
	err := tx.Create(&retail).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

/*find Retail object by ID*/
func (retailRepo *RetailRepository) FindByID(ID uint) (models.Retail, error) {
	var rtl models.Retail
	err := retailRepo.db.Where("id = ?", ID).First(&rtl).Error
	return rtl, err
}

/*find retail by email attribute*/
func (rtlRepo *RetailRepository) FindEmail(Email string) (models.Retail, error) {
	var rtl models.Retail
	err := rtlRepo.db.Where("email = ?", Email).First(&rtl).Error
	return rtl, err
}

/*retrieve entire record of Retail object*/
func (rtlRepo *RetailRepository) All() ([]models.Retail, error) {
	var rtl []models.Retail
	err := rtlRepo.db.Find(&rtl).Error
	return rtl, err
}

/*change Retail account attribute*/
func (rtlRepo *RetailRepository) Update(rtl *models.Retail) error {
	tx := rtlRepo.db.Begin()
	err := tx.Model(&models.Retail{}).Where("id = ?", rtl.ID).Updates(&rtl).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

/*decrease cashier quota where new cashier employee is registered*/
func (rtlRepo *RetailRepository) DecreaseCashier(rtl *models.Retail) error {
	tx := rtlRepo.db.Begin()
	err := tx.Model(&models.Retail{}).Where("id = ? AND cashieravb - 1 > ?", rtl.ID, 0).Update("cashieravb", rtl.CashierAvb-1).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

/*delete Retail record*/
func (rtlRepo *RetailRepository) Delete(ID uint) error {
	tx := rtlRepo.db.Begin()
	err := tx.Model(&models.Retail{}).Where("id = ?", ID).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
