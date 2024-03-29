package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"

	"gorm.io/gorm"

	"time"
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

/*update object porperty values*/
func (IR *ItemRepository) Update(item models.Item) error {
	tx := IR.Db.Begin()
	err := tx.Model(&models.Item{}).Where("id = ?", item.Id).Updates(&item).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (IR *ItemRepository) UpdateStock(id uint, stock int16) error {
	tx := IR.Db.Begin()
	err := tx.Model(&models.Item{}).Where("id = ?", id).Update("stock", abs(stock)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func abs(num int16) int16 {
	if num < 0 {
		return num * -1
	}
	return num
}

/*delete object*/
func (IR *ItemRepository) Delete(id uint) error {
	tx := IR.Db.Begin()
	err := tx.Model(&models.Item{}).Where("id = ?", id).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
