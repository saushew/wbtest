package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saushew/wb_testtask/internal/usecase"
)

type orderRoutes struct {
	orderUseCase usecase.Order
}

func newOrderRoutes(handler *mux.Router, o usecase.Order) {
	or := &orderRoutes{o}

	handler.HandleFunc("/order", or.handleGetByID()).Methods("GET")
}

func (or *orderRoutes) handleGetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")

		u, err := or.orderUseCase.GetByID(idStr)
		if err != nil {
			or.error(w, r, http.StatusNotFound, err)
			return
		}

		or.respond(w, r, http.StatusOK, u)
	}
}

func (or *orderRoutes) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	or.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (or *orderRoutes) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
