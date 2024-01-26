package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"
	"time"

	"gorm.io/gorm"
)

type ReturnRepository struct {
	DB *gorm.DB
}

func NewReturnRepository() *ReturnRepository {
	return &ReturnRepository{
		DB: configs.GetDB(),
	}
}

func (rtnRepo *ReturnRepository) Create(rtn models.Return) error {
	tx := rtnRepo.DB.Begin()
	err := tx.Create(&rtn).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (rtnRepo *ReturnRepository) FindById(id uint) (models.Return, error) {
	var rtn models.Return
	err := rtnRepo.DB.Where("id = ?", id).First(&rtn).Error
	return rtn, err
}

func (rtnRepo *ReturnRepository) FindDateGTE(tm time.Time) ([]models.Return, error) {
	var rtn []models.Return
	err := rtnRepo.DB.Where("created_at >= ?", tm).Find(&rtn).Error
	return rtn, err
}

func (rtnRepo *ReturnRepository) FindDateLTE(tm time.Time) ([]models.Return, error) {
	var rtn []models.Return
	err := rtnRepo.DB.Where("created_at <= ?", tm).Find(&rtn).Error
	return rtn, err
}

func (rtnRepo *ReturnRepository) All() ([]models.Return, error) {
	var rtn []models.Return
	err := rtnRepo.DB.Find(&rtn).Error
	return rtn, err
}

func (rtnRepo *ReturnRepository) Update(rtn *models.Return) error {
	tx := rtnRepo.DB.Begin()
	err := tx.Model(&models.Return{}).Where("id = ?", rtn.ID).Updates(&rtn).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (rtnRepo *ReturnRepository) Delete(id uint) error {
	tx := rtnRepo.DB.Begin()
	err := tx.Model(&models.Return{}).Where("id = ?", id).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
