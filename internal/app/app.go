package app

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/configs"
	"github.com/vpbuyanov/short-url/internal/server"
)

type App interface {
	Run(ctx context.Context) error
}

type app struct {
	cfg    *configs.Config
	logger *logrus.Logger
}

func New(config *configs.Config, log *logrus.Logger) App {
	return &app{
		cfg:    config,
		logger: log,
	}
}

func (app *app) Run(ctx context.Context) error {
	s := server.New(app.logger, app.cfg.Server)

	s.Start(ctx)
	return nil
}
