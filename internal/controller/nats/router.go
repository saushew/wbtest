package nats

import (
	"github.com/nats-io/stan.go"
	"github.com/saushew/wb_testtask/internal/usecase"
)

// NewHandler .
func NewHandler(o usecase.Order) stan.MsgHandler {
	return newOrderHandlers(o)
}
