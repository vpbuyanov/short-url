package app

import (
	"context"
	"fmt"

	"github.com/vpbuyanov/short-url/internal/configs"
	"github.com/vpbuyanov/short-url/internal/repos"
	"github.com/vpbuyanov/short-url/internal/server"
	"github.com/vpbuyanov/short-url/internal/usecase"
)

type App struct {
	cfg *configs.Config
}

func New(config *configs.Config) App {
	return App{
		cfg: config,
	}
}

func (app *App) Run(ctx context.Context) error {
	s := server.New(&app.cfg.Server)

	reposURL := repos.New()
	urlUC := usecase.New(reposURL, &app.cfg.Server)

	err := s.Start(ctx, urlUC)
	if err != nil {
		return fmt.Errorf("can't start server, err: %w", err)
	}

	return nil
}
