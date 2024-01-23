package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"

	"gorm.io/gorm"
)

type RetailRepository struct {
	db *gorm.DB
}

// initializing repository
func NewRetailRepository() *RetailRepository {
	return &RetailRepository{
		db: configs.GetDB(),
	}
}

// create new record
func (retailRepo *RetailRepository) CreateRetail(retail models.Retail) error {
	tx := retailRepo.db.Begin()
	err := tx.Create(&retail).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

// find by ID
func (retailRepo *RetailRepository) FindByID(ID uint) (models.Retail, error) {
	var rtl models.Retail

	err := retailRepo.db.Where("id = ?", ID).First(&rtl).Error

	return rtl, err
}

// get retail accounts from entire db record
func (rtlRepo *RetailRepository) FindAll() ([]models.Retail, error) {
	var rtl []models.Retail

	err := rtlRepo.db.Find(&rtl).Error

	return rtl, err
}

// change Retail account attribute
func Update() {}
