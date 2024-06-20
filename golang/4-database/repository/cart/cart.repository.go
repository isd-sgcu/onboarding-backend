package cart

type Repository interface {
	AddOrder(itemId int, quantity int)
	GetOrders() map[int]int
	RemoveOrder(itemId int)
}

type repositoryImpl struct {
	db map[int]int
}

func NewRepository() Repository {
	return &repositoryImpl{
		db: make(map[int]int),
	}
}

func (c *repositoryImpl) AddOrder(itemId int, quantity int) {
	c.db[itemId] += quantity
}

func (c *repositoryImpl) GetOrders() map[int]int {
	return c.db
}

func (c *repositoryImpl) RemoveOrder(itemId int) {
	delete(c.db, itemId)
}
