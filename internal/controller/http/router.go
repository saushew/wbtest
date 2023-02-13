package http

import (
	"github.com/gorilla/mux"
	"github.com/saushew/wb_testtask/internal/usecase"
)

// NewRouter .
func NewRouter(handler *mux.Router, o usecase.Order) {
	newOrderRoutes(handler, o)
}
