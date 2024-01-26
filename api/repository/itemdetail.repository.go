package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"

	"gorm.io/gorm"
)

type ItemDetailRepository struct {
	Db *gorm.DB
}

func NewItmDtlRepo() *ItemDetailRepository {
	return &ItemDetailRepository{
		Db: configs.GetDB(),
	}
}

// Create new Item Detail
func (ItDtlRepo *ItemDetailRepository) Create(itmDtl models.ItemDetail) error {
	tx := ItDtlRepo.Db.Begin()
	err := tx.Create(&itmDtl).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// Get All Item Detail Data
func (ItDtlRepo *ItemDetailRepository) All() ([]models.ItemDetail, error) {
	var itmDtl []models.ItemDetail
	err := ItDtlRepo.Db.Find(&itmDtl).Error
	return itmDtl, err
}

/*find data by uuid*/
func (ItDtlRepo *ItemDetailRepository) Uuid(uuid string) (models.ItemDetail, error) {
	var itdtl models.ItemDetail
	err := ItDtlRepo.Db.Where("uuid = ?", uuid).First(&itdtl).Error
	return itdtl, err
}

/*Sold Out method change avail property to false*/
func (ItDtlRepo *ItemDetailRepository) SoldOut(itmDtl models.ItemDetail) error {
	tx := ItDtlRepo.Db.Begin()
	err := tx.Model(&models.ItemDetail{}).Where("uuid = ?", itmDtl.UUID).Update("avail", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

/*update Item Detail object*/
func (ItDtlRepo *ItemDetailRepository) Update(itmDtl *models.ItemDetail) error {
	tx := ItDtlRepo.Db.Begin()
	err := tx.Model(&models.ItemDetail{}).Where("uuid = ?", itmDtl.UUID).Updates(&itmDtl).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

/*Delete Item Detail Object*/
func (ItDtlRepo *ItemDetailRepository) Delete(uuid string) error {
	tx := ItDtlRepo.Db.Begin()
	err := tx.Delete(&models.ItemDetail{}, uuid).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
