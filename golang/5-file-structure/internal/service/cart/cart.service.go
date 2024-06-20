package cart

import (
	"github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/model"
	"github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/repository/cart"
)

type Service interface {
	AddOrder(itemId int, quantity int) error
	RemoveOrder(id string) error
	Checkout() (*[]model.Order, error)
}

type serviceImpl struct {
	repo cart.Repository
}

func NewService(repo cart.Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (c *serviceImpl) AddOrder(itemId int, quantity int) error {
	order := model.Order{
		ItemId:   itemId,
		Quantity: quantity,
	}

	err := c.repo.AddOrder(&order)
	if err != nil {
		return err
	}

	return nil
}

func (c *serviceImpl) RemoveOrder(id string) error {
	err := c.repo.RemoveOrder(id)
	if err != nil {
		return err
	}

	return nil
}

func (c *serviceImpl) Checkout() (*[]model.Order, error) {
	var result []model.Order
	err := c.repo.GetOrders(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
