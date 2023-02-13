package main

import (
	"log"

	"github.com/saushew/wb_testtask/config"
	"github.com/saushew/wb_testtask/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
