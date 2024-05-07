package app

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/configs"
	"github.com/vpbuyanov/short-url/internal/repos"
	"github.com/vpbuyanov/short-url/internal/server"
	"github.com/vpbuyanov/short-url/internal/services"
)

type App struct {
	cfg *configs.Config
}

func New(config *configs.Config) App {
	return App{
		cfg: config,
	}
}

func (app *App) Run(ctx context.Context, log *logrus.Logger) error {
	s := server.New(&app.cfg.Server)

	url := services.New()
	reposURL := repos.New(url)

	err := s.Start(ctx, log, reposURL)
	if err != nil {
		return fmt.Errorf("can't start server, err: %w", err)
	}

	return nil
}
