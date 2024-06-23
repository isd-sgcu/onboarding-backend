package user

import (
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	AddUser(in *model.User) error
	GetUsers(result *[]model.User) error
	RemoveUser(id string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) AddUser(in *model.User) error {
	return r.db.Create(&in).Error
}

func (r *repositoryImpl) GetUsers(result *[]model.User) error {
	return r.db.Model(&model.User{}).Find(result).Error
}

func (r *repositoryImpl) RemoveUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.User{}).Error
}
