package user

import (
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(in *model.User) error
	FindOne(id string, user *model.User) error
	Delete(id string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) Create(in *model.User) error {
	return r.db.Create(&in).Error
}

func (r *repositoryImpl) FindOne(id string, user *model.User) error {
	return r.db.First(user, "id = ?", id).Error
}

func (r *repositoryImpl) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.User{}).Error
}
