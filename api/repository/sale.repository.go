package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"
	"time"

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

/*retrieve all record if transaction date is Greater or same with argument*/
func (SlRepo *SaleRepository) FindDateGTE(tArg time.Time) ([]models.Sale, error) {
	var sale []models.Sale
	err := SlRepo.Db.Where("created_at >= ?", tArg).Find(&sale).Error
	return sale, err
}

/*retrieve all record if transaction date is Less or same with argument*/
func (SlRepo *SaleRepository) FindDateLTE(tArg time.Time) ([]models.Sale, error) {
	var sale []models.Sale
	err := SlRepo.Db.Where("created_at <= ?", tArg).Find(&sale).Error
	return sale, err
}

/*get All*/
func (SlRepo *SaleRepository) All(id uint) ([]models.Sale, error) {
	var sale []models.Sale
	err := SlRepo.Db.Find(&sale).Error
	return sale, err
}

/*Update*/
func (SlRepo *SaleRepository) Update(sale *models.Sale) error {
	tx := SlRepo.Db.Begin()
	err := tx.Model(&models.Sale{}).Where("id = ?", sale.ID).Updates(&sale).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

/*Delete object*/
func (SlRepo *SaleRepository) Delete(id uint) error {
	tx := SlRepo.Db.Begin()
	err := tx.Model(&models.Sale{}).Where("id = ?", id).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
