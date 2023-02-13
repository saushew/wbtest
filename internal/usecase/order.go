package usecase

import (
	"fmt"

	"github.com/saushew/wb_testtask/internal/entity"
)

// OrderUseCase .
type OrderUseCase struct {
	repo  OrderRepo
	cache OrderCache
}

// New .
func New(r OrderRepo, c OrderCache) *OrderUseCase {
	return &OrderUseCase{
		repo:  r,
		cache: c,
	}
}

// Create .
func (uc *OrderUseCase) Create(order *entity.Order) error {
	err := uc.repo.Store(order)
	if err != nil {
		return fmt.Errorf("OrderUseCase - Create - uc.repo.Store: %w", err)
	}

	uc.cache.Append(order)

	return nil
}

// GetByID .
func (uc *OrderUseCase) GetByID(id string) (*entity.Order, error) {
	order, err := uc.cache.GetByID(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// Recover .
func (uc *OrderUseCase) Recover() error {
	orders, err := uc.repo.GetHistory()
	if err != nil {
		return fmt.Errorf("OrderUseCase - Recover - uc.repo.GetHistory: %w", err)
	}

	uc.cache.Refill(orders)

	return nil
}
