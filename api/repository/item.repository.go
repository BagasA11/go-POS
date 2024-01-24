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

type ItemRepository struct {
	Db *gorm.DB
}

/*create Item Repository instance*/
func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		Db: configs.GetDB(),
	}
}

/*create new record*/
func (IR *ItemRepository) Create(item models.Item) error {
	tx := IR.Db.Begin()
	err := tx.Create(&item).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

/*find item by ID method*/
func (IR *ItemRepository) FindByID(id uint) (models.Item, error) {
	var item models.Item
	err := IR.Db.Where("id = ?", id).First(&item).Error
	// SELECT * FROM items WHERE items.id = id;
	return item, err
}

/*find item by name*/
func (IR *ItemRepository) FindName(name string) ([]models.Item, error) {
	var item []models.Item
	//LIKE
	err := IR.Db.Where("name = ?", "%"+name+"%").Find(&item).Error
	// SELECT * FROM items WHERE items.name LIKE '%{name}%';
	return item, err
}

/*retrieve entire record of Item object*/
func (IR *ItemRepository) All() ([]models.Item, error) {
	var item []models.Item
	err := IR.Db.Find(&item).Error
	return item, err
}
