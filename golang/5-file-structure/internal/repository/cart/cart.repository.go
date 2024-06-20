package cart

import (
	"github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	AddOrder(in *model.Order) error
	GetOrders(result *[]model.Order) error
	RemoveOrder(id string) error
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
