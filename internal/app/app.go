package app

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/configs"
	"github.com/vpbuyanov/short-url/internal/server"
)

type App interface {
	Run(context.Context, *logrus.Logger) error
}

type app struct {
	cfg *configs.Config
}

func New(config *configs.Config) App {
	return &app{
		cfg: config,
	}
}

func (app *app) Run(ctx context.Context, log *logrus.Logger) error {
	s := server.New(&app.cfg.Server)

	s.Start(ctx, log)
	return nil
}
