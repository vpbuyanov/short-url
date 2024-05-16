package main

import (
	"context"

	"github.com/gofiber/fiber/v2/log"

	"github.com/vpbuyanov/short-url/internal/app"
	"github.com/vpbuyanov/short-url/internal/configs"
)

func main() {
	ctx := context.Background()
	cfg := configs.LoadConfig()

	a := app.New(cfg)

	err := a.Run(ctx)
	if err != nil {
		log.Error(err)
		return
	}
}
