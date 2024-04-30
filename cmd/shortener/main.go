package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/configs"
	"github.com/vpbuyanov/short-url/internal/app"
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

	a := app.New(cfg, logger)

	err = a.Run(ctx)
	if err != nil {
		panic(err)
	}
}
