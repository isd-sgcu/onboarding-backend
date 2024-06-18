package cart

import "github.com/isd-sgcu/onboarding-backend/golang/2-dependency-injection/repository/cart"

// We renamed Cart to Service since we moved the file to the cart package.
// This is group cart.Service and cart.Repository together under the cart package.
type Service interface {
	AddOrder(itemId int, quantity int)
	RemoveOrder(itemId int)
	Checkout() int
}

type serviceImpl struct {
	repo cart.Repository //via import "github.com/isd-sgcu/onboarding-backend/golang/2-dependency-injection/repository/cart"
}

func NewService(repo cart.Repository) Service {
	return &serviceImpl{ // We added & to return a pointer to the struct instead of the struct itself
		repo: repo,
	}
}

// We added * to tell Go that this function is a method of &serviceImpl
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
