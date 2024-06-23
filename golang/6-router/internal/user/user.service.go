package user

import (
	"github.com/isd-sgcu/onboarding-backend/golang/5-architecture/apperror"
	"github.com/isd-sgcu/onboarding-backend/golang/5-architecture/internal/model"
)

type Service interface {
	AddUser(itemId int, quantity int) *apperror.AppError
	RemoveUser(id string) *apperror.AppError
	Checkout() (*[]model.User, *apperror.AppError)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (c *serviceImpl) AddUser(itemId int, quantity int) *apperror.AppError {
	user := model.User{
		ItemId:   itemId,
		Quantity: quantity,
	}

	err := c.repo.AddUser(&user)
	if err != nil {
		return apperror.InternalServerError(err.Error())
	}

	return nil
}

func (c *serviceImpl) RemoveUser(id string) *apperror.AppError {
	err := c.repo.RemoveUser(id)
	if err != nil {
		return apperror.InternalServerError(err.Error())
	}

	return nil
}

func (c *serviceImpl) Checkout() (*[]model.User, *apperror.AppError) {
	var result []model.User
	err := c.repo.GetUsers(&result)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	return &result, nil
}
