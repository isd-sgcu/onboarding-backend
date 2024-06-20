package user

import (
	"github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/model"
	user "github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/repository/user"
)

type Service interface {
	AddUser(itemId int, quantity int) error
	RemoveUser(id string) error
	Checkout() (*[]model.User, error)
}

type serviceImpl struct {
	repo user.Repository
}

func NewService(repo user.Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (c *serviceImpl) AddUser(itemId int, quantity int) error {
	user := model.User{
		ItemId:   itemId,
		Quantity: quantity,
	}

	err := c.repo.AddUser(&user)
	if err != nil {
		return err
	}

	return nil
}

func (c *serviceImpl) RemoveUser(id string) error {
	err := c.repo.RemoveUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (c *serviceImpl) Checkout() (*[]model.User, error) {
	var result []model.User
	err := c.repo.GetUsers(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
