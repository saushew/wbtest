package repo

import (
	"encoding/json"
	"fmt"

	"github.com/saushew/wb_testtask/internal/entity"
	"github.com/saushew/wb_testtask/pkg/postgres"
)

// OrderRepo .
type OrderRepo struct {
	*postgres.Postgres
}

// New .
func New(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{pg}
}

// Store .
func (r *OrderRepo) Store(order *entity.Order) error {

	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("OrderRepo - Store - json.Marshal: %w", err)
	}

	_, err = r.DB.Exec(
		`INSERT INTO orders (model)
		VALUES ($1)`, data,
	)
	if err != nil {
		return fmt.Errorf("OrderRepo - Store - r.DB.Exec: %w", err)
	}

	return nil
}

// GetHistory .
func (r *OrderRepo) GetHistory() (map[string]entity.Order, error) {
	row := r.DB.QueryRow(
		`SELECT COUNT(*)
		FROM orders`,
	)

	var numOfRows int
	err := row.Scan(&numOfRows)
	if err != nil {
		return nil, fmt.Errorf("OrderRepo - GetHistory - row.Scan: %w", err)
	}

	rows, err := r.DB.Query(
		`SELECT model
		FROM orders`,
	)
	if err != nil {
		return nil, fmt.Errorf("OrderRepo - GetHistory - r.DB.Query: %w", err)
	}

	var result = make(map[string]entity.Order, numOfRows)

	var order entity.Order
	for rows.Next() {
		var byteOrder []byte
		if err := rows.Scan(&byteOrder); err != nil {
			return nil, fmt.Errorf("OrderRepo - GetHistory - rows.Scan: %w", err)
		}

		if err := json.Unmarshal(byteOrder, &order); err != nil {
			return nil, fmt.Errorf("OrderRepo - GetHistory -  json.Unmarshal: %w", err)
		}

		result[order.OrderUID] = order
	}

	return result, nil
}
