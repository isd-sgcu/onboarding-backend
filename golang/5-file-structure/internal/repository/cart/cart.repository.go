package cart

import (
	"github.com/isd-sgcu/onboarding-backend/golang/4-database/model"
	"gorm.io/gorm"
)

type Repository interface {
	// when using a database, so many errors can occur. Therefore, we need to return an error
	AddOrder(in *model.Order) error
	GetOrders(result *[]model.Order) error // orders are put into result
	RemoveOrder(id string) error           // use common model's ID field (uuid but we convert it to string)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) AddOrder(in *model.Order) error {
	return r.db.Create(&in).Error
}

func (r *repositoryImpl) GetOrders(result *[]model.Order) error {
	return r.db.Model(&model.Order{}).Find(result).Error
}

func (r *repositoryImpl) RemoveOrder(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.Order{}).Error
}
