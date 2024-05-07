package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/app"
	"github.com/vpbuyanov/short-url/internal/configs"
)

func main() {
	ctx := context.Background()
	cfg := configs.LoadConfig()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	err := logger.Level.UnmarshalText([]byte(cfg.Logger.LogLevel))
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal log level: %v", err))
	}

	a := app.New(cfg)

	err = a.Run(ctx, logger)
	if err != nil {
		panic(err)
	}
}
