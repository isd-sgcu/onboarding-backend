package service

// In Java, this may have been called ICart
// The 'C' in Cart is captialised because it is exported
type Cart interface {
	AddOrder(itemId int, quantity int)
	RemoveOrder(itemId int)
	Checkout() int
}

// This is not exported because it is not capitalised, we only want use the interface outside this package
type cartImpl struct {
	orders map[int]int
}

// We need to export this in order to make shopping carts
func NewEmptyCart() Cart { // The shopping cart MUST implement the Cart interface (have every method in the interface but can have other methods too)
	return cartImpl{
		orders: make(map[int]int), //make() is to init the map
	}
}

func NewCartWithOrders(orders map[int]int) Cart {
	return cartImpl{
		orders: orders,
	}
}

// (c cartImpl) tells Go that this function is a method of struct cartImpl
func (c cartImpl) AddOrder(itemId int, quantity int) {
	c.orders[itemId] += quantity
}

func (c cartImpl) RemoveOrder(itemId int) {
	delete(c.orders, itemId)
}

func (c cartImpl) Checkout() int {
	total := 0
	for _, quantity := range c.orders { // _ is a blank identifier, it is used to ignore the index
		// range returns the key, value of map as pairs
		total += quantity
	}
	return total
}
