package repository

import (
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
