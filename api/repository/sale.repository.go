package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"

	"gorm.io/gorm"
)

type SaleRepository struct {
	Db *gorm.DB
}

/*Sale Repository instance*/
func NewSaleRepository() *SaleRepository {
	return &SaleRepository{
		Db: configs.GetDB(),
	}
}

/*create new sale record*/
func (SlRepo *SaleRepository) Create(sale models.Sale) error {
	tx := SlRepo.Db.Begin()
	err := tx.Create(&sale).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

/*find sales by Id*/
func (SlRepo *SaleRepository) FindByID(id uint) (models.Sale, error) {
	var sale models.Sale
	err := SlRepo.Db.Where("id = ?", id).First(&sale).Error
	return sale, err
}
