package cart

import "github.com/isd-sgcu/onboarding-backend/golang/3-config-env/repository/cart"

type Service interface {
	AddOrder(itemId int, quantity int)
	RemoveOrder(itemId int)
	Checkout() int
}

type serviceImpl struct {
	repo cart.Repository
}

func NewService(repo cart.Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (c *serviceImpl) AddOrder(itemId int, quantity int) {
	c.repo.AddOrder(itemId, quantity)
}

func (c *serviceImpl) RemoveOrder(itemId int) {
	c.repo.RemoveOrder(itemId)
}

func (c *serviceImpl) Checkout() int {
	total := 0
	for _, quantity := range c.repo.GetOrders() {
		total += quantity
	}
	return total
}
