package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"

	"gorm.io/gorm"
)

type SaleDetailRepo struct {
	Db *gorm.DB
}

func NewSaleDetailRepo() *SaleDetailRepo {
	return &SaleDetailRepo{
		Db: configs.GetDB(),
	}
}

/*Create method to insert new data*/
func (SdtRepo *SaleDetailRepo) Create(sdt models.SaleDetail) error {
	tx := SdtRepo.Db.Begin()
	err := tx.Create(&sdt).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

/*create multiple record*/
func (SdtRepo *SaleDetailRepo) CreateMany(sdts []models.SaleDetail) error {
	tx := SdtRepo.Db.Begin()
	err := tx.Create(&sdts).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

/*Find Sale Detail data using id parameter*/
func (SdtRepo *SaleDetailRepo) FindID(id uint) (models.SaleDetail, error) {
	var sdt models.SaleDetail
	err := SdtRepo.Db.Where("id = ?", id).First(&sdt).Error
	return sdt, err
}
