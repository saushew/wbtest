package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/saushew/wb_testtask/config"
	"github.com/saushew/wb_testtask/internal/controller/http"
	"github.com/saushew/wb_testtask/internal/controller/nats"
	"github.com/saushew/wb_testtask/internal/usecase"
	"github.com/saushew/wb_testtask/internal/usecase/cache"
	"github.com/saushew/wb_testtask/internal/usecase/repo"
	"github.com/saushew/wb_testtask/pkg/httpserver"
	"github.com/saushew/wb_testtask/pkg/natsstreaming/subscription"
	"github.com/saushew/wb_testtask/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {

	// Repository
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	orderUseCase := usecase.New(
		repo.New(pg),
		cache.New(),
	)

	// Recover cache from db
	if err := orderUseCase.Recover(); err != nil {
		log.Fatal(fmt.Errorf("app - Run - orderUseCase.Recover: %w", err))
	}


	// Nats-streaming subscription
	natsHandler := nats.NewHandler(orderUseCase)
	natsSubscription, err := subscription.New(
		cfg.NatsStreamig.Subject,
		cfg.NatsStreamig.ClusterID,
		cfg.NatsStreamig.ClientID,
		cfg.NatsStreamig.URL,
		natsHandler,
	)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - natsstreaming.New: %w", err))
	}

	// HTTP Server
	handler := mux.NewRouter()
	http.NewRouter(handler, orderUseCase)
	httpServer := httpserver.New(handler, cfg.HTTP.Port)

	log.Println("app started")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-natsSubscription.Notify():
		log.Println(fmt.Errorf("app - Run - natsSubscription.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = natsSubscription.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - natsSubscription.Shutdown: %w", err))
	}
}
