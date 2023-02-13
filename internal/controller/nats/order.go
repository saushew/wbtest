package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
	"github.com/saushew/wb_testtask/internal/entity"
	"github.com/saushew/wb_testtask/internal/usecase"
)

type orderHandlers struct {
	orderUseCase usecase.Order
}

func newOrderHandlers(o usecase.Order) stan.MsgHandler {
	or := &orderHandlers{o}
	return or.create()
}

func (o *orderHandlers) create() stan.MsgHandler {
	return func(m *stan.Msg) {
		var newOrder entity.Order

		err := json.Unmarshal(m.Data, &newOrder)
		if err != nil {
			return
		}

		err = o.orderUseCase.Create(&newOrder)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
