package main

import (
	"context"

	"github.com/vpbuyanov/short-url/internal/services"

	"github.com/vpbuyanov/short-url/internal/app"
	"github.com/vpbuyanov/short-url/internal/configs"
)

func main() {
	ctx := context.Background()
	cfg := configs.LoadConfig()

	logger, err := services.InitLogger(cfg.Logger)
	if err != nil {
		panic(err)
	}

	a := app.New(cfg)

	err = a.Run(ctx, logger)
	if err != nil {
		panic(err)
	}
}
