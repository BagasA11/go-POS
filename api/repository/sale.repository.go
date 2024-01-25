package repository

import (
	"BagasA11/go-POS/configs"

	"gorm.io/gorm"
)

type SaleRepository struct {
	Db *gorm.DB
}

func NewSaleRepository() *SaleRepository {
	return &SaleRepository{
		Db: configs.GetDB(),
	}
}
