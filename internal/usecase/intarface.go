package usecase

import "github.com/saushew/wb_testtask/internal/entity"

type (
	// Order .
	Order interface {
		Create(*entity.Order) error
		GetByID(string) (*entity.Order, error)
		Recover() error
	}

	// OrderRepo .
	OrderRepo interface {
		Store(*entity.Order) error
		GetHistory() (map[string]entity.Order, error)
	}

	// OrderCache .
	OrderCache interface {
		Append(*entity.Order)
		Refill(map[string]entity.Order)
		GetByID(string) (*entity.Order, error)
	}
)
