package cache

import (
	"fmt"
	"sync"

	"github.com/saushew/wb_testtask/internal/entity"
)

// OrderCache .
type OrderCache struct {
	m     *sync.Mutex
	cache map[string]entity.Order
}

// New .
func New() *OrderCache {
	return &OrderCache{
		m:     new(sync.Mutex),
		cache: make(map[string]entity.Order),
	}
}

// Append .
func (c *OrderCache) Append(order *entity.Order) {
	c.cache[order.OrderUID] = *order
}

// Refill .
func (c *OrderCache) Refill(orders map[string]entity.Order) {
	c.m.Lock()
	for id, order := range orders {
		c.cache[id] = order
	}
	c.m.Unlock()
}

// GetByID .
func (c *OrderCache) GetByID(id string) (*entity.Order, error) {
	order, ok := c.cache[id]
	if !ok {
		return nil, fmt.Errorf("Order not found")
	}
	
	return &order, nil
}
